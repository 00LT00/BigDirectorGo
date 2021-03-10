package service

import (
	error2 "BigDirector/error"
	logger "BigDirector/log"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := getDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(error2.NewError(err.Error(), ""))
	}
	return db.Debug()
}

func getDsn() string {
	db := getDB()
	return fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.User, db.Pass, db.Port, db.DBName)
}

func getDB() database {
	if *logger.Mode {
		return Conf.DBDev
	} else {
		return Conf.DB
	}
}
