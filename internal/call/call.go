package call

import (
	"fmt"
	"gfAdmin/internal/protorpc"
	pc "gfAdmin/internal/protorpc"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
)

var call_id uint64

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

func generateID() uint64 {
	return atomic.AddUint64(&call_id, 1)
}

func (client *ProtoRpcClient) Call(conn *net.Conn, req *pc.ServerMessage, timeoutMs int) *pc.ClientMessage {
	//生成req.ID
	req.Id = generateID()
	client.CallsMutex.Lock()
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
	len_, _ := protorpc.ProtorpcPack(&msg, &buf)
	(*conn).Write(buf[:len_]) //not safe
	ctx.Wait(timeoutMs)

	client.CallsMutex.Lock()
	defer client.CallsMutex.Unlock()
	delete(client.Calls, req.Id)
	if ctx.Res != nil {
		return ctx.Res
	}
	return nil
}
