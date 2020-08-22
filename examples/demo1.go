package main

import (
	"fmt"
	"os"
	"log"
	"time"
	"context"
	"reflect"
	etcdcron "etcd-cron-job/cron"
	"github.com/coreos/etcd/clientv3"
)

func main() {
	fmt.Println("starting...")
	c, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{"etcd-host1:2379", "etcd-host2:2379"},
	})
	fmt.Println("client = ", c)
	fmt.Println("type = ", reflect.TypeOf(c))
	// etcdcron.(c)
	// etcdcron.WithEtcdMutexBuilder(c)
	cron, _ := etcdcron.New(etcdcron.WithEtcdMutexBuilder(c))
	cron.AddJob(etcdcron.Job{
		Name: "job0",
		Rhythm: "@every 3s",
		Func: func(ctx context.Context) error {
			// Handler
			log.Println("Every 10 seconds from", os.Getpid())
			return nil
		},
	})
	cron.Start(context.Background())
	time.Sleep(100 * time.Second)
	cron.Stop()
}