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

## TODOs
- `cmd/saga/start` should serve TemporalService gRPC API
- The TemporalService gRPC API should be reachable via `curl`

## Getting started

### Step 1: Temporal Server
```shell
git clone https://github.com/temporalio/docker-compose.git temporal-docker-compose
cd  temporal-docker-compose
docker-compose up
```
This may take a minute or two to pull all the Docker image layers.

### Step 2: Start Temporal Workflow
```shell
# Start Temporal Workflow
go run cmd/saga/start/main.go

# Start Temporal Worker
go run cmd/saga/worker/main.go
```