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

var (
	targetDir  string
	payloadDir string
	processDir string
)

func init() {
	CmdJob.Flags().StringVarP(&targetDir, "target-dir", "t", "app/job", "generate target directory")
	payloadDir = targetDir + "/job_payload"
	processDir = targetDir + "/job_process"
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the job file. Example: ego tpl job xxx")
		return
	}

	// 检查目录
	if _, err := os.Stat(payloadDir); os.IsNotExist(err) {
		// 创建目录
		if err := os.MkdirAll(payloadDir, 0700); err != nil {
			fmt.Fprintf(os.Stderr, "create %s directory err: %v\n", payloadDir, err)
			return
		}
	}

	if _, err := os.Stat(processDir); os.IsNotExist(err) {
		// 创建目录
		if err := os.MkdirAll(processDir, 0700); err != nil {
			fmt.Fprintf(os.Stderr, "create %s directory err: %v\n", processDir, err)
			return
		}
	}

	modName, err := base.ModulePath("go.mod")
	if modName == "" || err != nil {
		fmt.Printf("go.mod no exist.\nPlease make sure you operate in the go project root directory\n")
		return
	}
	name := args[0]

	AddPayload(modName, name)

	AddProcessor(modName, name)

	fmt.Println(color.WhiteString("Don't forget to add the job_payload/TypeName and job_process/Process to the " +
		"config/async_config/job.go->Jobs()"))
}

func AddPayload(appName, name string) bool {
	to := fmt.Sprintf("%s/%s.go", payloadDir, name)
	api := Task{AppName: appName, Name: name}

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s job payload already exists: %s\n", name, to)
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
	to := fmt.Sprintf("%s/%s.go", processDir, name)
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
