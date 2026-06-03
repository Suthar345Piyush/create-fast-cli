// shell completions using cobra

// this completions will wires up shell completion sub-commands for bash, zsh, fish via cobra built in generation helpers

// Bash will works on Windows
// for Zsh, Fish use WSL

package completions

import (
	"os"

	"github.com/spf13/cobra"
)

//'completion' sub-commmand under root - have three childrens: bash, zsh, fish

func AddTo(root *cobra.Command) {

	// completion command

	completionCommand := &cobra.Command{
		Use:   "completion [bash|zsh|fish]",
		Short: "Generate shell completion scripts",
		// Long: `To load completions:

		// Bash:
		//   source <(create-fast-cli completion bash)
		// 	#Persist:
		// 	create-fast-cli completion bash > /etc/bash_completion.d/create-fast-cli

		// Zsh:
		//   source <(create-fast-cli completion zsh)
		// 	#Persist
		// 	create-fast-cli completion zsh > "${fpath[1]}/_create-fast-cli"

		// Fish:
		//   create-fast-cli completion fish | source
		// 	#Persist
		// 	create-fast-cli completion fish > ~/.config/fish/completions/create-fast-cli.fish

		//   `,

		ValidArgs:             []string{"bash", "zsh", "fish"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return root.GenBashCompletion(os.Stdout)
			case "zsh":
				return root.GenZshCompletion(os.Stdout)
			case "fish":
				return root.GenFishCompletion(os.Stdout, true)
			}
			return nil
		},
	}

	root.AddCommand(completionCommand)
}
