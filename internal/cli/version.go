package cli

// Build metadata, injected at release time by GoReleaser via -ldflags -X.
// The defaults apply to `go install` and local `go build` invocations, where
// no linker flags are set; released binaries report the real tag and SHA.
var (
	// version is the release version, e.g. "1.2.3". "dev" for unreleased builds.
	version = "dev"
	// commit is the git SHA the binary was built from. Empty for local builds.
	commit = ""
)
