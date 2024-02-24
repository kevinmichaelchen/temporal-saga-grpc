---
title: Tour of Code
description: A tour through the code
---

![/temporal-saga-grpc/diagrams/architecture.svg](/temporal-saga-grpc/public/diagrams/architecture.svg)

There are 5 Go services total: [3 “microservices”][domain-services] responsible for
domain data, and [2 services][temporal-services] which interact with Temporal.

One of those services is the “starter” which clients pass data to. The “Starter”
[accepts all input data, starts a workflow, and immediately returns][starter]
some key data, such as the workflow ID and various primary keys of the created
data. This is actually a good user experience, because clients hear back from
the server as soon as possible, while the backend can chug along with complex
asynchronous work.

Once the workflow has started and been received by the Temporal Cluster, the 
message will be dispatched to our [“Worker”][worker], which has
[registered][register] a [workflow definition][workflow-definition], as well as
a [controller][controller] for making the network requests for of our 
[3 activities][execute-activity].

[worker]: https://github.com/kevinmichaelchen/temporal-saga-grpc/blob/v0.1.0/cmd/saga/worker/app/worker/worker.go#L109-L152
[register]: https://github.com/kevinmichaelchen/temporal-saga-grpc/blob/v0.1.0/cmd/saga/worker/app/worker/worker.go#L126-L132
[workflow-definition]: https://github.com/kevinmichaelchen/temporal-saga-grpc/blob/v0.1.0/pkg/saga/workflow.go
[controller]: https://github.com/kevinmichaelchen/temporal-saga-grpc/blob/v0.1.0/pkg/saga/activity.go
[execute-activity]: https://github.com/kevinmichaelchen/temporal-saga-grpc/blob/v0.1.0/pkg/saga/workflow.go#L39-L55
[starter]: https://github.com/kevinmichaelchen/temporal-saga-grpc/blob/v0.1.0/cmd/saga/start/service/service.go
[domain-services]: https://github.com/kevinmichaelchen/temporal-saga-grpc/tree/v0.1.0/cmd/svc
[temporal-services]: https://github.com/kevinmichaelchen/temporal-saga-grpc/tree/v0.1.0/cmd/saga