package admin

import (
	"context"

	"gfAdmin/api/admin/v1"
	"gfAdmin/internal/model"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) AddPermissions(ctx context.Context, req *v1.AddPermissionsReq) (res *v1.AddPermissionsRes, err error) {
	err = service.Permissions().AddPermission(ctx, model.Permissions{
		Name:   req.Name,
		Description: req.Description,
	})
	return
}
