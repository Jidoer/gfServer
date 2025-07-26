package rpc

import (
	//"fmt"
	"gfAdmin/internal/client"
	"sync"

	"github.com/gorilla/websocket"
)

type Control struct {
	ClientLock sync.RWMutex
	Clients    map[string]*client.Client
	Messages   chan *Message
}

func NewControl() *Control {
	c := &Control{
		Clients:  make(map[string]*client.Client),
		Messages: make(chan *Message, 1000),
	}
	c.StartBroadcast() // 创建广播
	return c
}

// 设备号订阅
func (c *Control) Subscribe(client *client.Client) *Control {
	c.ClientLock.Lock()
	c.Clients[client.Id] = client
	c.ClientLock.Unlock()
	//welecom :)
	//msg := fmt.Sprintf("<info>login</info><id>%s</id><zs>%d</zs>", client.Id, len(c.Clients))
	//c.Broadcast(NewMessage(client, []byte(msg)))

	return c
}

// 设备号取消订阅
func (c *Control) UnSubscribe(client *client.Client) *Control {
	c.ClientLock.Lock()
	if c.Clients != nil {
		delete(c.Clients, client.Id)
	}
	c.ClientLock.Unlock()
	//Baye :(
	//msg := fmt.Sprintf("<info>out</info><id>%s</id><zs>%d</zs>", client.Id, len(c.Clients))
	//c.Broadcast(NewMessage(client, []byte(msg)))

	return c
}

// 开始消息广播（全体）
func (c *Control) StartBroadcast() {
	go func() {
		for {
			message := <-c.Messages
			for _, client := range c.Clients {
				if client.WsConn != nil {
					(*(client.WsConn)).WriteMessage(websocket.BinaryMessage, message.Content)
				} else if client.TcpConn != nil {
					(*(client.TcpConn)).Write([]byte(message.Content))
				}
			}
		}
	}()
}

// 发送消息 广播
func (c *Control) Broadcast(message *Message) {
	c.Messages <- message
}

func (c *Control) SendTo(ClientID string, message *Message) {
	for _, client := range c.Clients {
		if client.Id == ClientID {
			//(*(client.Conn)).Write([]byte(message.Content))
			if client.WsConn != nil {
				(*(client.WsConn)).WriteMessage(websocket.BinaryMessage, message.Content)
			} else if client.TcpConn != nil {
				(*(client.TcpConn)).Write([]byte(message.Content))
			}
		}
	}
}

// func (c *Control) GetServerList() *pb.ServerList {
// 	var mList []*pb.Server
// 	list_ := c.GetServerListForNeb()
// 	for _, s_ := range *list_ {
// 		sub_list := pb.Server{
// 			IP:       s_.IP,
// 			Name:     s_.Name,
// 			Port:     int32(s_.Port),
// 			GameType: int32(s_.GameType),
// 			Version:  int32(s_.Version),
// 		}
// 		mList = append(mList, &sub_list)
// 	}
// 	return &pb.ServerList{
// 		List: mList,
// 	}
// }

func (c *Control) GetServerListForNeb() *[]client.SubServer {
	var s []client.SubServer
	for _, sub_c := range c.Clients {
		if sub_c.Auth == 2 && sub_c.Service != nil {
			sm := *sub_c.Service
			s = append(s, sm)
		}
	}
	return &s
}
