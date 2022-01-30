package thirdHttp

import (
	"github.com/go-resty/resty/v2"
	"github.com/mittacy/log"
	"github.com/pkg/errors"
	"strings"
	"time"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New("third_http")
}

// Client 获取新的resty.Client
func Client() *resty.Client {
	return resty.New()
}

// Get GET请求，返回数据为map结构
// @param host 域名，example: https://www.baidu.com
// @param uri example: /user
// @param timeout 超时控制  example: time.Second*5
// @return map[string]interface{} data结构数据
// @return int 返回的业务状态码
// @return error
func Get(host, uri string, timeout time.Duration) (map[string]interface{}, int, error) {
	url := fullUrl(host, uri)

	client := resty.New().SetTimeout(timeout)
	reply := Reply{}

	resp, err := client.R().SetResult(&reply).ForceContentType("application/json").Get(url)
	if err != nil {
		logger.Errorw(url, "resp", resp, "err", err)
		return nil, CodeUnknown, errors.WithStack(err)
	} else {
		logger.Infow(url, "resp", resp)
	}

	if !resp.IsSuccess() {
		return nil, CodeUnknown, errors.New(resp.String())
	}

	if reply.Code != CodeSuccess {
		return nil, reply.Code, errors.New(reply.Msg)
	}

	return resPackage(reply), reply.Code, nil
}

// GetParams GET请求，返回数据为map结构
// @param host 域名，example: https://www.baidu.com
// @param uri example: /user
// @param params 请求参数
// @param timeout 超时控制  example: time.Second*5
// @return map[string]interface{} data结构数据
// @return int 返回的业务状态码
// @return error
func GetParams(host, uri string, params map[string]string, timeout time.Duration) (map[string]interface{}, int, error) {
	url := fullUrl(host, uri)

	client := resty.New().SetTimeout(timeout)
	reply := Reply{}

	resp, err := client.R().SetQueryParams(params).SetResult(&reply).ForceContentType("application/json").Get(url)
	if err != nil {
		logger.Errorw(url, "res", resp, "err", err)
		return nil, CodeUnknown, errors.WithStack(err)
	} else {
		logger.Infow(url, "res", resp)
	}

	if !resp.IsSuccess() {
		return nil, CodeUnknown, errors.New(resp.String())
	}

	if reply.Code != CodeSuccess {
		return nil, reply.Code, errors.New(reply.Msg)
	}

	return resPackage(reply), reply.Code, nil
}

// Post POST请求
// @param host 域名，example: https://www.baidu.com
// @param uri example: /user
// @param body 请求体数据，struct/map/[]byte/……
// @param timeout 超时控制  example: time.Second*5
// @return map[string]interface{} data结构数据
// @return int 返回的业务状态码
// @return error
func Post(host, uri string, body interface{}, timeout time.Duration) (map[string]interface{}, int, error) {
	url := fullUrl(host, uri)

	client := resty.New().SetTimeout(timeout)
	reply := Reply{}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&reply).
		Post(url)
	if err != nil {
		logger.Errorw(host+uri, "res", resp, "err", err)
		return nil, CodeUnknown, errors.WithStack(err)
	} else {
		logger.Infow(host+uri, "res", resp)
	}

	if !resp.IsSuccess() {
		return nil, CodeUnknown, errors.New(resp.String())
	}

	if reply.Code != CodeSuccess {
		return nil, reply.Code, errors.New(reply.Msg)
	}

	return resPackage(reply), reply.Code, nil
}

func fullUrl(host, uri string) string {
	uri = strings.TrimLeft(uri, "/")
	host = strings.TrimRight(host, "/")
	return host + "/" + uri
}

func resPackage(res Reply) map[string]interface{} {
	if v, ok := res.Data.(map[string]interface{}); ok {
		return v
	}

	// 其他特殊类型包装
	return map[string]interface{}{
		"data": res.Data,
	}
}
