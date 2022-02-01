package async

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego/library/log"
)

type Job struct {
	TypeName string
	Handler  asynq.Handler
}

var (
	ErrNoInit = errors.New("asynq: please initialize with async.Init() method")
	isInit    bool
	logger    *log.Logger
	client    *asynq.Client
)

func InitLog() {
	logger = log.New("job")
}

// Init 初始化队列服务
func Init(opt asynq.RedisClientOpt) {
	InitLog()
	client = asynq.NewClient(opt)
	isInit = true
}

// Enqueue 任务加入队列，使用前需要先初始化队列服务: async.Init(...)
// @param data 任务结构体数据
// @param typeName 任务typeName
// @param opts 可选配置，失败重试、指定队列、优先级……
// @return error
func Enqueue(typeName string, data interface{}, opts ...asynq.Option) error {
	if !isInit {
		return ErrNoInit
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	task := asynq.NewTask(typeName, payload, opts...)
	return enqueueByTask(task, opts...)
}

// enqueueByTask 任务加入任务队列
// @param task 任务
// @param opts 可选配置，失败重试、指定队列、优先级……
// @return error
func enqueueByTask(task *asynq.Task, opts ...asynq.Option) error {
	// 加入队列
	info, err := client.Enqueue(task, opts...)
	if err != nil {
		msg := fmt.Sprintf("send task fail, type: %s", task.Type())
		logger.Errorw(msg, "payload", task.Payload(), "err", err)
		return err
	}

	logger.Infof("%s, taskId: %s 加入 %s 队列成功", task.Type(), info.ID, info.Queue)
	return nil
}
