package main

import (
	"fmt"
	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	hello "github.com/ivansli/rpc_gateway/protobuf_example"
)

func helloRequest() {
	req := &hello.HelloRequest{}
	fileDEsc, messageDesc := descriptor.MessageDescriptorProto(req)

	fmt.Println("== 文件级别概览信息 ================")

	fmt.Println(fileDEsc.GetName())       // proto 文件名
	fmt.Println(fileDEsc.GetPackage())    // 包名称
	fmt.Println(fileDEsc.GetSyntax())     // 语法协议
	fmt.Println(fileDEsc.GetService())    // 服务
	fmt.Println(fileDEsc.GetDependency()) // 依赖的包
	fmt.Println(fileDEsc.GetExtension())  // extension 扩展信息


	fmt.Println("== 文件级别：fileDEsc.GetExtension() ================")
	fmt.Println(fileDEsc.Options) // 包级别的扩展信息

	// 获取包内所有 扩展信息
	for _, v := range fileDEsc.GetExtension() {
		fmt.Println(v.GetName()) // 包内所有扩展字段

		// 当扩展字段是 author 时 shi file级别扩展字段
		if v.GetName() == "author" {
			if v, err := proto.GetExtension(fileDEsc.Options, hello.E_Author); err == nil {
				// 存在 option 信息，则获取对应值
				// 注意先进行断言，在获取值
				fmt.Printf("author default:%v \n", *(v.(*string)))
			} else {
				fmt.Println(err)
			}
		}
	}

	fmt.Println("== 消息级别：messageDesc.Options ================")
	fmt.Println(messageDesc.Options)

	// message信息中的 option 信息
	if v, err := proto.GetExtension(messageDesc.Options, hello.E_WhichMsg); err == nil {
		// 存在 option 信息，则获取对应值
		// 注意先进行断言，在获取值
		fmt.Printf("which_msg default: %v\n", *(v.(*string)))
	} else {
		fmt.Println(err)
	}

	// 遍历message中每一个字段
	// field（*FieldDescriptorProto） 字段信息
	for _, field := range messageDesc.Field {
		fmt.Println(*field.Name, "--", field.GetName())

		// *(field.Name) 字段名
		if field.GetName() == "age" {

			// 字段信息中的 option 信息
			if v, err := proto.GetExtension(field.GetOptions(), hello.E_DefaultInt); err == nil {
				// 存在 option 信息，则获取对应值
				// 注意先进行断言，在获取值
				fmt.Printf("age default:%v\n", *(v.(*int32)))
			} else {
				fmt.Println(err)
			}
		}

		if *(field.Name) == "name" {
			if v, err := proto.GetExtension(field.Options, hello.E_DefaultString); err == nil {
				fmt.Printf("name default:%v\n", *(v.(*string)))
			}
		}
	}

	fmt.Println("== 另外一种获取Field信息写法 ================")

	// 直接获取对应某个字段的 option 信息
	if v, err := proto.GetExtension(messageDesc.Field[0].Options, hello.E_DefaultString); err == nil {
		fmt.Printf("GetExtensions(usmf.Options, test.E_FileOpt1) file_opt1:%v", *(v.(*string)))
	} else {
		fmt.Printf("GetExtensions(usmf.Options, test.E_FileOpt1) failed: %v", err)
	}
}

func helloResponse() {

}

// 获取扩展信息
func getExtendInfos() {
	helloRequest()
}

func main() {
	getExtendInfos()
}
