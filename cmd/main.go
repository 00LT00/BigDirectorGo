package main

import (
	"BigDirector/service"
	"fmt"
)

var conf = service.Conf

func main() {
	//logger.InfoLog.Println("hello")
	fmt.Println(conf)
}
