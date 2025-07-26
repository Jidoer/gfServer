// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemMenus is the golang structure for table SystemMenus.
type SystemMenus struct {
	Id          uint        `json:"id"          orm:"id"            description:"主键 ID"`
	ParentId    uint        `json:"parentId"    orm:"parent_id"     description:"父级路由 ID，0 表示顶级路由"`
	Path        string      `json:"path"        orm:"path"          description:"路由路径"`
	Name        string      `json:"name"        orm:"name"          description:"路由名称"`
	Component   string      `json:"component"   orm:"component"     description:"组件路径"`
	Redirect    string      `json:"redirect"    orm:"redirect"      description:"重定向路径"`
	Title       string      `json:"title"       orm:"title"         description:"标题"`
	IsLink      string      `json:"isLink"      orm:"is_link"       description:"外链地址"`
	IsHide      int         `json:"isHide"      orm:"is_hide"       description:"是否隐藏"`
	IsKeepAlive int         `json:"isKeepAlive" orm:"is_keep_alive" description:"是否缓存"`
	IsAffix     int         `json:"isAffix"     orm:"is_affix"      description:"是否固定标签"`
	IsIframe    int         `json:"isIframe"    orm:"is_iframe"     description:"是否嵌套 iframe"`
	Roles       string      `json:"roles"       orm:"roles"         description:"允许访问的角色，JSON 存储"`
	Icon        string      `json:"icon"        orm:"icon"          description:"图标"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:"更新时间"`
}
