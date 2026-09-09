# Pior Labs Health Service

Pior Labs Health Service is a lightweight, platform-native health aggregation service for Pior Labs applications and infrastructure.

It is being built in **Go** for two reasons:

1. health aggregation is a natural fit for a small, concurrent, long-running Go service; and
2. the repository is intentionally being used to learn Go through hands-on implementation rather than one-shot AI generation.

## Purpose

The service will provide one place to understand the health of Pior Labs applications and their exposed readiness components.

Instead of only answering whether an application URL responds, it should eventually be able to represent richer state such as:

- application/API availability;
- database readiness;
- storage readiness;
- dependency readiness;
- response latency.

The intended first consumer is [`pior-labs/app-dashboard`](https://github.com/pior-labs/app-dashboard), which should be able to request one stable platform-health API rather than independently polling every application.

## Relationship to Uptime Kuma

This project is **not** initially intended to recreate Uptime Kuma feature-for-feature.

Uptime Kuma can continue to provide generic uptime monitoring and alerting. `service-health` focuses on Pior Labs-specific health aggregation and application readiness contracts. Their responsibilities can be revisited later if real usage makes the overlap unnecessary.

## Learning-first development

The implementation is deliberately incremental.

The core rule is:

> Write the first implementation of each important Go concept yourself. Use AI as a teacher, reviewer, and debugger before using it as an implementer.

See [`LEARNING.md`](./LEARNING.md) for the milestone-by-milestone learning plan and [`ROADMAP.md`](./ROADMAP.md) for the product/service roadmap.

The initial progression is:

1. minimal Go HTTP server;
2. one remote health check;
3. explicit service/check types;
4. multiple checks sequentially;
5. refactor to concurrent checks;
6. scheduled checks and in-memory state;
7. context cancellation and graceful shutdown;
8. external configuration;
9. aggregate health API;
10. testing with Go's standard tools;
11. Docker and Pior Labs deployment;
12. Dashboard integration.

Concurrency is intentionally introduced only after the sequential implementation works. The goal is to understand why Go's concurrency model improves this service rather than beginning with code that hides the comparison.

## Intended technology direction

Prefer Go's standard library while learning and while it remains sufficient:

- `net/http`
- `encoding/json`
- `context`
- `time`
- `sync`
- `os/signal`
- `testing`
- `net/http/httptest`

Third-party dependencies should earn their place through an actual requirement rather than being introduced to reproduce patterns from larger frameworks.

## Repository structure

The implementation structure should emerge as code is written. A likely small shape is:

```text
cmd/
  server/
internal/
  checker/
  model/
  config/
docs/
  DECISIONS/
```

Do not create empty abstraction layers solely to match this sketch.

## Current status

Planning and learning structure established. Application implementation has not started yet.
