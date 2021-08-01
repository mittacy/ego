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
	"strings"
)

// CmdApi the api command.
var CmdApi = &cobra.Command{
	Use:   "api",
	Short: "Generate the api template implementations",
	Long:  "Generate the api template implementations. Example: ego tpl api xxx",
	Run:   run,
}
var targetDir string
var wireDir string

func init() {
	//CmdApi.Flags().StringVarP(&targetDir, "target-dir", "t", "app", "generate target directory")
	targetDir = "app"
	wireDir = "router"
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

	modName, err := base.ModulePath("go.mod")
	if modName == "" || err != nil {
		fmt.Printf("go.mod no exist.\nPlease make sure you operate in the go project root directory\n")
		return
	}

	name := args[0]

	model.AddModel(name)

	AddValidator(name)

	AddTransform(modName, name)

	AddApi(name)

	service.AddService(modName, name)

	data.AddData(modName, name)

	AddWire(modName, name)

	fmt.Println("success!")
}

func AddApi(name string) bool {
	to := fmt.Sprintf("%s/api/%s.go", targetDir, name)
	api := Api{Name: name}

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
	dir := fmt.Sprintf("%s/validator/%sValidator", targetDir, name)
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
	dir := fmt.Sprintf("%s/transform/%sTransform", targetDir, name)
	transform := Transform{AppName: appName, Name: name}

	// 检查目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 创建目录
		if err := os.Mkdir(dir, 0711); err != nil {
			fmt.Fprintf(os.Stderr, "create %s transform directory err: %s\n", name, dir)
			return false
		}
	}

	// 检查文件
	to := fmt.Sprintf("%s/%s.go", dir, name)
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

func AddWire(appName, name string) bool {
	dir := wireDir

	// 检查目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Target directory: %s does not exist.\n"+
			"Please make sure you operate in the go project root directory\n", dir)
		return false
	}

	wirePath := fmt.Sprintf("%s/wire.go", dir)
	wire := NewWire(appName, name)

	// 检查是否已经包含函数
	b, err := ioutil.ReadFile(wirePath)
	if err != nil {
		fmt.Printf("Target wire: %s does not exist.\n", wirePath)
		return false
	}

	if strings.Contains(string(b), fmt.Sprintf("Init%sApi", wire.Name)) {
		fmt.Printf("function Init%sApi already exists in %s\n", wire.Name, wirePath)
		return false
	}

	// 打开 wire.go 文件
	file, err := os.OpenFile(wirePath, os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s no exists\n", dir)
		return false
	}
	defer file.Close()

	// 追加函数
	b, err = wire.execute()
	if err != nil {
		log.Fatal(err)
	}

	if _, err = file.Write(b); err != nil {
		fmt.Printf("create function Init%sApi err: %v\n", wire.Name, err)
	}

	fmt.Printf("create function Init%sApi in %s\n", wire.Name, wirePath)
	return true
}
