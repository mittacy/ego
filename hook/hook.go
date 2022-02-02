package hook

const (
	BeforeStart = "BeforeStart"
)

var (
	hooks      map[string]func()
	hooksQueue []string
)

func init() {
	hooks = map[string]func(){}
	hooksQueue = []string{BeforeStart}
}

// Register 注册事件
// @param hookName 事件名
// @param do
func Register(hookName string, do func()) {
	for _, v := range hooksQueue {
		if hookName == v {
			hooks[hookName] = do
		}
	}
}

// Run 启动事件
// @param hookName 事件名
func Run(hookName string) {
	if hook, ok := hooks[hookName]; ok {
		hook()
	}
}

// RunAll 启动全部事件
//func RunAll() {
//	for _, hookKey := range hooksQueue {
//		if hook, ok := hooks[hookKey]; ok {
//			hook()
//		}
//	}
//}
