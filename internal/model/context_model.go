package model

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gsession"
)

type Context struct {
	Session *ghttp.Session // Session in context.
	User    *ContextUser   // User in context.
	Request *ghttp.Request // Request in context.(include session...)
}

type ContextUser struct {
	Id       uint   // User ID.
	Passport string // User passport.
	Nickname string // User nickname.
	Role     int    // User role from database.
	Avatar   string // User avatar.
	Email    string // User email.
	Phone    string // User phone.
	Status   uint   // User status.
	//more
	RoleName string
	Auths    []string `json:"auths"` //permissions from role
}
// type ContextPrintServer struct {
// 	Session *PrintServer_Session
// }

type MySessionManager struct {
	Manager *gsession.Manager
	Ctx     *context.Context
}

// type User struct {
// 	Id       uint        `json:"id"       orm:"id"        description:"用户ID"`
// 	Passport string      `json:"passport" orm:"passport"  description:"账号uid"`
// 	Password string      `json:"password" orm:"password"  description:"密码"`
// 	Nickname string      `json:"nickname" orm:"nickname"  description:"用户昵称"`
// 	Role     int         `json:"role"     orm:"role"      description:"角色(0-普通用户 10管理员)"`
// 	Avatar   string      `json:"avatar"   orm:"avatar"    description:"头像url"`
// 	Email    string      `json:"email"    orm:"email"     description:"邮箱"`
// 	Phone    string      `json:"phone"    orm:"phone"     description:"手机号"`
// 	Status   uint        `json:"status"   orm:"status"    description:"账户状态 (1: Normal, 2: Blocked)"`
// 	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:""`
// 	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:""`
// }
