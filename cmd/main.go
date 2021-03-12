package main

import (
	"BigDirector/app"
)

//var conf = service.Conf

// @title BigDirector API
// @version 2.0
// @description 我是大导演API

// @contact.name 00LT00
// @contact.url http://blog.zerokirin.online
// @contact.email lightning@zerokirin.online

// @BasePath api.zerokirin.online/api/BigDirector

func main() {
	app.Register()
	app.Run()
}
