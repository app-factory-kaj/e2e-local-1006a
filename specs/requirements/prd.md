# greeter — PRD

## Problem Statement

Teams building and wiring services on this platform need a minimal, predictable HTTP service to exercise end-to-end flows (deployment, routing, contract checks) without the overhead of a real business domain. Today there is no small, reference-shaped service to stand in for that purpose.

## Solution

Greeter is a small Go HTTP service exposing a single endpoint, `GET /hello`, that returns a JSON greeting for a given name. It exists to be simple, predictable, and easy to integrate against — a clean reference implementation.

## Actors

- **API Consumer** — any client application or service that calls the greeter HTTP endpoint to obtain a greeting.

## User Stories

1. As an API consumer, I want to call `GET /hello?name=X`, so that I receive a JSON greeting personalized to that name.
2. As an API consumer, I want to call `GET /hello` without a `name` parameter and still get a successful response, so that the endpoint behaves predictably even when the caller omits the parameter.

## Product Decisions

- When `name` is omitted or empty, the service returns a default greeting using "World" as the name (e.g. "Hello, World!") rather than an error. *assumed*
- The JSON response shape is a single `message` field carrying the full greeting sentence, e.g. `{"message": "Hello, X!"}`. *assumed*
- The service requires no sign-in or authentication — it is a stateless backend endpoint with no end-user identity involved. *assumed*
- The service depends on no external third-party services.
- The service persists no data; each request is handled statelessly.

## Out of Scope

- User interface of any kind — this is an API-only service.
- Persistence, history, or storage of past requests.
- Authentication, authorization, or per-caller rate limiting.
- Any endpoint beyond `GET /hello`.

## Open Questions

None at this time.

E2E marker e2e-local-1006a.