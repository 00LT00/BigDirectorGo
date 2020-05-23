package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type Service struct {
	DB     *gorm.DB
	Router *gin.Engine
	Conf   conf
	Redis  *redis.Client
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
	Redis struct {
		Addr string
		Pass string
		DB   int
	}
	Wx struct {
		AppID     string
		AppSecret string
	}
}

func (s *Service) init() {
	s.initConfig()
	s.initDB()
	s.initRedis()
	s.initRouter()
}

func (s *Service) initConfig() {
	c := new(conf)
	_, err := toml.DecodeFile("./config/config_BigDirector.toml", c)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// 	加载小程序id和认证密钥
	_, err = toml.DecodeFile("./config/weixin.toml", c)
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
	db.AutoMigrate(&User{}, &Project{}, &Project_User{})

	s.DB = db
	//debug模式
	s.DB = s.DB.Debug()

}

func (s *Service) initRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     s.Conf.Redis.Addr,
		Password: s.Conf.Redis.Pass,
		DB:       s.Conf.Redis.DB, // use default DB
	})
	ctx := client.Context()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis", err.Error())
		panic(err)
	}
	//fmt.Println(pong, err)
	// Output: PONG <nil>
	s.Redis = client
}

func (s *Service) makeErrJSON(httpStatusCode int, errCode int, msg interface{}) (int, interface{}) {
	return httpStatusCode, &gin.H{"error": errCode, "msg": fmt.Sprint(msg)}
}

func (s *Service) makeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, &gin.H{"error": 0, "msg": "success", "data": data}
}
