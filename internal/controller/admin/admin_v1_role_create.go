package admin

import (
	"context"

	"gfAdmin/api/admin/v1"
	"gfAdmin/internal/model/do"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) RoleCreate(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleCreateRes, err error) {

	id, err := service.Permissions().RoleCreate(ctx, do.Roles{
		RolesID:     req.RolesID,
		Name:        req.Name,
		RoleSign:    req.RoleSign,
		Rank:        req.Rank,
		Status:      req.Status,
		Description: req.Description,
	})
	if err != nil {
		return res, err
	}
	role_permissions := []do.RolePermissions{}
	for _, v := range req.Permissions {
		role_permissions = append(role_permissions, do.RolePermissions{
			RoleId:       int(id), //this is real role id
			PermissionId: v,
		})
	}
	err = service.Permissions().RoleUpdateRelation(ctx, do.Roles{Id: int(id)}, role_permissions)
	return res, err
}
