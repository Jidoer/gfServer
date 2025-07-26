// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PrintServer is the golang structure for table PrintServer.
type PrintServer struct {
	Id             uint        `json:"id"             orm:"id"              description:""`
	Passport       string      `json:"passport"       orm:"passport"        description:"唯一登录秘钥"`
	Token          string      `json:"token"          orm:"token"           description:"登录秘钥 可以设定为序列号 可以相同"`
	Name           string      `json:"name"           orm:"name"            description:"客户端名称"`
	Type           int         `json:"type"           orm:"type"            description:"客户端类型"`
	LocationType   int         `json:"locationType"   orm:"location_type"   description:"位置类型"`
	Location       string      `json:"location"       orm:"location"        description:"地址信息"`
	UsbProduct     string      `json:"usbProduct"     orm:"usb_product"     description:"连接的设备"`
	Balance        float64     `json:"balance"        orm:"balance"         description:"余额"`
	WithdrawnMoney float64     `json:"withdrawnMoney" orm:"withdrawn_money" description:"已提现的金额"`
	IsOnline       int         `json:"isOnline"       orm:"is_online"       description:"是否在线"`
	Ban            int         `json:"ban"            orm:"ban"             description:"是否禁止登录"`
	ExpirationTime *gtime.Time `json:"expirationTime" orm:"expiration_time" description:"到期时间"`
}
