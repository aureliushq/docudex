# docudex

**Per-project, version-exact Go API reference docs.**

docudex aggregates version-exact API reference documentation for a Go project's
dependencies into a local store you can search from the terminal or browse in a
local web UI — scoped to exactly what your project uses, and fully available
offline.

Run `docudex init` in a Go repo and it reads `go.mod`, materializes API
reference docs for your direct dependencies (plus the Go standard library)
pinned to the exact versions you use, and makes them searchable (`docudex
search`) or browse-able (`docudex serve`).

## License

[Apache-2.0](./LICENSE)
