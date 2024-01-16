---
title: Buf
description: Using Buf for developing with Protocol Buffers
---

## What is Buf?

[Buf][buf] is a tool that makes working with [Protocol Buffers][protobuf] a lot nicer. (I would add a section called “Why Protobuf?” here, but this [blog post][why-protobuf] more than suffices).

[buf]: https://buf.build/
[protobuf]: https://protobuf.dev/
[why-protobuf]: https://buf.build/blog/the-real-reason-to-use-protobuf

It provides value in two forms:

1. The [Buf CLI][cli] for code generation, breaking change detection, linting, and formatting.
1. The [Buf Schema Registry][bsr] (BSR) for storing and managing Protobuf files as versioned modules, allowing you to consume and publish APIs without friction.
   - The BSR did to Protobufs what VCS systems (like GitHub) did to source code.

[bsr]: https://buf.build/docs/bsr/introduction
[cli]: https://buf.build/docs/ecosystem/cli-overview

## Buf Schema Registry

We've published our 4 modules to the BSR:

- [licenseapis]
- [orgapis]
- [profileapis]
- [temporalapis]

[licenseapis]: https://buf.build/kevinmichaelchen/licenseapis/docs/main:license.v1beta1
[orgapis]: https://buf.build/kevinmichaelchen/orgapis/docs/main:org.v1beta1
[profileapis]: https://buf.build/kevinmichaelchen/profileapis/docs/main:profile.v1beta1
[temporalapis]: https://buf.build/kevinmichaelchen/temporalapis/docs/main:temporal.v1beta1
