// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"gfAdmin/api/system/v1"
)

type ISystemV1 interface {
	GetMenu(ctx context.Context, req *v1.GetMenuReq) (res *v1.GetMenuRes, err error)
	SystemInit(ctx context.Context, req *v1.SystemInitReq) (res *v1.SystemInitRes, err error)
}
