package user

import (
	"context"

	"gfAdmin/api/user/v1"
	"gfAdmin/internal/service"
)

func (c *ControllerV1) PhoneRegister(ctx context.Context, req *v1.PhoneRegisterReq) (res *v1.PhoneRegisterRes, err error) {
	r, err := service.Sms().SendaSms(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	res = &v1.PhoneRegisterRes{
		CodeID: r,
	}
	return res, nil
}
