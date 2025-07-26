// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RolesDao is the data access object for the table roles.
type RolesDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns RolesColumns // columns contains all the column names of Table for convenient usage.
}

// RolesColumns defines and stores column names for the table roles.
type RolesColumns struct {
	Id          string //
	RolesID     string // 给用户使用唯一ID
	Name        string // 名称
	RoleSign    string // 标识
	Rank        string // 权重
	Status      string // 状态码 0正常 1禁用
	Description string // 描述
	CreatedAt   string //
	UpdatedAt   string //
}

// rolesColumns holds the columns for the table roles.
var rolesColumns = RolesColumns{
	Id:          "id",
	RolesID:     "rolesID",
	Name:        "name",
	RoleSign:    "role_sign",
	Rank:        "rank",
	Status:      "status",
	Description: "description",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewRolesDao creates and returns a new DAO object for table data access.
func NewRolesDao() *RolesDao {
	return &RolesDao{
		group:   "default",
		table:   "roles",
		columns: rolesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RolesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RolesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RolesDao) Columns() RolesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RolesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RolesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RolesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
