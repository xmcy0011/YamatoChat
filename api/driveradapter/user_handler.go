package driveradapter

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
	"github.com/xmcy0011/CoffeeChat/api/driveradapter/dto"
	"github.com/xmcy0011/CoffeeChat/api/interfaces"
	"github.com/xmcy0011/CoffeeChat/api/logics"
	"github.com/xmcy0011/CoffeeChat/pkg/rest"
)

// UserHandler 用户入站适配器
type UserHandler interface {
	RegisterPublicRouters(gin *gin.Engine)
}

var (
	//go:embed jsonschema/user/register.json
	userRegisterSchemaStr string
)

type userHandler struct {
	userLogic interfaces.UserLogic

	// schema
	userRegisterSchema *gojsonschema.Schema
}

// NewUserHandler 创建用户入站适配器
func NewUserHandler() *userHandler {
	return &userHandler{
		userLogic:          logics.NewUserLogic(),
		userRegisterSchema: MustNewSchema(userRegisterSchemaStr),
	}
}

func (h *userHandler) RegisterPublicRouters(gin *gin.Engine) {
	gin.POST("/v1/users", h.registerUser)
}

func (h *userHandler) registerUser(ctx *gin.Context) {
	var req dto.UserRegisterRequest
	if err := ValidateAndBindGin(ctx, h.userRegisterSchema, &req); err != nil {
		rest.ReplyError(ctx, err)
		return
	}

	resp, err := h.userLogic.RegisterUser(ctx, interfaces.UserRegisterRequest{
		UserName:     req.Username,
		UserNickName: req.NickName,
		UserPassword: req.Password,
	})
	if err != nil {
		rest.ReplyError(ctx, err)
		return
	}

	rest.ReplyOK(ctx, http.StatusOK, dto.UserRegisterResponse{UserID: resp.UserID})
}
