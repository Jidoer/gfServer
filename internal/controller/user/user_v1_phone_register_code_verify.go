package user

import (
	"context"

	// "github.com/gogf/gf/v2/errors/gerror"

	"gfAdmin/api/user/v1"
	// "gfAdmin/internal/model"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) PhoneRegisterCodeVerify(ctx context.Context, req *v1.PhoneRegisterCodeVerifyReq) (res *v1.PhoneRegisterCodeVerifyRes, err error) {
	phone, err := service.Sms().VerifySmsCode(ctx, req.CodeID, req.Code)
	if err != nil {
		return nil, err //gerror.New("验证码错误")
	}
	//Create and login
	token, user, err := service.User().AutoCreate(ctx, phone)
	if err == nil {
		res = &v1.PhoneRegisterCodeVerifyRes{
			Token: token,
			User:  user,
		}
	}
	return
}
