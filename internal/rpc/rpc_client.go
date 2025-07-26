package rpc

import (
	"fmt"
	"gfAdmin/internal/consts"
	"gfAdmin/internal/protorpc"
	"net"

	"github.com/golang/protobuf/proto"
)

func RpcClientGetState(conn net.Conn, parms *protorpc.GetStutusParam, result **protorpc.GetStutusResult) {
	fmt.Println("RpcClientGetState()")
	var pm [][]byte
	// 将结构体序列化为字节数组
	d1, _ := proto.Marshal(parms)
	pm = append(pm, d1)
	req := protorpc.ServerMessage{
		IsReq:  true,
		Method: consts.Method_RpcClientGetState,
		Params: pm,
	}
	r_ := call_client.Call(&conn, &req, 15000)
	if r_ == nil {
		fmt.Println("call_client Call error()")
		return
	}
	// 	//解码
	r := protorpc.GetStutusResult{}
	proto.Unmarshal(r_.Result, &r)
	fmt.Println("get ", r.Device_Name, " status ok")
	*result = &r
}

// func call_test(conn net.Conn, parms protorpc.LoginParam, result *protorpc.LoginResult) { //rpc_client_test
// 	//login and parms
// 	//parms to [][]byte
// 	var pm [][]byte
// 	// 将结构体序列化为字节数组
// 	d1, _ := proto.Marshal(&parms)
// 	pm = append(pm, d1)
// 	req := protorpc.ServerMessage{
// 		IsReq:  true,
// 		Method: "rpc_client_test",
// 		Params: pm,
// 	}
// 	fmt.Println("call_client Calling for server_call_test()")
// 	r_ := call_client.Call(&conn, &req, 15000)
// 	if r_ == nil {
// 		fmt.Println("call_client Call error()")
// 		return
// 	}
// 	//解码
// 	login_test_r := protorpc.LoginResult{}
// 	proto.Unmarshal(r_.Result, &login_test_r)
// 	fmt.Println("Login ok", login_test_r.Token)
// 	result.Token = login_test_r.Token
// 	result.UserId = login_test_r.UserId
// }
