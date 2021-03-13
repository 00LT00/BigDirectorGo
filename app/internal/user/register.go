package user

import (
	error2 "BigDirector/error"
	"github.com/gin-gonic/gin"
)

// 获取openID
// @Tags user
// @Summary get openID
// @Description get openID from Weixin
// @ID get-openID
// @Produce  json
// @Param code query string true "wx.login()获取的code"
// @Success 200 {object} utils.SuccessResp
// @Failure 500 {string} ""
// @Router /user [get]
func OpenID(c *gin.Context) interface{} {
	code := c.Query("code")
	if code == "" {
		panic(error2.NewHttpError(404, "40401", "code null"))
	}
	return getOpenID(code)
}
