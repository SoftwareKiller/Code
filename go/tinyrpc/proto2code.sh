#/bin/sh

if [ "$1" ]
then 
	# ηζ pb η¨εΊ
	protoc --go_out=plugins=grpc:./pb/ ./proto/hello.proto
	
else
	echo "Wrong parameter. It should be like this, ./proto2go.sh file_name"
fi
