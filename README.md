# temporal-saga-grpc

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

### Step 2: Start Temporal Workflow
```shell
# Start Temporal Workflow
go run saga/start/main.go

# Start Temporal Worker
go run saga/worker/main.go
```