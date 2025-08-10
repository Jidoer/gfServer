package cmd

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"

	// "github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gmode"

	"gfAdmin/internal/cache"
	"gfAdmin/internal/consts"
	"gfAdmin/internal/controller/admin"
	"gfAdmin/internal/controller/system"
	"gfAdmin/internal/controller/test_api"
	"gfAdmin/internal/controller/user"
	"gfAdmin/internal/mrpc"
	"gfAdmin/internal/service"

	"github.com/gorilla/websocket"
)

// var (
// 	// RouteGroupPrefix 路由统一前缀
// 	RouteGroupPrefix = "/v1"
// )

var logger = g.Log("cmd")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(r *ghttp.Request) {
	// p := r.Cookie.Get("p") //PassPort
	// t := r.Cookie.Get("t") //Token
	// service.BizCtx().Init()
	logger.Info(r.Context(), "WebSocket 验证...")
	service.Middleware().Ctx(r)
	service.Middleware().Auth(r)

	logger.Info(r.Context(), "WebSocket 认证成功 连接建立...")
	conn, err := upgrader.Upgrade(r.Response.ResponseWriter, r.Request, nil)
	if err != nil {
		logger.Error(r.Context(), "WebSocket Upgrade 失败:", err)
		return
	}
	unpackSetting := mrpc.UnpackSetting{
		Mode:              1, // 自定义模式
		PackageMaxLength:  mrpc.DEFAULT_PACKAGE_MAX_LENGTH,
		BodyOffset:        mrpc.PROTORPC_HEAD_LENGTH,
		LengthFieldOffset: mrpc.PROTORPC_HEAD_LENGTH_FIELD_OFFSET,
		LengthFieldBytes:  mrpc.PROTORPC_HEAD_LENGTH_FIELD_BYTES,
	}
	//r.Session.
	transport_conn := mrpc.TransportConn{
		TransportStack: "websocket",
		WsConn:         conn,
	}
	mrpc.HandleConnection(r.Context(), transport_conn, unpackSetting)
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetPort(5000)

			s.SetSessionIdName("token")
			// s.SetSessionCookieOutput(true)
			// HOOK, 开发阶段禁止浏览器缓存,方便调试
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}
			s.Use(ghttp.MiddlewareHandlerResponse)
			// s.BindHandler("/test", test)
			s.Group("/", func(group *ghttp.RouterGroup) {
				// custom group middlewares
				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareCORS,
				)
				// Register route handlers.
				var (
					userCtrl    = user.NewV1()
					systemCtrl  = system.NewV1()
					adminCtrl   = admin.NewV1()
					testApiCtrl = test_api.NewV1()
				)
				group.Bind(
					userCtrl,
					systemCtrl,
					// adminCtrl,   //危险行为 TODO 仅编码测试
					testApiCtrl, //test api
				)
				//Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/user/info":       userCtrl.UserInfo,
						"/system/menu/get": systemCtrl.GetMenu,
					})
				})
				adr := []string{consts.PERMISSION_btn_add,
					consts.PERMISSION_btn_edit,
					consts.PERMISSION_btn_del,
					consts.PERMISSION_btn_link,
					//partform permissions
					consts.PERMISSION_PARTFORM_ADMIN,
				}
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(func(r *ghttp.Request) {
						service.Middleware().AuthPath(r,
							&adr)
					})
					// group.ALLMap(g.Map{
					// 	"/admin/user/list": adminCtrl.UserList,
					// })
					group.Bind(adminCtrl) //平台管理员
				})
			})
			s.BindHandler("/api/ws", wsHandler) //websocket
			// Custom enhance API document.
			enhanceOpenAPIDoc(s)
			InitSystem()
			s.EnableHTTPS("server.pem","server.key")
			// Just run the server.
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "GoFrame",
			URL:  "https://goframe.org",
		},
	}
}

// func(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("WebSocket 连接失败:", err)
// 		return
// 	}
// 	// defer conn.Close()
// 	// log.Println("WebSocket 连接建立")
// 	transport_conn := TransportConn{
// 		TransportStack: "websocket",
// 		WsConn:         conn,
// 	}
// 	handleConnection(transport_conn, unpackSetting)
// }

func InitSystem() {
	cache.SetAdapter(context.Background())
}
