package cli

import "github.com/spf13/cobra"

func newAddCmd(_ configResolver) *cobra.Command {
	return &cobra.Command{
		Use:   "add <module>[@version]",
		Short: "Add a docset to the project manifest",
		Long: "Add a module docset to docudex.toml beyond what go.mod lists, or an\n" +
			"extra version of an existing one. With no @version, resolve @latest\n" +
			"from the module proxy.",
		Args: cobra.ArbitraryArgs,
		RunE: stubRunE("add"),
	}
}
