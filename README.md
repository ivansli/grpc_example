# grpc_example
grpc demo

# 目录结构

```
.
├── Makefile
├── README.md
├── cmd
│    ├── protobufExample
│    │   └── protobufExample.go # 定义的各种扩展信息获取方式
│    ├── client
│    │   └── client.go # 客户端
│    └── server
│        ├── httpserver
│        │   └── httpServer.go # gateway网关服务端
│        └── rpcserver
│            └── server.go # grpc服务端
├── generate.sh # 自动生成 .pb.go文件
├── go.mod
├── go.sum
└── proto
|   ├── gateway.pb.go # pb对应go
|   ├── gateway.pb.gw.go # pb对应gateway网关
|   └── gateway.proto # proto文件
└── protobuf_example
    ├── hello.pb.go # pb对应go
    ├── hello.pb.gw.go # pb对应gateway网关
    └── hello.proto # proto文件 定义了各种扩展信息
```

# proto类型
- 普通类型
- 其他类型 
  - Any
  - timestamp
  - Empty
  - FieldMask
  

## Any使用

## Timestamp使用

## FieldMask使用


## Google API Design Guide (谷歌API设计指南)中文版
https://www.bookstack.cn/read/API-design-guide/API-design-guide-README.md

## gRpc
https://www.bookstack.cn/read/topgoer/1085263d5a0e6e12.md
