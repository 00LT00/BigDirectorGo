package gin

import (
	_ "BigDirector/docs"
	logger "BigDirector/log"
	"BigDirector/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	s    = service.Service
	conf = service.Conf
)

func Register() {

	gin.SetMode(gin.TestMode)
	r := gin.New()
	//日志收集
	Log := gin.LoggerWithWriter(logger.InfoLog.Writer())
	Recover := gin.RecoveryWithWriter(logger.ErrLog.Writer())

	//注册swagger
	swagger := r.Group("/doc")
	swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//ping
	r.GET("/ping", f(ping))

	r.Use(Log, Recover, HttpRecover, cors.Default())
	//启用校验
	if !*logger.Mode {
		r.Use(check)
	}
	initRouter(r)

	s.Router = r
}
