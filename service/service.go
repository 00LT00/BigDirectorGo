package service

import (
	casbin2 "github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type service struct {
	DB     *gorm.DB
	Casbin *casbin2.Enforcer
	Router *gin.Engine
}

func initService() *service {
	s := new(service)
	s.DB = initDB()
	s.Casbin = initCasbin()
	return s
}
