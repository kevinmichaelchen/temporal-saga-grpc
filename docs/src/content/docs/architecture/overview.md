---
title: Architectural Overview
description: A guide to the overall architecture here.
---

It started from my exploration into _sagas_, a pattern for orchestrating distributed transactions that span across multiple services and databases. After discovering [**Temporal**][temporal], I wanted a go-to sandbox where I could learn about _durable workflows_. Instead of the typical car-hotel-flight example you find in microservices literature, I chose domains more applicable to SaaS products: an Org, a User, and a License.

[temporal]: /temporal-saga-grpc/tech-stack/temporal

:::note[Evolution of this project]

Over time, this project grew beyond just Temporal, into an amalgamation of valuable tools, libraries, frameworks, and patterns. Read more about the full tech-stack [**here**][tech-stack].

[tech-stack]: /temporal-saga-grpc/tech-stack/overview

:::

![/temporal-saga-grpc/diagrams/architecture.svg](/temporal-saga-grpc/diagrams/architecture.svg)
