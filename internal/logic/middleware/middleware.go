package middleware

import (
	"context"
	// "gfAdmin/internal/logic/session"
	"gfAdmin/internal/model"
	"gfAdmin/internal/service"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	// "github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gsession"
)

type (
	sMiddleware struct{}
)

var manager_ *model.MySessionManager //自定义session Manager
var logger = g.Log("middleware")

func NewSessionManager(m *gsession.Manager, ctx *context.Context) *model.MySessionManager {
	return &model.MySessionManager{
		Manager: m,
		Ctx:     ctx,
	}
}

func init() {
	service.RegisterMiddleware(New())
	//init MySessionManager
	g.Log("[init]").Info(context.Background(), "init MySessionManager")
	manager := gsession.New(time.Second * 180)
	manager.SetStorage(gsession.NewStorageMemory())
	ctx := gctx.New()
	manager_ = NewSessionManager(manager, &ctx)
}

func (s *sMiddleware) GetSessionManager(ctx context.Context) *model.MySessionManager {
	return manager_
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

// GetSessionId retrieves and returns session id from cookie or header.
// func (r *Request) GetSessionId() string {
// 	id := r.Cookie.GetSessionId()
// 	if id == "" {
// 		id = r.Header.Get(r.Server.GetSessionIdName())
// 	}
// 	return id
// }

// Ctx injects custom business context variable into context of current request.
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	this_session := manager_.Manager.New(*manager_.Ctx, r.GetSessionId())
	customCtx := &model.Context{
		Session: this_session, //r.Session, //auto inject session
		Request: r,            //auto inject request
	}
	sid, _ := this_session.Id() //r.Session.Id()

	logger.Info(r.Context(), "Session ID: ", sid)
	service.BizCtx().Init(r, customCtx)
	if user := service.Session().GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{ //拿出来 user 信息
			Id:       user.Id,
			Passport: user.Passport,
			Nickname: user.Nickname,
			Role:     user.Role,
			Avatar:   user.Avatar,
			Email:    user.Email,
			Phone:    user.Phone,
			Status:   user.Status,
			//更新角色信息
			RoleName: user.Roles[0],
			Auths:    user.Auths,
		}
		//更新登录信息到缓存 (redis or cache)
		//cache
		//过期时间:
		// this_session.
	} else {
		g.Log().Info(r.Context(), "GetUser is nil")
	}
	// Continue execution of next middleware.
	r.Middleware.Next()
	//mysession.Close()...
	logger.Info(context.Background(), "Session ", sid, " Close()")
	this_session.Close()
}

// 1.Auth 对请求进行验证，只允许已登录的用户访问
// 2.CORS 允许跨域资源共享
// 3.对path进行鉴权
func (s *sMiddleware) Auth(r *ghttp.Request) {
	if service.User().IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

func (s *sMiddleware) AuthPath(r *ghttp.Request, need_permissions *[]string) {
	if service.User().HavedAllPermissions(r.Context(),need_permissions) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
