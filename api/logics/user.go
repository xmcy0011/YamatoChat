package logics

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/xmcy0011/CoffeeChat/api/interfaces"
	"github.com/xmcy0011/CoffeeChat/pkg/logger"
	"github.com/xmcy0011/CoffeeChat/pkg/rest"
)

type userLogic struct {
	dbUser interfaces.DBUser
	logger *logrus.Logger
}

// NewUserLogic 创建实例
func NewUserLogic() interfaces.UserLogic {
	return &userLogic{dbUser: dbUser, logger: logger.GetLogger()}
}

func (l *userLogic) RegisterUser(ctx context.Context, request interfaces.UserRegisterRequest) (*interfaces.UserRegisterResponse, error) {
	ok, _, err := l.dbUser.QueryByUserName(ctx, request.UserName)
	if err != nil {
		l.logger.Error("GetUserByName", err)
		return nil, err
	}

	if ok {
		return nil, rest.NewHTTPError(rest.BadRequest, "user already exists")
	}

	userID, err := l.dbUser.Add(ctx, request.UserName, request.UserNickName, request.UserPassword)
	if err != nil {
		l.logger.Error("AddUser", err)
		return nil, err
	}

	return &interfaces.UserRegisterResponse{UserID: userID}, nil
}
