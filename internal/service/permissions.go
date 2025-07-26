// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gfAdmin/internal/model"
	"gfAdmin/internal/model/do"
)

type (
	IPermissions interface {
		GetRolesList(ctx context.Context, page int, limit int) (*[]model.Roles, int, error)
		RolesDelete(ctx context.Context, role model.Roles) error
		RoleCreate(ctx context.Context, role do.Roles) (int64, error)
		RoleUpdate(ctx context.Context, role do.Roles) error
		// 更新关系表
		RoleUpdateRelation(ctx context.Context, role do.Roles, role_permissions []do.RolePermissions) error
		GetPermissionsList(ctx context.Context, page int, limit int) (*[]model.Permissions, int, error)
		AddPermission(ctx context.Context, permission model.Permissions) error
		DeletePermission(ctx context.Context, permission model.Permissions) error
		UpdatePermission(ctx context.Context, permission model.Permissions) error
		GetRoleAndPermissions(ctx context.Context, role_id int) (string, *[]string, error)
	}
)

var (
	localPermissions IPermissions
)

func Permissions() IPermissions {
	if localPermissions == nil {
		panic("implement not found for interface IPermissions, forgot register?")
	}
	return localPermissions
}

func RegisterPermissions(i IPermissions) {
	localPermissions = i
}
