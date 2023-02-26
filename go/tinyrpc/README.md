# tinyrpc

rpc练习代码

## 工具版本

- 系统版本 `Ubuntu 20.04.5`
- Go版本 `go version go1.20.1 linux/amd64`
- pb版本 `libprotoc 3.6.1`

## 安装方式

pb

```sh
sudo apt install libprotobuf-dev protobuf-compiler

go install github.com/golang/protobuf/protoc-gen-go@latest
```

## 工具使用

文件要带`option go_package="./;pb";`
```sh
proto2code.sh $file.proto
```