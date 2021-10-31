package dao

import (
	"api/user/global"
	"api/user/repo"
	"api/user/repo/cachekey"
	"context"
	"github.com/go-redis/cache/v8"
	"time"
)

type UserAuthDao struct {
	Ctx context.Context
}

func GetUserAuthDao(ctx context.Context) *UserAuthDao {
	return &UserAuthDao{
		Ctx: ctx,
	}
}

func (u UserAuthDao) SaveVrifCode(auth *repo.UserAuth) error {
	return global.GetCache().Set(&cache.Item{
		Key:   cachekey.GetCacheKey(cachekey.UserVeriCode, auth.Phone),
		Value: auth,
		TTL:   30 * time.Minute,
		Ctx:   u.Ctx,
	})
}

func (u UserAuthDao) GetVrifCode(phone string) (auth *repo.UserAuth, err error) {
	auth = &repo.UserAuth{}
	err = global.GetCache().Get(u.Ctx, cachekey.GetCacheKey(cachekey.UserVeriCode, phone), auth)
	return auth, err
}
