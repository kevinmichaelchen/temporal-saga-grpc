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

## What's in the box?

Enough chatter. What's the tech stack look like?

### 🧰 Running things

- [**pkgx**][pkgx] is without a doubt the best package and command runner out there.
- [**Taskfile**][taskfile] is a superior version of Makefile.

Combine both for an optimal local DevEx (seamless, portable, invisible).

### 👽 APIs

- [**GraphQL**][graphql] is a pretty widely recognized standard. Great for aggregating data and lots of frontend support.
- GraphQL is hard to build in the backend, and [**Protocol Buffers**][protobuf] are the undisputed champion of schema modeling.
- If you’re using Protobufs, choosing not to use [**Buf**][buf] isn’t really a choice.
- To actually build an API from your Protobufs, [**ConnectRPC**][connect] is the way to go.
- The problem is that ConnectRPC isn’t GraphQL. To bridge that gap, use [**Vanguard**][vanguard] (for transcoding) and [**Tailcall**][tailcall] (for API composition).

### 🗃️ Data

- [**Atlas**][atlas] for declarative database migrations — because no one likes having to look at hundreds of migration files or having to guess what the schema currently looks like.
- [**Postgres**][postgres] — need I say more?
- [**SQLBoiler**][sqlboiler] as a great Go ORM.

### 🔭 Observability

- [**OpenTelemetry**][otel] (OTel) is the industry standard.
- [**Jaeger**][jaeger] lets you visualize distributed traces.
- [**HyperDX**][hyperdx] is another open-source platform I'm keen to try.

### 🎨 Frontend

- [**NextJS**][nextjs] as the React framework
- [**`shadcn/ui`**][shadcn-ui] for beautiful [Radix][radix]/[Tailwind][tailwind] components.

### 🤷 Misc

- [**Docker**][docker]
- [**`uber-go/fx`**][fx] for [dependency injection][di]
- [**Golang**][golang]

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
