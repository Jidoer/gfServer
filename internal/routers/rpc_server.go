package routers

import (
	"gfAdmin/internal/protorpc"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func Rpc_Server_AutoRegister(req *protorpc.ClientMessage, res *protorpc.ServerMessage) {
	fmt.Println("Rpc_Server_AutoRegister()")
	res.Id = req.Id
	res.IsReq = false
	var loginParam protorpc.LoginParam
	e:= proto.Unmarshal(req.Params[0],&loginParam)
	if e!=nil{
		res.Error.Code = 500
		res.Error.Message = e.Error()
		return
	}
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