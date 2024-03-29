# temporal-saga-grpc

[![Lines Of Code](https://aschey.tech/tokei/github/kevinmichaelchen/temporal-saga-grpc?category=code&style=for-the-badge)](https://github.com/kevinmichaelchen/temporal-saga-grpc)

<p align="center">

![./docs/public/diagrams/architecture.svg](./docs/public/diagrams/architecture.svg)

</p>

This project demonstrates using
<a target="_blank" href="https://temporal.io/">Temporal</a> to orchestrate a
<a target="_blank" href="https://microservices.io/patterns/data/saga.html">saga</a>
(effectively a distributed transaction) that interacts with multiple services
and has a robust, edge-case-proof rollback strategy, as well as durable function
execution. **Temporal abstracts away failures.**

## Getting started

### Step 1: Spin everything up

You can spin everything up with:

```shell
make
```

> [!NOTE]
>
> Under the hood, we use a bunch of tools (which you can read about [here][tech-stack]), but to run things seamlessly locally, the one tool you will need is [pkgx][pkgx].

[tech-stack]: https://kevinmichaelchen.github.io/temporal-saga-grpc/architecture/overview/
[pkgx]: https://pkgx.sh/

### Step 2: Observe the workflow

Let's get ready to observe this thing in action!

- View traces in Jaeger — [localhost:16686][jaeger-ui].
- View the workflow in Temporal's UI — [localhost:8233][temporal-ui].

[temporal-ui]: http://localhost:8233
[jaeger-ui]: http://localhost:16686

### Step 3: Start a Temporal Workflow

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
      "full_name": "Kevin Chen"
    }
  }
EOF
```
