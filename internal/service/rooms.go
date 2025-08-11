// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"gfAdmin/internal/model"
)

type (
	IRooms interface {
		CreateRoom(in *model.AllocationReq) (out *model.AllocationRes, err error)
	}
)

var (
	localRooms IRooms
)

func Rooms() IRooms {
	if localRooms == nil {
		panic("implement not found for interface IRooms, forgot register?")
	}
	return localRooms
}

func RegisterRooms(i IRooms) {
	localRooms = i
}
