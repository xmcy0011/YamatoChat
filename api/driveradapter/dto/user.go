package dto

// UserRegisterRequest 用户注册请求
type UserRegisterRequest struct {
	Username string `json:"user_name"`
	NickName string `json:"nick_name"`
	Password string `json:"password"`
}

// UserRegisterResponse 用户注册响应
type UserRegisterResponse struct {
	UserID int64 `json:"user_id"`
}
