package dbaccess

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xmcy0011/CoffeeChat/api/interfaces"
	"github.com/xmcy0011/CoffeeChat/pkg/logger"
)

type dbUser struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewDBUser(sql *sql.DB) interfaces.DBUser {
	return &dbUser{db: sql, logger: logger.GetLogger()}
}

func (u *dbUser) Add(ctx context.Context, userName, nickName, pwd string) (id int64, err error) {
	var tx *sql.Tx
	if tx, err = u.db.Begin(); err != nil {
		u.logger.Warn("Begin error:", err.Error())
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// check exist
	count := int64(0)
	sql := "select count(1) from im_user where user_name=?"
	if err = tx.QueryRowContext(ctx, sql, userName).Scan(&count); err != nil {
		u.logger.Warnf("QueryRow userName: %s, error: %s", userName, err.Error())
		return 0, err
	}

	if count > 0 {
		u.logger.Warn("user already exist, userName:", userName)
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
	pwdHash := getPwdHash(pwd, salt)

	// insert
	sql = "insert into im_user(user_name, pwd_salt, pwd_hash, nick_name, created, updated) values(?,?,?,?,?,?)"
	r, err := u.db.ExecContext(ctx, sql, userName, salt, pwdHash, nickName, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		u.logger.Warn("Exec error:", err.Error())
		return 0, err
	}
	userId, err := r.LastInsertId()
	if err != nil {
		u.logger.Warnf("RowsAffected error:%s ", err.Error())
		return 0, err
	}
	return userId, nil
}

// 返回一个32位md5加密后的字符串
func getPwdHash(pwd, salt string) string {
	// md5(md5(pwd)+salt)

	h := md5.New()
	h.Write([]byte(pwd))
	tempHash := hex.EncodeToString(h.Sum(nil))

	h.Reset()
	h.Write([]byte(tempHash + salt))
	tempHash = hex.EncodeToString(h.Sum(nil))

	return tempHash
}

func (u *dbUser) QueryByID(ctx context.Context, userId int64) (bool, *interfaces.UserModel, error) {
	userInfo := &interfaces.UserModel{}
	sqlStr := "select id,nick_name,user_name,created,updated from im_user where id = ?"
	err := u.db.QueryRowContext(ctx, sqlStr, userId).Scan(&userInfo.ID, &userInfo.NickName, &userInfo.UserName, &userInfo.Created, &userInfo.Updated)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil, nil
		}
		return false, nil, err
	}
	return true, userInfo, nil
}

func (u *dbUser) QueryByUserName(ctx context.Context, userName string) (bool, *interfaces.UserModel, error) {
	sqlStr := "select id,nick_name,user_name,created,updated from im_user where user_name = ?"
	userInfo := &interfaces.UserModel{}
	err := u.db.QueryRowContext(ctx, sqlStr, userName).Scan(&userInfo.ID, &userInfo.NickName, &userInfo.UserName, &userInfo.Created, &userInfo.Updated)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil, nil
		}
		return false, nil, err
	}
	return true, userInfo, nil
}

func (u *dbUser) BatchQueryByIDs(ctx context.Context, ids []int64) ([]*interfaces.UserModel, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	placeholders := make([]string, 0)
	args := make([]interface{}, 0)
	for _, v := range ids {
		placeholders = append(placeholders, "?")
		args = append(args, v)
	}

	sql := fmt.Sprintf("select id,user_nick_name,created,updated from im_user where id in(%s)", strings.Join(placeholders, ","))
	rows, err := u.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	users := make([]*interfaces.UserModel, 0)
	for rows.Next() {
		user := &interfaces.UserModel{}
		err := rows.Scan(&user.ID, &user.NickName, &user.Created, &user.Updated)
		if err != nil {
			u.logger.Warn("Scan error:", err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *dbUser) Update(ctx context.Context, userName, userNickName, userPwd string) (int64, error) {
	sqlStr := "update im_user set nick_name = ?, user_pwd = ? where user_name = ?"
	result, err := u.db.ExecContext(ctx, sqlStr, userNickName, userPwd, userName)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
