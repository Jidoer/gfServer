package admin

import (
	"context"

	// "github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gfAdmin/api/admin/v1"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) GetRolesList(ctx context.Context, req *v1.GetRolesListReq) (res *v1.GetRolesListRes, err error) {
	// return nil, gerror.NewCode(gcode.CodeNotImplemented)
	r, total, err := service.Permissions().GetRolesList(ctx,req.Page,req.Limit)
	if err != nil {
		return nil, err //!直接response了数据库错误 gerror.NewCode(gcode.CodeNotImplemented)
	}
	if(r == nil){
		return nil, gerror.New("r为空")
	}
	res = &v1.GetRolesListRes{
		List:  *r,
		Total: total,
	}
	return res, nil
}
