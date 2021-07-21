package main

import (
	"context"
	"time"

	"github.com/gogf/gf-tracing/tracing"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/gtrace"
	"github.com/gogf/gf/util/gutil"
)

const (
	ServiceName       = "tracing-inprocess"
	JaegerUdpEndpoint = "localhost:6831"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tp, err := tracing.InitJaeger(ServiceName, JaegerUdpEndpoint)
	if err != nil {
		g.Log().Ctx(ctx).Fatal(err)
	}

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			g.Log().Ctx(ctx).Fatal(err)
		}
	}(ctx)
	ctx, span := gtrace.NewSpan(ctx, "main")
	defer span.End()

	user1 := GetUser(ctx, 1)
	g.Dump(user1)

	user100 := GetUser(ctx, 100)
	g.Dump(user100)
}

// GetUser retrieves and returns hard coded user data for demonstration.
func GetUser(ctx context.Context, id int) g.Map {
	ctx, span := gtrace.NewSpan(ctx, "GetUser")
	defer span.End()
	m := g.Map{}
	gutil.MapMerge(
		m,
		GetInfo(ctx, id),
		GetDetail(ctx, id),
		GetScores(ctx, id),
	)
	return m
}

// GetInfo retrieves and returns hard coded user info for demonstration.
func GetInfo(ctx context.Context, id int) g.Map {
	ctx, span := gtrace.NewSpan(ctx, "GetInfo")
	defer span.End()
	if id == 100 {
		return g.Map{
			"id":     100,
			"name":   "john",
			"gender": 1,
		}
	}
	return nil
}

// GetDetail retrieves and returns hard coded user detail for demonstration.
func GetDetail(ctx context.Context, id int) g.Map {
	ctx, span := gtrace.NewSpan(ctx, "GetDetail")
	defer span.End()
	if id == 100 {
		return g.Map{
			"site":  "https://goframe.org",
			"email": "john@goframe.org",
		}
	}
	return nil
}

// GetScores retrieves and returns hard coded user scores for demonstration.
func GetScores(ctx context.Context, id int) g.Map {
	ctx, span := gtrace.NewSpan(ctx, "GetScores")
	defer span.End()
	if id == 100 {
		return g.Map{
			"math":    100,
			"english": 60,
			"chinese": 50,
		}
	}
	return nil
}
