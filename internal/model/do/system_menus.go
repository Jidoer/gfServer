// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemMenus is the golang structure of table SystemMenus for DAO operations like Where/Data.
type SystemMenus struct {
	g.Meta      `orm:"table:SystemMenus, do:true"`
	Id          interface{} // 主键 ID
	ParentId    interface{} // 父级路由 ID，0 表示顶级路由
	Path        interface{} // 路由路径
	Name        interface{} // 路由名称
	Component   interface{} // 组件路径
	Redirect    interface{} // 重定向路径
	Title       interface{} // 标题
	IsLink      interface{} // 外链地址
	IsHide      interface{} // 是否隐藏
	IsKeepAlive interface{} // 是否缓存
	IsAffix     interface{} // 是否固定标签
	IsIframe    interface{} // 是否嵌套 iframe
	Roles       interface{} // 允许访问的角色，JSON 存储
	Icon        interface{} // 图标
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
