package service

import (
	error2 "BigDirector/error"
	"fmt"
	casbin2 "github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
)

func initCasbin() *casbin2.Enforcer {
	db := getDB()
	//gorm-adapter 历史遗留问题
	a, err := gormadapter.NewAdapter("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/", db.User, db.Pass, db.Port), db.DBName)
	if err != nil {
		panic(error2.NewError(err.Error(), ""))
	}
	e, err := casbin2.NewEnforcer(Conf.Casbin.Model, a)
	if err != nil {
		panic(error2.NewError(err.Error(), ""))
	}
	return e
}
