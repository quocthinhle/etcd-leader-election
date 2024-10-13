package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/quocthinhle/leader-election-v2/election"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	POD_NAME  = os.Getenv("POD_NAME")
	ETCD_HOST = os.Getenv("ETCD_HOST")
)

func main() {
	sigs := make(chan os.Signal, 1)
	defer close(sigs)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Hello, world! I am ", POD_NAME)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{ETCD_HOST},
		DialTimeout:          1 * time.Second,
		DialKeepAliveTimeout: 2 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		panic(err)
	}

	// test the etcd connection
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()
	_, err = cli.Get(ctx, "test")
	if err != nil {
		fmt.Println("connect failed, err:", err)
		panic(err)
	}

	election := election.NewElection(cli, POD_NAME)
	err = election.ElectLeader()
	if err != nil {
		fmt.Println("failed to elect leader, err:", err)
		return
	}

	<-sigs
	fmt.Println("Gracefully shutting down")
}
