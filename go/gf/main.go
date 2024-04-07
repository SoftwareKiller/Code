package main

import (
	"fmt"
	"gf/example"
	"net"
)

func main() {
	//example.DefaultCase()
	//example.SetPort()
	//example.MutiInstance()
	//example.BindDomain()
	/*ln, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal("Listen failed, err :", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accetp failed, err :", err)
			continue
		}

		go proc(conn)
	}*/
	example.StrBuildCost()
}

func proc(conn net.Conn) {
	defer conn.Close()

	var buf [1024]byte

	len, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("Read Failed, err:", err)
		return
	}

	fmt.Println(len)

	conn.Write([]byte("I am server"))
}
