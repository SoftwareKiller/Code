package main

import (
	"fmt"
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/hashicorp/consul/api"
)

func Register(prefixs []string) error {
	if len(prefixs) == 0 {
		panic("Must have param")
	}

	registration := new(api.AgentServiceRegistration)
	registration.Address = "127.0.0.1:8500"
	registration.Port = 8500
	registration.ID = fmt.Sprintf("http-%s", registration.Address)
	registration.Name = "Hello"
	registration.Tags = []string{"hello"}

	check := new(api.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://127.0.0.1:8888%s", "/ping")
	check.Timeout = "1s"                         //设置超时 1s
	check.Interval = "3s"                        //设置间隔 3s
	check.DeregisterCriticalServiceAfter = "10s" //check失败后30秒删除本服务，注销时间，相当于过期时间

	registration.Check = check

	config := api.DefaultConfig()
	config.Address = registration.Address
	config.Scheme = "http"
	client, err := api.NewClient(config)
	if err != nil {
		panic("Get Client Failed")
	}
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		g.Log().Errorf("Err msg:%s", err.Error())
		panic("ServiceRegister Failed")
	}
	return nil
}

func main() {
	s := g.Server()
	s.SetPort(8888)

	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！")
	})

	//增加健康检测
	s.BindHandler("/ping", func(r *ghttp.Request) {
		r.Response.Status = http.StatusOK
		r.Response.WriteExit("pong")
	})

	s.Start()

	Register([]string{"nginx"})

	g.Wait()
}

/*func main() {
	//example.DefaultCase()
	//example.SetPort()
	//example.MutiInstance()
<<<<<<< HEAD
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
=======
	example.BindDomain()
}*/
>>>>>>> a1a40cb921b31919c852069e6133b5235ccb3ea4
