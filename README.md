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

The upstream microservices that are called during the workflow all use gRPC.

Inspiration:
* <a target="_blank" href="https://github.com/temporalio/samples-go/blob/main/saga/workflow.go">temporalio/samples-go</a>
* <a target="_blank" href="https://github.com/temporalio/money-transfer-project-template-go/blob/main/workflow.go">money-transfer-project-template-go</a>

## Getting started

### Step 1: Temporal Server
```shell
git clone https://github.com/temporalio/docker-compose.git temporal-docker-compose
cd  temporal-docker-compose
docker-compose up
```
This may take a minute or two to pull all the Docker image layers.

### Step 2: Start Temporal Worker
```shell
go run cmd/saga/worker/main.go
```

### Step 3: Start Temporal Workflow gRPC Server
```shell
go run cmd/saga/start/main.go
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