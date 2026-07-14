# 3. Global content store + single FTS index

- Status: Accepted
- Date: 2026-07-14
- Spec: [ENGG-5](https://linear.app/docudex/issue/ENGG-5)

## Context

Stored docs and the search index could be scoped per-project or held globally.
Per-project storage duplicates identical `module@version` docs across every project that
uses them. For search, options were a per-project index, an ATTACH-per-docset scheme, or
one shared database; the grilling (2026-07-11) contested this and resolved on a single
global index.

## Decision

**Storage:** a single global content store under home at
`store/<module>@<version>/`, content-addressed by `module@version`. **Markdown is
canonical** (one file per package, one heading per exported symbol), plus metadata
sidecars. Projects sharing a `module@version` share one stored copy; different versions
of the same module coexist. No store garbage collection in the MVP.

**Search:** one global **SQLite FTS5** database, indexed at fetch time, keyed by docset,
with BM25 ranking. Project scoping is a **query-time filter** on the manifest's docset
list. The same index backs both CLI `search` and the web UI. The driver is **pure-Go
SQLite (no CGO)** to preserve single-binary distribution (see ADR-0004). The index is
**derived, rebuildable state** — `docudex rebuild-index` regenerates it from the Markdown
store, so a corrupted index never means re-fetching content.

## Consequences

- Disk isn't filled with duplicate docsets; multiple versions never conflict.
- Home is overridable (`DOCUDEX_HOME` / `--home`) so each test gets an isolated store —
  the secondary test seam.
- Markdown-on-disk is greppable and lets coding agents consume docs later with no
  integration.
- A single writer to one global index needs care around concurrent fetches (later work).
- No GC means the store grows unbounded until a future refcounting/GC feature.
