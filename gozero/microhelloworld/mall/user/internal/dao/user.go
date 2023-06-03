package dao

import (
	"context"
	"fmt"
	"user/database"
	"user/internal/model"
)

type UserDao struct {
	conn *database.DBConn
}

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{
		conn: conn,
	}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {
	sql := fmt.Sprintf("insert into %s (name, gender) values (?, ?)", user.TableName(), user.Name, user.Gender)
	d.conn.Conn.ExecCtx(ctx, sql)
	return nil
}
