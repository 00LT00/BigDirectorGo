package gin

import (
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
	}
	//performanceR :=r.Group("/performance")
	//{
	//	performanceR.PUT("/info",f())
	//	performanceR.GET("/info",f())
	//}

}

func f(m func(c *gin.Context) interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := m(c)
		c.JSON(utils.MakeSuccessJSON(data))
	}
}
