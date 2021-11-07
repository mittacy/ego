package tpl

import (
	"github.com/mittacy/ego/internal/tpl/api"
	"github.com/mittacy/ego/internal/tpl/data"
	"github.com/mittacy/ego/internal/tpl/model"
	"github.com/mittacy/ego/internal/tpl/service"
	"github.com/mittacy/ego/internal/tpl/task"
	"github.com/spf13/cobra"
)

// CmdTpl represents the proto command.
var CmdTpl = &cobra.Command{
	Use:   "tpl",
	Short: "Generate the template files",
	Long:  "Generate the template files.",
	Run:   run,
}

func init() {
	CmdTpl.AddCommand(api.CmdApi)
	CmdTpl.AddCommand(service.CmdService)
	CmdTpl.AddCommand(data.CmdData)
	CmdTpl.AddCommand(model.CmdModel)
	CmdTpl.AddCommand(task.CmdTask)
}

func run(cmd *cobra.Command, args []string) {

}
