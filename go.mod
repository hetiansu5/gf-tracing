module github.com/gogf/gf-tracing

go 1.11

require (
	github.com/gogf/gcache-adapter v0.1.2
	github.com/gogf/gf v1.16.5-0.20210715141900-88009ee2781f
	github.com/gogf/katyusha v0.1.3-0.20210402080039-f9c0f380474e
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.4.3
	go.opentelemetry.io/otel v1.0.0-RC1.0.20210720160618-63dfe64aaea1 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC1.0.20210720160618-63dfe64aaea1 // indirect
	go.opentelemetry.io/otel/sdk v1.0.0-RC1.0.20210720160618-63dfe64aaea1 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/grpc v1.36.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20201103155942-6e800b9b0161
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20201103155942-6e800b9b0161
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)
