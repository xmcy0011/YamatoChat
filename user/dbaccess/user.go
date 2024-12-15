package dbaccess

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/xmcy0011/CoffeeChat/user/interfaces"
	"go.uber.org/zap"
)

type userRepository struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewUserRepository(sql *sql.DB) interfaces.UserRepository {
	return &userRepository{db: sql, logger: u.logger}
}

func (u *userRepository) Add(ctx context.Context, userName, userNickName, userPwd string) (int64, error) {
	// check exist
	sql := fmt.Sprintf("select count(1) from im_user where user_name=?", userName)
	row := u.db.QueryRow(sql)

	count := int64(0)
	if err := row.Scan(&count); err != nil {
		u.logger.Warn("QueryRow", zap.String("userName", userName), zap.Error(err))
		return 0, err
	} else if count > 0 {
		u.logger.Warn("user already exist", zap.String("user_name", userName))
		return 0, errors.New("user already exist")
	}

	// build 32 bytes salt
	saltArr := make([]byte, 32)
	if _, err := rand.Reader.Read(saltArr); err != nil {
		u.logger.Warn("build random salt error:", err.Error())
		return 0, err
	}
	salt := hex.EncodeToString(saltArr)

	// calc pwdHash,md5(md5(pwd)+salt)
	pwdHash := GetPwdHash(userPwd, salt)

	u.logger.Infof("userName:%s,userPwdSalt:%s,userPwdHash:%s", userName, salt, pwdHash)

	// insert
	sql = fmt.Sprintf("insert into im_user(user_name,user_pwd_salt,user_pwd_hash,user_nick_name,"+
		"user_token,user_attach,created,updated) values('%s','%s','%s','%s','','',%d,%d)",
		userName, salt, pwdHash, userNickName, time.Now().Unix(), time.Now().Unix())
	r, err := u.db.Exec(sql)
	if err != nil {
		u.logger.Warn("Exec error:", err.Error())
		return 0, err
	}
	userId, err := r.LastInsertId()
	if err != nil {
		u.logger.Warnf("RowsAffected error:%s ", err.Error())
		return 0, err
	} else {
		if userId > 0 {
			return userId, nil
		} else {
			u.logger.Warn("unknown error, no effect row")
			return 0, errors.New("unknown error, no effect row")
		}
	}
}
