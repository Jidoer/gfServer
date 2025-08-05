package call

import (
	"gfAdmin/internal/client"
	// "gfAdmin/internal/protorpc"
	pc "gfAdmin/internal/protorpc"
	"fmt"
	// "net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

type ContextPtr struct {
	Req *pc.ServerMessage //*Request
	Res *pc.ClientMessage //*Response
	// Cond *sync.Cond
	Done chan struct{}
}

func (ctx *ContextPtr) Notify() {
	fmt.Println(ctx.Req.Id, "is Notify()")
	//ctx.Cond.Broadcast()
	close(ctx.Done)
}

func (ctx *ContextPtr) Wait(timeoutMs int) {
	// ctx.Cond.L.Lock()
	// defer ctx.Cond.L.Unlock()
	timeout := time.After(time.Duration(timeoutMs) * time.Millisecond)
	select {
	case <-ctx.Done:
		fmt.Println("Wait was successful")
		return
	case <-timeout:
		fmt.Println("Timeout occurred")
		return
	}
}

type ProtoRpcClient struct {
	// connectState int
	Calls      map[uint64]*ContextPtr
	CallsMutex sync.Mutex
}

func NewProtoRpcClient() *ProtoRpcClient {
	client := &ProtoRpcClient{
		Calls: make(map[uint64]*ContextPtr),
	}
	return client
}

var id uint64 = 0 //递增

func (client *ProtoRpcClient) Call(cli *client.Client, req *pc.ServerMessage, timeoutMs int) *pc.ClientMessage {
	client.CallsMutex.Lock()
	req.Id = atomic.AddUint64(&id, 1)
	ctx := &ContextPtr{
		Req: req,
		//Cond: sync.NewCond(&sync.Mutex{}),
		Done: make(chan struct{}),
	}
	client.Calls[req.Id] = ctx // Add request to the map
	client.CallsMutex.Unlock()

	req.IsReq = true
	serverMessageBuf, _ := proto.Marshal(req)
	msg := pc.ProtoRPCMessage{}
	pc.ProtorpcMessageInit(&msg)
	msg.Head.Length = uint32(len(serverMessageBuf))
	msg.Body = serverMessageBuf
	buf := make([]byte, 1024)
	len_, _ := pc.ProtorpcPack(&msg, &buf)
	if(cli != nil){
		if(cli.WsConn != nil) {
			// WebSocket connection
			err := (*cli.WsConn).WriteMessage(websocket.BinaryMessage, buf[:len_])
			if err != nil {
				fmt.Println("Error writing to WebSocket:", err)
				return nil
			}
		}
		if(cli.TcpConn != nil) {
			// TCP connection
			_, err := (*cli.TcpConn).Write(buf[:len_])
			if err != nil {
				fmt.Println("Error writing to TCP connection:", err)
				return nil
			}
		}
	}
	//(*conn).Write(buf[:len_])
	ctx.Wait(timeoutMs)
	client.CallsMutex.Lock()
	defer client.CallsMutex.Unlock()
	delete(client.Calls, req.Id)
	if ctx.Res != nil {
		return ctx.Res
	}
	return nil
}
