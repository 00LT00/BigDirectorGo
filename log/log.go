package logger

import (
	error2 "BigDirector/error"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var (
	ErrLog  *log.Logger
	InfoLog *log.Logger
	//开发模式
	Mode = flag.Bool("mode", false, "dev mode")
)

func init() {
	flag.Parse()
	if *Mode {
		ErrLog = log.New(os.Stderr, "[ERROR]", log.LstdFlags|log.Llongfile)
		InfoLog = log.New(os.Stdout, "[INFO]", log.LstdFlags|log.Llongfile)
	} else {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			panic(error2.NewError(err.Error(), ""))
		}
		timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
		FilePath := filepath.Join(dir, "log", timeStamp)

		err = os.MkdirAll(FilePath, os.ModePerm)
		switch runtime.GOOS {
		case "windows":
			break
		case "linux":
			_ = os.Chmod(FilePath, os.ModePerm)
		}

		if err != nil {
			fmt.Println(err.Error())
			panic(error2.NewError(err.Error(), ""))
		}

		//创建日志文件
		ErrLogFile, err := os.Create(filepath.Join(FilePath, "error.log"))
		defer ErrLogFile.Close()
		LogFile, err := os.Create(filepath.Join(FilePath, "info.log"))
		defer LogFile.Close()

		ErrLog = log.New(ErrLogFile, "[ERROR]", log.LstdFlags|log.Llongfile)
		InfoLog = log.New(LogFile, "[INFO]", log.LstdFlags|log.Llongfile)
	}

	InfoLog.Println("Info logger init successful")
	ErrLog.Println("Error logger init successful")

}
