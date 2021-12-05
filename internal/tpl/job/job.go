package job

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mittacy/ego/internal/base"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

// CmdJob the api command.
var CmdJob = &cobra.Command{
	Use:   "job",
	Short: "Generate the async job template implementations",
	Long:  "Generate the async job template implementations. Example: ego tpl job xxx",
	Run:   run,
}

var targetDir string

func init() {
	CmdJob.Flags().StringVarP(&targetDir, "target-dir", "t", "internal", "generate target directory")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the api file. Example: ego tpl job xxx")
		return
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Target directory: %s does not exist.\n"+
			"Please make sure you operate in the go project root directory\n", targetDir)
		return
	}

	modName, err := base.ModulePath("go.mod")
	if modName == "" || err != nil {
		fmt.Printf("go.mod no exist.\nPlease make sure you operate in the go project root directory\n")
		return
	}

	name := args[0]

	// 检查目录
	dir := fmt.Sprintf("%s/job/%sJob", targetDir, name)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 创建目录
		if err := os.Mkdir(dir, 0711); err != nil {
			fmt.Fprintf(os.Stderr, "create %sJob directory err: %s\n", name, dir)
			return
		}
	}

	AddTask(modName, name)

	AddProcessor(modName, name)

	fmt.Println(color.WhiteString("Don't forget to add the %sJob.NewProcessor() to the app/job/job.go", name))
}

func AddTask(appName, name string) bool {
	to := fmt.Sprintf("%s/job/%sJob/task.go", targetDir, name)
	api := Task{AppName: appName, Name: name}

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s job task already exists: %s\n", name, to)
		return false
	}

	b, err := api.execute()
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(to, b, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("create file %s\n", to)
	return true
}

func AddProcessor(appName, name string) bool {
	to := fmt.Sprintf("%s/job/%sJob/processor.go", targetDir, name)
	api := Processor{AppName: appName, Name: name}

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s job task already exists: %s\n", name, to)
		return false
	}

	b, err := api.execute()
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(to, b, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("create file %s\n", to)
	return true
}
