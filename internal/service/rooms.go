// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gfAdmin/internal/model"
)

type (
	IRooms interface {
		CreateRoom(in *model.Room_CreateRoomReq) (out *model.Room_CreateRoomRes, err error)
		QueryRoom(ctx context.Context) (list *model.Room_ListRoomsRes, err error)
		GetRoomByUuid(ctx context.Context, uuid string) (r *model.Room, err error)
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
