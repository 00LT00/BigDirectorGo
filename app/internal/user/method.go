package user

import (
	"BigDirector/app/database"
	error2 "BigDirector/error"
	"BigDirector/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

var s = service.Service

// 获取openID
// @Tags user
// @Summary get openID
// @Description get openID from Weixin
// @ID get-openID
// @Produce  json
// @Param code query string true "wx.login()获取的code"
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=string} "openID"
// @Failure 400 {object} utils.FailureResponse "code null"
// @Failure 500 {object} utils.FailureResponse "service error"
// @Router /user/openID [get]
func OpenID(c *gin.Context) interface{} {
	code := c.Query("code")
	if code == "" {
		panic(error2.NewHttpError(400, "40001", "code null"))
	}
	return getOpenID(code)
}

// 新建或更新用户信息
// @Tags user
// @Summary create or update user information
// @Description create or update user information
// @ID set-Info
// @Accept json
// @Produce  json
// @Param openID body database.User true "用户的openID" "openID"
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=string} "success"
// @Failure 400 {object} utils.FailureResponse "40001 param error"
// @Failure 500 {object} utils.FailureResponse "service error"
// @Router /user/info [put]
func SetInfo(c *gin.Context) interface{} {
	u := new(database.User)
	err := c.ShouldBindJSON(u)
	if err != nil {
		panic(error2.NewHttpError(400, "40001", err.Error()))
	}
	if err := s.DB.Clauses(clause.OnConflict{ //冲突时更新除主键外的所有值
		UpdateAll: true,
	}).Create(u).Error; err != nil {
		panic(err.Error())
	}
	return "success"
}
