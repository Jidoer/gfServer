package v1

import "github.com/gogf/gf/v2/frame/g"

type SystemInitReq struct {
	g.Meta `path:"/system/init" method:"get" tags:"System" summary:"系统初始化"`
}

type SystemInitRes struct{}
