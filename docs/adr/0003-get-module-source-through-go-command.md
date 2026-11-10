# 0003 - Get module source through Go command

- 2026-11-10
- Status: Accepted

## Context

DocuDex needs the source code for each dependency before it can build docs. We could write our own client for the Go module proxy or use
`go mod download`, which already does that. Writing our own proxy client means DocuDex runs without the Go toolchain installed. But we have
to handle private module support, run checksum checks, etc. We're only targeting Go developers for early access release and they already have
Go installed.

## Decision

We will fetch the module source code through the `go mod download -json m@v` command.

## Consequences

- Users should have the Go toolchain installed
- Checksum checks and private modules come for free since users will have configured it already
- Any proxy that Go supports already works, including https:// and file://. Tests use the file:// proxy and never need the network.
