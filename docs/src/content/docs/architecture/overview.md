---
title: Architectural Overview
description: A guide to the overall architecture here.
---

Given the name of the project, we have to mention [Temporal][temporal] as an ideal solution for implementing not only “distributed transactions” that span across multiple microservices, but also most asynchronous, relatively long-lived processes, especially ones that involve human interaction.

## Running things

- [pkgx][pkgx] and [Taskfile][taskfile] for optimal local DevEx (seamless, portable, invisible)

## APIs

- [Tailcall][tailcall] and [GraphQL][graphql]
- [Buf][buf] for [Protocol Buffer][protobuf] tooling
- [ConnectRPC][connect] for backing APIs

## Data

- [Atlas][atlas] for declarative database migrations
- [Postgres][postgres]

## Observability

- [OpenTelemetry][otel] (OTel)
- [Jaeger][jaeger] (soon to be replaced by [HyperDX][hyperdx])

## Frontend

- [NextJS][nextjs] React framework
- [shadcn/ui][shadcn-ui] for beautiful [Radix][radix]/[Tailwind][tailwind] components

## Misc

- [Docker][docker]
- [Golang][golang]

[atlas]: https://atlasgo.io/
[buf]: https://buf.build/
[connect]: https://connectrpc.com/
[docker]: https://www.docker.com/
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
[tailcall]: https://tailcall.run/
[tailwind]: https://tailwindcss.com/
[taskfile]: https://taskfile.dev/
[temporal]: https://temporal.io/
