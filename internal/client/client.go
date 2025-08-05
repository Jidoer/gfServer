package client

import (
	// "gfAdmin/internal/model"
	"context"
	"gfAdmin/internal/model"
	// "gfAdmin/internal/model"
	// "gfAdmin/internal/model/entity"
	"net"

	"github.com/gorilla/websocket"
	//"net"
)

type Client struct {
	//Conn *net.Conn
	WsConn  *websocket.Conn
	TcpConn *net.Conn
	Id      string
	//*Userinfo
	// Userinfo *model.User_Session // *UserInfo 游戏客户端和后台通用。。。可空
	// PrintServer *entity.PrintServer
	// Userinfo *entity.User // *UserInfo
	// Context *model.Context
	User    *model.ContextUser
	Context *context.Context
	Tokan    string
	Auth     int //0 unauthorized, 1 client authorized 2 server authorized
	Service  *SubServer //!=nil -->is sub server || 子服务器通信可以使用http api!
}

type SubServer struct {
	Id       string
	IP       string
	Name     string
	Port     int
	GameType int // Match_TYPE ...
	Version  int
}
