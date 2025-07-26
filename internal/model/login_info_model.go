package model

import "time"

type LoginInfo_db struct {
	Id       uint   `json:"id"`
	Passport string `json:"passport"`
	Ip       string `json:"ip"`
	Device   string `json:"device"`
	// Agent    string `json:"agent"`
	Status   int    `json:"status"`
	Token    string `json:"token"` //?!
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

// ip == ip update || ip != ip create db
type LoginInfo_last struct {
	Id       uint   `json:"id"`
	Device   string `json:"device"`
	Passport string `json:"passport"`
	Ip       string `json:"ip"`
	// Agent    string `json:"agent"`
	Token     string `json:"token"`
	LoginTime time.Time `json:"login_time"`
}
