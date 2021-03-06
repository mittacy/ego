package main

import (
	"github.com/mittacy/ego/internal/plugin"
	"github.com/mittacy/ego/internal/project"
	"github.com/mittacy/ego/internal/tpl"
	"github.com/mittacy/ego/internal/upgrade"
	"github.com/spf13/cobra"
	"log"
)

const version = "v1.6.23"

var rootCmd = &cobra.Command{
	Use:     "ego",
	Short:   "ego: An elegant toolkit for Gin.",
	Long:    `ego: An elegant toolkit for Gin.`,
	Version: version,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(tpl.CmdTpl)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(plugin.CmdPlugin)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
