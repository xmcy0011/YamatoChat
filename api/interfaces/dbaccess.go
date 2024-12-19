package interfaces

import "context"

// UserModel 用户信息
type UserModel struct {
	// 用户ID
	ID int64
	// 用户昵称
	NickName string
	// 用户名
	UserName string
	// 密码盐
	PwdSalt string
	// 密码哈希
	PwdHash string
	// 创建时间
	Created int64
	// 更新时间
	Updated int64
}

// DBUser 用户数据库操作接口
type DBUser interface {
	// Add 添加用户
	Add(ctx context.Context, userName, nickName, pwd string) (int64, error)
	// QueryByID 根据用户ID查询用户
	QueryByID(ctx context.Context, userId int64) (bool, *UserModel, error)
	// QueryByUserName 根据用户名查询用户
	QueryByUserName(ctx context.Context, userName string) (bool, *UserModel, error)
	// BatchQueryByIDs 根据用户ID批量查询用户
	BatchQueryByIDs(ctx context.Context, ids []int64) ([]*UserModel, error)
}
