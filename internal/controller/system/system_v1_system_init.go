package system

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gfAdmin/api/system/v1"
	"gfAdmin/internal/consts"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) SystemInit(ctx context.Context, req *v1.SystemInitReq) (res *v1.SystemInitRes, err error) {
	if(consts.Test){
		err = service.System().InitMenu(ctx)
		return nil,err
	}
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
