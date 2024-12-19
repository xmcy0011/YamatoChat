package dbaccess

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"github.com/xmcy0011/CoffeeChat/pkg/logger"
)

func TestUser(t *testing.T) {
	convey.Convey("TestUser", t, func() {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		logger := logger.GetLogger()

		user := &dbUser{db: db, logger: logger}
		convey.Convey("query error", func() {
			mock.ExpectBegin()
			mock.ExpectQuery("select count(1) from im_user where user_name=?").WithArgs("test").WillReturnError(errors.New("query error"))
			mock.ExpectRollback()

			_, err := user.Add(context.Background(), "test", "test", "test")
			assert.Error(t, err)
		})

		convey.Convey("user already exist", func() {
			mock.ExpectBegin()
			mock.ExpectQuery("select count(1) from im_user where user_name=?").WithArgs("test").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			mock.ExpectRollback()

			_, err = user.Add(context.Background(), "test", "test", "test")
			assert.Error(t, errors.New("user already exist"))
		})

		convey.Convey("insert error", func() {
			mock.ExpectBegin()
			mock.ExpectQuery("select count(1) from im_user where user_name=?").WithArgs("test").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
			mock.ExpectExec("insert into im_user(.*)").WithArgs("test", "test", "test").WillReturnError(errors.New("insert error"))
			mock.ExpectRollback()

			_, err = user.Add(context.Background(), "test", "test", "test")
			assert.Error(t, err)
		})

		convey.Convey("success", func() {
			mock.ExpectBegin()
			mock.ExpectQuery("^select").WithArgs("test").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
			mock.ExpectExec("^insert into im_user").WithArgs("test", sqlmock.AnyArg(), sqlmock.AnyArg(), "test", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			_, err := user.Add(context.Background(), "test", "test", "test")
			assert.NoError(t, err)
		})
	})
}

func TestUserQueryByID(t *testing.T) {
	convey.Convey("TestUserQueryByID", t, func() {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		logger := logger.GetLogger()

		user := &dbUser{db: db, logger: logger}
		ctx := context.Background()

		convey.Convey("query error", func() {
			mock.ExpectQuery("select").WithArgs(1).WillReturnError(errors.New("query error"))
			_, _, err := user.QueryByID(ctx, 1)
			assert.Equal(t, err, errors.New("query error"))
		})

		convey.Convey("no row", func() {
			mock.ExpectQuery("select").WithArgs(1).WillReturnError(sql.ErrNoRows)
			ok, _, err := user.QueryByID(ctx, 1)
			assert.False(t, ok)
			assert.NoError(t, err)
		})

		convey.Convey("success", func() {
			now := time.Now()
			mock.ExpectQuery("select").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "nick_name", "user_name", "created", "updated"}).
				AddRow(1, "test", "test", now.Unix(), now.Unix()))
			ok, userInfo, err := user.QueryByID(ctx, 1)
			assert.True(t, ok)
			assert.Equal(t, userInfo.ID, int64(1))
			assert.Equal(t, userInfo.NickName, "test")
			assert.Equal(t, userInfo.UserName, "test")
			assert.Equal(t, userInfo.Created, now.Unix())
			assert.Equal(t, userInfo.Updated, now.Unix())
			assert.NoError(t, err)
		})
	})
}

func TestUserBatchQueryByIDs(t *testing.T) {
	convey.Convey("TestUserBatchQueryByIDs", t, func() {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		logger := logger.GetLogger()

		user := &dbUser{db: db, logger: logger}
		ctx := context.Background()

		convey.Convey("query error", func() {
			mock.ExpectQuery("select").WithArgs(1).WillReturnError(errors.New("query error"))
			_, err := user.BatchQueryByIDs(ctx, []int64{1})
			assert.Equal(t, err, errors.New("query error"))
		})

		convey.Convey("success", func() {
			now := time.Now()
			mock.ExpectQuery("select").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "nick_name", "created", "updated"}).
				AddRow(1, "test", now.Unix(), now.Unix()))
			users, err := user.BatchQueryByIDs(ctx, []int64{1})
			assert.NoError(t, err)
			assert.Equal(t, users[0].ID, int64(1))
			assert.Equal(t, users[0].NickName, "test")
			assert.Equal(t, users[0].Created, now.Unix())
			assert.Equal(t, users[0].Updated, now.Unix())
		})
	})
}

func TestUserQueryByUserName(t *testing.T) {
	convey.Convey("TestUserQueryByUserName", t, func() {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		logger := logger.GetLogger()

		user := &dbUser{db: db, logger: logger}
		ctx := context.Background()

		convey.Convey("query error", func() {
			mock.ExpectQuery("select").WithArgs("test").WillReturnError(errors.New("query error"))
			_, _, err := user.QueryByUserName(ctx, "test")
			assert.Equal(t, err, errors.New("query error"))
		})

		convey.Convey("no row", func() {
			mock.ExpectQuery("select").WithArgs("test").WillReturnError(sql.ErrNoRows)
			ok, _, err := user.QueryByUserName(ctx, "test")
			assert.False(t, ok)
			assert.NoError(t, err)
		})

		convey.Convey("success", func() {
			now := time.Now()
			mock.ExpectQuery("select").WithArgs("test").WillReturnRows(sqlmock.NewRows([]string{"id", "nick_name", "user_name", "created", "updated"}).
				AddRow(1, "test", "test", now.Unix(), now.Unix()))
			ok, userInfo, err := user.QueryByUserName(ctx, "test")
			assert.True(t, ok)
			assert.Equal(t, userInfo.ID, int64(1))
			assert.Equal(t, userInfo.NickName, "test")
			assert.Equal(t, userInfo.UserName, "test")
			assert.Equal(t, userInfo.Created, now.Unix())
			assert.Equal(t, userInfo.Updated, now.Unix())
			assert.NoError(t, err)
		})
	})
}
