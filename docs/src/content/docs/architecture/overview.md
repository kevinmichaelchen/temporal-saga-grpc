---
title: Architectural Overview
description: A guide to the overall architecture here.
---

:::caution[Be warned]

This project is by no means a production-ready application. Nor does it attempt to model anything of real-world substance.

:::

It is on par to many “demo applications” you'll find in microservices literature that use the common “book a car, book a hotel, book a flight” scenario.

The only difference being that here I'm using a different scenario, one that SaaS products might be more familiar with: “create an Org, create a User, create a License”.

The question many ask is: “How do you perform these steps atomically across microservices?”

In the literature, you'll see solutions for this scenarios commonly referred to as “distributed transactions” or “sagas”. That's where my interest here started. I wanted to explore the most cutting-edge way of implementing such patterns, that avoided the complexity and effort of solutions of old, such as Two-Phase Commits and [Transactional Outboxes][outbox].

[outbox]: https://microservices.io/patterns/data/transactional-outbox.html

Then I stumbled upon [Temporal][temporal] — an ideal solution for implementing not only “distributed transactions” that span across multiple microservices, but also most asynchronous, relatively long-lived processes, especially ones that involve human interaction.

:::note

Over time, I kept adding more and more cool tools and libraries. An apt name for the project is no longer `temporal-saga-grpc`, but rather `kevins-ideal-tech-stack`.

:::

## Running things

- [pkgx][pkgx] is without a doubt the best package and command runner out there. and [Taskfile][taskfile] for optimal local DevEx (seamless, portable, invisible)

## APIs

- [Tailcall][tailcall] and [GraphQL][graphql]
- [Buf][buf] for [Protocol Buffer][protobuf] tooling
- [ConnectRPC][connect] for backing APIs
- [Vanguard][vanguard] to transcode REST to ConnectRPC

## Data

- [Atlas][atlas] for declarative database migrations
- [Postgres][postgres]
- [SQLBoiler][sqlboiler] as a Go ORM

## Observability

- [OpenTelemetry][otel] (OTel)
- [Jaeger][jaeger] (soon to be replaced by [HyperDX][hyperdx])

## Frontend

- [NextJS][nextjs] React framework
- [shadcn/ui][shadcn-ui] for beautiful [Radix][radix]/[Tailwind][tailwind] components

## Misc

- [Docker][docker]
- [FX][fx] for [dependency injection][di]
- [Golang][golang]

[atlas]: https://atlasgo.io/
[buf]: https://buf.build/
[connect]: https://connectrpc.com/
[di]: https://en.wikipedia.org/wiki/Dependency_injection
[docker]: https://www.docker.com/
[fx]: https://github.com/uber-go/fx
[golang]: https://go.dev/
[graphql]: https://graphql.org/
[hyperdx]: https://www.hyperdx.io/
[jaeger]: https://www.jaegertracing.io/
[nextjs]: https://nextjs.org/
[otel]: https://opentelemetry.io/
[postgres]: https://www.postgresql.org/
[pkgx]: https://pkgx.sh/
[protobuf]: https://protobuf.dev/
[radix]: https://www.radix-ui.com/
[shadcn-ui]: https://ui.shadcn.com/
[sqlboiler]: https://github.com/volatiletech/sqlboiler
[tailcall]: https://tailcall.run/
[tailwind]: https://tailwindcss.com/
[taskfile]: https://taskfile.dev/
[temporal]: https://temporal.io/
[vanguard]: https://github.com/connectrpc/vanguard-go
