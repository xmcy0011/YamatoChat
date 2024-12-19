package interfaces

import "context"

//go:generate mockgen -source=logics.go -destination=./mocks/logics_mock.go -package=mocks

type AuthLogic interface {
}

// UserRegisterRequest 用户注册请求
type UserRegisterRequest struct {
	UserName     string
	UserNickName string
	UserPassword string
}

// UserRegisterResponse 用户注册响应
type UserRegisterResponse struct {
	UserID int64
}

// UserLogic 用户逻辑
type UserLogic interface {
	RegisterUser(ctx context.Context, request UserRegisterRequest) (*UserRegisterResponse, error)
}
