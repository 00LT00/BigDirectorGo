package service

import (
	error2 "BigDirector/error"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

type database struct {
	Addr   string
	Port   string
	User   string
	Pass   string
	DBName string
}

type casbin struct {
	Model string
}

type server struct {
	Port string
	Sign string
}

type config struct {
	DB     database
	DBDev  database
	Casbin casbin
	Server server
	Wx     weixin
}

type weixin struct {
	AppID     string
	AppSecret string
}

func initConfig() *config {
	dir, err := os.Getwd()
	conf := new(config)
	if err != nil {
		panic(error2.NewError(err.Error(), ""))
	}
	_, err = toml.DecodeFile(filepath.Join(dir, *configFilePath), conf)
	if err != nil {
		panic(error2.NewError(err.Error(), ""))
	}
	wx := new(struct { //toml第一层结构体不被解析
		WeiXin weixin
	})
	wxAppID := os.Getenv("WX_APPID")
	wxAppSecret := os.Getenv("WX_APPSECRET")
	if wxAppID != "" && wxAppSecret != "" {
		wx.WeiXin.AppID = wxAppID
		wx.WeiXin.AppSecret = wxAppSecret
	} else {
		_, err := toml.DecodeFile(filepath.Join(dir, *weixinConfigFilePath), wx)
		if err != nil {
			panic(error2.NewError(err.Error(), ""))
		}
	}
	conf.Wx = wx.WeiXin
	return conf
}
