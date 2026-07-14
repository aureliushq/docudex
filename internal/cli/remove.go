package cli

import "github.com/spf13/cobra"

func newRemoveCmd(_ configResolver) *cobra.Command {
	return &cobra.Command{
		Use:   "remove <module>",
		Short: "Remove a docset from the project manifest",
		Long:  "Drop a module docset from docudex.toml so the project's doc scope stays relevant.",
		Args:  cobra.ArbitraryArgs,
		RunE:  stubRunE("remove"),
	}
}
