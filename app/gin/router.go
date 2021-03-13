package gin

import (
	"BigDirector/app/internal/user"
	"BigDirector/utils"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		// 获取openID
		// @Tags user
		// @Summary get openID
		// @Description get openID from Weixin
		// @ID get-openID
		// @Produce  json
		// @Param code query string true "wx.login()获取的code"
		// @Param sign header string true "spppk"
		// @Success 200 {object} utils.SuccessResponse{data=string} "openID"
		// @Failure 500 {object} utils.FailureResponse "error request"
		// @Router /user/openID [get]
		userRouter.GET("/openID", f(user.OpenID))

	}

}

func f(m func(c *gin.Context) interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := m(c)
		c.JSON(utils.MakeSuccessJSON(data))
	}
}
