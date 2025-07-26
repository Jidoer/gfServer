package rpc

import (
	"gfAdmin/internal/client"
	"net"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

/*move to MasterServer/client/client.go
type Client struct {
	Conn *net.Conn
	Id   string
	//*Userinfo
	Userinfo *dbase.UserInfo // *UserInfo
	Tokan    string
	Auth     int //0 unauthorized, 1 client authorized 2 server authorized
}*/

// 生成唯一标识（设备号）
func generateId() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

func NewWsClient(wsConn *websocket.Conn) *client.Client {
	return &client.Client{
		WsConn:   wsConn,
		Id:       generateId(),
		Tokan:    "",
		// Userinfo: nil,
		Auth:     0,
	}
}
func NewTcpClient(conn *net.Conn) *client.Client {
	return &client.Client{
		TcpConn:     conn,
		Id:       generateId(),
		Tokan:    "",
		// Userinfo: nil,
		Auth:     0,
	}
}
func NewClient(conn *TransportConn) *client.Client {
	if(conn.WsConn != nil){
		return NewWsClient(conn.WsConn)
	}
	return NewTcpClient(conn.TcpConn)
}

func New(id string) *client.Client{
	return &client.Client{
		WsConn:   nil,
		TcpConn: nil,
		Id:       id,
		Tokan:    "",
		Auth:     0,
	}
}
