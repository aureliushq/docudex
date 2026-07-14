package cli

import "github.com/spf13/cobra"

func newServeCmd(_ configResolver) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the local web UI for browsing the project's docs",
		Long: "Serve a local web UI (embedded SPA) backed by the same index as\n" +
			"terminal search, for browsing this project's docsets by module,\n" +
			"package, and symbol.",
		Args: cobra.NoArgs,
		RunE: stubRunE("serve"),
	}
}
