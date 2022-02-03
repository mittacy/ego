package plugin

import (
	"github.com/mittacy/ego/internal/plugin/git"
	"github.com/spf13/cobra"
)

// CmdPlugin represents the proto command.
var CmdPlugin = &cobra.Command{
	Use:   "plugin",
	Short: "Introduces the specified plug-in to the current project.",
	Long:  "Introduces the specified plug-in to the current project.",
	Run:   run,
}

func init() {
	CmdPlugin.AddCommand(git.CmdGit)
}

func run(cmd *cobra.Command, args []string) {
}
