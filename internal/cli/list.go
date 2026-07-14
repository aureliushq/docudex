package cli

import "github.com/spf13/cobra"

func newListCmd(_ configResolver) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List the project's docsets, versions, and fetch state",
		Long:  "Show every docset in docudex.toml, its version(s), and whether each is fetched into the store.",
		Args:  cobra.NoArgs,
		RunE:  stubRunE("list"),
	}
}
