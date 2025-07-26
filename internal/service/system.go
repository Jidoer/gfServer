// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gfAdmin/internal/model/entity"
)

type (
	ISystem interface {
		GetMenu(ctx context.Context) (*[]entity.SystemMenus, int, error)
		GetMenuByte() []byte
		SaveMenu(ctx context.Context, menu *entity.SystemMenus) error
		InitMenu(ctx context.Context) error
	}
)

var (
	localSystem ISystem
)

func System() ISystem {
	if localSystem == nil {
		panic("implement not found for interface ISystem, forgot register?")
	}
	return localSystem
}

func RegisterSystem(i ISystem) {
	localSystem = i
}
