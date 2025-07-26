// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gfAdmin/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		GetSessionManager(ctx context.Context) *model.MySessionManager
		// Ctx injects custom business context variable into context of current request.
		Ctx(r *ghttp.Request)
		// 1.Auth 对请求进行验证，只允许已登录的用户访问
		// 2.CORS 允许跨域资源共享
		// 3.对path进行鉴权
		Auth(r *ghttp.Request)
		AuthPath(r *ghttp.Request, need_permissions *[]string)
		// CORS allows Cross-origin resource sharing.
		CORS(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
