// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ABTokens is the golang structure of table ABTokens for DAO operations like Where/Data.
type ABTokens struct {
	g.Meta    `orm:"table:ABTokens, do:true"`
	Id        interface{} //
	Token     interface{} //
	AllowType interface{} // 允许类型
	Remark    interface{} // 备注
}
