package vars

import (
	"context"
	"gfAdmin/internal/model"
	"gfAdmin/internal/service"

	"github.com/gogf/gf/v2/frame/g"

)

var Roles_Auths []model.Roles

func init() {
	ctx := context.Background()
	total := 0 
	Roles_Auths, total, err := service.Permissions().GetRolesList(ctx, 1, 1000)
	if(err != nil){
		panic(err)
	}
	g.Log().Info(ctx,"Roles_Auths total: %d", total)
	g.Log().Info(ctx,"Roles_Auths: %v", Roles_Auths)
}

