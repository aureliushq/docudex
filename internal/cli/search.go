package cli

import "github.com/spf13/cobra"

func newSearchCmd(_ configResolver) *cobra.Command {
	return &cobra.Command{
		Use:   "search <query>",
		Short: "Full-text search across the project's docsets",
		Long: "Run full-text search over only this project's docsets, ranked by\n" +
			"relevance with highlighted snippets and the owning package/symbol.",
		Args: cobra.ArbitraryArgs,
		RunE: stubRunE("search"),
	}
}
