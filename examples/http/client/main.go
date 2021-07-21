package main

import (
	"context"

	"github.com/gogf/gf-tracing/tracing"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/net/gtrace"
)

const (
	ServiceName       = "tracing-http-client"
	JaegerUdpEndpoint = "localhost:6831"
)

func main() {
	tp, err := tracing.InitJaeger(ServiceName, JaegerUdpEndpoint)
	if err != nil {
		g.Log().Fatal(err)
	}
	ctx := context.TODO()
	defer tp.Shutdown(ctx)

	StartRequests(ctx)
}

func StartRequests(ctx context.Context) {
	ctx, span := gtrace.NewSpan(ctx, "StartRequests")
	defer span.End()

	ctx = gtrace.SetBaggageValue(ctx, "name", "john")

	client := g.Client().Use(ghttp.MiddlewareClientTracing)

	content := client.Ctx(ctx).GetContent("http://127.0.0.1:8199/hello")
	g.Log().Ctx(ctx).Print(content)
}
