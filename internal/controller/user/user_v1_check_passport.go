package user

import (
	"context"

	// "github.com/gogf/gf/v2/errors/gcode"
	// "github.com/gogf/gf/v2/errors/gerror"

	"gfAdmin/api/user/v1"
	"gfAdmin/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) CheckPassport(ctx context.Context, req *v1.CheckPassportReq) (res *v1.CheckPassportRes, err error) {
	var ok bool
	ok,err = service.User().CheckPassport(ctx, req.Passport)
	if(err != nil){
		return
	}
	if !ok {
		return nil, gerror.Newf(`Passport "%s" is already token by others`, req.Passport)
	}
	return
}
