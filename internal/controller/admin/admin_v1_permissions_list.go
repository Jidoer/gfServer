package admin

import (
	"context"

	"gfAdmin/api/admin/v1"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) PermissionsList(ctx context.Context, req *v1.PermissionsListReq) (res *v1.PermissionsListRes, err error) {
	// return nil, gerror.NewCode(gcode.CodeNotImplemented)
	r_, total, err := service.Permissions().GetPermissionsList(ctx,req.Page,req.Limit)
	if err != nil {
		return nil, err //!直接response了数据库错误 gerror.NewCode(gcode.CodeNotImplemented)
	}
	res = &v1.PermissionsListRes{
		List:  *r_,
		Total: total,
	}
	return res, nil
}
