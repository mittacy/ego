package upgrade

import (
	"fmt"
	"github.com/mittacy/ego/internal/base"
	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the ego tools",
	Long:  "Upgrade the ego tools. Example: ego upgrade",
	Run:   Run,
}

// Run upgrade the ego tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoInstall()
	if err != nil {
		fmt.Println(err)
	}
}
