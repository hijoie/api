package entity

import (
	"api/pkg/bizerr"
	"api/pkg/util/response"
	"api/user/config"
	"api/user/dao"
	"api/user/repo"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"strconv"
	"time"
)

type UserAuthEntity struct {
	repo.UserAuth
	Ctx context.Context
}

func (s *UserAuthEntity) genVrifCode() string {
	return strconv.Itoa(rand.Intn(8999) + 1000)
}

func (s *UserAuthEntity) CreateVrifCode() error {
	//TODO 发送短信
	if s.Phone == "" {
		return errors.New("create verifycode error phone is empty")
	}
	s.VrifCode = s.genVrifCode()
	err := dao.GetUserAuthDao(s.Ctx).SaveVrifCode(&s.UserAuth)
	return err
}

func (s *UserAuthEntity) AuthVrifCode() (ok bool, err error) {
	vricode, err := dao.GetUserAuthDao(s.Ctx).GetVrifCode(s.Phone)
	if err != nil {
		return
	}
	if s.VrifCode != vricode.VrifCode {
		return false, bizerr.Error{Code: response.VrifiCodeError, Msg: "验证码错误"}
	}
	return true, nil
}

func (s *UserAuthEntity) GenToken(userid string) (tokenStr string, err error) {
	claims := make(jwt.MapClaims)
	now := time.Now().Unix()
	claims["exp"] = now + int64(config.GetInt("jwt.access_expire"))
	claims["iat"] = now
	claims["userId"] = userid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(config.GetString("jwt.access_secret")))
}
