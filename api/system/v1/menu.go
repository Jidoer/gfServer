package v1

import (
	// "gfAdmin/internal/model"
	"gfAdmin/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetMenuReq struct {
	g.Meta `path:"/system/menu/get" method:"get" tags:"System" summary:"获取菜单列表"`
}

type GetMenuRes []entity.SystemMenus


