package v1

import (
	// "gfAdmin/internal/model"

	"gfAdmin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// type FormatData struct {
// 	Code int         `json:"code"`
// 	Msg  string      `json:"msg"`
// }

type RegisterReq struct {
	g.Meta    `path:"/user/register" method:"post" tags:"UserService" summary:"注册一个新用户"`
	Passport  string `v:"required|length:6,16"`
	Password  string `v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Nickname  string
}
type RegisterRes struct{}

type CheckPassportReq struct {
	g.Meta   `path:"/user/check_passport" method:"post" tags:"UserService" summary:"检查账号是否已经注册"`
	Passport string `v:"required|length:6,16"`
}
type CheckPassportRes struct {
}

type LoginReq struct {
	g.Meta   `path:"/user/login" method:"post" tags:"UserService" summary:"用户登录"`
	Passport string `v:"required|length:6,16"`
	Password string `v:"required|length:6,16"`
}

// type LoginRes struct {
// 	Token string
// 	// Expire
// }
type LoginRes struct {
	Token string `json:"token"`
	//Expire int64  `json:"expire"`
}

type SignOutReq struct {
	g.Meta `path:"/user/signout" method:"post" tags:"UserService" summary:"用户登出"`
}
type SignOutRes struct{}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"UserService" summary:"获取用户信息"`
}

// type UserInfoRes model.User_role_db
	
type UserInfoRes model.User_Session

// type UserInfoRes struct {
// 	Id        uint     `json:"id"`
// 	Passport  string   `json:"passport"`
// 	Nickname  string   `json:"nickname"`
// 	Role      int      `json:"role"`
// 	Avatar    string   `json:"avatar"`
// 	Roles     []string `json:"roles"`
// 	Auths     []string `json:"auths"`
// 	Phone     string   `json:"phone"`
// 	Email     string   `json:"email"`
// 	Status    int      `json:"status"`
// 	CreatedAt string   `json:"created_at"`
// }


type UserUpdateReq struct {
	g.Meta   `path:"/user/update" method:"post" tags:"UserService" summary:"更新用户信息"`
	Id       uint   `v:"required|min:1"`
	Passport string `v:"required|length:6,16"`
	Nickname string
	Password string `v:"length:6,16"`
	// Password2 string `v:"length:6,16|same:Password"`
}
type UserUpdateRes struct{}



type UserChangePasswordReq struct {
	g.Meta   `path:"/user/changePassword" method:"post" tags:"UserService" summary:"修改密码"`
	Id       uint   `v:"required|min:1"`
	Password string `v:"required|length:6,16"`
	//Password2 string `
	//v:"required|length:6,16|same:Password"`
}
type UserChangePasswordRes struct{}

