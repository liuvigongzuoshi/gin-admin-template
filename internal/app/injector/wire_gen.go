// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"gin-admin-template/internal/app/api"
	"gin-admin-template/internal/app/bll/impl/bll"
	"gin-admin-template/internal/app/model/impl/gorm/model"
	"gin-admin-template/internal/app/module/adapter"
	"gin-admin-template/internal/app/router"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	auther, cleanup, err := InitAuth()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := InitGormDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	resource := &model.Resource{
		DB: db,
	}
	role := &model.Role{
		DB: db,
	}
	roleMenu := &model.RoleMenu{
		DB: db,
	}
	menuActionResource := &model.MenuActionResource{
		DB: db,
	}
	menuResource := &model.MenuResource{
		DB: db,
	}
	user := &model.User{
		DB: db,
	}
	userRole := &model.UserRole{
		DB: db,
	}
	casbinAdapter := &adapter.CasbinAdapter{
		ResourceModel:           resource,
		RoleModel:               role,
		RoleMenuModel:           roleMenu,
		MenuActionResourceModel: menuActionResource,
		MenuResourceModel:       menuResource,
		UserModel:               user,
		UserRoleModel:           userRole,
	}
	syncedEnforcer, cleanup3, err := InitCasbin(casbinAdapter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	demo := &model.Demo{
		DB: db,
	}
	bllDemo := &bll.Demo{
		DemoModel: demo,
	}
	apiDemo := &api.Demo{
		DemoBll: bllDemo,
	}
	mock := &api.Mock{}
	menu := &model.Menu{
		DB: db,
	}
	menuAction := &model.MenuAction{
		DB: db,
	}
	login := &bll.Login{
		Auth:            auther,
		UserModel:       user,
		UserRoleModel:   userRole,
		RoleModel:       role,
		RoleMenuModel:   roleMenu,
		MenuModel:       menu,
		MenuActionModel: menuAction,
	}
	apiLogin := &api.Login{
		LoginBll: login,
	}
	trans := &model.Trans{
		DB: db,
	}
	bllMenu := &bll.Menu{
		TransModel:              trans,
		MenuModel:               menu,
		MenuResourceModel:       menuResource,
		MenuActionModel:         menuAction,
		MenuActionResourceModel: menuActionResource,
		ResourceModel:           resource,
	}
	apiMenu := &api.Menu{
		MenuBll: bllMenu,
	}
	bllRole := &bll.Role{
		Enforcer:      syncedEnforcer,
		TransModel:    trans,
		RoleModel:     role,
		RoleMenuModel: roleMenu,
		UserModel:     user,
	}
	apiRole := &api.Role{
		RoleBll: bllRole,
	}
	bllUser := &bll.User{
		Enforcer:      syncedEnforcer,
		TransModel:    trans,
		UserModel:     user,
		UserRoleModel: userRole,
		RoleModel:     role,
	}
	apiUser := &api.User{
		UserBll: bllUser,
	}
	sys := &api.Sys{}
	bllResource := &bll.Resource{
		Enforcer:      syncedEnforcer,
		ResourceModel: resource,
	}
	apiResource := &api.Resource{
		ResourceBll: bllResource,
	}
	routerRouter := &router.Router{
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		DemoAPI:        apiDemo,
		MockAPI:        mock,
		LoginAPI:       apiLogin,
		MenuAPI:        apiMenu,
		RoleAPI:        apiRole,
		UserAPI:        apiUser,
		SysAPI:         sys,
		ResourceAPI:    apiResource,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine:         engine,
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		MenuBll:        bllMenu,
	}
	return injector, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
