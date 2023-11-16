# temporal-saga-grpc

[![Lines Of Code](https://aschey.tech/tokei/github/kevinmichaelchen/temporal-saga-grpc?category=code&style=for-the-badge)](https://github.com/kevinmichaelchen/temporal-saga-grpc)

<p align="center">
<a href="https://raw.githubusercontent.com/kevinmichaelchen/temporal-saga-grpc/main/docs/design.png">
<img width="400" src="./docs/design.png" />
</a>
</p>

This project demonstrates using
<a target="_blank" href="https://temporal.io/">Temporal</a> to orchestrate a
<a target="_blank" href="https://microservices.io/patterns/data/saga.html">saga</a>
(effectively a distributed transaction) that interacts with multiple services
and has a robust, edge-case-proof rollback strategy, as well as durable function
execution.

**Temporal abstracts away failures.**

The upstream microservices that are called during the workflow all use gRPC.

## Resources:

- <a target="_blank" href="https://github.com/temporalio/samples-go/blob/main/saga/workflow.go">temporalio/samples-go</a>
- <a target="_blank" href="https://github.com/temporalio/money-transfer-project-template-go/blob/main/workflow.go">money-transfer-project-template-go</a>
- <a target="_blank" href="https://www.swyx.io/why-temporal/">swyx — Why
  Temporal?</a>
- <a target="_blank" href="https://youtu.be/-KWutSkFda8">YouTube — Intro to
  Temporal with Go SDK</a>

## Getting started

### Step 0: Spin up Temporal and Jaeger

We use Docker Compose for [Temporalite][temporalite] (a lighter-weight version 
of Temporal's [docker-compose][temporal-docker-compose] repo), as well as
for [Jaeger][jaeger] (a telemetry backend).

[temporalite]: https://github.com/temporalio/temporalite
[temporal-docker-compose]: https://github.com/temporalio/docker-compose
[jaeger]: https://www.jaegertracing.io

You can spin everything up with:

```shell
docker-compose up -d
```

Temporalite runs the Temporal server on 7233, and the Temporal Web UI on 8233.

### Step 1: Start Temporal Worker

```shell
go run cmd/saga/worker/main.go
```

### Step 2: Start Temporal Workflow gRPC Server

```shell
env GRPC_CONNECT_PORT=8081 go run cmd/saga/start/main.go
```

### Step 3: Start upstream gRPC Services

```shell
GRPC_CONNECT_PORT=9090 go run cmd/svc/license/main.go
GRPC_CONNECT_PORT=9091 go run cmd/svc/org/main.go
GRPC_CONNECT_PORT=9092 go run cmd/svc/profile/main.go
```

### Step 4: Start Temporal Workflow

```shell
curl -v http://localhost:8081/temporal.v1beta1.TemporalService/CreateOnboardingWorkflow \
  -H "Content-Type: application/json" \
  -d '{"license": {"name": "L1"}, "org": {"name": "Org1"}, "profile": {"name": "Kevin Chen"}}'

http POST \
  http://localhost:8081/temporal.v1beta1.TemporalService/CreateOnboardingWorkflow \
    license:='{"name": "L1"}' \
    org:='{"name": "Org1"}' \
    profile:='{"name": "Kevin Chen"}'
```

### Step 5: Check the UIs

#### Jaeger

See traces in Jaeger at [localhost:16686](http://localhost:16686).

#### Temporal Web UI

See the Temporal Workflow at [localhost:8233](http://localhost:8233).
