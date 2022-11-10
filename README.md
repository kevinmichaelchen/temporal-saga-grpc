# temporal-saga-grpc

[![Lines Of Code](https://tokei.rs/b1/github/kevinmichaelchen/temporal-saga-grpc?category=code)](https://github.com/kevinmichaelchen/temporal-saga-grpc)

<p align="center">
<img width="400" src="./docs/design.png" />
</p>

This project demonstrates using <a target="_blank" href="https://temporal.io/">Temporal</a>
to orchestrate a <a target="_blank" href="https://microservices.io/patterns/data/saga.html">saga</a>
(effectively a distributed transaction) that interacts with multiple services
and has a robust, edge-case-proof rollback strategy, as well as durable function
execution.

**Temporal abstracts away failures.**

The upstream microservices that are called during the workflow all use gRPC.

## Resources:
* <a target="_blank" href="https://github.com/temporalio/samples-go/blob/main/saga/workflow.go">temporalio/samples-go</a>
* <a target="_blank" href="https://github.com/temporalio/money-transfer-project-template-go/blob/main/workflow.go">money-transfer-project-template-go</a>
* <a target="_blank" href="https://www.swyx.io/why-temporal/">swyx — Why Temporal?</a>
* <a target="_blank" href="https://youtu.be/-KWutSkFda8">YouTube — Intro to Temporal with Go SDK</a>

## Getting started

### Step 0: Spin up Temporal and Jaeger
We use Docker Compose for
[Temporalite](https://github.com/temporalio/temporalite)
(a lighter-weight version of Temporal's
[docker-compose](https://github.com/temporalio/docker-compose) repo),
as well as for [Jaeger](https://www.jaegertracing.io/) (a telemetry backend).

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
go run cmd/saga/start/main.go
```

### Step 3: Start upstream gRPC Services
```shell
go run cmd/svc/license/main.go
go run cmd/svc/org/main.go
go run cmd/svc/profile/main.go
```

### Step 4: Start Temporal Workflow
```shell
curl -v http://localhost:8081/com.teachingstrategies.temporal.v1beta1.TemporalService/CreateLicense \
  -H "Content-Type: application/json" \
  -d '{"license": {"name": "L1"}, "org": {"name": "Org1"}, "profile": {"name": "Kevin Chen"}}'

http POST \
  http://localhost:8081/com.teachingstrategies.temporal.v1beta1.TemporalService/CreateLicense \
    license:='{"name": "L1"}' \
    org:='{"name": "Org1"}' \
    profile:='{"name": "Kevin Chen"}'
```

### Step 5: Check the UIs
#### Jaeger
See traces in Jaeger at [localhost:16686](http://localhost:16686).

#### Temporal Web UI
See the Temporal Workflow at [localhost:8233](http://localhost:8233).