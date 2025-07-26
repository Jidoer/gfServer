package user

import (
	"context"
	"gfAdmin/internal/cache"
	"gfAdmin/internal/dao"
	"gfAdmin/internal/model"
	"gfAdmin/internal/model/do"
	"gfAdmin/internal/model/entity"
	"gfAdmin/internal/service"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/gogf/gf/v2/os/gcache"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

var logger = g.Log("user")

// Create creates user account.
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}
	var (
		available bool
	)
	// Passport checks.
	available, err = s.CheckPassport(ctx, in.Passport)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
	}
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Passport: in.Passport,
			Password: in.Password,
			Nickname: in.Nickname,
		}).Insert()
		return err
	})
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// SignIn creates session for given user account.
func (s *sUser) SignIn(ctx context.Context, in model.UserSignInInput) (err error) {
	//var user *entity.User
	var user *model.User_role_db
	err = g.Model(model.User_role_db{}).Where(do.User{
		Passport: in.Passport,
		Password: in.Password,
	}).WithAll().Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New(`Passport or Password not correct`)
	}
	var user_session model.User_Session
	var permissions []string
	for _, v := range user.Roles.RolePermissions {
		permissions = append(permissions, v.Permissions.Name)
	}
	// if len(user.Roles.RolePermissions) > 0 {
	user_session.Roles = make([]string, 1)
	user_session.Roles[0] = user.Roles.Name
	user_session.Auths = permissions
	if err = service.Session().SetUser(ctx, &user_session); err != nil {
		return err
	}
	var us model.LoginInfo_last
	u, _ := cache.Instance().Get(context.Background(), "user_"+user.Passport)
	// u := cache.MustGet("user_" + user.Passport)
	if err = u.Scan(&us); err != nil {
		//panic(err)
		logger.Error(context.Background(), "Get cache error: %v", err)
	}
	token, _ := service.BizCtx().Get(ctx).Session.Id()
	us = model.LoginInfo_last{
		Id:        user.Id,
		LoginTime: time.Now(),
		Passport:  user.Passport,
		// Ip: r,
		// Device: ,
		Token: token,
	}
	err = cache.Instance().Set(context.Background(), "user_"+user.Passport, us, time.Second*180)
	// err = cache.CacheSet("user_"+user.Passport, us, time.Second*180)
	if err != nil {
		logger.Error(context.Background(), "Set cache error: %v", err)
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
		Role:     user.Role,
		Status:   user.Status,
		// SessionID: service.Session().GetSessionId(ctx),
	})
	return nil
}

// SignOut removes the session for current signed-in user.
func (s *sUser) SignOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

// IsPassportAvailable checks and returns given passport is available for signing up.
func (s *sUser) CheckPassport(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Nickname: nickname,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetProfile retrieves and returns current user info in session.
//
//	func (s *sUser) GetProfile(ctx context.Context) *entity.User {
//		return service.Session().GetUser(ctx)
//	}
func (s *sUser) GetProfile(ctx context.Context) *model.User_Session {
	return service.Session().GetUser(ctx)
}

// func (s *sUser) GetUserInfo(ctx context.Context) *model.User_role_db{

// }

func (s *sUser) GetUsersList(ctx context.Context, page int, limit int) (*[]entity.User, int, error) {
	var total int
	var users []entity.User
	err := dao.User.Ctx(ctx).Where(do.User{}).Limit(limit).Offset((page-1)*limit).ScanAndCount(&users, &total, false)
	return &users, total, err
}

func (s *sUser) HavedAllPermissions(ctx context.Context, need_permissions *[]string) bool {
	user := service.BizCtx().Get(ctx).User
	if user == nil || len(user.Auths) == 0 {
		return false
	}
	// if(need_permissions == nil){
	// 	return false
	// }
	for _, v := range *need_permissions {
		haved := false
		for _, v2 := range user.Auths {
			if v2 == v {
				haved = true
				break
			}
		}
		if !haved {
			return false
		}
	}
	return true
}
