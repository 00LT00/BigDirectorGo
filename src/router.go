package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Service) initRouter() {
	r := gin.Default()
	//校验码和cors头
	r.Use(cors.Default(),s.Check())

	//用户组路由
	user:= r.Group("/user")
	user.PUT("/", func(c *gin.Context) {
		c.JSON(s.Registered(c))
	})

	user.GET("/:userid", func(c *gin.Context) {
		c.JSON(s.GetUser(c))
	})

	user.PATCH("/:userid", func(c *gin.Context) {
		c.JSON(s.UpdateUser(c))
	})

	//
	////项目组路由
	//pjt:=r.Group("/project")
	//
	////环节组路由
	//pcs:=r.Group("process")


	s.Router = r
	err:=s.Router.Run(s.Conf.Server.Port)
	panic(err)
}
