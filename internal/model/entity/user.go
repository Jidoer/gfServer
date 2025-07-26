// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table User.
type User struct {
	Id       uint        `json:"id"       orm:"id"        description:"用户ID"`
	Passport string      `json:"passport" orm:"passport"  description:"账号uid"`
	Password string      `json:"password" orm:"password"  description:"密码"`
	Nickname string      `json:"nickname" orm:"nickname"  description:"用户昵称"`
	Role     int         `json:"role"     orm:"role"      description:"角色(0-普通用户 10管理员)"`
	Avatar   string      `json:"avatar"   orm:"avatar"    description:"头像url"`
	Email    string      `json:"email"    orm:"email"     description:"邮箱"`
	Phone    string      `json:"phone"    orm:"phone"     description:"手机号"`
	Status   uint        `json:"status"   orm:"status"    description:"账户状态 (1: Normal, 2: Blocked)"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:""`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:""`
}
