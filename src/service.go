package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type Service struct {
	DB     *gorm.DB
	Router *gin.Engine
	Conf   conf
}

type conf struct {
	Server struct {
		Port string
		Key  string
	}
	DB struct {
		Addr   string
		User   string
		Pass   string
		DBName string
	}
	DBDev struct {
		Addr   string
		User   string
		Pass   string
		DBName string
	}
}

func (s *Service) init() {
	s.initConfig()
	s.initDB()
	s.initRouter()
}

func (s *Service) initConfig() {
	c := new(conf)
	_, err := toml.DecodeFile("./config/config_BigDirector.toml", c)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	s.Conf = *c
}

func (s *Service) initDB() {
	var db *gorm.DB
	var err error
	if os.Getenv("ZERO_PROD") == "ZEROKIRIN" {
		db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.Conf.DB.User, s.Conf.DB.Pass, s.Conf.DB.Addr, s.Conf.DB.DBName))
	} else {
		db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.Conf.DBDev.User, s.Conf.DBDev.Pass, s.Conf.DBDev.Addr, s.Conf.DBDev.DBName))
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("success connect to DB")
	//自动建表
	db.AutoMigrate(&User{})

	s.DB = db
	//debug模式
	s.DB = s.DB.Debug()

}

func (s *Service) makeErrJSON(httpStatusCode int, errCode int, msg interface{}) (int, interface{}) {
	return httpStatusCode, &gin.H{"error": errCode, "msg": fmt.Sprint(msg)}
}

func (s *Service) makeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, &gin.H{"error": 0, "msg": "success", "data": data}
}
