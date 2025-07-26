package admin

import (
	"context"

	// "github.com/gogf/gf/v2/errors/gcode"
	// "github.com/gogf/gf/v2/errors/gerror"

	"gfAdmin/api/admin/v1"
	"gfAdmin/internal/model"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) PermissionsDelete(ctx context.Context, req *v1.PermissionsDeleteReq) (res *v1.PermissionsDeleteRes, err error) {
	p := model.Permissions{Id: req.Id}
	err = service.Permissions().DeletePermission(ctx, p)
	return
}
