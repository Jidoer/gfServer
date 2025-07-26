package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"gfAdmin/internal/call"
	"gfAdmin/internal/service"

	// "gfAdmin/internal/model/entity"
	"gfAdmin/internal/protorpc"
	"gfAdmin/internal/routers"

	"github.com/gogf/gf/v2/frame/g"
	proto "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)
var logger = g.Log("RPC SERVER")

const (
	DEFAULT_PACKAGE_MAX_LENGTH        = 1024
	PROTORPC_HEAD_LENGTH              = 8
	PROTORPC_HEAD_LENGTH_FIELD_OFFSET = 4
	PROTORPC_HEAD_LENGTH_FIELD_BYTES  = 4
)

type UnpackSetting struct {
	Mode              int
	PackageMaxLength  int
	BodyOffset        int
	LengthFieldOffset int
	LengthFieldBytes  int
}

var call_client call.ProtoRpcClient
var control = NewControl()

func HandleConnection(ctx context.Context,conn TransportConn, setting UnpackSetting) { //加入个session指针传入 *Client_info 可空。。
	//defer conn.Close()
	// client := NewClient(&conn) // 创建设备信息
	// service.PrintServer().SetClientID(ctx,client.Id)
	id := service.PrintServer().Get(ctx).Session.Passport
	client := New(id)
	client.WsConn = conn.WsConn
	
	ps := service.PrintServer().Get(ctx)
	if(ps == nil){
		logger.Error(ctx,"Connect error: Get PrintServer info error")
		return
	}
	control.Subscribe(client)  // 设备连接中控台
	defer func() {
		if conn.WsConn != nil {
			control.UnSubscribe(client)
			(*conn.WsConn).Close()
			return
		}
		if conn.TcpConn != nil {
			control.UnSubscribe(client)
			(*conn.TcpConn).Close()
		}
	}()

	buffer := make([]byte, setting.PackageMaxLength)
	for {
		var read_err error
		n := 0

		if conn.TransportStack == "tcp" {
			n, read_err = (*conn.TcpConn).Read(buffer)
		} else if conn.TransportStack == "websocket" {
			_, p, err0 := (*conn.WsConn).ReadMessage()
			if err0 != nil {
				fmt.Printf("Client(%s)读取失败!\n", conn.TransportStack)
				break
			}
			n = len(p)
			buffer = p
		}
		if read_err != nil {
			fmt.Printf("Client(%s)读取失败!\n", conn.TransportStack)
			break
		}

		if n < setting.BodyOffset+setting.LengthFieldBytes {
			fmt.Println("Received data is too short")
			continue
		}
		fmt.Println("[me<-]Received data:")
		//print hex
		for i := 0; i < setting.BodyOffset; i++ {
			fmt.Printf("%02x ", buffer[i])
		}
		fmt.Printf("\n")
		fmt.Printf("input<-Hex: %s\n", hex.Dump(buffer[:n]))
		msg := protorpc.ProtoRPCMessage{}
		len_, err := protorpc.ProtorpcUnpack(&msg, buffer[:n])
		fmt.Println("[?]Unpacked length:", len_)
		if err != nil {
			fmt.Println("Error unpacking:", err)
			continue
		}
		ckerr := protorpc.ProtorpcHeadCheck(&msg.Head)
		if ckerr != nil {
			fmt.Println("Check error:", ckerr)
			continue
		}
		var client_req protorpc.ClientMessage
		err_ := proto.Unmarshal([]byte(msg.Body), &client_req)
		if err_ != nil {
			fmt.Println("Error proto unpacking:", err_)
			continue
		}
		//fmt.Printf("Received client message: %v\n", client_req)
		//res := protorpc.ServerMessage{}
		if client_req.IsReq {
			go onReqMessage(&conn, &client_req)
		} else {
			go onResMessage(&conn, &client_req)
		}

	}
}

func onReqMessage(conn *TransportConn, client_req *protorpc.ClientMessage) { //客户端请求
	res := protorpc.ServerMessage{}
	fmt.Println("[?]Received client request:", client_req.Id)
	// 处理客户端请求
	found := false
	for _, route := range routers.Router {
		if route.Method == client_req.Method {
			found = true
			route.Handler(client_req, &res)
			break
		}
	}
	if !found {
		routers.ServerNotFound(client_req, &res)
	}
	// 发送响应
	serverBuf, _ := proto.Marshal(&res)
	msg := protorpc.ProtoRPCMessage{}
	protorpc.ProtorpcMessageInit(&msg)
	msg.Head.Length = uint32(len(serverBuf))
	msg.Body = serverBuf
	buf := make([]byte, 1024)
	len_, _ := protorpc.ProtorpcPack(&msg, &buf)
	if conn.TransportStack == "tcp" {
		(*conn.TcpConn).Write(buf[:len_])
		return
	}
	if conn.TransportStack == "websocket" {
		(*conn.WsConn).WriteMessage(websocket.BinaryMessage, buf[:len_])
	}

	// 服务器请求客户端测试
	// result := protorpc.LoginResult{}
	// p := protorpc.LoginParam{
	// 	Username: "user",
	// 	Password: "pass",
	// }
	// server_call_test(*conn, p, &result)
}

func onResMessage(conn *TransportConn, client_req *protorpc.ClientMessage) { //客户端响应
	// 处理客户端响应
	fmt.Println("[!]Received client response:", client_req.Id)
	call_client.CallsMutex.Lock()
	if call_client.Calls[client_req.Id] != nil {
		call_client.Calls[client_req.Id].Res = client_req
		call_client.Calls[client_req.Id].Notify()
	}
	call_client.CallsMutex.Unlock()

}

// func server_call_test(conn net.Conn, parms *protorpc.LoginParam, result *protorpc.LoginResult) { //rpc_client_test
// 	//login and parms
// 	//parms to [][]byte
// 	var pm [][]byte
// 	// 将结构体序列化为字节数组
// 	d1, _ := proto.Marshal(parms)
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
