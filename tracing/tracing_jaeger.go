package tracing

import (
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/trace"
	"strings"
)

// InitJaeger initializes and registers jaeger to global TracerProvider.
//
// The output parameter `flush` is used for waiting exported trace spans to be uploaded,
// which is useful if your program is ending and you do not want to lose recent spans.
func InitJaeger(serviceName, endpoint string) (flush func(), err error) {
	var endpointOption jaeger.EndpointOption
	if strings.HasPrefix(endpoint, "http") {
		// HTTP.
		endpointOption = jaeger.WithCollectorEndpoint(endpoint)
	} else {
		// UDP.
		endpointOption = jaeger.WithAgentEndpoint(endpoint)
	}
	return jaeger.InstallNewPipeline(
		endpointOption,
		jaeger.WithProcess(jaeger.Process{
			ServiceName: serviceName,
		}),
		jaeger.WithSDK(&trace.Config{
			DefaultSampler: trace.AlwaysSample(),
		}),
	)
}
