package hook

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// AddGitCommitMsg 增加 git commit 注释规范
// 规范参考: http://www.ruanyifeng.com/blog/2016/01/commit_message_change_log.html
func AddGitCommitMsg() {
	// 检查 .git 目录
	gitDir := ".git"
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr,"Git directory: %s does not exist\n", gitDir)
		fmt.Fprintf(os.Stderr,"Make sure that you are currently a git repository and that you are currently in the project root directory\n")
		return
	}

	// 检查 .git/hooks 目录
	hookDir := gitDir + "/hooks"
	if _, err := os.Stat(hookDir); os.IsNotExist(err) {
		if err := os.MkdirAll(hookDir, 0700); err != nil {
			fmt.Fprintf(os.Stderr, "create %s directory err: %s\n", hookDir, err)
		}
	}

	// 检查 commit-msg 文件
	msgFileDir := hookDir + "/commit-msg"
	if _, err := os.Stat(msgFileDir); os.IsNotExist(err) {
		// 写入hooks/commit-msg
		if err := ioutil.WriteFile(msgFileDir, []byte(commitMsg), 0755); err != nil {
			log.Fatalf("写入%s commit-msg失败, err: %+v", msgFileDir, err)
		}
	}
}
