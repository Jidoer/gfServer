package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	// "github.com/gogf/gf/v2/frame/g"

	"gfAdmin/api/user/v1"
	// "gfAdmin/internal/consts"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	// g.Log().Debug(ctx,"handle:session:",ctx.Value("session"))
	user := service.User().GetProfile(ctx)
	if user == nil {
		return nil, gerror.New("user not found")
	}
	res = (*v1.UserInfoRes)(user)

	// res = &v1.UserInfoRes{
	// 	Id:       user.Id,
	// 	Passport: user.Passport,
	// 	Nickname: user.Nickname,
	// 	Role: func() int {
	// 		if user.Role == 10 {
	// 			return 1 //防止从前端获取
	// 		}
	// 		return 0
	// 	}(),
	// 	Phone:     user.Phone,
	// 	Email:     user.Email,
	// 	Status:    int(user.Status),
	// 	CreatedAt: user.CreateAt.Format("2006-01-02 15:04:05"),
	// 	Roles: func() []string {
	// 		if user.Role == 10 {
	// 			return consts.Admin_roles
	// 		}
	// 		return consts.User_roles
	// 	}(),
	// 	Auths: func() []string {
	// 		if user.Role == 10 {
	// 			return consts.Admin_authBtnList
	// 		}
	// 		return consts.User_authBtnList
	// 	}(),
	// }
	return res, nil
}
