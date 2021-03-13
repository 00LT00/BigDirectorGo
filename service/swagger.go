package service

import (
	logger "BigDirector/log"
	"os/exec"
)

func initSwagger() {
	cmd := exec.Command("swag", "init","--parseDependency","-g","./main.go","-o","../docs")
	cmd.Stdout = logger.InfoLog.Writer()
	cmd.Stderr = logger.InfoLog.Writer()
	err := cmd.Run()
	if err != nil {
		logger.ErrLog.Println(err)
	}
}
