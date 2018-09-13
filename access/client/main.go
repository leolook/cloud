package main

import (
	"cloud/lib/pool"
	"cloud/pb"
	"context"
	"fmt"
	log "github.com/alecthomas/log4go"
	"google.golang.org/grpc"
	"time"
)

type Call func(ctx context.Context, req interface{}) (interface{}, error)

func main() {

	p, _ := pool.NewConnPool(func() (pool.ConnRes, error) {
		return grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	}, 1, 10*time.Second)

	ch := make(chan int, 0)
	go func() {
		conn, _ := p.Get()
		n, _ := conn.(*grpc.ClientConn)

		for i := 0; i < 10; i++ {
			client := pb.NewDemoServiceClient(n)
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			rsp, err := client.Demo(ctx, &pb.DemoReq{Id: 28})
			if err != nil {
				log.Error(err)
				return
			}
			log.Info(fmt.Sprintf("%d,%d", 1, rsp.Id))

			//time.Sleep(5 * time.Second)
		}

	}()

	<-ch
}
