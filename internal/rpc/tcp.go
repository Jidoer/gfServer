package rpc

// import (
// 	"gfAdmin/internal/call"
// 	"gfAdmin/internal/config"
// 	"fmt"
// 	"net"
// )

// func StartTCPServer(cfg *config.Config) {
// 	call_client = *call.NewProtoRpcClient()
	
// 	unpackSetting := UnpackSetting{
// 		Mode:              1, // 自定义模式
// 		PackageMaxLength:  DEFAULT_PACKAGE_MAX_LENGTH,
// 		BodyOffset:        PROTORPC_HEAD_LENGTH,
// 		LengthFieldOffset: PROTORPC_HEAD_LENGTH_FIELD_OFFSET,
// 		LengthFieldBytes:  PROTORPC_HEAD_LENGTH_FIELD_BYTES,
// 	}
// 	listener, err := net.Listen("tcp", ":7070")
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 		return
// 	}
// 	defer listener.Close()

// 	fmt.Println("Server listening on port 7070")
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting connection:", err)
// 			continue
// 		}
// 		transport_conn := TransportConn{
// 			TransportStack: "tcp",
// 			TcpConn: &conn,
// 		}
// 		go HandleConnection(ctx,transport_conn, unpackSetting)
// 	}
// }
