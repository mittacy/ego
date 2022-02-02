package hook

const commitMsg = `#!/bin/sh

# 目录 merge request
MERGE_MSG=`+"`" +`cat $1 | egrep '^Merge branch*'`+"`" +`

if [ "$MERGE_MSG" != "" ]; then
	exit 0
fi

COMMIT_MSG=`+"`" + `cat $1 | egrep "^(feat|fix|docs|style|refactor|test|chore)(\(\w+\))?:\s(\S|\w)+"` + "`" + `

if [ "$COMMIT_MSG" = "" ]; then
	echo "INVALID COMMIT MSG: does not match "\<type>(\<scope\>): \<subject\>" !"
	echo "Commit message 格式错误，请参照: http://www.ruanyifeng.com/blog/2016/01/commit_message_change_log.html\n"
	exit 1
fi

if [ ${#COMMIT_MSG} -lt 10 ]; then
	echo "Commit Message Too Shot, Please show me more detail!(Greater than 10 characters)\n"
	exit 1
fi
`