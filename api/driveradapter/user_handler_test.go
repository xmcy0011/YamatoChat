package driveradapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"github.com/xmcy0011/CoffeeChat/api/driveradapter/dto"
	"github.com/xmcy0011/CoffeeChat/api/interfaces"
	"github.com/xmcy0011/CoffeeChat/api/interfaces/mocks"
	"go.uber.org/mock/gomock"
)

func TestUserHandler_RegisterUser(t *testing.T) {
	convey.Convey("test register user", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserLogic := mocks.NewMockUserLogic(ctrl)
		userHandler := NewUserHandler()
		userHandler.userLogic = mockUserLogic

		convey.Convey("bind error", func() {
			ginCtx, _ := gin.CreateTestContext(httptest.NewRecorder())
			ginCtx.Request = httptest.NewRequest("POST", "/api/v1/user/register", nil)

			userHandler.registerUser(ginCtx)

			assert.Equal(t, ginCtx.Writer.Status(), http.StatusBadRequest)
		})

		convey.Convey("validate error", func() {
			convey.Convey("invalid user_name", func() {
				req := dto.UserRegisterRequest{
					Username: "",
					NickName: "test",
					Password: "123456",
				}
				reqData, _ := json.Marshal(req)

				w := httptest.NewRecorder()
				ginCtx, _ := gin.CreateTestContext(w)
				ginCtx.Request = httptest.NewRequest("POST", "/api/v1/user/register", bytes.NewBuffer(reqData))

				userHandler.registerUser(ginCtx)
				assert.Equal(t, ginCtx.Writer.Status(), http.StatusBadRequest)
				res := make(map[string]interface{})
				_ = json.Unmarshal(w.Body.Bytes(), &res)
				assert.Equal(t, res["code"], float64(40000))
				assert.Equal(t, res["description"], "user_name: String length must be greater than or equal to 1")
			})

			convey.Convey("invalid nick name", func() {
				req := dto.UserRegisterRequest{
					Username: "test",
					NickName: "",
					Password: "123456",
				}
				reqData, _ := json.Marshal(req)

				w := httptest.NewRecorder()
				ginCtx, _ := gin.CreateTestContext(w)
				ginCtx.Request = httptest.NewRequest("POST", "/api/v1/user/register", bytes.NewBuffer(reqData))

				userHandler.registerUser(ginCtx)
				assert.Equal(t, ginCtx.Writer.Status(), http.StatusBadRequest)
				res := make(map[string]interface{})
				_ = json.Unmarshal(w.Body.Bytes(), &res)
				assert.Equal(t, res["code"], float64(40000))
				assert.Equal(t, res["description"], "nick_name: String length must be greater than or equal to 1")
			})

			convey.Convey("invalid password", func() {
				req := dto.UserRegisterRequest{
					Username: "test",
					NickName: "test",
					Password: "",
				}
				reqData, _ := json.Marshal(req)

				w := httptest.NewRecorder()
				ginCtx, _ := gin.CreateTestContext(w)
				ginCtx.Request = httptest.NewRequest("POST", "/api/v1/user/register", bytes.NewBuffer(reqData))

				userHandler.registerUser(ginCtx)
				assert.Equal(t, ginCtx.Writer.Status(), http.StatusBadRequest)
				res := make(map[string]interface{})
				_ = json.Unmarshal(w.Body.Bytes(), &res)
				assert.Equal(t, res["code"], float64(40000))
				assert.Equal(t, res["description"], "password: String length must be greater than or equal to 6")
			})
		})

		convey.Convey("logic error", func() {
			req := dto.UserRegisterRequest{
				Username: "test",
				NickName: "test",
				Password: "123456",
			}
			reqData, _ := json.Marshal(req)

			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = httptest.NewRequest("POST", "/api/v1/user/register", bytes.NewBuffer(reqData))

			mockUserLogic.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("logic error"))

			userHandler.registerUser(ginCtx)
			assert.Equal(t, http.StatusInternalServerError, ginCtx.Writer.Status())
			res := make(map[string]interface{})
			_ = json.Unmarshal(w.Body.Bytes(), &res)
			assert.Equal(t, res["code"], float64(50000))
			assert.Equal(t, res["description"], "logic error")
		})

		convey.Convey("success", func() {
			req := dto.UserRegisterRequest{
				Username: "test",
				NickName: "test",
				Password: "123456",
			}
			reqData, _ := json.Marshal(req)

			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = httptest.NewRequest("POST", "/api/v1/user/register", bytes.NewBuffer(reqData))

			mockUserLogic.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(&interfaces.UserRegisterResponse{
				UserID: 1,
			}, nil)

			userHandler.registerUser(ginCtx)
			assert.Equal(t, http.StatusOK, ginCtx.Writer.Status())
			res := make(map[string]interface{})
			_ = json.Unmarshal(w.Body.Bytes(), &res)
			assert.Equal(t, res["user_id"], float64(1))
			t.Log("body: ", w.Body.String())
		})
	})
}
