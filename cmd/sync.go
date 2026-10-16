/*
Copyright © 2026 Ilango Rajagopal <hey@i4o.dev>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newSyncCommand() *cobra.Command {
	// syncCmd represents the init command
	syncCmd := &cobra.Command{
		Use:   "sync",
		Args:  cobra.NoArgs,
		Short: "Fetch everything in the manifest that's missing from the store",
		Long: "Fetch every docset listed in docudex.toml that isn't already in the\n" +
			"store, making a fresh clone fully browsable. With --from-gomod, update\n" +
			"the manifest's default versions from go.mod (preserving extra pinned\n" +
			"versions) instead.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("sync called")
		},
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return syncCmd
}
