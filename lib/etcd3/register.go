package etcdv3

import (
	"fmt"
	"strings"
	"time"

	log "cloud/lib/log"
	etcd3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"golang.org/x/net/context"
)

// Prefix should start and end with no slash
var Prefix = "etcd3_naming"
var client etcd3.Client
var serviceKey string
var stopSignal = make(chan bool, 1)

// Register
func Register(name string, addr string, target string, interval time.Duration, ttl int64) error {

	endpoints := make([]string, 0)
	if strings.Contains(target, ",") {
		endpoints = strings.Split(target, ",")
	} else {
		endpoints = append(endpoints, target)
	}

	client, err := etcd3.New(
		etcd3.Config{
			Endpoints: endpoints,
		})
	if err != nil {
		return err
	}

	serviceKey = fmt.Sprintf("/%s/%s/%s", Prefix, name, addr)

	go func() {
		//定时器
		ticker := time.NewTicker(interval)
		for {

			resp, err := client.Grant(context.TODO(), ttl)
			if err != nil {
				log.Error(fmt.Sprintf("Failed to grant,err=%v", err))
				continue
			}

			_, err = client.Get(context.Background(), serviceKey)
			if err != nil && err == rpctypes.ErrKeyNotFound {

				_, err = client.Put(context.TODO(), serviceKey, addr, etcd3.WithLease(resp.ID))
				if err != nil {
					log.Error(fmt.Sprintf("Failed to put,serviceKey=%+v,serviceValue=%+v,resp=%+v",
						serviceKey, addr, resp))
				}

			} else {

				_, err = client.Put(context.Background(), serviceKey, addr, etcd3.WithLease(resp.ID))
				if err != nil {
					log.Error(fmt.Sprintf("Failed to put,serviceKey=%+v,serviceValue=%+v,resp=%+v",
						serviceKey, addr, resp))
				}
			}

			select {
			case <-stopSignal:
				return
			case <-ticker.C:
			}
		}
	}()
	return nil
}

// UnRegister delete registered service from etcd
func UnRegister() error {

	stopSignal <- true
	stopSignal = make(chan bool, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := client.Delete(ctx, serviceKey)

	return err
}
