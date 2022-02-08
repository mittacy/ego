package data

import (
	"fmt"
	"github.com/mittacy/ego/internal/base"
	"github.com/mittacy/ego/internal/tpl/model"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

// CmdData the data command.
var CmdData = &cobra.Command{
	Use:   "data",
	Short: "Generate the data template implementations",
	Long:  "Generate the data template implementations. Example: ego tpl data xxx -t=app/internal",
	Run:   run,
}

const (
	InjectMysql = "mysql"
	InjectMongo = "mongo"
	InjectRedis = "redis"
	InjectHttp  = "http"
)

var (
	targetDir      string
	databaseHandle []string
)

func init() {
	CmdData.Flags().StringVarP(&targetDir, "target-dir", "t", "app/internal", "generate target directory")
	CmdData.Flags().StringArrayVarP(&databaseHandle, "database-handle", "d", []string{InjectMysql, InjectRedis}, "inject database handle:mysql,redis,mongo,http")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the data file. Example: ego tpl data xxx")
		return
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Target directory: %s does not exist\n", targetDir)
		return
	}

	modName, err := base.ModulePath("go.mod")
	if modName == "" || err != nil {
		fmt.Printf("go.mod no exist.\nPlease make sure you operate in the go project root directory\n")
		return
	}

	AddData(modName, args[0], databaseHandle)

	model.AddModel(args[0])

	fmt.Println("success!")
}

func AddData(appName, name string, databaseHandle []string) bool {
	to := fmt.Sprintf("%s/data/%s.go", targetDir, name)

	data := Data{
		AppName: appName,
		Name:    name,
	}
	data.parseInject(databaseHandle)

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s data already exists: %s\n", name, to)
		return false
	}

	b, err := data.execute()
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(to, b, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("create file %s\n", to)
	return true
}
