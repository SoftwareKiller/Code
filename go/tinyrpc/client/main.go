package main

import (
	"context"
	"log"
	"time"
	"tinyrpc/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%s", err.Error())
	}
	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Hello(ctx, &pb.String{Value: " rpc client"})
	if err != nil {
		log.Fatalf("could not hello, err :%s", err.Error())
	}

	log.Printf("Hello Server:%s", r.GetValue())
}
