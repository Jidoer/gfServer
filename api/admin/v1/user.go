package v1

import (
	"gfAdmin/internal/model"
	"gfAdmin/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	g.Meta `path:"/admin/user/list" method:"get" tags:"AdminService" summary:"用户列表"`
	Page   int `v:"required|min:1"`
	Limit  int `v:"required|min:1"`
}

type UserListRes struct {
	List  []entity.User `json:"list"`
	Total int           `json:"total"`
}

type UserDeleteReq struct {
	g.Meta `path:"/admin/user/delete" method:"post" tags:"AdminService" summary:"删除用户"`
	Id     uint `v:"required|min:1"`
}
type UserDeleteRes struct{}

type GetRolesListReq struct {
	g.Meta `path:"/admin/permissions/getRolesList" method:"get" tags:"AdminService" summary:"获取用户角色列表"`
	Page   int `v:"required|min:1"`
	Limit  int `v:"required|min:1"`
}

//	type GetRolesListRes struct {
//		List []entity.Roles `json:"list"`
//		Total int `json:"total"`
//	}
type GetRolesListRes struct {
	List  []model.Roles `json:"list"`
	Total int           `json:"total"`
}

///admin/permissions/rolesDelete
type RolesDeleteReq struct {
	g.Meta `path:"/admin/permissions/rolesDelete" method:"post" tags:"AdminService" summary:"删除角色"`
	Id     int `v:"required|min:1"`
	RoleSign string `v:"required"`
}

type RolesDeleteRes struct{}

// export interface RoleCreateReq {
// 	rolesID: number;
// 	name: string;
// 	roleSign: string;
// 	rank: number;
// 	status:number;
// 	description: string;
// 	permissions : string[] //SignName
// }
type RoleCreateReq struct {
	g.Meta `path:"/admin/permissions/roleCreate" method:"post" tags:"AdminService" summary:"添加角色"`
	ID          int    `v:"required|min:0"`
	RolesID     int    `v:"required|min:0"`
	Name        string `v:"required|length:1,16"`
	RoleSign    string `v:"required"`
	Rank        int    `v:"required"`
	Status      int    `v:"required"`
	Description string `v:"required"`
	Permissions []int `v:"required"` //permissions id
}

type RoleCreateRes struct{}

type RoleUpdateReq struct {
	g.Meta `path:"/admin/permissions/roleUpdate" method:"post" tags:"AdminService" summary:"更新角色"`
	ID          int    `v:"required|min:0"`
	RolesID     int    `v:"required|min:0"`
	Name        string `v:"required|length:1,16"`
	RoleSign    string `v:"required"`
	Rank        int    `v:"required"`
	Status      int    `v:"required"`
	Description string `v:"required"`
	Permissions []int `v:"required"` //permissions id
}
type RoleUpdateRes struct{}



type PermissionsListReq struct {
	g.Meta `path:"/admin/permissions/getPermissionsList" method:"get" tags:"AdminService" summary:"获取平台权限列表"`
	Page   int `v:"required|min:1"`
	Limit  int `v:"required|min:1"`
}

type PermissionsListRes struct {
	List []model.Permissions `json:"list"`
	Total int `json:"total"`
}


// export interface Permission {
// 	id: number;
// 	name: string;
// 	description: string;
// }
type AddPermissionsReq struct {
	g.Meta `path:"/admin/permissions/addPermissions" method:"post" tags:"AdminService" summary:"添加平台权限"`
	Name   string `v:"required|length:1,16"`
	Description string `v:"required"`
}

type AddPermissionsRes struct {
	Id int `json:"id"`
}

type PermissionsDeleteReq struct {
	g.Meta `path:"/admin/permissions/permissionsDelete" method:"post" tags:"AdminService" summary:"删除平台权限"`
	Id     int `v:"required|min:1"`
}

type  PermissionsDeleteRes struct{}

type PermissionsUpdateReq struct {
	g.Meta `path:"/admin/permissions/permissionsUpdate" method:"post" tags:"AdminService" summary:"更新平台权限"`
	Id     int `v:"required|min:1"`
	Name   string `v:"required|length:1,16"`
	Description string `v:"required"`
}

type PermissionsUpdateRes struct{}


