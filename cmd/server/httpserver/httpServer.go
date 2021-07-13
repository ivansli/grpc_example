package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/ivansli/rpc_gateway/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

const (
	_server_port = "localhost:8080"
)

// 启动的时候
// 1. 先启动grpc服务
// 2. 在启动gateway网关服务
func main() {
	// 生成新的mux
	gwmux := runtime.NewServeMux()
	// 使用 gateway.pb.gw.go 提供的注册方法
	// _server_port 为grpc的地址
	// 注意：
	// 	无证书时需要添加 []grpc.DialOption{grpc.WithInsecure()}，否则会有问题
	//
	// 注意：
	// 		如果存在多个服务，则每一个都需要注册一下

	// 1. 注册 Prod 服务
	pb.RegisterProdServiceHandlerFromEndpoint(context.Background(),
		gwmux, _server_port, []grpc.DialOption{grpc.WithInsecure()},
	)
	// 2. 注册 order 服务
	pb.RegisterOrderHandlerFromEndpoint(context.Background(),
		gwmux, _server_port, []grpc.DialOption{grpc.WithInsecure()},
	)

	// 通过http对外提供服务
	// http提供的服务端口地址
	s := &http.Server{
		Addr: ":8081",
		Handler: gwmux,
	}

	// http方式监听并提供服务
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
