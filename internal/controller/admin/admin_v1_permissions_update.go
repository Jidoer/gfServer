package admin

import (
	"context"

	"gfAdmin/api/admin/v1"
	"gfAdmin/internal/model"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) PermissionsUpdate(ctx context.Context, req *v1.PermissionsUpdateReq) (res *v1.PermissionsUpdateRes, err error) {
	err = service.Permissions().UpdatePermission(ctx, model.Permissions{Id: req.Id, Name: req.Name, Description: req.Description})
	return
}
