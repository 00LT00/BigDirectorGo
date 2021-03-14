package database

import (
	logger "BigDirector/log"
	"BigDirector/service"
)

var s = service.Service

func Register() {
	err := s.DB.AutoMigrate(&Performance{}, &Process{}, &User{}, &Group{})
	if err != nil {
		logger.ErrLog.Fatalln(err.Error())
	}
}
