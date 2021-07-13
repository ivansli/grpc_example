package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/ivansli/rpc_gateway/proto"
	"google.golang.org/grpc"
	"math/rand"
	"time"
)

type MyData struct {
	Name string
	Age  int
}

// grpc客户端
func main() {
	// context对象
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	// 简历连接
	conn, _ := grpc.DialContext(ctx, ":8080", []grpc.DialOption{grpc.WithInsecure()}...)
	// 注意：关闭连接
	defer conn.Close()

	// 随机数生成器
	r := rand.New(rand.NewSource(time.Now().Unix()))


	// 通过连接生成Prod的grpc客户端
	client := pb.NewProdServiceClient(conn)
	// 客户端调用对应接口
	res, err := client.GetProdStock(ctx, &pb.ProdRequest{
		ProdId: r.Int31(),
	})

	fmt.Println(res, err)


	// 通过连接生成Order的grpc客户端
	client2 := pb.NewOrderClient(conn)
	re2, err := client2.GetOrder(ctx, &pb.OrderRequest{
		OrderId: r.Int63(),
	})

	fmt.Println(re2)

	var m MyData
	if err == nil {
		// 从re2获取Any字段的字节切片
		// re2.Any.GetValue() 为 []byte
		b := re2.Any.GetValue()
		// json反序列化到MyData结构体
		json.Unmarshal(b, &m)

		fmt.Println(m)
	}

	//t := re2.Time.AsTime().In(time.Local).Format("2006-01-02 15:04:05")
}
