---
title: Why Temporal?
description: Why you might want to adopt Temporal
---

A single request from a web or mobile client needs to be ultimately result in
the creation of several entities: an org, a user, a license, etc.

Each of these domain models has a corresponding _business logic layer_. (For
some shops, it might be the databases themselves that are completely separate).
There isn't an easy way to atomically coordinate all of these operations.

There are several patterns that exist to address these kinds of distributed
coordination problems, such as
[sagas](https://microservices.io/patterns/data/saga.html) or two-phase commits.

At a high level, we know we want these operations to execute in sequence, pretty
quickly, and with rollbacks after a little while if they can't all succeed
together.

## Context cancellation

One way we could try to accomplish this is with plain old _context cancellation_
in Golang, using the [`errgroup`](https://pkg.go.dev/golang.org/x/sync/errgroup)
package. This gets the job done most of the time, but it's not as durable as
we'd like. For example, what if the container/pod/node spontaneously crashes in
the middle of that sequence of events? The result would be out-of-sync,
inconsistent state — something we definitely want to avoid as much as possible.

## Monolith

Another option is to move toward a monolith. In this context, that would mean
merging all of those disparate _business logic layers_ into one layer. In sum:
one business logic layer, one database, completely ACID (atomic, consistent,
isolated, and durable).

## Choreography

Lots of shops adopt message brokers to durably pass data around. One service
publishes an event to a queue or topic. That event gets consumed by another
service, which does something. If an error occurs, _compensating transactions_
(a fancy term for rollbacks) are triggered the same way: via queues.

Usually queues and databases are coupled together with the
[Transactional Outbox](https://microservices.io/patterns/data/transactional-outbox.html)
pattern.

This all works fine. But it's a lot of work. The high-level flow is hard to
grok. Logic is scattered when you do choreography. It's not clear what messages
are passed around, which queues/topics exist, etc. You can consult a diagram,
but that can easily go stale. Overall, it's complex (lots of parts) and
convoluted (hard to follow).

## Orchestration

With orchestration, there's a central component that's conducting and
controlling the entire operation. This can be an improvement in terms of
readability and faster developer onboarding. It's less convoluted. But the
complexity doesn't go away.

You still have to manage that ad-hoc database for the Transactional Outbox. You
still have the message brokers to manage and maintain. You've got boilerplate
logic for retries, timeouts, and heartbeats.

> State preservation means developers write far less of the tedious, boilerplate
> code that was initially introduced by microservices. It also means that the
> ad-hoc infrastructure — like standalone queues, caches, and databases — is no
> longer needed. This reduces ops burden and the overhead of adding new
> features. It also makes onboarding new hires much easier because they no
> longer need to ramp up on the messy and domain-specific state management code.
>
> Say that your application relies on a bank’s API that is often unavailable.
> With a traditional application, you would need to wrap any code that calls the
> bank API with numerous try/catch statements, retry logic, and timeouts. \[With
> Temporal\], all of these things are provided out of the box.

## Durable Workflows

Just like an event-driven system built on queues, a system built on Temporal is
able to resume workflow execution after a spontaneous crash thanks to its
underlying stateful execution model.

Temporal Workflow Executions are _Reentrant Processes_ that are resumable,
recoverable, and reactive.

- Resumable: Ability of a process to continue execution after execution was
  suspended on an awaitable.
- Recoverable: Ability of a process to continue execution after execution was
  suspended on a failure.
- Reactive: Ability of a process to react to external events.

Temporal offers:

- Durability: functions execute effectively once and to completion, regardless
  of whether your code executes for seconds or years.
- Reliability: as shown above, functions will fully recover after a failure or
  outage.
- Scalable: the platform is capable of supporting millions to billions of
  Workflow Executions executing concurrently.

Each Temporal Workflow Execution has exclusive access to its local state,
executes concurrently to all other Workflow Executions, and communicates with
other Workflow Executions and the environment via message passing.

A Temporal Application can consist of millions to billions of Workflow
Executions. Workflow Executions are lightweight components. A Workflow Execution
consumes few compute resources; in fact, if a Workflow Execution is suspended,
such as when it is in a waiting state, the Workflow Execution consumes no
compute resources at all.

## Reservations

### Is it mature enough?

It's used by some [big names](https://temporal.io/use-cases):

- [Snapchat](https://eng.snap.com/build_a_reliable_system_in_a_microservices_world_at_snap/)
- [Netflix](https://www.youtube.com/watch?v=LliBP7YMGyA)
- [Doordash](https://doordash.engineering/2020/08/14/workflows-cadence-event-driven-processing/)
- [Stripe](https://www.youtube.com/watch?v=Crkcr1S-NSc)
- [Coinbase](https://temporal.io/case-studies/reliable-crypto-transactions-at-coinbase)
- [Datadog](https://www.youtube.com/watch?v=Hz7ZZzafBoE)
- [Hashicorp](https://www.youtube.com/watch?v=kDlrM6sgk2k&t=1188s)
- [Airbyte](https://temporal.io/case-studies/airbyte-case-study)
- [Box](https://temporal.io/case-studies/temporal-a-central-brain-for-box)
- [Checkr](https://temporal.io/case-studies/how-temporal-simplified-checkr-workflows/)
- [Descript](https://temporal.io/case-studies/descript-case-study)
- [Zebra Medical Vision](https://temporal.io/case-studies/zebra-medical-case-study/)

### Does it have a Cloud offering?

- Supports US East region
- SOC2 compliant
- Encryption means they don't see any of our data
- VPC peering
- Scalable (150,000 actions per second)
- Public uptime reporting
- Good SLA (99.9% uptime)
- Latency SLO of 200ms per region for p99
- Pricing
  - simple, transparent, and controllable
  - automatically adjusts based on volume within a namespace
  - pay only for what you use, not what you might need.
  - spend less on infrastructure and staff to run your applications
