#/bin/sh

if [ "$1" ]
then 
	# 生成 pb 程序
	protoc --go_out=./pb/ --go-grpc_out=./pb/ ./proto/hello.proto
	
else
	echo "Wrong parameter. It should be like this, ./proto2go.sh file_name"
fi
