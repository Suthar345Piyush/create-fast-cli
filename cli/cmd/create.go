// the create subcommand that will launches the full application and scaffold

package cmd

import (
	"fmt"
	"os"

	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/prompt"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/scaffold"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Launch the interactive project wizard",
	Long:  `create launches the step-by-step interactive wizard that collects your project preferences and scaffolds a ready-to-build Go CLI project.`,
	RunE:  runCreate,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

// run create function - it will run the create command

func runCreate(_ *cobra.Command, _ []string) error {

	cfg, err := prompt.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error collecting answers:", err)
		return err
	}

	// scaffold pipeline(render -> write -> ide open) with the tui

	if err := scaffold.Generate(cfg); err != nil {
		fmt.Fprintln(os.Stderr, "\nScaffold failed:", err)
		return err
	}

	return nil

}
