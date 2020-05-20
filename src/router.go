package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Service) initRouter() {
	r := gin.Default()
	//校验码和cors头
	r.Use(cors.Default(), s.Check())

	//用户组路由
	user := r.Group("/user")
	//创建
	user.PUT("/", func(c *gin.Context) {
		c.JSON(s.Registered(c))
	})
	//获取
	user.GET("/:userid", func(c *gin.Context) {
		c.JSON(s.GetUser(c))
	})
	//修改（小程序无法使用PATCH）
	user.PUT("/:userid", func(c *gin.Context) {
		c.JSON(s.UpdateUser(c))
	})
	//获取项目列表
	user.GET("/:userid/*project", func(c *gin.Context) {
		c.JSON(s.GetUserProject(c))
	})

	//项目组路由
	project := r.Group("/project")
	//创建
	project.PUT("/", func(c *gin.Context) {
		c.JSON(s.AddProject(c))
	})
	//获取详情
	project.GET("/:projectid/*userid", func(c *gin.Context) {
		c.JSON(s.GetProject(c))
	})
	//增加成员
	project.POST("/member/", func(c *gin.Context) {
		c.JSON(s.AddMember(c))
	})
	// 获取项目的用户
	project.GET("/", func(c *gin.Context) {
		c.JSON(s.GetProjectUser(c))
	})
	//更改
	project.PUT("/:projectid/*userid", func(c *gin.Context) {
		c.JSON(s.UpdateProject(c))
	})

	s.Router = r
	err := s.Router.Run(s.Conf.Server.Port)
	panic(err)
}
