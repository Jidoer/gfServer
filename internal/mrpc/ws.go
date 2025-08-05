package mrpc

// import (
// 	"GoSyncServe/config"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool { return true },
// }

// func StartWebSocketServer(cfg *config.Config) {
// 	unpackSetting := UnpackSetting{
// 		Mode:              1, // 自定义模式
// 		PackageMaxLength:  DEFAULT_PACKAGE_MAX_LENGTH,
// 		BodyOffset:        PROTORPC_HEAD_LENGTH,
// 		LengthFieldOffset: PROTORPC_HEAD_LENGTH_FIELD_OFFSET,
// 		LengthFieldBytes:  PROTORPC_HEAD_LENGTH_FIELD_BYTES,
// 	}

// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		conn, err := upgrader.Upgrade(w, r, nil)
// 		if err != nil {
// 			log.Println("WebSocket 连接失败:", err)
// 			return
// 		}
// 		// defer conn.Close()
// 		// log.Println("WebSocket 连接建立")
// 		transport_conn := TransportConn{
// 			TransportStack: "websocket",
// 			WsConn:         conn,
// 		}
// 		handleConnection(transport_conn, unpackSetting)
// 	})

// 	log.Println("WebSocket 服务器启动:", cfg.Listen)
// 	// http.ListenAndServe(cfg.Listen, nil)
// 	err := http.ListenAndServeTLS(cfg.Listen, "cert.pem", "key.pem", nil)
//     if err != nil {
//         log.Fatal("ListenAndServeTLS error: ", err)
//     }
// }
