#!/bin/bash

GOOGLE_APIS=$GOPATH"/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.5/third_party/googleapis"
PATH_PROTO=$(pwd)

echo "current path: $PATH_PROTO"


# 生成 gRpc 对应文件
# --go_out=plugins=grpc:.  最后一个点为生成的go文件的输出路径
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOOGLE_APIS --go_out=plugins=grpc:.  ./proto/*.proto
# 生成 gRpc-gateway 对应文件
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOOGLE_APIS --grpc-gateway_out=logtostderr=true:. ./proto/*.proto
