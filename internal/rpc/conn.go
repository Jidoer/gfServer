package rpc

import (
	"net"

	"github.com/gorilla/websocket"
)

type TransportConn struct {
	TransportStack string //supported tcp ws
	TcpConn        *net.Conn
	WsConn         *websocket.Conn
}
