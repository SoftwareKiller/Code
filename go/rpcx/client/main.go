package main

import (
	"context"
	"fmt"

	"github.com/smallnest/rpcx/client"
)

type Request struct {
	S string
}

type Reply struct {
	S string
}

func main() {
	d, err := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8972", "")
	if err != nil {
		fmt.Printf("failed to create discovery: %v\n", err)
		return
	}
	xclient := client.NewXClient("Hello", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &Request{
		S: " rpcx client",
	}

	reply := &Reply{}
	err = xclient.Call(context.Background(), "Hello", args, reply)
	if err != nil {
		fmt.Printf("failed to call: %v\n", err)
	} else {
		fmt.Printf("req %s, reply %s\n", args.S, reply.S)
	}
}
