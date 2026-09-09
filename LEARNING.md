# Go Learning Plan

This repository is intentionally both a useful Pior Labs service and a structured Go learning project.

The goal is not to have an AI agent generate the service end to end. The goal is to learn Go by implementing the important parts directly, using AI for explanation, review, debugging, and targeted examples.

## Learning rule

For each milestone:

1. Read the milestone and the Go concepts involved.
2. Learn the concepts before asking for a complete implementation.
3. Write the first working attempt yourself.
4. Run and test it.
5. Ask AI to review the code, explain mistakes, and suggest idiomatic improvements.
6. Make the revision yourself where practical.
7. Record what you learned before moving on.

Unless explicitly requested, AI should prefer hints, explanations, code review, debugging help, and small examples over complete milestone implementations.

## What AI should help with

AI is useful for:

- explaining unfamiliar Go syntax and standard-library APIs;
- comparing approaches and tradeoffs;
- reviewing code for correctness and idiomatic Go;
- explaining compiler and runtime errors;
- suggesting tests and edge cases;
- identifying race conditions, resource leaks, or missing cancellation;
- reviewing package boundaries after a working version exists;
- generating small isolated examples when a concept is unclear.

## What AI should not do by default

Avoid delegating these wholesale while the corresponding concept is being learned:

- generating the entire service scaffold;
- implementing all health-check logic in one prompt;
- introducing frameworks before the standard library has been explored;
- adding abstractions simply because they are common in larger codebases;
- rewriting a working milestone before its behavior is understood;
- hiding concurrency behind a library before goroutines, channels, synchronization, and context are understood.

## Milestone 1 — Minimal Go HTTP service

### Build

Create a Go module and a minimal HTTP server exposing:

- `GET /health`

The endpoint can initially return a small JSON response indicating that `service-health` itself is running.

### Learn

- `go mod init`
- packages
- imports
- `package main`
- functions
- variables
- `net/http`
- handlers
- `encoding/json`
- basic error handling

### Implement yourself

- module initialization;
- `main.go`;
- HTTP server startup;
- `/health` handler.

### Done when

- the service starts locally;
- `curl` can reach `/health`;
- the response is valid JSON;
- you can explain every line in the implementation.

### Reflection

Record notes here after completing the milestone:

- What was new?
- What differed from JavaScript/TypeScript?
- What was confusing?
- What would you write differently now?

---

## Milestone 2 — Check one remote service

### Build

Add a checker that requests one known Pior Labs health/readiness endpoint and returns a result containing at minimum:

- service name;
- status;
- HTTP status code where applicable;
- latency;
- timestamp;
- error information when the request fails.

Start with one hard-coded service. Configuration comes later.

### Learn

- structs
- struct literals
- methods
- `http.Client`
- `time.Duration`
- timeouts
- `defer`
- error values
- multiple return values

### Important lesson

Do not use `http.Get` as the final production pattern. Learn why a reusable `http.Client` with an explicit timeout is preferable for a long-running service.

### Done when

- a healthy endpoint produces a healthy result;
- an unreachable endpoint produces a useful failure result;
- a slow endpoint cannot block forever;
- latency is measured;
- you can explain where resources such as response bodies are released.

---

## Milestone 3 — Domain model

### Build

Introduce explicit types for concepts such as:

- service definition;
- check result;
- health status.

Keep the model small and driven by actual requirements.

### Learn

- named types
- constants
- methods
- zero values
- slices
- JSON struct tags
- pointers versus values

### Done when

- check results serialize into a clean API shape;
- status values cannot accidentally become arbitrary strings without a conscious choice;
- you can explain why each type is a value or pointer.

---

## Milestone 4 — Multiple services, sequentially

### Build

Check several configured service definitions one after another and return all results.

Do this sequentially first, even though concurrency will eventually be better.

### Learn

- slices
- `for ... range`
- append
- composition of functions
- separating orchestration from individual health checks

### Why sequential first

The sequential version creates a correctness baseline. The later concurrent implementation should produce the same logical results faster, rather than combining concurrency learning with basic orchestration bugs.

### Done when

- several services can be checked;
- the result ordering is understood and intentional;
- one failed service does not prevent the others from being checked.

---

## Milestone 5 — Concurrent health checks

### Build

Refactor the multi-service checker so independent checks can run concurrently.

Do not add concurrency until the sequential version works and is tested.

### Learn

- goroutines
- `sync.WaitGroup`
- channels
- mutexes where appropriate
- race conditions
- bounded versus unbounded concurrency

### Learning exercise

Implement or experiment with at least two synchronization approaches in a small isolated example before choosing the production approach.

Understand the tradeoff between:

- one goroutine per service;
- a bounded worker pool.

For the expected small Pior Labs service count, prefer the simplest correct design unless measurements justify more complexity.

### Done when

- independent checks overlap in time;
- results remain complete and deterministic enough for the API contract;
- `go test -race ./...` reports no race in covered code;
- you can explain how goroutines terminate.

---

## Milestone 6 — Scheduled checks and in-memory state

### Build

Run checks on a configurable interval and retain the latest result for each monitored service in memory.

The service should be able to answer API requests without performing every downstream check synchronously.

### Learn

- `time.Ticker`
- long-running goroutines
- shared state
- synchronization
- lifecycle management

### Done when

- checks run periodically;
- latest results can be read safely while checks are updating;
- stopping the scheduler releases its ticker and goroutines.

---

## Milestone 7 — Context and graceful shutdown

### Build

Add cancellation and graceful process shutdown.

The service should stop accepting new work, cancel outstanding checks where practical, stop scheduled work, and exit cleanly when the process receives a termination signal.

### Learn

- `context.Context`
- cancellation propagation
- deadlines
- `os/signal`
- graceful `http.Server` shutdown

### Done when

- Ctrl+C causes an orderly shutdown;
- outbound requests honor cancellation;
- scheduled goroutines stop;
- there are no obvious leaked background loops.

---

## Milestone 8 — Configuration

### Build

Replace hard-coded monitor definitions with configuration.

Start with the smallest format that satisfies actual deployment requirements. Environment variables or a small config file are both acceptable; choose deliberately.

A monitor definition will likely need concepts such as:

- stable ID/name;
- URL;
- check type;
- timeout;
- optional expected status;
- optional readiness interpretation.

### Learn

- configuration parsing
- validation
- defaults
- startup errors

### Done when

- invalid configuration fails clearly at startup;
- defaults are explicit;
- secrets are not required for the initial public-health checks unless a real endpoint requires them.

---

## Milestone 9 — Aggregate health API

### Build

Expose a stable JSON API intended for `app-dashboard`.

Potential endpoints include:

- `GET /health` — health of `service-health` itself;
- `GET /api/services` — latest result for each monitored service;
- `GET /api/health` — aggregate platform state.

Do not finalize endpoint names until the implemented model is clear.

### Learn

- API design
- handler composition
- status codes
- JSON response contracts
- read-only shared state access

### Product direction

The service should understand Pior Labs readiness responses rather than only answering whether a URL returned HTTP 200.

Examples may eventually include application components such as:

- web;
- API;
- database;
- image/file storage;
- authentication dependencies.

### Done when

- the dashboard can consume one stable endpoint without knowing each application's individual health URL;
- unhealthy downstream services do not make the health-service API itself unusable.

---

## Milestone 10 — Testing

### Build

Add tests as the implementation matures rather than leaving all testing until the end.

By this milestone, explicitly cover the service using Go's standard testing tools.

### Learn

- `testing`
- table-driven tests
- `httptest.Server`
- dependency injection through simple interfaces/functions
- race testing

### Important scenarios

Test at least:

- healthy response;
- non-2xx response;
- timeout;
- connection failure;
- malformed readiness payload if parsing application-specific readiness;
- several concurrent checks;
- cancellation;
- aggregate degraded/unhealthy rules.

### Done when

- `go test ./...` passes;
- `go test -race ./...` passes for the covered concurrent code;
- downstream services can be simulated without calling real production endpoints.

---

## Milestone 11 — Containerization and Pior Labs deployment

### Build

Only after the service works locally, add the production packaging needed to run on Pior Labs.

Potential work:

- multi-stage Dockerfile;
- minimal runtime image or appropriate static-binary strategy;
- container health check;
- private container networking;
- `platform-deploy` integration;
- GitHub Actions validation/deployment.

### Learn

- `go build`
- cross-compilation basics
- static binaries and CGO implications
- build versus runtime image boundaries
- signal behavior as PID 1 in containers

### Constraint

Do not introduce Kubernetes, service meshes, or other orchestration solely for this service. It should fit the existing Pior Labs Docker/Caddy platform.

---

## Milestone 12 — Dashboard integration

### Build

Update `app-dashboard` to consume the health-service API and display platform/application health.

The health service owns health aggregation. The dashboard owns presentation.

### Learn

This milestone is less about Go syntax and more about API boundary design:

- consumer-driven API design;
- backward compatibility;
- avoiding presentation-specific logic in the service.

### Done when

- Dashboard can render health without independently polling every application;
- adding a new monitored Pior Labs application normally requires changing health-service configuration rather than dashboard application code.

---

# Future learning topics

These should be introduced only when a real requirement appears:

- persistence/history;
- metrics and Prometheus exposition;
- alerting;
- TLS certificate expiry checks;
- TCP checks;
- authenticated checks;
- retry/backoff policies;
- circuit breakers;
- worker pools;
- generic interfaces and type parameters;
- profiling with `pprof`;
- OpenTelemetry.

Do not add these simply to make the service resemble a commercial monitoring system.

# Learning journal

After meaningful milestones, append a short dated entry containing:

- what was implemented;
- what was learned;
- what needed AI help;
- a mistake or misconception corrected;
- one concept to revisit.

The journal should be concise. The purpose is to preserve evidence of the learning process and make future review easier, not to create daily status bureaucracy.
