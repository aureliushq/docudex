# Docudex — Context

Docudex is a **CLI-first, per-project, version-exact Go API reference aggregator**.
Running `docudex init` in a Go repo reads `go.mod` and materializes API reference
documentation for the project's direct dependencies (plus the Go standard library)
into a local store, pinned to the exact versions the project uses. Docs are then
searchable from the terminal (`docudex search`) or browsable in a local web UI
(`docudex serve`) — scoped to exactly what the project uses, and fully available
offline.

Spec of record: Linear **ENGG-5**. MVP scope is **Go-only** and **API-reference-only**;
guides/prose, other ecosystems (Rust is the planned second adapter), a TUI reader,
agent/MCP features, and hosted/team features are all out of scope for the MVP.

## Glossary (ubiquitous language)

Use these terms exactly; don't drift to synonyms.

- **Docset** — the documentation for one `module@version`. The unit of everything:
  fetched, stored, indexed, listed, searched, scoped.
- **Manifest** — the committed `docudex.toml` at the project root. The **source of
  truth** for a project's doc scope. Lists docsets and their pinned version(s); the
  first/marked version of an entry is the default for browse/search.
- **Store** — the global content store under **home** at `store/<module>@<version>/`.
  Markdown is canonical (plus metadata sidecars), content-addressed by `module@version`
  so projects share one copy and multiple versions coexist. No garbage collection in MVP.
- **Home** — docudex's root directory, default `~/.docudex`, overridable via the
  `DOCUDEX_HOME` environment variable (or `--home`). Holds the store and the index.
- **Module proxy** — the upstream serving module source zips via the Go module proxy
  protocol (`/@v/list`, `/@v/<version>.info`, `/@v/<version>.zip`, `/@latest`). Base URL
  is configurable (default `https://proxy.golang.org`), via `DOCUDEX_PROXY_URL` or
  `--proxy-url`. Acquisition is protocol-based — **never HTML scraping**.
- **Index** — a single global **SQLite FTS5** database (BM25 ranking) keyed by docset.
  Derived, rebuildable state (`docudex rebuild-index`); project scoping is a query-time
  filter on the manifest's docset list. Shared by CLI search and the web UI.
- **Adapter** — the ecosystem-specific boundary (proxy client + extractor + renderer).
  Go is the only adapter in the MVP; the boundary is the seam Rust (and later
  ecosystems) will implement.
- **Drift** — divergence between `go.mod` versions and the manifest's versions.
  Manifest-reading commands warn about drift and name the fix command; drift is
  **never auto-healed**. Only manifest-*writing* commands (`init`, `sync --from-gomod`)
  read `go.mod`.

## CLI surface (MVP)

`init`, `add`, `remove`, `sync` (with `--from-gomod`), `list`, `search`, `serve`,
`rebuild-index`. Cobra command tree. See `docs/adr/` for load-bearing decisions.

## Architecture (module layout intent)

Names, not paths — the intended internal packages as behaviour lands:

- **Go adapter** — proxy client, source extractor (`go/parser`, `go/doc`,
  `go/doc/comment`), and Markdown renderer (pkgsite is the fidelity reference).
- **Store** — reads/writes the content-addressed Markdown store.
- **Manifest** — reads/writes `docudex.toml`; computes drift against `go.mod`.
- **Index** — the SQLite FTS5 database; indexing and querying.
- **Server** — the `serve` HTTP server exposing a private/unstable JSON API; the
  React + React Router v7 SPA is embedded via `go:embed` for single-binary distribution.

Current state (ticket ENGG-6): the module skeleton, the Cobra command tree with stub
commands, CI, and configuration resolution exist. Real behaviour lands in later tickets.

## The two tested seams

Docudex's test strategy exercises external behaviour through the CLI and the serve JSON
API — never mocked internals. Two configuration values are the seams that make this
possible, and both are resolved in `internal/config`:

- **Module-proxy URL (primary seam)** — tests point docudex at a `file://` fixture
  directory laid out in module-proxy format, so the fetch pipeline runs against
  purpose-built fixture modules with no network. Set via `DOCUDEX_PROXY_URL` / `--proxy-url`.
- **Home directory (secondary seam)** — every test gets an isolated temporary store.
  Set via `DOCUDEX_HOME` / `--home`. Config hygiene, not a code seam.

Precedence for both: **flag → environment variable → default**.
