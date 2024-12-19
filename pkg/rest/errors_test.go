package rest

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func Test_NewHTTPError(t *testing.T) {
	convey.Convey("TestNewHTTPError", t, func() {
		err := NewHTTPError(BadRequest, "test")
		assert.Equal(t, err.Code, BadRequest)
		assert.Equal(t, err.Error(), `{"code":40000,"cause":"errors_test.go:func1:12","description":"test"}`)
	})
}

func Test_Error(t *testing.T) {
	convey.Convey("TestError", t, func() {
		err := NewHTTPError(BadRequest, "test")
		assert.Equal(t, err.Error(), `{"code":40000,"cause":"errors_test.go:func1:20","description":"test"}`)
	})
}
