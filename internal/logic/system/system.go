package menu

import (
	"context"
	"gfAdmin/internal/consts"
	"gfAdmin/internal/dao"
	"gfAdmin/internal/model"
	"time"

	// "gfAdmin/internal/model"

	// "gfAdmin/internal/model"
	"gfAdmin/internal/model/do"
	"gfAdmin/internal/model/entity"
	"gfAdmin/internal/service"

	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/database/gdb"
	// "github.com/gogf/gf/frame/g"
	// "github.com/gogf/gf/v2/net/ghttp"
)

type (
	sSystem struct{}
)

func init() {
	service.RegisterSystem(New())
}

func New() service.ISystem {
	return &sSystem{}
}

//just dev test web

func (s *sSystem) GetMenu(ctx context.Context) (*[]entity.SystemMenus, int, error) {
	routers := make([]entity.SystemMenus, 0)
	//with dbase...
	// return &consts.Menus
	count := 0
	err := dao.SystemMenus.Ctx(ctx).Cache(
		gdb.CacheOption{
			Name: "system_menu",
			Duration: time.Minute * 10,
		},
	).Where(do.SystemMenus{}).Scan(&routers)
	count = len(routers)
	//.ScanAndCount(&routers, &count, false)
	//Order(in.OrderBy, in.OrderDirection).
	//Limit(in.Offset, in.Limit).
	//ScanAndCount(&items, &total, false)
	// s ,err := g.Model("SystemMenus")6
	if err != nil {
		return nil, 0, err
	}
	return &routers, count, nil
}

func (s *sSystem) GetMenuByte() []byte {
	return consts.Menus_byte
}

func (s *sSystem) SaveMenu(ctx context.Context, menu *entity.SystemMenus) error {
	_, err := dao.SystemMenus.Ctx(ctx).Save(do.SystemMenus{
		ParentId:    menu.ParentId,
		Path:        menu.Path,
		Name:        menu.Name,
		Component:   menu.Component,
		Redirect:    menu.Redirect,
		Title:       menu.Title,
		IsLink:      menu.IsLink,
		IsHide:      menu.IsHide,
		IsKeepAlive: menu.IsKeepAlive,
		IsAffix:     menu.IsAffix,
		IsIframe:    menu.IsIframe,
		Roles:       menu.Roles,
		Icon:        menu.Icon,
	})
	return err
}

func (s *sSystem) InitMenu(ctx context.Context) error {
	//队列
	q := gqueue.New()
	for _, menu := range consts.Menus {
		//添加/节点
		me := do.SystemMenus{
			ParentId:    0,
			Path:        menu.Path,
			Name:        menu.Name,
			Component:   menu.Component,
			Redirect:    menu.Redirect,
			Title:       menu.Meta.Title,
			IsLink:      menu.Meta.IsLink,
			IsHide:      menu.Meta.IsHide,
			IsKeepAlive: menu.Meta.IsKeepAlive,
			IsAffix:     menu.Meta.IsAffix,
			IsIframe:    menu.Meta.IsIframe,
			Roles:       menu.Meta.Roles,
			Icon:        menu.Meta.Icon,
		}
		d, err := dao.SystemMenus.Ctx(ctx).Save(me)
		menu.M_id, _ = d.LastInsertId()
		if err != nil {
			return err
		}
		q.Push(menu)
	}
	for q.Len() > 0 {
		m := q.Pop()
		kk, _ := m.(model.Menus)

		if len(kk.Children) > 0 {
			for _, v := range kk.Children {
				switched := do.SystemMenus{
					ParentId:    kk.M_id,
					Path:        v.Path,
					Name:        v.Name,
					Component:   v.Component,
					Redirect:    v.Redirect,
					Title:       v.Meta.Title,
					IsLink:      v.Meta.IsLink,
					IsHide:      v.Meta.IsHide,
					IsKeepAlive: v.Meta.IsKeepAlive,
					IsAffix:     v.Meta.IsAffix,
					IsIframe:    v.Meta.IsIframe,
					Roles:       v.Meta.Roles,
					Icon:        v.Meta.Icon,
				}
				lsatid, err := dao.SystemMenus.Ctx(ctx).Save(switched)
				v.M_id, _ = lsatid.LastInsertId()
				if err != nil {
					return err
				}
				q.Push(v)
			}
		}
	}
	return nil

}
