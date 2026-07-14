package cli

import "github.com/spf13/cobra"

func newInitCmd(_ configResolver) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Create docudex.toml from the project's go.mod",
		Long: "Read the project's go.mod, materialize a docudex.toml listing the\n" +
			"direct dependencies (plus the Go standard library) at their exact\n" +
			"versions, and fetch their docs into the store.",
		Args: cobra.NoArgs,
		RunE: stubRunE("init"),
	}
}
