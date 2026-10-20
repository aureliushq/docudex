/*
Copyright © 2026 Ilango Rajagopal <hey@i4o.dev>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func NewRootCommand() *cobra.Command {
	root := &cobra.Command{
		Aliases: []string{"dd"},
		Use:     "docudex",
		Short:   "Per-project, version-exact Go API reference docs",
		Long: "docudex aggregates version-exact API reference documentation for a\n" +
			"Go project's dependencies into a local store you can search from the\n" +
			"terminal or browse in a local web UI — scoped to exactly what your\n" +
			"project uses, and fully available offline.",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.docudex.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	root.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	root.AddCommand(
		newInitCommand(),
		newSyncCommand(),
	)

	return root
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := NewRootCommand()
	if err := fang.Execute(context.Background(), cmd); err != nil {
		fmt.Fprintln(os.Stderr, "docudex error: ", err)
		os.Exit(1)
	}
}
