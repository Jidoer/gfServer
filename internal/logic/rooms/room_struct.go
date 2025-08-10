package rooms

import "time"

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
	UUID            string   `json:"uuid"`
	IP              string   `json:"ip"`
	GameServerPorts []string `json:"gameServerPorts"`
	Status          string   `json:"status"`
	Msg             string   `json:"msg"`
	UosAppID        string   `json:"uosAppId"`
	ProfileID       string   `json:"profileId"`
	RegionID        string   `json:"regionId"`
}