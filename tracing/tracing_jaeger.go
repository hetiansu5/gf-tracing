package tracing

import (
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// InitJaeger initializes and registers jaeger to global TracerProvider.
//
// The output parameter `tp` is used for waiting exported trace spans to be uploaded,
// which is useful if your program is ending and you do not want to lose recent spans.
func InitJaeger(serviceName, endpoint string) (tp *trace.TracerProvider, err error) {
	var endpointOption jaeger.EndpointOption
	if strings.HasPrefix(endpoint, "http") {
		// HTTP.
		endpointOption = jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint))
	} else {
		// UDP.
		agentConfig := strings.SplitN(endpoint, ":", 2)
		if len(agentConfig) == 2 {
			endpointOption = jaeger.WithAgentEndpoint(jaeger.WithAgentHost(agentConfig[0]), jaeger.WithAgentPort(agentConfig[1]))
		} else {
			endpointOption = jaeger.WithAgentEndpoint(jaeger.WithAgentHost(agentConfig[0]))
		}
	}

	// Create the Jaeger exporter
	exp, err := jaeger.New(endpointOption)
	if err != nil {
		return nil, err
	}
	tp = trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in an Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}
