package handler

import (
	"api/pkg/util"
	"api/pkg/util/response"
	"api/user/entity"
	"api/user/pb"
	"api/user/repo"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserHandler struct {
	pb.UnimplementedUserServer
}

//此server 不认证。
func (u UserHandler) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	fmt.Println(fullMethodName)
	return ctx, nil
}

func (u UserHandler) CreateVerifyCode(ctx context.Context, req *pb.CreateVerifyCodeReq) (*pb.CreateVerifyCodeRsp, error) {
	ue := entity.UserAuthEntity{
		Ctx: ctx,
	}
	ue.UserAuth = repo.UserAuth{Phone: req.PhoneNum}
	err := ue.CreateVrifCode()
	if err != nil {
		return nil, response.ErrorHandle(err)
	}
	return &pb.CreateVerifyCodeRsp{}, nil
}

func (u UserHandler) LoginByPhone(ctx context.Context, req *pb.LoginByPhoneReq) (*pb.TokenResp, error) {
	//check 验证码
	uae := entity.UserAuthEntity{
		Ctx: ctx,
	}
	uae.Phone = req.PhoneNum
	uae.VrifCode = req.VerifyCode
	if ok, err := uae.AuthVrifCode(); err != nil {
		return nil, response.ErrorHandle(err)
	} else if !ok {
		return nil, status.Error(codes.InvalidArgument, "验证码错误")
	}

	//检查用户是否存在
	user := &repo.User{
		Phone: req.PhoneNum,
	}
	userentity := entity.UserEntity{
		User: user,
	}
	err := userentity.GetUserByPhone()

	//未注册自动注册
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userentity.UserId = util.GetUniqueIdString()
		userentity.UserName = req.UserName
		err = userentity.Register()
		if err != nil {
			return nil, response.ErrorHandle(err)
		}
	} else {
		return nil, response.ErrorHandle(err)
	}

	//生成token
	token, err := uae.GenToken(userentity.UserId)
	if err != nil {
		return nil, response.ErrorHandle(err)
	}
	return &pb.TokenResp{
		Token:    token,
		PhoneNum: userentity.Phone,
		UserName: userentity.UserName,
	}, nil
}
