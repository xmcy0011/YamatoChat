package interfaces

import "context"

// DBUser 用户信息
type DBUser struct {
	UserId       int64
	UserNickName string
	UserName     string
	UserPwdSalt  string
	UserPwdHash  string
	Created      int64
	Updated      int64
}

// UserRepository 用户数据库操作接口
type UserRepository interface {
	Add(ctx context.Context, userName, userNickName, userPwd string) (int64, error)
}
