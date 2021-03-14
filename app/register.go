package app

import (
	"BigDirector/app/database"
	"BigDirector/app/gin"
	logger "BigDirector/log"
	"BigDirector/service"
)

var (
	s    = service.Service
	conf = service.Conf
)

func Register() {
	database.Register()
	gin.Register()
}

func Run() {
	err := s.Router.Run(conf.Server.Port)
	logger.ErrLog.Println(err)
}
