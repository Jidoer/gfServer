// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gfAdmin/internal/model"
	"gfAdmin/internal/model/entity"
)

type (
	IUser interface {
		// Create creates user account.
		Create(ctx context.Context, in model.UserCreateInput) (err error)
		AutoCreate(ctx context.Context, phone string) (token string, user *model.User_role_db, err error)
		GetRandomPassport() string
		GetRandomPassword() string
		GetRandomNickname() string
		// IsSignedIn checks and returns whether current user is already signed-in.
		IsSignedIn(ctx context.Context) bool
		// SignIn creates session for given user account.
		SignIn(ctx context.Context, in model.UserSignInInput) (err error)
		// SignOut removes the session for current signed-in user.
		SignOut(ctx context.Context) error
		// IsPassportAvailable checks and returns given passport is available for signing up.
		CheckPassport(ctx context.Context, passport string) (bool, error)
		CheckPhone(ctx context.Context, phone string) (bool, error)
		// IsNicknameAvailable checks and returns given nickname is available for signing up.
		IsNicknameAvailable(ctx context.Context, nickname string) (bool, error)
		// GetProfile retrieves and returns current user info in session.
		//
		//	func (s *sUser) GetProfile(ctx context.Context) *entity.User {
		//		return service.Session().GetUser(ctx)
		//	}
		GetProfile(ctx context.Context) *model.User_Session
		GetUsersList(ctx context.Context, page int, limit int) (*[]entity.User, int, error)
		HavedAllPermissions(ctx context.Context, need_permissions *[]string) bool
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
