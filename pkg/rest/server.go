package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReplyOK 成功响应
func ReplyOK(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

// ReplyError 错误响应
func ReplyError(c *gin.Context, err error) {
	var statusCode int
	var body interface{}

	switch e := err.(type) {
	case *HTTPError:
		statusCode = e.Code / 100
		body = e
	default:
		statusCode = http.StatusInternalServerError
		body = NewHTTPError(InternalServerError, e.Error())
	}

	c.JSON(statusCode, body)
}
