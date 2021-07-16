
```protobuf
message M1 {
    int64 m = 1;
}

message M2 {
    string m = 1;
}

// 定义 oneof 类型
message Params {
// 注意：位置标识
    oneof Param { // 最终oneof结构在go代码中的类型是接口类型，各个字段都实现了该接口
        M1 m1 = 1;
        M2 m2 = 2;
        int64 m3 = 3;
    }
    string p1 = 4;
}
```
proto生成的go代码结构


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


