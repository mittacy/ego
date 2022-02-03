package git

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

// CmdGit
var CmdGit = &cobra.Command{
	Use:   "git",
	Short: "the git plugin",
	Long:  "the git plugin. Example: ego plugin git xxx -t=./git",
	Run:   run,
}

var (
	targetDir string
	pluginKey string
)

func init() {
	CmdGit.Flags().StringVarP(&targetDir, "target-dir", "t", "./.git", "generate target directory")
	CmdGit.Flags().StringVarP(&pluginKey, "plugin-key", "k", CommitLint, "plugin key: commitLint/……")
}

func run(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		log.Printf("git plugin: %s directory not exists\n", targetDir)
		return
	}

	switch pluginKey {
	case CommitLint:
		AddGitCommitMsg(targetDir)
	default:
		log.Printf("git plugin: %s git-plugin does not exist", pluginKey)
	}
}
