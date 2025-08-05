package mrpc

import (
	"gfAdmin/internal/client"
	"fmt"
	"net"

	"github.com/gorilla/websocket"
)

type TransportConn struct {
	TransportStack string //supported tcp ws
	TcpConn        *net.Conn
	WsConn         *websocket.Conn
	//加入设备信息 客户端信息...
	Client  *client.Client
	Control *Control
}

func (t *TransportConn) Close() {
	if t.TcpConn != nil {
		(*t.TcpConn).Close()
	}
	if t.WsConn != nil {
		t.WsConn.Close()
	}
	// if t.Client != nil {
	// 	t.Control.UnSubscribe(t.Client)
	// }
}

func (t *TransportConn) GetIp() (string,error) {
	if t.TcpConn != nil {
		return (*t.TcpConn).RemoteAddr().String(),nil
	} else if t.WsConn != nil {
		return t.WsConn.RemoteAddr().String(),nil
	}
	return "",fmt.Errorf("TransportConn has no valid connection")
}
