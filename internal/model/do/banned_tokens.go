// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BannedTokens is the golang structure of table BannedTokens for DAO operations like Where/Data.
type BannedTokens struct {
	g.Meta     `orm:"table:BannedTokens, do:true"`
	Id         interface{} //
	Token      interface{} // 被阻止登录的token
	BannedTime *gtime.Time // 禁用时间
	Allowed    interface{} // 是否已允许
	Remark     interface{} // 备注
}
