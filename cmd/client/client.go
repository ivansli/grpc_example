package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	pb "github.com/ivansli/rpc_gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"math/rand"
	"time"
)

type MyData struct {
	Name string
	Age  int
}


func getFileOptions() {
	e := &pb.Error{}
	// 文件描述信息, message描述信息
	fileDEsc, messageDesc := descriptor.MessageDescriptorProto(e)

	fmt.Println(fileDEsc.GetName()) // proto 文件名
	fmt.Println(fileDEsc.GetPackage()) // 包名称
	fmt.Println(fileDEsc.GetSyntax()) // 语法协议
	fmt.Println(fileDEsc.GetService()) // 服务
	fmt.Println(fileDEsc.GetDependency()) // 依赖的包
	fmt.Println(fileDEsc.GetExtension()) // extension 扩展信息

	// 遍历message中每一个字段
	// field（*FieldDescriptorProto） 字段信息
	for _, field := range messageDesc.Field {
		fmt.Println(*field.Name, "--",field.GetDefaultValue())

		// *(field.Name) 字段名
		if *(field.Name) == "code" {
			// 字段信息中的 option 信息
			if v, err := proto.GetExtension(field.Options, pb.E_DefaultInt); err == nil {
				// 存在 option 信息，则获取对应值
				// 注意先进行断言，在获取值
				fmt.Printf("default:%v\n", *(v.(*int64)))
			}
		}

		if *(field.Name) == "reason" {
			if v, err := proto.GetExtension(field.Options, pb.E_DefaultString); err == nil {
				fmt.Printf("default:%v\n", *(v.(*string)))
			}
		}
	}

	// 直接获取对应某个字段的 option 信息
	if v, err := proto.GetExtension(messageDesc.Field[0].Options, pb.E_DefaultInt); err == nil {
		fmt.Printf("GetExtensions(usmf.Options, test.E_FileOpt1) file_opt1:%d", *(v.(*int64)))
	} else {
		fmt.Printf("GetExtensions(usmf.Options, test.E_FileOpt1) failed: %v", err)
	}
}



// grpc客户端
func main() {
	//getFileOptions()
	//return
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

	s, ok := status.FromError(err)
	fmt.Println(s.Code(), s.Message(), s.Details(), ok)

	fmt.Printf("%#v", err)

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
