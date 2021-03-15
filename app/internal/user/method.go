package user

import (
	"BigDirector/app/database"
	error2 "BigDirector/error"
	"BigDirector/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

var db = service.Service.DB

// 获取openID
// @Tags user
// @Summary 获取openID
// @Description get openID from Weixin
// @ID get-OpenID
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
// @Summary 创建或更改用户信息
// @Description create or update user information
// @ID set-User-Info
// @Accept json
// @Produce  json
// @Param openID body database.User true "用户的openID"
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
	if err := db.Clauses(clause.OnConflict{ //冲突时更新除主键外的所有值
		UpdateAll: true,
	}).Create(u).Error; err != nil {
		panic(err.Error())
	}
	return "success"
}

// 获取用户信息
// @Tags user
// @Summary 获取用户信息
// @Description get user information by openID
// @ID get-User-Info
// @Produce  json
// @Param openID query string true "openID"
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=database.User} "UserInfo"
// @Failure 400 {object} utils.FailureResponse "openID null"
// @Failure 500 {object} utils.FailureResponse "service error"
// @Router /user/info [get]
func GetInfo(c *gin.Context) interface{} {
	openID := c.Query("openID")
	if openID == "" {
		panic(error2.NewHttpError(400, "40001", "openID null"))
	}
	u := new(database.User)
	u.OpenID = openID
	if err := db.First(u).Error; err != nil {
		panic(err.Error())
	}
	return u
}
