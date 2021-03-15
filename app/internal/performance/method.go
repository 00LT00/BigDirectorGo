package performance

import (
	"BigDirector/app/database"
	error2 "BigDirector/error"
	"BigDirector/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var s = service.Service

// 新建或更新演出信息
// @Tags performance
// @Summary 创建或更改演出信息
// @Description createPerformance or update performance information
// @ID set-Performance-Info
// @Accept json
// @Produce  json
// @Param performance body database.Performance true "演出信息，创建时ID为空, name必填"
// @Param openID query string true "创建人或者导演组一员的openID 必填"
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=string} "只返回performanceID"
// @Failure 400 {object} utils.FailureResponse "40001 param error"
// @Failure 403 {object} utils.FailureResponse "40301 can't set performance"
// @Failure 500 {object} utils.FailureResponse "service error"
// @Router /performance/info [put]
func SetInfo(c *gin.Context) interface{} {
	//获取用户信息
	openID := c.Query("openID")
	if openID == "" {
		panic(error2.NewHttpError(400, "40001", "openID null"))
	}
	u := new(database.User)
	u.OpenID = openID
	if err := s.DB.Where(u).First(u).Error; err != nil {
		panic(err.Error())
	}

	//获取演出信息
	p := new(database.Performance)
	err := c.ShouldBindJSON(p)
	if err != nil {
		panic(error2.NewHttpError(400, "40001", err.Error()))
	}
	if p.PerformanceID == "" {
		createPerformance(p, u)
	} else {
		//校验casbin权限
		ok, err := s.Casbin.Enforce(u.OpenID, p.PerformanceID, "performance data", "get")
		if err != nil {
			panic(err.Error())
		} else if !ok {
			panic(error2.NewHttpError(403, "40301", "can't set performance"))
		}
		//保存
		if err := s.DB.Save(p).Error; err != nil {
			panic(err.Error())
		}
	}
	return p.PerformanceID
}

func createPerformance(p *database.Performance, u *database.User) {
	p.PerformanceID = uuid.NewV4().String()
	p.Users = []*database.User{u}
	tx := s.DB.Begin()
	//创建演出
	if err := tx.Create(p).Error; err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	//创建导演组
	g := &database.Group{
		GroupID:       uuid.NewV4().String(),
		PerformanceID: p.PerformanceID,
		Name:          database.DirectorRole,
		Users:         []*database.User{u}, //用户默认为导演组人员
	}
	if err := tx.Create(g).Error; err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	//casbin中设置导演角色
	if _, err := s.Casbin.AddRoleForUserInDomain(u.OpenID, g.GroupID, p.PerformanceID); err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	if _, err := s.Casbin.AddPolicy(g.GroupID, p.PerformanceID, "*", ".*"); err != nil { //.* 可以正则匹配所有操作，表示最高权限
		tx.Rollback()
		panic(err.Error())
	}
	tx.Commit()
	return
}

// 获取演出详情
// @Tags performance
// @Summary 获取演出信息
// @Description get performance information
// @ID get-Performance-Info
// @Produce  json
// @Param performanceID query string true "performanceID 必填"
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=database.Performance} "演出详情"
// @Failure 400 {object} utils.FailureResponse "40001 param error"
// @Failure 500 {object} utils.FailureResponse "service error"
// @Router /performance/info [get]
func GetInfo(c *gin.Context) interface{} {
	performanceID := c.Query("performanceID")
	if performanceID == "" {
		panic(error2.NewHttpError(400, "40001", "performanceID null"))
	}
	p := new(database.Performance)
	p.PerformanceID = performanceID
	if err := s.DB.First(p).Error; err != nil {
		panic(err.Error())
	}
	return p
}
