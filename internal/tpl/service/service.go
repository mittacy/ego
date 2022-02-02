package service

import (
	"fmt"
	"github.com/mittacy/ego/internal/base"
	"github.com/mittacy/ego/internal/tpl/data"
	"github.com/mittacy/ego/internal/tpl/model"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

// CmdService the service command.
var CmdService = &cobra.Command{
	Use:   "service",
	Short: "Generate the service template implementations",
	Long:  "Generate the service template implementations. Example: ego tpl service xxx -t=app/internal",
	Run:   run,
}
var targetDir string

func init() {
	CmdService.Flags().StringVarP(&targetDir, "target-dir", "t", "app/internal", "generate target directory")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the service file. Example: ego tpl service xxx")
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

	AddService(modName, args[0])

	data.AddData(modName, args[0])

	model.AddModel(args[0])

	fmt.Println("success!")
}

func AddService(appName, name string) bool {
	to := fmt.Sprintf("%s/service/%s.go", targetDir, name)
	service := Service{
		AppName: appName,
		Name:    name,
	}

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s service already exists: %s\n", name, to)
		return false
	}

	b, err := service.execute()
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(to, b, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("create file %s\n", to)
	return true
}
