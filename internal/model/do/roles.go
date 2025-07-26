// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure of table roles for DAO operations like Where/Data.
type Roles struct {
	g.Meta      `orm:"table:roles, do:true"`
	Id          interface{} //
	RolesID     interface{} // 给用户使用唯一ID
	Name        interface{} // 名称
	RoleSign    interface{} // 标识
	Rank        interface{} // 权重
	Status      interface{} // 状态码 0正常 1禁用
	Description interface{} // 描述
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
