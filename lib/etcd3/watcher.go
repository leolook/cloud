package etcdv3

import (
	"fmt"
	etcd3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/naming"
	"time"
)

// watcher is the implementaion of grpc.naming.Watcher
type watcher struct {
	re            *resolver // re: Etcd Resolver
	client        etcd3.Client
	isInitialized bool
}

// Close do nothing
func (w *watcher) Close() {
}

// Next to return the updates
func (w *watcher) Next() ([]*naming.Update, error) {
	// prefix is the etcd prefix/value to watch
	prefix := fmt.Sprintf("/%s/%s/", Prefix, w.re.serviceName)
	// check if is initialized
	if !w.isInitialized {
		// query addresses from etcd
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		resp, err := w.client.Get(ctx, prefix, etcd3.WithPrefix())
		w.isInitialized = true
		if err == nil {
			addr := extractAddr(resp)
			//if not empty, return the updates or watcher new dir
			if l := len(addr); l != 0 {
				updates := make([]*naming.Update, l)
				for i := range addr {
					updates[i] = &naming.Update{Op: naming.Add, Addr: addr[i]}
				}
				return updates, nil
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rch := w.client.Watch(ctx, prefix, etcd3.WithPrefix())
	for v := range rch {
		for _, ev := range v.Events {
			switch ev.Type {
			case mvccpb.PUT:
				return []*naming.Update{{Op: naming.Add, Addr: string(ev.Kv.Value)}}, nil
			case mvccpb.DELETE:
				return []*naming.Update{{Op: naming.Delete, Addr: string(ev.Kv.Value)}}, nil
			}
		}
	}
	return nil, nil
}

func extractAddr(resp *etcd3.GetResponse) []string {

	addr := make([]string, 0)

	if resp == nil || resp.Kvs == nil {
		return addr
	}

	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			addr = append(addr, string(v))
		}
	}

	return addr
}
