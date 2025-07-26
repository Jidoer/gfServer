package session

import (
	"context"
	"gfAdmin/internal/consts"
	"gfAdmin/internal/model"
	// "gfAdmin/internal/model/entity"

	// "gfAdmin/internal/model/entity"
	"gfAdmin/internal/service"

	// "gfAdmin/internal/vars"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sSession struct{}
)

var logger = g.Log("session")

func init() {
	service.RegisterSession(New())
}

func New() service.ISession {
	return &sSession{}
}

// SetUser sets user into the session.
func (s *sSession) SetUser(ctx context.Context, user *model.User_Session) error {
	return service.BizCtx().Get(ctx).Session.Set(consts.UserSessionKey, user)
}

// GetUser retrieves and returns the user from session.
// It returns nil if the user did not sign in.
func (s *sSession) GetUser(ctx context.Context) *model.User_Session {
	logger.Printf(ctx, "[GetUser] %+v", ctx)
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.MustGet(consts.UserSessionKey); !v.IsNil() { //!
			logger.Printf(ctx, "[GetUser] %+v", v)
			var user *model.User_Session //*entity.User //*entity.User
			_ = v.Struct(&user)
			// user_session := model.User_Session{
			// 	Id:       user.Id,
			// 	Passport: user.Passport,
			// 	Nickname: user.Nickname,
			// 	Avatar:   user.Avatar,
			// 	// Roles:    user.Roles,
			// 	// Auths:    user.Auths,
			// 	Email:    user.Email,
			// 	Phone:    user.Phone,
			// 	Status:   user.Status,
			// 	CreateAt: user.CreateAt,
			// 	UpdateAt: user.UpdateAt,
			// }
			//get user roles and auths from db (or cache)
			// name, permiss, err := service.Permissions().GetRoleAndPermissions(ctx, user.Role)
			// if err != nil {
			// 	logger.Error(ctx, "GetRoleAndPermissions error: %v", err)
			// 	return nil
			// }
			// user.Roles = make([]string, 1)
			// user.Roles[0] = name
			// user.Auths = *permiss
			return user
		}
	}
	return nil
}

// RemoveUser removes user rom session.
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.UserSessionKey)
	}
	return nil
}

// func (s *sSession) GetSessionManager() *ghttp.Session {

// }
