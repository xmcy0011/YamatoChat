package driveradapter

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

type testKey struct {
	Key string `json:"key"`
}

func TestValidateAndBindGin(t *testing.T) {
	convey.Convey("TestValidateAndBindGin", t, func() {

		loader := gojsonschema.NewStringLoader(`{"type": "object", "properties": {"key": {"type": "string"}}}`)
		schema, err := gojsonschema.NewSchema(loader)
		assert.NoError(t, err)

		req := testKey{Key: "value"}
		data, _ := json.Marshal(req)

		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		ctx.Request = httptest.NewRequest("POST", "/api/v1/user/register", bytes.NewBuffer(data))

		result := map[string]interface{}{}
		err = ValidateAndBindGin(ctx, schema, &result)
		assert.NoError(t, err)
		assert.Equal(t, result["key"], "value")
	})
}

func TestBindGin(t *testing.T) {
	convey.Convey("TestBindGin", t, func() {
		req := testKey{Key: "value"}
		data, _ := json.Marshal(req)

		ctx := &gin.Context{Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer(data))}}

		result := map[string]interface{}{}
		err := BindGin(ctx, &result)
		assert.NoError(t, err)
		assert.Equal(t, result["key"], "value")
	})
}
