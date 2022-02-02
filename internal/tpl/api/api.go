package api

import (
	"fmt"
	"github.com/mittacy/ego/internal/base"
	"github.com/mittacy/ego/internal/tpl/data"
	"github.com/mittacy/ego/internal/tpl/model"
	"github.com/mittacy/ego/internal/tpl/service"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

// CmdApi the api command.
var CmdApi = &cobra.Command{
	Use:   "api",
	Short: "Generate the api template implementations",
	Long:  "Generate the api template implementations. Example: ego tpl api xxx",
	Run:   run,
}
var (
	targetDir   string
	internalDir string
)

func init() {
	CmdApi.Flags().StringVarP(&targetDir, "target-dir", "t", "app", "generate target directory")
	CmdApi.Flags().StringVarP(&internalDir, "internal-dir", "i", "app/internal", "generate internal directory")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the api file. Example: ego tpl api xxx")
		return
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Target directory: %s does not exist.\n"+
			"Please make sure you operate in the go project root directory\n", targetDir)
		return
	}

	if _, err := os.Stat(internalDir); os.IsNotExist(err) {
		fmt.Printf("Target directory: %s does not exist.\n"+
			"Please make sure you operate in the go project root directory\n", internalDir)
		return
	}

	modName, err := base.ModulePath("go.mod")
	if modName == "" || err != nil {
		fmt.Printf("go.mod no exist.\nPlease make sure you operate in the go project root directory\n")
		return
	}

	name := args[0]

	model.AddModel(name)

	AddValidator(name)

	AddTransform(modName, name)

	AddApi(modName, name)

	service.AddService(modName, name)

	data.AddData(modName, name)

	fmt.Println("success!")
}

func AddApi(appName, name string) bool {
	to := fmt.Sprintf("%s/api/%s.go", targetDir, name)
	api := Api{AppName: appName, Name: name}

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s api already exists: %s\n", name, to)
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

func AddValidator(name string) bool {
	dir := fmt.Sprintf("%s/validator/%sValidator", internalDir, name)
	validator := Validator{Name: name}

	// 检查目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 创建目录
		if err := os.Mkdir(dir, 0711); err != nil {
			fmt.Fprintf(os.Stderr, "create %s validator directory err: %s\n", name, dir)
			return false
		}
	}

	// 检查文件
	to := fmt.Sprintf("%s/%s.go", dir, name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s validator already exists: %s\n", name, to)
		return false
	}

	b, err := validator.execute()
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(to, b, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("create file %s\n", to)
	return true
}

func AddTransform(appName, name string) bool {
	to := fmt.Sprintf("%s/transform/%s.go", internalDir, name)
	transform := Transform{AppName: appName, Name: name}

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s transform already exists: %s\n", name, to)
		return false
	}

	b, err := transform.execute()
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(to, b, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("create file %s\n", to)
	return true
}
