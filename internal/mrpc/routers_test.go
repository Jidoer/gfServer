package mrpc

import (
	"gfAdmin/internal/protorpc"
	"testing"
)

func Test_Routers(t *testing.T) {
	req := protorpc.ClientMessage{Method: "login"}
	res := &protorpc.ServerMessage{}

	for _, route := range Router {
		if route.Method == req.Method {
			route.Handler(nil,&req, res)
			t.Log("Response:", res)
			return
		}
	}
	ServerNotFound(&req, res)
	t.Log("Response Server NotFound", res)
}