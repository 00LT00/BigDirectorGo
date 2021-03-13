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
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=string} "openID"
// @Failure 500 {object} utils.FailureResponse "error request"
// @Router /user/openID [get]
func OpenID(c *gin.Context) interface{} {
	code := c.Query("code")
	if code == "" {
		panic(error2.NewHttpError(404, "40401", "code null"))
	}
	return getOpenID(code)
}
