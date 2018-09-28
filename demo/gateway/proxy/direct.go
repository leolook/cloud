package proxy

import (
	"cloud/lib/log"
	"cloud/lib/pool"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/transport"
	"strings"
	"time"
)

type Frame struct {
	Buf []byte
}

func Option() []grpc.ServerOption {

	poolMp := initPool()

	option := make([]grpc.ServerOption, 2)
	option[0] = grpc.CustomCodec(&Code{}) //设置数据解析码
	option[1] = grpc.UnknownServiceHandler(func(srv interface{}, stream grpc.ServerStream) error {

		sync := &SyncRequest{
			Stream: stream,
		}

		err := sync.init(poolMp)
		if err != nil {
			return err
		}

		resChan := sync.Request()
		for i := 0; i < 1; i++ {
			select {
			case v := <-resChan:
				{
					if v.Err != nil {
						return v.Err
					}

					err := stream.SendMsg(v.Buf)
					if err != nil {
						log.Errorf("Failed to send msg from server stream,err=%v", err)
						return err
					}
				}
			}
		}

		return nil
	}) //数据流拦截器

	return option
}

//初始化连接池
func initPool() map[string]*pool.ConnPool {

	p, err := pool.NewConnPool(func() (pool.ConnRes, error) {
		return grpc.Dial("127.0.0.1:8002", grpc.WithInsecure(), grpc.WithCodec(&Code{}))
	}, 10, 10*time.Second)

	mp := make(map[string]*pool.ConnPool)
	if err != nil {
		log.Errorf("Failed to new connect pool,err=%v", err)
		return mp
	}

	mp["demo"] = p

	return mp
}

type SyncRequest struct {
	Stream     grpc.ServerStream
	Method     string
	ConnPool   *pool.ConnPool
	Ctx        context.Context
	Buf        *Frame
	ClientName string
}

type SyncResponse struct {
	Buf *Frame
	Err error
}

func (s *SyncRequest) init(poolMp map[string]*pool.ConnPool) error {

	ctx := s.Stream.Context()
	ctxStream, ok := transport.StreamFromContext(ctx)
	if !ok {
		return fmt.Errorf("failed to stream from context,ok=%v", ok)
	}
	method := ctxStream.Method()

	var key string
	if strings.Contains(method, "/") {
		key = strings.Split(method, "/")[1]
	}

	key = strings.ToLower(key)
	start, end := strings.LastIndex(key, "."), strings.Index(key, "service")
	if start != -1 && end != -1 {
		tmp := []rune(key)
		tmp = tmp[start+1 : end]
		key = string(tmp)
	}

	connPool, ok := poolMp[key]
	if !ok {
		return fmt.Errorf("not found connect pool,ok=%v", ok)
	}

	s.Ctx, s.Method, s.ConnPool = ctx, method, connPool
	s.ClientName = key

	return nil
}

//异步请求
func (s *SyncRequest) Request() chan *SyncResponse {

	res := make(chan *SyncResponse, 1)

	go func() {
		//从连接池中获取连接
		conn, err := s.ConnPool.Get()
		if err != nil {
			res <- &SyncResponse{
				Err: err,
			}
			return
		}
		defer s.ConnPool.Put(conn)

		//获取服务器数据
		var buf Frame
		err = s.Stream.RecvMsg(&buf)
		if err != nil {
			log.Errorf("Failed to recv msg from server stream,err=%v", err)
			res <- &SyncResponse{
				Err: err,
			}
			return
		}

		//创建客户端流
		client, _ := conn.(*grpc.ClientConn)
		clientStream, err := client.NewStream(s.Ctx,
			&grpc.StreamDesc{
				ServerStreams: true,
				ClientStreams: true,
			}, s.Method)

		if err != nil {
			res <- &SyncResponse{
				Err: err,
			}
			return
		}

		//异步往客户端发送数据
		go func() {
			err = clientStream.SendMsg(&buf)
			if err != nil {
				log.Errorf("Failed to send msg by client stream(%s),err=%v", s.ClientName, err)
				res <- &SyncResponse{
					Err: err,
				}
				return
			}
		}()

		//接收客户端返回的数据
		err = clientStream.RecvMsg(&buf)
		if err != nil {
			log.Errorf("Failed to recv msg from client stream(%s),err=%v", s.ClientName, err)
			res <- &SyncResponse{
				Err: err,
			}
			return
		}

		res <- &SyncResponse{
			Buf: &buf,
		}

	}()

	return res
}
