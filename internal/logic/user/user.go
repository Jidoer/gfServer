package user

import (
	"context"
	"gfAdmin/internal/cache"
	"gfAdmin/internal/dao"
	"gfAdmin/internal/model"
	"gfAdmin/internal/model/do"
	"gfAdmin/internal/model/entity"
	"gfAdmin/internal/service"
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	// "github.com/gogf/gf/v2/os/gcache"
	uuid "github.com/satori/go.uuid"
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

func (s *sUser) AutoCreate(ctx context.Context, phone string) (token string, user *model.User_role_db, err error) {
	auto_input := model.UserCreateInput{
		Passport: service.User().GetRandomPassport(),
		Phone:    phone,
		Password: service.User().GetRandomPassword(),
		Nickname: service.User().GetRandomNickname(),
	}
	var (
		available    bool
		us           model.LoginInfo_last
		permissions  []string
		user_session *model.User_Session
		u            *gvar.Var
	)

	// 检查手机号是否已注册
	//sign in
	// var user *model.User_role_db
	err = g.Model(model.User_role_db{}).Where(do.User{
		Phone: auto_input.Phone,
	}).WithAll().Scan(&user)
	if err != nil {
		logger.Error(ctx, "检查手机号是否已注册 error: %v", err)
		return
	}
	if user != nil {
		logger.Print(ctx, "手机号已注册 -->goto 登录")
		goto AutoRet
	}

	available, err = s.CheckPassport(ctx, auto_input.Passport)
	if !available {
		err = gerror.Newf(`Passport "%s" is already token by others`, auto_input.Passport)
	}
	if err != nil {
		return
	}
	err = dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Passport: auto_input.Passport,
			Password: auto_input.Password,
			Nickname: auto_input.Nickname,
			Phone:    auto_input.Phone,
		}).Insert()
		return err
	})
	if err != nil {
		return
	}
	//sign in
	//var user *model.User_role_db
	err = g.Model(model.User_role_db{}).Where(do.User{
		Passport: auto_input.Passport,
		Password: auto_input.Password,
	}).WithAll().Scan(&user)
	if err != nil {
		return
	}
	if user == nil {
		err = gerror.New(`Passport or Password not correct`)
		return
	}

AutoRet:
	user_session = &model.User_Session{
		Passport: user.Passport,
		Nickname: user.Nickname,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Status:   user.Status,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
	for _, v := range user.Roles.RolePermissions {
		permissions = append(permissions, v.Permissions.Name)
	}
	// if len(user.Roles.RolePermissions) > 0 {
	user_session.Roles = make([]string, 1)
	user_session.Roles[0] = user.Roles.Name
	user_session.Auths = permissions
	if err = service.Session().SetUser(ctx, user_session); err != nil {
		return
	}
	// var us model.LoginInfo_last

	u, _ = cache.Instance().Get(context.Background(), "user_"+user.Passport)
	// u := cache.MustGet("user_" + user.Passport)
	if err = u.Scan(&us); err != nil {
		//panic(err)
		logger.Error(context.Background(), "Get cache error: %v", err)
	}

	token, _ = service.BizCtx().Get(ctx).Session.Id()
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
	return
}

func (s *sUser) GetRandomPassport() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}
func (s *sUser) GetRandomPassword() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}
func (s *sUser) GetRandomNickname() string {
	prefixes := []string{"炫酷", "狂野", "疾风", "暗影", "魔法", "钢铁", "疾速", "雷霆"}
	names := []string{"战士", "刺客", "猎人", "法师", "骑士", "领主", "剑魂", "影刃"}
	suffixes := []string{"001", "X", "Pro", "王者", "之魂", "大帝", "", "", ""} // 有些后缀为空，模拟可选项
	rand.Seed(time.Now().UnixNano())
	prefix := prefixes[rand.Intn(len(prefixes))]
	name := names[rand.Intn(len(names))]
	suffix := suffixes[rand.Intn(len(suffixes))]
	return prefix + name + suffix
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
	// return token, nil
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

func (s *sUser) CheckPhone(ctx context.Context, phone string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Phone: phone,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	// count, err := dao.User.Ctx(ctx).Where(do.User{
	// 	Nickname: nickname,
	// }).Count()
	// if err != nil {
	// 	return false, err
	// }
	// return count == 0, nil
	return true, nil // always available
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
