---
title: Why Atlas?
description: Why you might want to adopt Atlas
---

Why use [Atlas][atlas]?

[atlas]: https://atlasgo.io/

Rephrased: Do developers need to be writing execution plans in 2024? Or is it time to embrace declarative schema migrations? There was some debate about this amongst the general public in this Hacker Noon [thread][ycomb].

[ycomb]: https://news.ycombinator.com/item?id=30020288

As the thread notes, declarative migrations are nothing new. Facebook has been using them for its MySQL fleet — one of the largest in the world — for over a decade. [Ariel][a8m], one of the engineers who built Facebook's entity framework, [`ent`][ent], is now helping to build Atlas.

[a8m]: https://github.com/a8m
[ent]: https://github.com/ent/ent

## Features

- [Versioned authoring][versioned-authoring] to support hybrid-style workflows
- [Drift detection][v018]
- [Breaking change prevention][breaking-change-prevention]
- [Linting for destructive changes][lint]
- [Kubernetes Operator Support][k8s]
- [GitHub Action support][gh-actions]
- [SQL Triggers][v017]
- [ERDs][erd]
- [Legacy migrations and Mermaid diagram support][v016]
- [Functions and procedures][v015]
- [Checkpoints][v014]
- [Views][v013]
- [Plain SQL support][plain-sql]
- [Terraform support][terraform] with [plain SQL][tf-plain-sql]

[k8s]: https://atlasgo.io/integrations/kubernetes/operator
[v018]: https://atlasgo.io/blog/2024/01/09/atlas-v0-18
[v017]: https://atlasgo.io/blog/2024/01/01/atlas-v-0-17
[erd]: https://atlasgo.io/blog/2022/12/06/erd-and-json
[v016]: https://atlasgo.io/blog/2023/12/19/atlas-v-0-16
[v015]: https://atlasgo.io/blog/2023/10/19/atlas-v-0-15
[v014]: https://atlasgo.io/blog/2023/08/31/atlas-v-0-14
[v013]: https://atlasgo.io/blog/2023/08/06/atlas-v-0-13
[plain-sql]: https://atlasgo.io/blog/2023/01/05/atlas-v090
[tf-plain-sql]: https://atlasgo.io/blog/2023/04/21/terraform-v050
[terraform]: https://atlasgo.io/blog/2022/05/04/announcing-terraform-provider
[breaking-change-prevention]: https://atlasgo.io/blog/2023/03/31/preventing-breaking-changes
[gh-actions]: https://atlasgo.io/blog/2022/08/22/atlas-ci-github-actions
[lint]: https://atlasgo.io/blog/2022/07/14/announcing-atlas-lint
[versioned-authoring]: https://atlasgo.io/blog/2022/08/11/announcing-versioned-migration-authoring
