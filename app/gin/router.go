package gin

import (
	"BigDirector/app/internal/performance"
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

}

func f(m func(c *gin.Context) interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := m(c)
		c.JSON(utils.MakeSuccessJSON(data))
	}
}
