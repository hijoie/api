package entity

import (
	"api/user/dao"
	"api/user/repo"
	"context"
	"errors"
	"fmt"
)

type UserEntity struct {
	*repo.User
	ctx context.Context
}

//注册
func (u *UserEntity) Register() error {
	if u.Phone == "" {
		return errors.New("手机号错误")
	}
	err := dao.GetUserDao(u.ctx).FirstOrCreateByPhone(u.User)
	if err != nil {
		return fmt.Errorf("FirstOrCreate user error : %w", err)
	}
	return nil
}

//获取用户信息
func (u *UserEntity) GetUserByPhone() error {
	user, err := dao.GetUserDao(u.ctx).FindOneByPhone(u.Phone)
	if err != nil {
		return fmt.Errorf("get user error %w", err)
	}
	u.User = user
	return nil
}

// 修改用户信息
func (u *UserEntity) UpdateUserInfo() error {
	if u.ID == 0 {
		return fmt.Errorf("update user info error : userid is 0")
	}
	err := dao.GetUserDao(u.ctx).Update(u.User)
	if err != nil {
		return fmt.Errorf("update user error %w", err)
	}
	return nil
}
