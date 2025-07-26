// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"gfAdmin/api/admin/v1"
)

type IAdminV1 interface {
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	UserDelete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error)
	GetRolesList(ctx context.Context, req *v1.GetRolesListReq) (res *v1.GetRolesListRes, err error)
	RolesDelete(ctx context.Context, req *v1.RolesDeleteReq) (res *v1.RolesDeleteRes, err error)
	RoleCreate(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleCreateRes, err error)
	RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error)
	PermissionsList(ctx context.Context, req *v1.PermissionsListReq) (res *v1.PermissionsListRes, err error)
	AddPermissions(ctx context.Context, req *v1.AddPermissionsReq) (res *v1.AddPermissionsRes, err error)
	PermissionsDelete(ctx context.Context, req *v1.PermissionsDeleteReq) (res *v1.PermissionsDeleteRes, err error)
	PermissionsUpdate(ctx context.Context, req *v1.PermissionsUpdateReq) (res *v1.PermissionsUpdateRes, err error)
}
