package rest

import (
	"fmt"
	"runtime"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

const (
	// BadRequest 请求错误
	BadRequest = 40000
	// InternalServerError 服务器错误
	InternalServerError = 50000
	// Unauthorized 未授权
	Unauthorized = 40100
	// NotFound 未找到
	NotFound = 40400
	// Forbidden 禁止访问
	Forbidden = 40300
	// Conflict 冲突
	Conflict = 40900
)

// HTTPError 错误信息
type HTTPError struct {
	Code        int    `json:"code"`
	Cause       string `json:"cause"`
	Description string `json:"description"`
}

// NewHTTPError 创建错误信息
func NewHTTPError(code int, description string) *HTTPError {
	// 获取调用堆栈
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return &HTTPError{Code: code, Description: description}
	}
	fileName := file[strings.LastIndex(file, "/")+1:]
	funcName := runtime.FuncForPC(pc).Name()
	funcName = funcName[strings.LastIndex(funcName, ".")+1:]

	return &HTTPError{Code: code, Description: description, Cause: fmt.Sprintf("%s:%s:%d", fileName, funcName, line)}
}

// Error 返回错误信息
func (e *HTTPError) Error() string {
	json, _ := jsoniter.Marshal(e)
	return string(json)
}
