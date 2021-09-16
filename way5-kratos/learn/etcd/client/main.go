package main

import (
	"context"
	v1 "etcd/api/helloworld/v1"
	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	clientv3 "go.etcd.io/etcd/client/v3"
	srcgrpc "google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"192.168.5.130:2379"},
	})
	if err != nil {
		panic(err)
	}
	r := registry.New(cli)

	connGrpc, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///etcd"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer connGrpc.Close()

	connHttp, err := http.NewClient(
		context.Background(),
		http.WithEndpoint("discovery:///etcd"),
		http.WithDiscovery(r),
		http.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer connHttp.Close()

	for {
		callGRPC(r, connGrpc)
		time.Sleep(time.Second)
	}
}

func callGRPC(r *registry.Registry, conn *srcgrpc.ClientConn) {
	client := v1.NewGreeterClient(conn)

	reply, err := client.SayHello(context.Background(), &v1.HelloRequest{Name: "Kratos"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v\n", reply)
}
