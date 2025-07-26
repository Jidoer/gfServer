// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure for table roles.
type Roles struct {
	Id          int         `json:"id"          orm:"id"          description:""`
	RolesID     int         `json:"rolesID"     orm:"rolesID"     description:"给用户使用唯一ID"`
	Name        string      `json:"name"        orm:"name"        description:"名称"`
	RoleSign    string      `json:"roleSign"    orm:"role_sign"   description:"标识"`
	Rank        int         `json:"rank"        orm:"rank"        description:"权重"`
	Status      int         `json:"status"      orm:"status"      description:"状态码 0正常 1禁用"`
	Description string      `json:"description" orm:"description" description:"描述"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
}
