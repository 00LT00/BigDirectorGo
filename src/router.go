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
	//删除项目中的用户
	project.POST("/delete/:projectid", func(c *gin.Context) {
		c.JSON(s.DeleteProjectUser(c))
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
	//删除项目
	project.DELETE("/:projectid", func(c *gin.Context) {
		c.JSON(s.DeleteProject(c))
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
	//设置负责人
	process.POST("/:userid", func(c *gin.Context) {
		c.JSON(s.SetManager(c))
	})
	////修改
	//process.POST("/:userid", func(c *gin.Context) {
	//	c.JSON(s.UpdateProcess(c))
	//})

	//灯光，音效，后台，道具
	worker := r.Group("/worker")
	//设置除了环节以外的所有负责人
	worker.PUT("/:userid", func(c *gin.Context) {
		c.JSON(s.SetWorker(c))
	})
	//查看项目的所有相关负责人,包括环节负责人
	worker.GET("/", func(c *gin.Context) {
		c.JSON(s.GetWorker(c))
	})

	//项目的状态
	status := r.Group("/status")
	//设置状态
	status.POST("/", func(c *gin.Context) {
		c.JSON(s.SetProjectStatus(c))
	})
	//获取状态
	status.GET("/", func(c *gin.Context) {
		c.JSON(s.GetProjectStatus(c))
	})

	//推送服务的路由组
	send := r.Group("/send")
	send.POST("/", func(c *gin.Context) {
		c.JSON(s.ActionStart(c))
	})

	//文件服务
	file := r.Group("/file")
	//静态图片
	file.GET("/picture/:filename", s.GetPicture)
	//二维码
	file.GET("/wxacode/:projectid", func(c *gin.Context) {
		c.JSON(s.GetWxacodeBuffer(c))
	})
	file.POST("/excel", func(c *gin.Context) {
		c.JSON(s.GetExcel(c))
	})

	/*测试区*/
	//fmt.Println(s.GetOpenID("043VcuII1MHmF30qFcGI1YM5II1VcuI3"))
	//fmt.Println(s.GetToken())
	/**/

	s.Router = r
	err := s.Router.Run(s.Conf.Server.Port)
	panic(err)
}
