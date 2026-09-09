# Architecture Decisions

Use this directory for durable decisions that are specific to `service-health` and worth preserving beyond an individual pull request.

Do not create an ADR for every implementation detail. Platform conventions that are already established elsewhere do not need to be restated unless this service intentionally deviates from them.

## Suggested format

```text
0001-short-decision-title.md
```

Each ADR should answer:

- **Context** — what problem or constraint forced a decision?
- **Decision** — what was chosen?
- **Alternatives considered** — what reasonable options were rejected?
- **Consequences** — what does this make easier or harder?
- **Status** — proposed, accepted, superseded, or deprecated.

## Likely future decisions

Create ADRs only when these become concrete decisions, for example:

- configuration format and discovery model;
- concurrency strategy if it becomes more than a trivial implementation detail;
- in-memory versus persistent health history;
- how Pior Labs readiness payloads are interpreted;
- authentication strategy for internal checks;
- alerting ownership versus Uptime Kuma;
- production deployment/runtime choices that differ from standard Pior Labs conventions.

The learning plan lives in [`../../LEARNING.md`](../../LEARNING.md). Learning notes should not be turned into ADRs unless they also represent a durable architectural choice.
