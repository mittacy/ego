package task

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

// CmdTask the service command.
var CmdTask = &cobra.Command{
	Use:   "task",
	Short: "Generate the task template implementations",
	Long:  "Generate the task template implementations. Example: ego tpl task xxx -t=app",
	Run:   run,
}
var targetDir string

func init() {
	CmdTask.Flags().StringVarP(&targetDir, "target-dir", "t", "app", "generate target directory")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the task file. Example: ego tpl task xxx")
		return
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Target directory: %s does not exist\n", targetDir)
		return
	}

	AddTask(args[0])

	fmt.Println("success!")
}

func AddTask(name string) bool {
	to := fmt.Sprintf("%s/task/%s.go", targetDir, name)
	model := Model{Name: name}

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s task already exists: %s\n", name, to)
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