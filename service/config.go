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
	return conf
}
