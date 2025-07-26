package admin

import (
	"context"

	// "github.com/gogf/gf/v2/errors/gcode"
	// "github.com/gogf/gf/v2/errors/gerror"

	"gfAdmin/api/admin/v1"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	users, total, err := service.User().GetUsersList(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err //!直接response了数据库错误
	}
	res = &v1.UserListRes{
		Total: total,
		List:  *users,
	}
	return res, nil
}
