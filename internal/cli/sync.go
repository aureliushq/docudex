package cli

import "github.com/spf13/cobra"

func newSyncCmd(_ configResolver) *cobra.Command {
	var fromGoMod bool

	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Fetch everything in the manifest that's missing from the store",
		Long: "Fetch every docset listed in docudex.toml that isn't already in the\n" +
			"store, making a fresh clone fully browsable. With --from-gomod, update\n" +
			"the manifest's default versions from go.mod (preserving extra pinned\n" +
			"versions) instead.",
		Args: cobra.NoArgs,
		RunE: stubRunE("sync"),
	}

	cmd.Flags().BoolVar(&fromGoMod, "from-gomod", false,
		"update manifest default versions from go.mod (preserving extra pinned versions)")

	return cmd
}
