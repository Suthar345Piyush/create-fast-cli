// root cobra command for create-fast-cli

package cmd

import (
	"fmt"
	"os"

	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/logger"
	"github.com/spf13/cobra"
)

// our cli version

const version = "0.1.0"

/*
"verbose" is typically a persistent boolean flag defined in the cmd/root.go file. It is used to enable detailed console output (logs, debug info, or progress updates) across the entire CLI application.
*/

var verbose bool

var rootCmd = &cobra.Command{
	Use:     "create-fast-cli",
	Version: version,
	Short:   "Scaffold production-ready Go CLI projects in seconds",
	Long: `
⚡ FastCLI Starter

create-fast-cli scaffolds opinionated, production-ready  CLI projects
for Go - features included (Cobra, Viper, Zap, Bubbletea and more).

Run without arguments to launch the interactive wizard:

npx create-fast-cli   (or)   create-fast-cli create`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		if err := logger.Init(verbose); err != nil {
			return err
		}

		return config.InitViper()

	},

	RunE: func(cmd *cobra.Command, args []string) error {
		return runCreate(cmd, args)
	},
}

// this execute is called by main.go - single entry point

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable debug logging")
}
