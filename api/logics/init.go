package logics

import (
	"github.com/xmcy0011/CoffeeChat/api/interfaces"
)

// 逻辑层全局依赖
var (
	dbUser interfaces.DBUser
)

// SetDBUser 注入 dbUser 依赖
func SetDBUser(dbUser interfaces.DBUser) {
	dbUser = dbUser
}
