package project

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/mittacy/ego/internal/base"
	"github.com/mittacy/ego/internal/utils"
	"os"
	"path"
)

// Project is a project template.
type Project struct {
	Name string
	Path string
}

const (
	replaceStr = "ego-layout"
)

// New new a project from remote repo.
func (p *Project) New(ctx context.Context, dir string, layout string, branch string) error {
	to := path.Join(dir, p.Name)

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("š« %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "š Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		survey.AskOne(prompt, &override)
		if !override {
			return err
		}
		os.RemoveAll(to)
	}

	fmt.Printf("š Creating service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)

	repo := base.NewRepo(layout, branch)

	if err := repo.CopyTo(ctx, to, p.Path, []string{".git", ".github"}); err != nil {
		return err
	}

	os.Rename(
		path.Join(to, "cmd", "server"),
		path.Join(to, "cmd", p.Name),
	)
	base.Tree(to, dir)

	fmt.Printf("\nšŗ Project creation succeeded %s\n", color.GreenString(p.Name))

	fmt.Printf("Wait a moment, the program is in the final configuration work\n")

	// ęæę¢é”¹ē®äø­ēå­ē¬¦äø²
	base.Replace(to, replaceStr, p.Name)

	// å¤å¶éē½®ęä»¶
	developEnv := fmt.Sprintf("%s/.env.development", to)
	localEnv := fmt.Sprintf("%s/.env", to)
	_ = utils.Copy(localEnv, developEnv)

	fmt.Print("š» Use the following command to start the project š:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ go mod download "))
	fmt.Println(color.WhiteString("edit the .env.* configuration file"))
	fmt.Println(color.WhiteString("$ go build -o ./bin/server ."))
	fmt.Println(color.WhiteString("$ ./bin/server start http -c=.env.development -e=development -p=8080\n"))
	fmt.Println("		š¤ Thanks for using ego")
	fmt.Println("	š Tutorial: http://www.mittacy.com/column/1633512445750")
	return nil
}
