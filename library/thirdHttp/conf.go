package thirdHttp

import "time"

const (
	CodeSuccess = 0
	CodeUnknown = 1
)

type Reply struct {
	Code int
	Msg string
	Data interface{}
}

var defaultResTimeFormat = []string{
	"2006-01-02 15:04:05",
	time.RFC3339,
	time.RFC3339Nano,
}
