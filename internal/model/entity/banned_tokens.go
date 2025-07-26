// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BannedTokens is the golang structure for table BannedTokens.
type BannedTokens struct {
	Id         uint        `json:"id"         orm:"id"          description:""`
	Token      string      `json:"token"      orm:"token"       description:"被阻止登录的token"`
	BannedTime *gtime.Time `json:"bannedTime" orm:"banned_time" description:"禁用时间"`
	Allowed    int         `json:"allowed"    orm:"allowed"     description:"是否已允许"`
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`
}
