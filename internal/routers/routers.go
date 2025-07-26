package routers

import (
	"fmt"
	"gfAdmin/internal/protorpc"

	"github.com/golang/protobuf/proto"
)

// type ClientMessage protorpc.ClientMessage
// type ServerMessage protorpc.ServerMessage

// 定义处理器函数类型
type ProtoRPCHandler func(req *protorpc.ClientMessage, res *protorpc.ServerMessage)

// 路由结构体
type ProtoRPCRouter struct {
	Method  string
	Handler ProtoRPCHandler
}

// 错误响应函数
func ErrorServerResponse(res *protorpc.ServerMessage, code int, message string) {
	res.Error.Code = int32(code)
	res.Error.Message = message
}

// 404 找不到请求
func ServerNotFound(req *protorpc.ClientMessage, res *protorpc.ServerMessage) {
	ErrorServerResponse(res, 404, "Not Found")
}

// 400 错误请求
func ServerBadRequest(req *protorpc.ClientMessage, res *protorpc.ServerMessage) {
	ErrorServerResponse(res, 400, "Bad Request")
}

// -----------------登录处理函数------------------------
func RpcServerLogin(req *protorpc.ClientMessage, res *protorpc.ServerMessage) {
	// 实现登录逻辑
	// res.Error.Code = 200
	// res.Error.Message = "Login successful"
	fmt.Println("RpcServerLogin()")
	res.Id = req.Id
	res.IsReq = false
	r := &protorpc.LoginResult{
		UserId: 1000,
		Token:  "123456",
	}

	data, err := proto.Marshal(r)
	if err != nil {
		res.Error.Code = 500
		res.Error.Message = err.Error()
		return
	}
	res.Result = data
}

// 路由表
var Router = []ProtoRPCRouter{
	{Method: "login", Handler: RpcServerLogin},
	{Method: "rpc_client_test", Handler: RpcServerLogin},
}
