package example

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func DefaultCase() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！")
	})
	s.Run()
}

func SetPort() {
	fmt.Println("Set Port")
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("go frame!")
	})
	s.SetPort(8100, 8200, 8300)
	s.Run()
}

func MutiInstance() {
	s1 := g.Server("s1")
	s1.SetPort(8080)
	s1.SetIndexFolder(true)
	s1.SetServerRoot("D:/master/Code/go/gf")
	s1.Start()

	s2 := g.Server("s2")
	s2.SetPort(8088)
	s2.SetIndexFolder(true)
	s2.SetServerRoot("D:/master/Code/go/gf")
	s2.Start()

	g.Wait()
}

func Hello1(r *ghttp.Request) {
	r.Response.Write("127.0.0.1: Hello1!")
}

func Hello2(r *ghttp.Request) {
	r.Response.Write("localhost: Hello2!")
}

func BindDomain() {
	s := g.Server()
	s.Domain("127.0.0.1").BindHandler("/", Hello1)
	s.Domain("localhost").BindHandler("/", Hello2)
	s.Run()
}
