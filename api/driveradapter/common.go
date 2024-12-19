package driveradapter

import (
	"io"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/xeipuuv/gojsonschema"
	"github.com/xmcy0011/CoffeeChat/pkg/rest"
)

// MustNewSchema 创建schema
func MustNewSchema(jsonData string) *gojsonschema.Schema {
	schema, err := gojsonschema.NewSchema(gojsonschema.NewStringLoader(jsonData))
	if err != nil {
		panic(err)
	}
	return schema
}

// ValidateAndBindGin 验证并绑定请求体
func ValidateAndBindGin(c *gin.Context, schema *gojsonschema.Schema, data interface{}) error {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return rest.NewHTTPError(rest.BadRequest, err.Error())
	}

	result, err := schema.Validate(gojsonschema.NewBytesLoader(body))
	if err != nil {
		return rest.NewHTTPError(rest.BadRequest, err.Error())
	}

	if !result.Valid() {
		firstErr := "unknown error"
		if len(result.Errors()) > 0 {
			firstErr = result.Errors()[0].String()
		}
		return rest.NewHTTPError(rest.BadRequest, firstErr)
	}

	if err := jsoniter.Unmarshal(body, data); err != nil {
		return rest.NewHTTPError(rest.BadRequest, err.Error())
	}
	return nil
}

func BindGin(c *gin.Context, data interface{}) error {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(body, data)
}
