// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RolePermissions is the golang structure for table role_permissions.
type RolePermissions struct {
	RoleId       int `json:"roleId"       orm:"role_id"       description:""`
	PermissionId int `json:"permissionId" orm:"permission_id" description:""`
}
