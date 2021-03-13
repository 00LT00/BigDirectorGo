package service

import (
	error2 "BigDirector/error"
	logger "BigDirector/log"
	"flag"
)

var (
	Conf    *config
	Service *service
)

var (
	//配置文件地址
	configFilePath       = flag.String("c", "config/config_BigDirector.toml", "config file")
	weixinConfigFilePath = flag.String("wx", "config/wx.toml", "config file")
)

func init() {
	defer func() {
		if r := recover(); r != nil {
			if r, ok := r.(error2.Error); ok {
				logger.ErrLog.Fatalln(r.Error())
			} else {
				logger.ErrLog.Fatalln(r.Error())
			}
		}
	}()

	Conf = initConfig()
	logger.InfoLog.Println("load config successful")
	Service = initService()
	logger.InfoLog.Println("load service successful")
	//initSwagger()

}
