package dao

import (
	"api/user/global"
	"api/user/repo"
	"context"
	"fmt"
)

type UserDao struct {
	ctx context.Context
}

func GetUserDao(ctx context.Context) *UserDao {
	return &UserDao{
		ctx: ctx,
	}
}

func (m *UserDao) FirstOrCreateByPhone(data *repo.User) error {
	return global.GetMyDB().WithContext(m.ctx).FirstOrCreate(data, "phone = ?", data.Phone).Error
}

func (m *UserDao) FindOne(id int64) (u *repo.User, err error) {
	err = global.GetMyDB().WithContext(m.ctx).First(u, id).Error
	return
}

func (m *UserDao) FindOneByPhone(phone string) (u *repo.User, err error) {
	u = &repo.User{}
	err = global.GetMyDB().WithContext(m.ctx).First(u, "phone = ?", phone).Error
	return
}

func (m *UserDao) Update(data *repo.User) error {
	if data.ID == 0 {
		return fmt.Errorf("update primary key is zero")
	}
	return global.GetMyDB().WithContext(m.ctx).Where("id = ?", data.ID).Updates(data).Error
}

func (m *UserDao) Delete(id int64) error {
	return global.GetMyDB().WithContext(m.ctx).Where("id = ?", id).Delete(&repo.User{}).Error
}
