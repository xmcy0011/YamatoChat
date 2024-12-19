package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestReplyOK(t *testing.T) {
	convey.Convey("TestReplyOK", t, func() {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		ReplyOK(c, http.StatusOK, "test")
		assert.Equal(t, c.Writer.Status(), http.StatusOK)
	})
}

func TestReplyError(t *testing.T) {
	convey.Convey("TestReplyError", t, func() {
		convey.Convey("http error", func() {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			ReplyError(c, NewHTTPError(BadRequest, "test"))
			assert.Equal(t, c.Writer.Status(), http.StatusBadRequest)
		})

		convey.Convey("other error", func() {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			ReplyError(c, errors.New("test"))
			assert.Equal(t, c.Writer.Status(), http.StatusInternalServerError)
		})
	})
}
