package user

import (
	"context"

	"gfAdmin/api/user/v1"
	"gfAdmin/internal/model"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	if err == nil {
		sid, _ := service.BizCtx().Get(ctx).Session.Id()
		res = &v1.LoginRes{
			Token: sid,
		}
		//set cookies for sessionid
		request := service.BizCtx().Get(ctx).Request
		request.Cookie.SetSessionId(sid) //set sessionid
	}
	return
}
