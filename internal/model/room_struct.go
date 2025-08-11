package model

import (
	"time"

)

type AllocationReq struct {
	UUID          string            `json:"uuid"`
	ProfileID     string            `json:"profileId"`
	RegionID      string            `json:"regionId"`
	Envs          map[string]string `json:"envs"`
	AllocationTTL string            `json:"allocationTTL"`
}

type AllocationRes struct {
	Allocation struct {
		UUID            string `json:"uuid"`
		GameID          string `json:"gameId"`
		ProfileID       string `json:"profileId"`
		RegionID        string `json:"regionId"`
		IP              string `json:"ip"`
		GameServerPorts []struct {
			Port     int    `json:"port"`
			Protocol string `json:"protocol"`
			Name     string `json:"name"`
		} `json:"gameServerPorts"`
		GameServerName string    `json:"gameServerName"`
		CreatedAt      time.Time `json:"createdAt"`
		CreatedByUser  string    `json:"createdByUser"`
		ModifiedAt     time.Time `json:"modifiedAt"`
		ModifiedByUser string    `json:"modifiedByUser"`
		FulfilledAt    time.Time `json:"fulfilledAt"`
		DeletedAt      time.Time `json:"deletedAt"`
		DeletedByUser  string    `json:"deletedByUser"`
		Status         string    `json:"status"`
		Msg            string    `json:"msg"`
		AllocationTTL  string    `json:"allocationTTL"`
		ProfileName    string    `json:"profileName"`
		RegionName     string    `json:"regionName"`
	} `json:"allocation"`
}

type AllocatedReq struct {
	UUID            string   `json:"uuid"`            // uuid: allocation uuid
	IP              string   `json:"ip"`              // ip: 分配的服务器ip地址
	GameServerPorts []string `json:"gameServerPorts"` // gameServerPorts: 分配的服务器端口信息
	Status          string   `json:"status"`          // status: 分配状态，allocated 或 failed
	Msg             string   `json:"msg"`             // msg: 分配失败的原因
	UosAppID        string   `json:"uosAppId"`        // uosAppId: UosApp ID
	ProfileID       string   `json:"profileId"`       // profileId: 启动配置ID
	RegionID        string   `json:"regionId"`        // regionId: 地域ID
}
