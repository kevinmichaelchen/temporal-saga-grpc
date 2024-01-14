---
title: Architectural Overview
description: A guide to the overall architecture here.
---

It started from my exploration into _sagas_, a pattern for orchestrating distributed transactions that span across multiple services and databases. After discovering [Temporal][temporal], I wanted a go-to sandbox where I could learn about _durable workflows_. Instead of the typical car-hotel-flight example you find in microservices literature, I chose domains more applicable to SaaS products: an Org, a User, and a License.

:::note[Evolution of this project]

Over time, this project grew beyond just Temporal, into an amalgamation of valuable tools, libraries, frameworks, and patterns.

:::

## What's in the box?

Enough chatter. What does the tech stack look like?

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

- [**Bun**][bun] — an all-in-one JS runtime and toolkit
- [**NextJS**][nextjs] as the React framework
- [**`shadcn/ui`**][shadcn-ui] for beautiful [Radix][radix]/[Tailwind][tailwind] components.

### 🤷 Misc

- [**Docker**][docker]
- [**`uber-go/fx`**][fx] for [dependency injection][di]
- [**Golang**][golang]

[atlas]: https://atlasgo.io/
[buf]: https://buf.build/
[bun]: https://bun.sh/
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
