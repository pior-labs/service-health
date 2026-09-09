# Service Health Roadmap

`service-health` is a Pior Labs-native health aggregation service. It is intentionally narrower than a generic monitoring platform such as Uptime Kuma.

The service should understand Pior Labs application health/readiness contracts, aggregate the latest state, and expose one stable API for consumers such as `app-dashboard`.

## Product goals

- Aggregate application and platform health in one service.
- Understand richer readiness state rather than only HTTP reachability.
- Provide stable read APIs for the Pior Labs dashboard.
- Keep the service lightweight and aligned with the existing Docker/Caddy platform.
- Use the project as a deliberate Go learning exercise.

## Non-goals for the initial version

- Rebuilding Uptime Kuma feature-for-feature.
- Public multi-tenant monitoring.
- Full incident management.
- Advanced alert routing.
- Long-term metrics storage.
- Kubernetes-specific service discovery.
- A standalone monitoring UI.

## Delivery milestones

| Milestone | Capability | Status |
| --- | --- | --- |
| 1 | Minimal Go server and self `/health` | Planned |
| 2 | Check one Pior Labs endpoint | Planned |
| 3 | Explicit service/check domain model | Planned |
| 4 | Sequential multi-service checking | Planned |
| 5 | Concurrent health checking | Planned |
| 6 | Scheduled checks and in-memory latest state | Planned |
| 7 | Context cancellation and graceful shutdown | Planned |
| 8 | External monitor configuration | Planned |
| 9 | Aggregate health API | Planned |
| 10 | Comprehensive standard-library tests | Planned |
| 11 | Docker and Pior Labs production integration | Planned |
| 12 | `app-dashboard` integration | Planned |

The learning details and completion criteria for each milestone live in [`LEARNING.md`](./LEARNING.md).

## Initial API direction

The exact contract should be designed from the implementation rather than fixed prematurely, but the service will likely expose concepts similar to:

- `GET /health` — whether `service-health` itself is functioning;
- `GET /api/services` — latest state of all monitored services;
- `GET /api/health` — aggregate Pior Labs platform state.

A downstream failure should be represented in the response data. It should not ordinarily make the health aggregation API itself return unusable output.

## Health states

The initial service should keep health states small and explicit. A likely starting point is:

- `healthy`
- `degraded`
- `unavailable`
- `unknown`

The exact rules for aggregate state should be documented once implemented.

## Monitoring direction

The first monitors should focus on existing Pior Labs application endpoints.

Rather than only checking that `https://cookbook.szarans.ca` returns a response, the service should be capable of representing richer state when an application exposes it, for example:

- application/API availability;
- database readiness;
- file/image storage readiness;
- authentication dependency readiness;
- response latency.

Each application remains responsible for defining truthful health/readiness endpoints. `service-health` aggregates and interprets them; it should not duplicate application-specific business checks unnecessarily.

## Relationship to Uptime Kuma

Uptime Kuma may continue to handle generic uptime monitoring and alerting.

`service-health` is intended to provide Pior Labs-specific understanding and a platform API. If overlap becomes unnecessary later, responsibilities can be revisited based on real usage rather than making replacement of Kuma an initial requirement.

## Future possibilities

Only pursue these after the core service is useful:

- health history/persistence;
- deployment/version metadata;
- Prometheus metrics;
- SSL expiry monitoring;
- notification/alert integrations;
- retry and backoff policies;
- authenticated/internal checks;
- richer infrastructure checks;
- MCP access if conversational platform-health queries become useful.
