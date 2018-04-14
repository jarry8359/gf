package main

import "gitee.com/johng/gf/g/net/ghttp"

func Hello1(r *ghttp.Request) {
    r.Response.Write("Hello World1!")
}

func Hello2(r *ghttp.Request) {
    r.Response.Write("Hello World2!")
}

func main() {
    s := ghttp.GetServer()
    s.Domain("127.0.0.1").BindHandler("/", Hello1)
    s.Domain("localhost").BindHandler("/", Hello2)
    s.Run()
}
