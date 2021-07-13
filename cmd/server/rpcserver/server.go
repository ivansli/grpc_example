package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/ivansli/rpc_gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	_server_port = ":8080"
)

// 具体对外提供服务的对象
type Servers struct {
}

func (s *Servers) GetProdStock(ctx context.Context, in *pb.ProdRequest) (*pb.ProdResponse, error) {
	fmt.Println(in.ProdId)

	// 加入睡眠时间，模拟超时
	// 客户端通过ctx设置超时时间为1秒
	//time.Sleep(2 *time.Second)
	return &pb.ProdResponse{ProdStock: in.ProdId}, nil
}

// 实现pb中定义的 GetOrder 服务的结构体
type Order struct {
	Name string
	Age  int64
}

// 自定义结构体
type MyData struct {
	Name string
	Age  int
}

// 实现pb中定义的 GetOrder 服务
func (o *Order) GetOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
	fmt.Println(in.OrderId)

	// 加入睡眠时间，模拟超时
	// 客户端通过ctx设置超时时间为1秒
	//time.Sleep(2 *time.Second)

	//
	// 对于Any类型存储其他任意类型, 使用有几种方法
	// 1. Any 包含 TypeUrl string, Value []byte
	// 		Value一般用来存储数据，是字节切片
	//		可以直接将字符串赋值给Value字段 &any.Any{Value:[]byte("字符串")}
	//
	// 2. 先将准备传递的结构使用json.Marshal变成字节切片，然后赋值给Value字段
	//

	m := MyData{Name: "ivansli", Age: 35}
	b, _ := json.Marshal(m)

	// 返回的数据
	return &pb.OrderResponse{
		OrderId: in.OrderId,
		Name:    "order" + strconv.FormatInt(in.OrderId, 10),
		Time:    timestamppb.New(time.Now()), // 直接获取timestamp对象
		M:       map[string]string{"hello": "world"},
		//Any:     &any.Any{Value: []byte("hello world")},
		Any:   &any.Any{Value: b}, // Any是pb中Any类型
		Empty: &empty.Empty{},
	}, nil

	// 返回grpc错误code码以及信息
	//return nil, status.Error(codes.Unavailable, "error happend")
}

// grpc对外提供的服务
// 启动过程：
// 	1. 先启动grpc
//	2. 再启动http
func main() {
	l, err := net.Listen("tcp", _server_port)
	if err != nil {
		log.Fatal(err)
	}

	// 注册 Servers 对外提供的服务
	s := grpc.NewServer()

	// 注意：
	// 如果存在多个服务，则每一个都需要注册一下
	pb.RegisterProdServiceServer(s, &Servers{})
	pb.RegisterOrderServer(s, &Order{})

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
