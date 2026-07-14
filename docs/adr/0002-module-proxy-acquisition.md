# 2. Acquire docs via the Go module proxy protocol

- Status: Accepted
- Date: 2026-07-14
- Spec: [ENGG-5](https://linear.app/docudex/issue/ENGG-5)

## Context

Version-exact API docs can be obtained by scraping pkg.go.dev, or by fetching module
source and extracting docs locally. Scraping is brittle (breaks on site redesigns),
hard to pin to exact versions, and unfriendly to private modules and offline use.

## Decision

Fetch `module@version` **source zips** via the documented Go module proxy protocol
(`/@v/list`, `/@v/<version>.info`, `/@v/<version>.zip`, `/@latest`) — **never HTML
scraping**. Extract and render docs locally from the source with the Go stdlib
machinery (`go/parser`, `go/doc`, `go/doc/comment`); pkgsite is the fidelity reference.

The proxy **base URL is configurable** (`DOCUDEX_PROXY_URL` env var or `--proxy-url`
flag), defaulting to `https://proxy.golang.org`. The protocol's `file://` support is
the primary test seam: tests point docudex at a `file://` fixture directory in
module-proxy layout and run the real pipeline with no network.

## Consequences

- Fetching is reliable and won't rot when a website is redesigned.
- Exact version pinning is native to the protocol.
- The same configurable-base-URL mechanism serves private proxies later (a monetization
  path) and fixture proxies in tests — one seam, three uses.
- Rendering fidelity (doc-comment → Markdown) becomes docudex's own responsibility and
  its highest-uncertainty component (spec's known risk #1).
- We use a dedicated `DOCUDEX_PROXY_URL` rather than `GOPROXY` to avoid GOPROXY's
  list / `direct` / `off` semantics; GOPROXY interop is deferred.
