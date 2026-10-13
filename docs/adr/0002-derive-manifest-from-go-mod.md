# 0002 - Derive a manifest from go.mod

- 2026-09-26
- Status: Accepted

## Context

We need a source of truth for fetching and storing docsets. This source will act as the one place where the DocuDex CLI will look for versions of the packages for building documentation.

## Decision

We will generate a `docudex.toml` manifest file that will act as our source of truth, instead of using the `go.mod` file. 

## Consequences

- Having a manifest means we explicitly state the source of truth for the docset
- Since `go.mod` is decoupled from the manifest, users can add multiple versions of the same package for comparison or upgrades
- We need to watch for drifts between `docudex.toml` and `go.mod` as users add/update/remove dependencies
- We need to notify the user when drift happens and the user has to run a `sync` command to update the manifest and the docset
