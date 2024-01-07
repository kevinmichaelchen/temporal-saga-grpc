module github.com/kevinmichaelchen/temporal-saga-grpc

go 1.21.5

require (
	buf.build/gen/go/kevinmichaelchen/licenseapis/connectrpc/go v1.14.0-20240106205504-60998e437501.1
	buf.build/gen/go/kevinmichaelchen/licenseapis/protocolbuffers/go v1.32.0-20240106205504-60998e437501.1
	buf.build/gen/go/kevinmichaelchen/orgapis/connectrpc/go v1.14.0-20240106023354-7971463d1bab.1
	buf.build/gen/go/kevinmichaelchen/orgapis/protocolbuffers/go v1.32.0-20240106023354-7971463d1bab.1
	buf.build/gen/go/kevinmichaelchen/profileapis/connectrpc/go v1.14.0-20240106023355-4074133bf516.1
	buf.build/gen/go/kevinmichaelchen/profileapis/protocolbuffers/go v1.32.0-20240106023355-4074133bf516.1
	buf.build/gen/go/kevinmichaelchen/temporalapis/connectrpc/go v1.14.0-20240105233422-8054bf72d102.1
	buf.build/gen/go/kevinmichaelchen/temporalapis/protocolbuffers/go v1.32.0-20240105233422-8054bf72d102.1
	connectrpc.com/connect v1.14.0
	connectrpc.com/grpchealth v1.3.0
	connectrpc.com/grpcreflect v1.2.0
	connectrpc.com/otelconnect v0.6.0
	connectrpc.com/vanguard v0.1.0
	github.com/bufbuild/protovalidate-go v0.4.3
	github.com/friendsofgo/errors v0.9.2
	github.com/google/uuid v1.5.0
	github.com/jackc/pgx/v5 v5.5.1
	github.com/rs/cors v1.10.1
	github.com/sethvargo/go-envconfig v0.9.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.4
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.15.0
	github.com/volatiletech/strmangle v0.0.6
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.21.0
	go.opentelemetry.io/otel/sdk v1.21.0
	go.temporal.io/sdk v1.25.1
	go.temporal.io/sdk/contrib/opentelemetry v0.3.0
	go.uber.org/fx v1.20.1
	golang.org/x/net v0.19.0
	google.golang.org/protobuf v1.32.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.32.0-20231115204500-e097f827e652.1 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gogo/status v1.1.1 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/cel-go v0.18.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.temporal.io/api v1.24.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230913181813-007df8e322eb // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230913181813-007df8e322eb // indirect
	google.golang.org/grpc v1.59.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
