// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BannedTokensDao is the data access object for the table BannedTokens.
type BannedTokensDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns BannedTokensColumns // columns contains all the column names of Table for convenient usage.
}

// BannedTokensColumns defines and stores column names for the table BannedTokens.
type BannedTokensColumns struct {
	Id         string //
	Token      string // 被阻止登录的token
	BannedTime string // 禁用时间
	Allowed    string // 是否已允许
	Remark     string // 备注
}

// bannedTokensColumns holds the columns for the table BannedTokens.
var bannedTokensColumns = BannedTokensColumns{
	Id:         "id",
	Token:      "token",
	BannedTime: "banned_time",
	Allowed:    "allowed",
	Remark:     "remark",
}

// NewBannedTokensDao creates and returns a new DAO object for table data access.
func NewBannedTokensDao() *BannedTokensDao {
	return &BannedTokensDao{
		group:   "default",
		table:   "BannedTokens",
		columns: bannedTokensColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BannedTokensDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BannedTokensDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BannedTokensDao) Columns() BannedTokensColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BannedTokensDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BannedTokensDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *BannedTokensDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
