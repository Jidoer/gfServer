// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table User for DAO operations like Where/Data.
type User struct {
	g.Meta   `orm:"table:User, do:true"`
	Id       interface{} // 用户ID
	Passport interface{} // 账号uid
	Password interface{} // 密码
	Nickname interface{} // 用户昵称
	Role     interface{} // 角色(0-普通用户 10管理员)
	Avatar   interface{} // 头像url
	Email    interface{} // 邮箱
	Phone    interface{} // 手机号
	Status   interface{} // 账户状态 (1: Normal, 2: Blocked)
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
}
