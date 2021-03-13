package service

import (
	logger "BigDirector/log"
	"os/exec"
)

func initSwagger() {
	cmd := exec.Command("swag", "init","--parseDependency","-g","./cmd/main.go")
	cmd.Stdout = logger.InfoLog.Writer()
	cmd.Stderr = logger.InfoLog.Writer()
	err := cmd.Start()
	err = cmd.Wait()
	if err != nil {
		logger.ErrLog.Println(err)
	}
}
