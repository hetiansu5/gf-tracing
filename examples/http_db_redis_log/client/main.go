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

	client := g.Client().Use(ghttp.MiddlewareClientTracing)
	// Add user info.
	idStr := client.Ctx(ctx).PostContent(
		"http://127.0.0.1:8199/user/insert",
		g.Map{
			"name": "john",
		},
	)
	if idStr == "" {
		g.Log().Ctx(ctx).Print("retrieve empty id string")
		return
	}
	g.Log().Ctx(ctx).Print("insert:", idStr)

	// Query user info.
	userJson := client.Ctx(ctx).GetContent(
		"http://127.0.0.1:8199/user/query",
		g.Map{
			"id": idStr,
		},
	)
	g.Log().Ctx(ctx).Print("query:", idStr, userJson)

	// Delete user info.
	deleteResult := client.Ctx(ctx).PostContent(
		"http://127.0.0.1:8199/user/delete",
		g.Map{
			"id": idStr,
		},
	)
	g.Log().Ctx(ctx).Print("delete:", idStr, deleteResult)
}
