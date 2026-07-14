package cli

import "github.com/spf13/cobra"

func newRebuildIndexCmd(_ configResolver) *cobra.Command {
	return &cobra.Command{
		Use:   "rebuild-index",
		Short: "Rebuild the search index from the Markdown store",
		Long: "Rebuild the derived full-text search index from the canonical Markdown\n" +
			"store, so a corrupted index never means re-fetching content.",
		Args: cobra.NoArgs,
		RunE: stubRunE("rebuild-index"),
	}
}
