module github.com/kevinmichaelchen/temporal-saga-grpc

go 1.21

require (
	buf.build/gen/go/kevinmichaelchen/licenseapis/connectrpc/go v1.11.1-20230620011624-257a6358889c.1
	buf.build/gen/go/kevinmichaelchen/licenseapis/grpc/go v1.3.0-20230620011624-257a6358889c.1
	buf.build/gen/go/kevinmichaelchen/licenseapis/protocolbuffers/go v1.31.0-20230620011624-257a6358889c.1
	buf.build/gen/go/kevinmichaelchen/orgapis/connectrpc/go v1.11.1-20221121195307-c877b91e28b6.1
	buf.build/gen/go/kevinmichaelchen/orgapis/grpc/go v1.3.0-20221121195307-c877b91e28b6.1
	buf.build/gen/go/kevinmichaelchen/orgapis/protocolbuffers/go v1.31.0-20221121195307-c877b91e28b6.1
	buf.build/gen/go/kevinmichaelchen/profileapis/connectrpc/go v1.11.1-20221121195307-54c7cc0bb136.1
	buf.build/gen/go/kevinmichaelchen/profileapis/grpc/go v1.3.0-20221121195307-54c7cc0bb136.1
	buf.build/gen/go/kevinmichaelchen/profileapis/protocolbuffers/go v1.31.0-20221121195307-54c7cc0bb136.1
	buf.build/gen/go/kevinmichaelchen/temporalapis/connectrpc/go v1.11.1-20230620011625-99cb60d4ce70.1
	buf.build/gen/go/kevinmichaelchen/temporalapis/protocolbuffers/go v1.31.0-20230620011625-99cb60d4ce70.1
	connectrpc.com/connect v1.11.1
	github.com/bufbuild/connect-grpchealth-go v1.1.1
	github.com/bufbuild/protovalidate-go v0.3.4
	github.com/google/uuid v1.3.1
	github.com/rs/cors v1.10.1
	github.com/sethvargo/go-envconfig v0.9.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.45.0
	go.opentelemetry.io/otel v1.19.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.19.0
	go.opentelemetry.io/otel/sdk v1.19.0
	go.opentelemetry.io/otel/trace v1.19.0
	go.temporal.io/sdk v1.25.1
	go.temporal.io/sdk/contrib/opentelemetry v0.3.0
	go.uber.org/fx v1.20.1
	golang.org/x/net v0.17.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231016165738-49dd2c1f3d0b
	google.golang.org/grpc v1.59.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.31.0-20230914171853-63dfe56cc2c4.1 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230512164433-5d1fd1a340c9 // indirect
	github.com/bufbuild/connect-go v1.10.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gogo/status v1.1.1 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/cel-go v0.18.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.19.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.temporal.io/api v1.24.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20231012201019-e917dd12ba7a // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
