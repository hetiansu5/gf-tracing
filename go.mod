module gftracing

go 1.11

require (
	github.com/gogf/gcache-adapter v0.1.0
	github.com/gogf/gf v1.15.5-0.20210311152939-41f2138b3938
	github.com/gogf/katyusha v0.1.2-0.20210203072924-61d9cc2addf5
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.4.3
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.16.0
	go.opentelemetry.io/otel/sdk v0.16.0
	google.golang.org/grpc v1.35.0
)

replace (
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20201103155942-6e800b9b0161
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20201103155942-6e800b9b0161
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)
