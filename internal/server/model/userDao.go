package model

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var MyUserDao *UserDao

type UserDao struct {
	rdb *redis.Client
}

func NewUserDao(rdb *redis.Client) *UserDao {
	userDao := &UserDao{
		rdb: rdb,
	}
	return userDao
}

func (userdao *UserDao) Login(userName string, userPwd string) (user *User, err error) {
	ctx := context.Background()
	val, err := userdao.rdb.Get(ctx, userName).Result()
	if err != nil {
		err = ErrUserNotFound
		return nil, err
	}
	if val != userPwd {
		err = ErrInvalidPassword
		return nil, err
	}
	return &User{
		Username: userName,
		Password: userPwd,
	}, nil
}

func (userdao *UserDao) Register(userName string, userPwd string) error {
	ctx := context.Background()
	// if userdao.rdb.Get(ctx, userName) != nil {
	// 	return ErrUserExists
	// }
	err := userdao.rdb.Set(ctx, userName, userPwd, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
