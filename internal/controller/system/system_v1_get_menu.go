package system

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gfAdmin/api/system/v1"
	"gfAdmin/internal/service"
)

// func get_menu_from_db() (*model.Menus,error){
// 	menus, err := g.DB().Model("sys_menu").Order("parent_id, id").All()
// 	if err != nil {
// 		return nil, err
// 	}
// }

func (c *ControllerV1) GetMenu(ctx context.Context, req *v1.GetMenuReq) (res *v1.GetMenuRes, err error) {
	// return nil, gerror.NewCode(gcode.CodeNotImplemented)
	user := service.User().GetProfile(ctx)
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized)
	}
	//if user.Role == 10 {
		menu, _, err := service.System().GetMenu(ctx)
		if err != nil {
			return nil, err
		}
		res = (*v1.GetMenuRes)(menu)
		return res, nil
	//}
}
