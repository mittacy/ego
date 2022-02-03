package hook

import "github.com/mittacy/ego/internal/plugin/git"

// AddGitCommitMsg 增加 git commit 注释规范
// 规范参考: http://www.ruanyifeng.com/blog/2016/01/commit_message_change_log.html
func AddGitCommitMsg() {
	// 检查 .git 目录
	git.AddGitCommitMsg(".git")
}
