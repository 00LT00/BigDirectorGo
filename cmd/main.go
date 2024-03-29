package main

import (
	"BigDirector/app"
	_ "BigDirector/error"
	logger "BigDirector/log"
	_ "BigDirector/service"
	_ "BigDirector/utils"
	"os"
)

// @title BigDirector API
// @version 2.0
// @description 我是大导演API

// @contact.name 00LT00
// @contact.url http://blog.zerokirin.online
// @contact.email lightning@zerokirin.online

// @host api.zerokirin.online
// @BasePath /BigDirector
func main() {

	app.Register()
	app.Run()
	code := <-logger.Exit
	os.Exit(code)
}
