# grpc_example
grpc demo

# 目录结构

```
.
├── Makefile
├── README.md
├── cmd
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
    ├── gateway.pb.go # pb对应go
    ├── gateway.pb.gw.go # pb对应gateway网关
    └── gateway.proto # proto文件
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
