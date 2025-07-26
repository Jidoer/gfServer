package permissions

import (
	"context"
	"time"
	// "time"

	// "gfAdmin/internal/dao"
	"gfAdmin/internal/cache"
	"gfAdmin/internal/model"
	"gfAdmin/internal/model/do"
	// "gfAdmin/internal/model/entity"

	// "gfAdmin/internal/model/do"
	// "gfAdmin/internal/model/entity"
	"gfAdmin/internal/service"

	// "github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var logger = g.Log("permissions")

type (
	sPermissions struct{}
)

func init() {
	service.RegisterPermissions(New())
}

func New() service.IPermissions {
	return &sPermissions{}
}

func (s *sPermissions) GetRolesList(ctx context.Context, page int, limit int) (*[]model.Roles, int, error) {
	var total int
	// var roles []entity.Roles
	// err := dao.Roles.Ctx(ctx).Where(do.Roles{}).Limit(limit).Offset((page-1)*limit).ScanAndCount(&roles, &total, false)
	// dao.Roles.Ctx(ctx).Where()
	var roles []model.Roles
	err := g.Model(model.Roles{}).WithAll().Where(do.Roles{}).ScanAndCount(&roles, &total, false) //Scan(&roles)

	return &roles, total, err
}

func (s *sPermissions) RolesDelete(ctx context.Context, role model.Roles) error {
	//暂时只需要id
	_, err := g.Model(model.Roles{}).Delete("role_sign=?", role.RoleSign) //通过role_sign
	return err
}

func (s *sPermissions) RoleCreate(ctx context.Context, role do.Roles) (int64, error) {
	//根据id更新
	m, err := g.Model(do.Roles{}).Insert(role)
	if err != nil {
		return 0, err
	}
	id, _ := m.LastInsertId()
	return id, err
}
func (s *sPermissions) RoleUpdate(ctx context.Context, role do.Roles) error {
	//根据id更新
	_, err := g.Model(do.Roles{}).Where("id=?", role.Id).Update(role) //暂时不用Save
	return err
}

// 更新关系表
func (s *sPermissions) RoleUpdateRelation(ctx context.Context, role do.Roles, role_permissions []do.RolePermissions) error {
	//根据id更新
	_, err := g.Model(do.RolePermissions{}).Where("role_id=?", role.Id).Delete()
	if err != nil {
		return err
	}
	if len(role_permissions) == 0 {
		return nil
	}
	_, err = g.Model(do.RolePermissions{}).Insert(role_permissions)
	return err
}

func (s *sPermissions) GetPermissionsList(ctx context.Context, page int, limit int) (*[]model.Permissions, int, error) {
	var total int
	var permissions []model.Permissions
	err := g.Model(model.Permissions{}).Where(do.Permissions{}).ScanAndCount(&permissions, &total, false) //Scan(&roles)
	return &permissions, total, err
}

func (s *sPermissions) AddPermission(ctx context.Context, permission model.Permissions) error {
	_, err := g.Model(model.Permissions{}).Insert(permission)
	return err
}

func (s *sPermissions) DeletePermission(ctx context.Context, permission model.Permissions) error {
	//暂时只需要id
	_, err := g.Model(model.Permissions{}).Delete("id=?", permission.Id) //(permission)
	return err
}

func (s *sPermissions) UpdatePermission(ctx context.Context, permission model.Permissions) error {
	//根据id更新
	_, err := g.Model(do.Permissions{}).Where("id=?", permission.Id).Update(permission)
	return err
}

func (s *sPermissions) GetRoleAndPermissions(ctx context.Context, role_id int) (string, *[]string, error) {
	var role *model.Roles
	r, err := cache.Instance().Get(ctx, "role_"+gconv.String(role_id))
	if err != nil {
		return "", nil, gerror.New("获取缓存失败！")
	}
	r.Struct(&role)
	if role != nil {
		goto ret
	}
	err = g.Model(model.Roles{}).Where(do.Roles{RolesID: role_id}).WithAll().Scan(&role)
	if err != nil {
		return "", nil, err
	}
	if role == nil {
		return "", nil, gerror.New("角色不存在！")
	}
	cache.Instance().Set(ctx, "role_"+gconv.String(role_id), role, time.Hour*24)
ret:
	var permissions_ []string
	for _, v := range role.RolePermissions {
		permissions_ = append(permissions_, v.Permissions.Name)
	}
	return role.Name, &permissions_, err
}
