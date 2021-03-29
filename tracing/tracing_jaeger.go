package tracing

import (
	"github.com/gogf/gf/os/genv"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/trace"
	"strings"
)

const (
	// The service name.
	jaegerEnvServiceName = "JAEGER_SERVICE_NAME"
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
	if err := genv.Set(jaegerEnvServiceName, serviceName); err != nil {
		return nil, err
	}
	return jaeger.InstallNewPipeline(
		endpointOption,
		jaeger.WithProcessFromEnv(),
		jaeger.WithSDKOptions(
			trace.WithSampler(trace.AlwaysSample()),
		),
	)
}
