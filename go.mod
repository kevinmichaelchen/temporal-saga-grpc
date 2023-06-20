module github.com/kevinmichaelchen/temporal-saga-grpc

go 1.19

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.30.0-20230530223247-ca37dc8895db.1
	github.com/bufbuild/connect-go v1.8.0
	github.com/bufbuild/connect-grpchealth-go v1.1.1
	github.com/google/go-cmp v0.5.9
	github.com/rs/cors v1.9.0
	github.com/sethvargo/go-envconfig v0.9.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.4
	go.buf.build/bufbuild/connect-go/kevinmichaelchen/orgapis v1.10.2
	go.buf.build/bufbuild/connect-go/kevinmichaelchen/profileapis v1.10.2
	go.buf.build/bufbuild/connect-go/kevinmichaelchen/temporalapis v1.10.2
	go.buf.build/grpc/go/kevinmichaelchen/licenseapis v1.4.2
	go.buf.build/grpc/go/kevinmichaelchen/orgapis v1.4.2
	go.buf.build/grpc/go/kevinmichaelchen/profileapis v1.4.2
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.42.0
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/exporters/jaeger v1.16.0
	go.opentelemetry.io/otel/sdk v1.16.0
	go.opentelemetry.io/otel/trace v1.16.0
	go.temporal.io/sdk v1.23.1
	go.temporal.io/sdk/contrib/opentelemetry v0.2.0
	go.uber.org/fx v1.20.0
	golang.org/x/net v0.10.0
	google.golang.org/genproto v0.0.0-20230525154841-bd750badd5c6
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gogo/status v1.1.1 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	go.buf.build/grpc/go/envoyproxy/protoc-gen-validate v1.4.8 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.temporal.io/api v1.21.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
