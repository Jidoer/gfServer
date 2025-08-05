package mrpc

import (
	"gfAdmin/internal/client"
	// "gfAdmin/internal/dbase"
	"gfAdmin/internal/protorpc"
	// "gfAdmin/internal/tool"
	"fmt"

	"github.com/golang/protobuf/proto"
)

// type ClientMessage protorpc.ClientMessage
// type ServerMessage protorpc.ServerMessage

// 定义处理器函数类型
type ProtoRPCHandler func(TransportConn *TransportConn, req *protorpc.ClientMessage, res *protorpc.ServerMessage)

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

func HandleRPC[T proto.Message, R proto.Message](
	req *protorpc.ClientMessage,
	res *protorpc.ServerMessage,
	param T,
	handler func(param T) (R, error), //handle user func
) {
	res.Id = req.Id
	res.IsReq = false

	err := proto.Unmarshal(req.Params, param)
	if err != nil {
		res.Error = &protorpc.Error{
			Code:    500,
			Message: fmt.Sprintf("Unmarshal error: %v", err),
		}
		return
	}

	result, err := handler(param)
	if err != nil {
		res.Error = &protorpc.Error{
			Code:    500,
			Message: fmt.Sprintf("Handler error: %v", err),
		}
		return
	}

	data, err := proto.Marshal(result)
	if err != nil {
		res.Error = &protorpc.Error{
			Code:    500,
			Message: fmt.Sprintf("Marshal error: %v", err),
		}
		return
	}
	res.Error = &protorpc.Error{
		Code:    200,
		Message: "OK",
	}
	res.Result = data
}

// -----------------登录处理函数------------------------
// func RpcServerLogin(req *protorpc.ClientMessage, res *protorpc.ServerMessage) {
// 	// 实现登录逻辑
// 	// res.Error.Code = 200
// 	// res.Error.Message = "Login successful"
// 	fmt.Println("RpcServerLogin()")
// 	res.Id = req.Id
// 	res.IsReq = false
// 	r := &protorpc.LoginResult{
// 		UserId: 1000,
// 		Token:  "123456",
// 	}

// 	data, err := proto.Marshal(r)
// 	if err != nil {
// 		res.Error.Code = 500
// 		res.Error.Message = err.Error()
// 		return
// 	}
// 	res.Result = data
// }

// func rpc_func_server_login(conn *TransportConn, req *protorpc.ClientMessage, res *protorpc.ServerMessage) {
// 	HandleRPC(req, res, &protorpc.LoginParam{}, func(param *protorpc.LoginParam) (*protorpc.LoginResult, error) {
// 		// if dbase.CheckUsername(param.Username) == false {
// 		// 	return nil, fmt.Errorf("[L]username %s not found", param.Username)
// 		// }
// 		// if dbase.CheckPassword(param.Password) == false {
// 		// 	return nil, fmt.Errorf("[L]password for user %s is incorrect", param.Username)
// 		// }
// 		pwd := tool.GetStringMd5(param.Password)
// 		u, tkn, err := dbase.Login(param.Username, pwd) // 登录时还会检查用户名和密码的合法性
// 		if err != nil || u == nil {
// 			return nil, fmt.Errorf("[L]login failed: %v", err)
// 		}
// 		conn.Client.Userinfo = u //设置用户信息到客户端

// 		r := &protorpc.LoginResult{
// 			Result: true,
// 			User: &protorpc.UserInfo{
// 				Uid:      uint64(u.ID),
// 				Username: u.Username,
// 			},
// 			Token: tkn,
// 		}
// 		return r, nil
// 	})
// }

// func rpc_func_server_register(conn *TransportConn, req *protorpc.ClientMessage, res *protorpc.ServerMessage) {
// 	HandleRPC(req, res, &protorpc.RegisterParam{}, func(param *protorpc.RegisterParam) (*protorpc.RegisterResult, error) {
// 		// if dbase.CheckUsername(param.Username) {
// 		// 	return nil, fmt.Errorf("[R]username %s invalid", param.Username)
// 		// }
// 		// if len(param.Password) < 6 {
// 		// 	return nil, fmt.Errorf("[R]password for user %s is too short", param.Username)
// 		// }
// 		// if len(param.Password) > 32 {
// 		// 	return nil, fmt.Errorf("[R]password for user %s is too long", param.Username)
// 		// }
// 		// if dbase.CheckEmail(param.Email) {
// 		// 	return nil, fmt.Errorf("[R]email %s already exists", param.Email)
// 		// }
// 		// 注册用户
// 		newUser := dbase.UserInfo{
// 			Username: param.Username,
// 			Password: param.Password,
// 		}
// 		u, err := dbase.Register(&newUser) //注册时还会检查用户名和密码的合法性
// 		if err != nil {
// 			return nil, fmt.Errorf("[R]register failed: %v", err)
// 		}
// 		r := &protorpc.RegisterResult{
// 			Result: true,
// 			User: &protorpc.UserInfo{
// 				Uid:      uint64(u.ID),
// 				Username: u.Username,
// 			},
// 		}
// 		return r, nil
// 	})
// }

// func rpc_func_server_login_with_token(conn *TransportConn, req *protorpc.ClientMessage, res *protorpc.ServerMessage) {
// 	HandleRPC(req, res, &protorpc.LoginWithTokenParam{}, func(param *protorpc.LoginWithTokenParam) (*protorpc.LoginWithTokenResult, error) {
// 		u, err := dbase.LoginWithToken(param.Username, param.Token)
// 		if err != nil || u == nil {
// 			return nil, fmt.Errorf("[LWT]login with token failed: %v", err)
// 		}
// 		conn.Client.Userinfo = u //设置用户信息到客户端
// 		r := &protorpc.LoginWithTokenResult{
// 			Result: true,
// 			User: &protorpc.UserInfo{
// 				Uid:      uint64(u.ID),
// 				Username: u.Username,
// 			},
// 			Token: param.Token,
// 		}
// 		return r, nil
// 	})
// }

func rpc_func_server_match(conn *TransportConn, req *protorpc.ClientMessage, res *protorpc.ServerMessage) {
	HandleRPC(req, res, &protorpc.MatchParam{}, func(param *protorpc.MatchParam) (*protorpc.MatchResult, error) {
		//收到后返回等待匹配
		r := &protorpc.MatchResult{
			Status: protorpc.MatchStatus_MATCHING,
		}
		go match_normal(*conn.Client, int(param.MatchType)) //开启匹配协程
		return r, nil
	})
}

//AuthKey = "_kMK,dk(Ml*kd&e+k#Kc=$dK;Kn,d=e#4dc=s.@dld-lss^ss~HtuP"

func rpc_func_server_register_server(conn *TransportConn, req *protorpc.ClientMessage, res *protorpc.ServerMessage) {
	HandleRPC(req, res, &protorpc.RegisterServerParam{}, func(param *protorpc.RegisterServerParam) (*protorpc.RegisterServerResult, error) {
		if param.AuthKey != "_kMK,dk(Ml*kd&e+k#Kc=$dK;Kn,d=e#4dc=s.@dld-lss^ss~HtuP" {
			conn.Close() //close the session
			return nil, fmt.Errorf("[RSS]server info is nil")
		}

		conn.Client.Auth = 2
		// _ip, _ := conn.GetIp()
		fmt.Println("Port:", int(param.ServerInfo.Port))
		conn.Client.Service = &client.SubServer{
			Id:       param.ServerInfo.Id,
			Name:     param.ServerInfo.Name,
			IP:       "tun.irnn.cn", //_ip,
			Port:     7000,          //int(param.ServerInfo.Port), //int(param.Port),
			GameType: int(param.ServerInfo.GameType),
			Version:  int(param.ServerInfo.Version),
		}
		r := &protorpc.RegisterServerResult{
			Result: true,
			// Server: param.Server,
		}
		return r, nil
	})
}

// 路由表
var Router = []ProtoRPCRouter{
	// {Method: "rpc_func_server_login", Handler: rpc_func_server_login},
	// {Method: "rpc_func_server_register", Handler: rpc_func_server_register},
	// {Method: "rpc_func_server_login_with_token", Handler: rpc_func_server_login_with_token},
	{Method: "rpc_func_server_match", Handler: rpc_func_server_match},
	{Method: "rpc_func_server_register_server", Handler: rpc_func_server_register_server},
}
