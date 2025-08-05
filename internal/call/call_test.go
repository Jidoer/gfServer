package call

import (
	"gfAdmin/internal/protorpc"
	"testing"
)

func TestCall(t *testing.T) {
	req := protorpc.ServerMessage{Method: "login", Id: 1}

	client := NewProtoRpcClient()
	client.Call(nil, &req, 10000)
	t.Log("done")

}
