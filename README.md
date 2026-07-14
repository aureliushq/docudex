# docudex

**Per-project, version-exact Go API reference docs.**

docudex aggregates version-exact API reference documentation for a Go project's
dependencies into a local store you can search from the terminal or browse in a
local web UI — scoped to exactly what your project uses, and fully available
offline.

Run `docudex init` in a Go repo and it reads `go.mod`, materializes API
reference docs for your direct dependencies (plus the Go standard library)
pinned to the exact versions you use, and makes them searchable (`docudex
search`) or browsable (`docudex serve`).

> **Status:** early development. The CLI surface exists; real behaviour lands
> ticket by ticket. See [`CONTEXT.md`](./CONTEXT.md) for the design of record.

## Install

### Homebrew (macOS)

```sh
brew install --cask aureliushq/tap/docudex
```

On Linux, use the install script or a manual download below.

### Install script (macOS / Linux)

```sh
curl -fsSL https://raw.githubusercontent.com/aureliushq/docudex/main/install.sh | sh
```

The script detects your OS and architecture, downloads the matching release
binary from GitHub, verifies its SHA-256 checksum, and installs it. Override the
install directory with `DOCUDEX_INSTALL_DIR` and pin a version with
`DOCUDEX_VERSION`:

```sh
curl -fsSL https://raw.githubusercontent.com/aureliushq/docudex/main/install.sh \
  | DOCUDEX_VERSION=v0.1.0 DOCUDEX_INSTALL_DIR="$HOME/.local/bin" sh
```

### Manual download

Grab the prebuilt binary for your platform from the
[releases page](https://github.com/aureliushq/docudex/releases)
(`docudex_<version>_<os>_<arch>`, or the `.exe` on Windows), rename it to
`docudex`, make it executable (`chmod +x docudex`), and put it somewhere on your
`PATH`.

### From source

```sh
go install github.com/aureliushq/docudex@latest
```

Binaries built this way report `dev` as their version; released binaries carry
the real tag and commit (`docudex --version`).

## Usage

```sh
docudex init            # create docudex.toml from the project's go.mod
docudex search <query>  # search the docs your project uses
docudex serve           # browse the docs in a local web UI
```

Run `docudex --help` for the full command tree.

## License

[Apache-2.0](./LICENSE)
