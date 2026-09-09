# Agent Instructions

This repository is intentionally a **Go learning project** as well as a Pior Labs service.

## Read first

Before making non-trivial implementation changes, read:

1. `README.md`
2. `LEARNING.md`
3. `ROADMAP.md`
4. relevant ADRs under `docs/DECISIONS/`

## Learning-first rule

Do not default to implementing an entire milestone for the developer.

For milestones that introduce a new Go concept, prefer this order:

1. explain the concept and relevant standard-library APIs;
2. provide hints or a small isolated example;
3. let the developer write the first implementation;
4. review the implementation for correctness and idiomatic Go;
5. explain errors and tradeoffs;
6. only provide a complete replacement implementation when explicitly requested or when the developer has already attempted the task and wants one.

The purpose is to preserve hands-on learning rather than maximize implementation speed.

## Implementation principles

- Prefer the Go standard library while it is sufficient.
- Keep the service small; add abstractions only when the code earns them.
- Do not introduce a web framework merely to avoid learning `net/http`.
- Do not introduce concurrency before the sequential implementation of the same behavior works.
- When concurrency is introduced, explain goroutine lifecycle, synchronization, cancellation, and race-safety.
- Use `context.Context` for request cancellation and lifecycle propagation where appropriate.
- Use explicit HTTP client timeouts for outbound checks.
- Add tests alongside behavior; use `httptest` instead of production endpoints in automated tests.
- Run `go test ./...` and, once concurrency exists, `go test -race ./...` before considering a milestone complete.

## Pior Labs boundaries

`service-health` should fit the existing Pior Labs Docker/Caddy deployment model. Do not introduce Kubernetes, another edge proxy, or unrelated infrastructure for this service.

The service owns health aggregation and interpretation. Individual applications remain responsible for exposing truthful health/readiness endpoints. `app-dashboard` owns presentation.

## Documentation

When a milestone is completed:

- update its status in `ROADMAP.md`;
- append a concise reflection to `LEARNING.md` if useful;
- create an ADR only for durable architectural decisions, not ordinary implementation details.

If implementation changes the agreed service direction, update the documentation in the same PR.
