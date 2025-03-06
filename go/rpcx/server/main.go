package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/smallnest/rpcx/server"
)

type Request struct {
	S string
}

type Reply struct {
	S string
}

type Hello int

func (t *Hello) Hello(ctx context.Context, req *Request, reply *Reply) error {
	reply.S = fmt.Sprintf("%s, Hello -- %s", time.Now().String(), req.S)
	return nil
}

func main() {
	s := server.NewServer()
	s.RegisterName("Hello", new(Hello), "")
	err := s.Serve("tcp", ":8972")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
