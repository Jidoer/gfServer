package mrpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"gfAdmin/internal/call"
	"gfAdmin/internal/client"
	"gfAdmin/internal/protorpc"
	"gfAdmin/internal/service"

	proto "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

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

func GetControl() *Control {
	return control
}
func init() {
	fmt.Println("Initializing rpc call_client...")
	call_client = *call.NewProtoRpcClient()
}
func HandleConnection(ctx context.Context, conn TransportConn, setting UnpackSetting) {
	//defer conn.Close()

	//获取已登录用户信息
	user := service.BizCtx().Get(ctx).User
	if user == nil {
		fmt.Println("err 用户未登录")
		return
	}
	client := NewClient(&conn) // 创建设备信息
	control.Subscribe(client)  // 设备连接中控台
	conn.Client = client       // 连接信息中添加设备信息
	conn.Control = control     // 连接信息中添加控制器
	client.Context = &ctx
	client.User = user
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
	if conn.Client == nil {
		return
	}
	if conn.Client.User == nil {
		if conn.Client.Auth == 2 {
			goto label_qq
		}
		if client_req.Method != "rpc_func_server_login" && client_req.Method != "rpc_func_server_register" && client_req.Method != "rpc_func_server_login_with_token" && client_req.Method != "rpc_func_server_register_server" {
			// ServerNotFound(client_req, &res)
			if conn.WsConn != nil {
				(*conn.WsConn).Close()
			}
			if conn.TcpConn != nil {
				(*conn.TcpConn).Close()
			}
			fmt.Println("[!]Client not logged in, closing connection.")
			return
		}
	}
label_qq:
	// 处理客户端请求
	found := false
	for _, route := range Router {
		if route.Method == client_req.Method {
			found = true
			route.Handler(conn, client_req, &res)
			break
		}
	}
	if !found {
		ServerNotFound(client_req, &res)
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

func RpcCall(cli *client.Client, method string, param proto.Message, result proto.Message, timeoutMs int) error { //common sampler call ->call()
	// 序列化参数
	data, err := proto.Marshal(param)
	if err != nil {
		return fmt.Errorf("marshal param failed: %v", err)
	}
	// 构造请求
	req := protorpc.ServerMessage{
		IsReq:  true,
		Method: method,
		Params: data, //[][]byte{data},
	}
	// 发起调用
	resp := call_client.Call(cli, &req, timeoutMs)
	if resp == nil {
		return fmt.Errorf("call_client.Call() returned nil")
	}
	// 反序列化结果
	err = proto.Unmarshal(resp.Result, result)
	if err != nil {
		return fmt.Errorf("unmarshal result failed: %v", err)
	}
	return nil
}
