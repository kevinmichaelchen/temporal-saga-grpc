# temporal-saga-grpc

[![Lines Of Code](https://aschey.tech/tokei/github/kevinmichaelchen/temporal-saga-grpc?category=code&style=for-the-badge)](https://github.com/kevinmichaelchen/temporal-saga-grpc)

<p align="center">

![./docs/diagrams/architecture.svg](./docs/diagrams/architecture.svg)

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

### Step 1: Spin everything up

You can spin everything up with:

```shell
make
```

> [!NOTE]
> 
> Under the hood, we use [pkgx][pkgx] to run Temporal's [dev
> server][temporal-cli], and Docker to run [Jaeger][jaeger] (a telemetry
> backend).

[temporal-cli]: https://github.com/temporalio/cli
[pkgx]: https://pkgx.sh/
[jaeger]: https://www.jaegertracing.io

### Step 2: Observe the workflow

Let's get ready to observe this thing in action!

- View traces in Jaeger — [localhost:16686][jaeger-ui].
- View the workflow in Temporal's UI — [localhost:8233][temporal-ui].

[temporal-ui]: http://localhost:8233
[jaeger-ui]: http://localhost:16686

### Step 3: Start a Temporal Workflow

With `curl`:

```shell
curl -v http://localhost:8081/temporal.v1beta1.TemporalService/CreateOnboardingWorkflow \
  -H "Content-Type: application/json" \
  --data-binary @- <<EOF
  {
    "license": {
      "start": "2023-11-16T12:00:00Z",
      "end": "2024-01-16T12:00:00Z"
    },
    "org": {
      "name": "Org 1"
    },
    "profile": {
      "name": "Kevin Chen"
    }
  }
EOF
```

With [`HTTPie`](https://httpie.io/):

```shell
pkgx http POST \
  http://localhost:8081/temporal.v1beta1.TemporalService/CreateOnboardingWorkflow <<<'
  {
    "license": {
      "start": "2023-11-16T12:00:00Z",
      "end": "2024-01-16T12:00:00Z"
    },
    "org": {
      "name": "Org 1"
    },
    "profile": {
      "name": "Kevin Chen"
    }
  }
  '
```

With [`grpcurl`](https://github.com/fullstorydev/grpcurl):

```shell
pkgx grpcurl \
  -use-reflection \
  -plaintext \
  -d @ localhost:8081 \
  temporal.v1beta1.TemporalService/CreateOnboardingWorkflow <<EOM
{
  "license": {
    "start": "2023-11-16T12:00:00Z",
    "end": "2024-01-16T12:00:00Z"
  },
  "org": {
    "name": "Org 1"
  },
  "profile": {
    "name": "Kevin Chen"
  }
}
EOM
```
