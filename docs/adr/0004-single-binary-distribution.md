# 4. Single self-contained binary

- Status: Accepted
- Date: 2026-07-14
- Spec: [ENGG-5](https://linear.app/docudex/issue/ENGG-5)

## Context

Docudex ships a CLI and a web UI (`serve`). The web UI is a React + React Router v7 SPA.
Distribution could bundle the SPA and any native dependencies separately, or embed
everything into one artifact. A frictionless install ("one download or one `go install`")
and reliable offline operation both push toward a single self-contained binary.

## Decision

Distribute docudex as a **single self-contained binary with no runtime dependencies**.
The React SPA is built and embedded into the binary via `go:embed`, so `docudex serve`
needs no external assets. All native-feeling dependencies must be **pure Go** — notably
the SQLite FTS5 driver (see ADR-0003) is CGO-free — so cross-compilation and
`go install` stay trivial.

## Consequences

- Installation is one download (or `go install`); nothing to configure at runtime.
- Everything already fetched stays fully browsable and searchable offline.
- The SPA build is a prerequisite of the Go build for release; CI/release tooling
  (GoReleaser, Homebrew tap — ENGG-7) must build the SPA before `go build`.
- Ruling out CGO constrains library choices (e.g. the SQLite driver) for the life of
  the project.
- The `serve` JSON API is treated as private/unstable in the MVP; embedding the SPA
  avoids freezing an API contract before real use shapes it.
