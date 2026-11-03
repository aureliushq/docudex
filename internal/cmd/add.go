/*
Copyright © 2026 Ilango Rajagopal <hey@i4o.dev>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newAddCommand() *cobra.Command {
	// syncCmd represents the init command
	syncCmd := &cobra.Command{
		Use:   "add",
		Args:  cobra.ArbitraryArgs,
		Short: "Add a docset to the project manifest",
		Long: "Add a module docset to docudex.toml beyond what go.mod lists, or an\n" +
			"extra version of an existing one. With no @version, resolve @latest\n" +
			"from the module proxy.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add called")
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
