// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ABTokensDao is the data access object for the table ABTokens.
type ABTokensDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns ABTokensColumns // columns contains all the column names of Table for convenient usage.
}

// ABTokensColumns defines and stores column names for the table ABTokens.
type ABTokensColumns struct {
	Id        string //
	Token     string //
	AllowType string // 允许类型
	Remark    string // 备注
}

// aBTokensColumns holds the columns for the table ABTokens.
var aBTokensColumns = ABTokensColumns{
	Id:        "id",
	Token:     "token",
	AllowType: "allow_type",
	Remark:    "remark",
}

// NewABTokensDao creates and returns a new DAO object for table data access.
func NewABTokensDao() *ABTokensDao {
	return &ABTokensDao{
		group:   "default",
		table:   "ABTokens",
		columns: aBTokensColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ABTokensDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ABTokensDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ABTokensDao) Columns() ABTokensColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ABTokensDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ABTokensDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ABTokensDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
