package mrpc

import (
	"gfAdmin/internal/protorpc"
	"gfAdmin/internal/client"
	"fmt"
	// "net"
)

// func Rpc_client_test(conn net.Conn, parms *protorpc.LoginParam, result *protorpc.LoginResult) {
// 	err := RpcCall(conn, "rpc_client_test", parms, result, 15000)
// 	if err != nil {
// 		fmt.Println("RPC failed:", err)
// 		return
// 	}
// 	fmt.Println("Login ok", result.Token)
// }

func Rpc_client_match_ok(cli *client.Client, parms *protorpc.MatchOKParam) (*protorpc.MatchOKResult, error) {
	result := new(protorpc.MatchOKResult)
	err := RpcCall(cli, "rpc_client_match_ok", parms, result, 15000)
	if err != nil {
		fmt.Println("RPC failed:", err)
		return nil, err
	}
	fmt.Println("Match OK", result)
	return result, nil
}

func Rpc_client_create_room(cli *client.Client, parms *protorpc.CreateRoomParam) (*protorpc.CreateRoomResult, error) {
	result := new(protorpc.CreateRoomResult)
	err := RpcCall(cli, "rpc_client_create_room", parms, result, 15000)
	if err != nil {
		fmt.Println("RPC failed:", err)
		return nil, err
	}
	fmt.Println("Create Room OK", result)
	return result, nil
}


