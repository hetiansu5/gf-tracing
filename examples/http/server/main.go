package main

import (
	"github.com/gogf/gf-tracing/tracing"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/net/gtrace"
)

const (
	ServiceName       = "tracing-http-server"
	JaegerUdpEndpoint = "localhost:6831"
)

func main() {
	_, err := tracing.InitJaeger(ServiceName, JaegerUdpEndpoint)
	if err != nil {
		g.Log().Fatal(err)
	}

	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareServerTracing)
		group.GET("/hello", HelloHandler)
	})
	s.SetPort(8199)
	s.Run()
}

func HelloHandler(r *ghttp.Request) {
	ctx, span := gtrace.NewSpan(r.Context(), "HelloHandler")
	defer span.End()

	value := gtrace.GetBaggageVar(ctx, "name").String()

	r.Response.Write("hello:", value)
}
