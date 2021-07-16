package hello

import (
	"fmt"
	pb "github.com/ivansli/rpc_gateway/proto"
	"testing"
)

// 定义 oneof 类型
//type Params struct {
//	// 注意：位置标识
//	//
//	// Types that are valid to be assigned to Param:
//	//	*Params_M1
//	//	*Params_M2
//	//	*Params_M3
//	Param                isParams_Param `protobuf_oneof:"Param"` // isParams_Param 为接口类型
//	P1                   string         `protobuf:"bytes,4,opt,name=p1,proto3" json:"p1,omitempty"`
//	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
//	XXX_unrecognized     []byte         `json:"-"`
//	XXX_sizecache        int32          `json:"-"`
//}

func TestOneof(t *testing.T) {
	// oneof 中 第一个参数
	// 注意：oneof中的参数是组合而成的，即：message的名称+oneof的字段名称，使用时都需要加上pb前缀
	params := &pb.Params{
		// Param 是oneof的名称
		// Param是 isParams_Param接口类型
		// &pb.Params_M1 实现了该接口
		Param: &pb.Params_M1{M1: &pb.M1{M: 100}},
		P1:    "oneof",
	}
	// 在判断当前
	switchType(params)
	getOneofValue(params)

	// oneof 中 第二个参数
	params1 := &pb.Params{
		// Param 是oneof的名称
		// Param是 isParams_Param接口类型
		// &pb.Params_M2 实现了该接口
		Param: &pb.Params_M2{M2: &pb.M2{M: "hello"}},
		P1:    "oneof",
	}
	switchType(params1)
	getOneofValue(params1)

	// oneof 中 第三个参数
	params2 := &pb.Params{
		// Param 是oneof的名称
		// Param是 isParams_Param接口类型
		// &pb.Params_M3 实现了该接口
		Param: &pb.Params_M3{M3: 10},
		P1:    "oneof",
	}
	switchType(params2)
	getOneofValue(params2)
}

// 通过断言确定使用了oneof中的哪个字段
func switchType(params *pb.Params) {
	// 断言 switch.type
	switch v := params.Param.(type) {
	case *pb.Params_M1:
		fmt.Println("is *pb.Params_M1 :", v)
	case *pb.Params_M2:
		fmt.Println("is *pb.Params_M2 :", v)
	case *pb.Params_M3:
		fmt.Println("is *pb.Params_M3 :", v)
	default:
		fmt.Println("unKnow")
	}
}

// 使用proto封装方法获取oneof值
// 如果为nil 可以认为是空，该字段没有被使用
func getOneofValue(params *pb.Params) {
	m1 := params.GetM1()
	m2 := params.GetM2()
	m3 := params.GetM3()

	fmt.Println(m1, m2, m3)
}
