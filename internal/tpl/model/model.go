package model

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

// CmdModel the service command.
var CmdModel = &cobra.Command{
	Use:   "model",
	Short: "Generate the model template implementations",
	Long:  "Generate the model template implementations. Example: ego tpl model xxx -t=internal",
	Run:   run,
}
var targetDir string

func init() {
	CmdModel.Flags().StringVarP(&targetDir, "target-dir", "t", "internal", "generate target directory")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the model file. Example: ego tpl model xxx")
		return
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Target directory: %s does not exist\n", targetDir)
		return
	}

	AddModel(args[0])

	fmt.Println("success!")
}

func AddModel(name string) bool {
	to := fmt.Sprintf("%s/model/%s.go", targetDir, name)
	model := Model{Name: name}

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s model already exists: %s\n", name, to)
		return false
	}

	b, err := model.execute()
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(to, b, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("create file %s\n", to)
	return true
}
