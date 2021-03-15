package gin

import (
	"BigDirector/app/internal/group"
	"BigDirector/app/internal/performance"
	"BigDirector/app/internal/process"
	"BigDirector/app/internal/user"
	"BigDirector/utils"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	userR := r.Group("/user")
	{
		//获取openID
		userR.GET("/openID", f(user.OpenID))
		//更新用户详情
		userR.PUT("/info", f(user.SetInfo))
		//获取用户信息
		userR.GET("/info", f(user.GetInfo))
	}
	performanceR := r.Group("/performance")
	{
		//更新演出信息
		performanceR.PUT("/info", f(performance.SetInfo))
		//获取演出详情
		performanceR.GET("/info", f(performance.GetInfo))
	}
	groupR := r.Group("/group")
	{
		//创建小组
		groupR.PUT("/info", f(group.SetInfo))
		//获取小组详情
		groupR.GET("/info", f(group.GetInfo))
		//设置权限（roles）
		groupR.PUT("/roles", f(group.SetRoles))
	}
	processR := r.Group("/process")
	{
		processR.PUT("/info", f(process.SetInfo))
	}

}

func f(m func(c *gin.Context) interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := m(c)
		c.JSON(utils.MakeSuccessJSON(data))
	}
}
