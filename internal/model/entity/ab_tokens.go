// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ABTokens is the golang structure for table ABTokens.
type ABTokens struct {
	Id        uint   `json:"id"        orm:"id"         description:""`
	Token     string `json:"token"     orm:"token"      description:""`
	AllowType int    `json:"allowType" orm:"allow_type" description:"允许类型"`
	Remark    string `json:"remark"    orm:"remark"     description:"备注"`
}
