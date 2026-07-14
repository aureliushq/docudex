# 1. Manifest (`docudex.toml`) is the source of truth

- Status: Accepted
- Date: 2026-07-14
- Spec: [ENGG-5](https://linear.app/docudex/issue/ENGG-5)

## Context

A project's doc scope could be derived live from `go.mod`, or held in a committed
manifest. Deriving from `go.mod` keeps a single source but couples doc scope to build
dependencies (no way to add an extra version to compare against, no offline-stable
record) and makes the scope invisible to teammates until they run a build. The design
grilling (2026-07-11) contested this and resolved in favour of an explicit, committed
manifest, with `go.mod` as an advisory input.

## Decision

The committed `docudex.toml` at the project root is the **source of truth** for a
project's doc scope. It is discovered by an upward directory walk (git-style), so
commands work from any subdirectory. `sync` fetches exactly what the manifest lists —
no more, no less. A dependency entry may pin multiple versions; the first/marked one is
the default for browse and search.

`go.mod` is read **only by manifest-writing commands** — `init` (to create the manifest)
and `sync --from-gomod` (to update default versions, preserving extra pinned versions).

Manifest-reading commands (`serve`, `search`, `sync`, `list`) cheaply compare against
`go.mod` and print a **drift warning** naming the drifted module and the fix command.
Drift is **never auto-healed** — the committed manifest changes only when the user
explicitly asks (`sync --from-gomod`).

## Consequences

- Teammates get identical doc scope from a fresh clone with one `docudex sync`.
- Users can add versions beyond `go.mod` (e.g. compare current vs. an upgrade target).
- Reading stale docs is guarded by an explicit, non-silent warning rather than silent
  correction — the committed file never changes behind the user's back.
- Two code paths touch versions (read vs. write); drift detection must stay cheap enough
  to run on every read command.
