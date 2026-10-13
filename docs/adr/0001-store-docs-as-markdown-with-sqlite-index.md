# 0001 - Store docs as Markdown with a SQLite index 

- 2026-09-26
- Status: Accepted

## Context

Docsets need to be stored locally so humans and agents can use it to build software. We have to pick a storage format that works for both and have the following features:

- We should be able to run `grep` and `find` commands to search from the command line
- We need to render the documentation in the local web UI
- We need to have full-text search

## Decision

We will store the documentation in a standardized Markdown format, index those documents and store the index in a SQLite database. We will do this instead of picking either - Markdown-only or SQLite-only. Markdown will be the source of truth for the docset and we'll use SQLite as a cache that can be deleted safely. `sync` will rebuild the cache. We will use `modernc.org/sqlite` as our SQLite driver, instead of `mattn/go-sqlite3`.

## Consequences

- Agents can run `grep` and `find` commands to search the Markdown
- We have to rebuild the cache when user runs a `sync`
- Since we're using the `modernc.org/sqlite` driver, a large `add` or `sync` will be slow, so we need to benchmark a real, large monorepo
- Since we're using the `modernc.org/sqlite` driver, we don't need to ship a C toolchain and users who are building from source don't need to have it installed. We can ship a single binary and have a simple CI/CD.
