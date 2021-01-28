package main

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/gtrace"
	"github.com/gogf/gf/util/gutil"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/trace"
)

const (
	JaegerEndpoint = "http://localhost:14268/api/traces"
	ServiceName    = "tracing-inprocess"
)

func main() {
	flush := initTracer()
	defer flush()

	ctx, span := gtrace.NewSpan(context.Background(), "main")
	defer span.End()

	user1 := GetUser(ctx, 1)
	g.Dump(user1)

	user100 := GetUser(ctx, 100)
	g.Dump(user100)
}

// initTracer creates a new trace provider instance and registers it as global trace provider.
func initTracer() func() {
	// Create and install Jaeger export pipeline.
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(JaegerEndpoint),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: ServiceName,
		}),
		jaeger.WithSDK(&trace.Config{DefaultSampler: trace.AlwaysSample()}),
	)
	if err != nil {
		g.Log().Fatal(err)
	}
	return flush
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
