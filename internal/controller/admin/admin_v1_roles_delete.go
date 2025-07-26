package admin

import (
	"context"

	"gfAdmin/api/admin/v1"
	"gfAdmin/internal/model"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) RolesDelete(ctx context.Context, req *v1.RolesDeleteReq) (res *v1.RolesDeleteRes, err error) {
	err = service.Permissions().RolesDelete(ctx, model.Roles{Id: req.Id,RoleSign: req.RoleSign})
	return
}
