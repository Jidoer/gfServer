package model

import (
	// "gfAdmin/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Roles struct {
	g.Meta      `orm:"table:roles, do:true"`
	Id          int         `json:"id"          orm:"id"          description:""`
	RolesID     int         `json:"rolesID"     orm:"rolesID"     description:"给用户使用唯一ID"`
	Name        string      `json:"name"        orm:"name"        description:"名称"`
	RoleSign    string      `json:"roleSign"    orm:"role_sign"   description:"标识"`
	Rank        int         `json:"rank"        orm:"rank"        description:"权重"`
	Status      int         `json:"status"      orm:"status"      description:"状态码 0正常 1禁用"`
	Description string      `json:"description" orm:"description" description:"描述"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
	RolePermissions []RolePermissions `json:"role_permissions" orm:"with:role_id=id"`
}

type Permissions struct {
	g.Meta      `orm:"table:permissions, do:true"`
	Id          int    `json:"id"          orm:"id"          description:""`
	Name        string `json:"name"        orm:"name"        description:""`
	Description string `json:"description" orm:"description" description:""`
}

type RolePermissions struct {
	g.Meta       `orm:"table:role_permissions, do:true"`
	RoleId       int `json:"roleId"       orm:"role_id"       description:""`
	PermissionId int `json:"permissionId" orm:"permission_id" description:""`
	Permissions Permissions `json:"permissions" orm:"with:id=permission_id"`
}

