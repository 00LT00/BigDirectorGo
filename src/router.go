package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Service) initRouter() {
	r := gin.Default()
	//校验码和cors头
	r.Use(cors.Default(), s.Check())

	//获取用户的openid
	openid := r.Group("/openid")
	openid.GET("/:code", func(c *gin.Context) {
		c.JSON(s.OpenID(c))
	})

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
	project.GET("/", func(c *gin.Context) {
		c.JSON(s.GetProject(c))
	})
	//增加成员
	project.POST("/member/", func(c *gin.Context) {
		c.JSON(s.AddMember(c))
	})
	// 获取项目的用户
	project.GET("/user/", func(c *gin.Context) {
		c.JSON(s.GetProjectUser(c))
	})
	//更改项目名
	project.PUT("/pnm/:projectid", func(c *gin.Context) {
		c.JSON(s.UpdateProjectName(c))
	})
	//更改导演
	project.PUT("/uid/:projectid", func(c *gin.Context) {
		c.JSON(s.UpdateProjectUserid(c))
	})
	//查看项目环节
	project.GET("/process/:projectid", func(c *gin.Context) {
		c.JSON(s.GetProjectProcess(c))
	})

	// 环节路由
	process := r.Group("/process")
	////增加环节
	//process.PUT("/:userid", func(c *gin.Context) {
	//	c.JSON(s.AddProcess(c))
	//})
	//获取
	process.GET("/", func(c *gin.Context) {
		c.JSON(s.GetProcess(c))
	})

	process.PUT("/:userid", func(c *gin.Context) {
		c.JSON(s.UpdateProcess(c))
	})
	////修改
	//process.POST("/:userid", func(c *gin.Context) {
	//	c.JSON(s.UpdateProcess(c))
	//})

	//推送服务的路由组
	send := r.Group("/send")
	send.POST("/", func(c *gin.Context) {
		c.JSON(s.ActionStart(c))
	})

	//图片服务
	file := r.Group("/picture")
	//静态图片
	file.GET("/file/:filename", s.GetPicture)

	/*测试区*/
	//fmt.Println(s.GetOpenID("043VcuII1MHmF30qFcGI1YM5II1VcuI3"))
	//fmt.Println(s.GetToken())
	/**/

	s.Router = r
	err := s.Router.Run(s.Conf.Server.Port)
	panic(err)
}
