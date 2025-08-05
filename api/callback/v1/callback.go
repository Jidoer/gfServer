package v1


//Romm created
type RoomCreatedReq struct {
	RoomId string `json:"roomId"`
}

type RoomCreatedRes struct {
	Success bool `json:"success"`
}