// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PrintServerDao is the data access object for the table PrintServer.
type PrintServerDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns PrintServerColumns // columns contains all the column names of Table for convenient usage.
}

// PrintServerColumns defines and stores column names for the table PrintServer.
type PrintServerColumns struct {
	Id             string //
	Passport       string // 唯一登录秘钥
	Token          string // 登录秘钥 可以设定为序列号 可以相同
	Name           string // 客户端名称
	Type           string // 客户端类型
	LocationType   string // 位置类型
	Location       string // 地址信息
	UsbProduct     string // 连接的设备
	Balance        string // 余额
	WithdrawnMoney string // 已提现的金额
	IsOnline       string // 是否在线
	Ban            string // 是否禁止登录
	ExpirationTime string // 到期时间
}

// printServerColumns holds the columns for the table PrintServer.
var printServerColumns = PrintServerColumns{
	Id:             "id",
	Passport:       "passport",
	Token:          "token",
	Name:           "name",
	Type:           "type",
	LocationType:   "location_type",
	Location:       "location",
	UsbProduct:     "usb_product",
	Balance:        "balance",
	WithdrawnMoney: "withdrawn_money",
	IsOnline:       "is_online",
	Ban:            "ban",
	ExpirationTime: "expiration_time",
}

// NewPrintServerDao creates and returns a new DAO object for table data access.
func NewPrintServerDao() *PrintServerDao {
	return &PrintServerDao{
		group:   "default",
		table:   "PrintServer",
		columns: printServerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PrintServerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PrintServerDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PrintServerDao) Columns() PrintServerColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PrintServerDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PrintServerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PrintServerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
