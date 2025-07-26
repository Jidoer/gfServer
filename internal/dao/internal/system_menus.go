// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemMenusDao is the data access object for the table SystemMenus.
type SystemMenusDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SystemMenusColumns // columns contains all the column names of Table for convenient usage.
}

// SystemMenusColumns defines and stores column names for the table SystemMenus.
type SystemMenusColumns struct {
	Id          string // 主键 ID
	ParentId    string // 父级路由 ID，0 表示顶级路由
	Path        string // 路由路径
	Name        string // 路由名称
	Component   string // 组件路径
	Redirect    string // 重定向路径
	Title       string // 标题
	IsLink      string // 外链地址
	IsHide      string // 是否隐藏
	IsKeepAlive string // 是否缓存
	IsAffix     string // 是否固定标签
	IsIframe    string // 是否嵌套 iframe
	Roles       string // 允许访问的角色，JSON 存储
	Icon        string // 图标
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// systemMenusColumns holds the columns for the table SystemMenus.
var systemMenusColumns = SystemMenusColumns{
	Id:          "id",
	ParentId:    "parent_id",
	Path:        "path",
	Name:        "name",
	Component:   "component",
	Redirect:    "redirect",
	Title:       "title",
	IsLink:      "is_link",
	IsHide:      "is_hide",
	IsKeepAlive: "is_keep_alive",
	IsAffix:     "is_affix",
	IsIframe:    "is_iframe",
	Roles:       "roles",
	Icon:        "icon",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewSystemMenusDao creates and returns a new DAO object for table data access.
func NewSystemMenusDao() *SystemMenusDao {
	return &SystemMenusDao{
		group:   "default",
		table:   "SystemMenus",
		columns: systemMenusColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemMenusDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemMenusDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemMenusDao) Columns() SystemMenusColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemMenusDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemMenusDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SystemMenusDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
