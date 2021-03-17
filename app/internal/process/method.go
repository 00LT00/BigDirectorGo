package process

import (
	"BigDirector/app/database"
	error2 "BigDirector/error"
	"BigDirector/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
)

var s = service.Service

// 新建或更新演出信息
// @Tags process
// @Summary 创建或更改环节信息，包括全部环节
// @Description create or update process list
// @ID set-Process-List
// @Accept json
// @Produce  json
// @Param Authorization header string true "格式为: token OPENID 这里替换成使用者的openID" default(token OPENID)
// @Param process body []database.Process true "process结构体数组 performanceID必须一致且存在"
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=string} "success"
// @Failure 400 {object} utils.FailureResponse "40001 param error"
// @Failure 500 {object} utils.FailureResponse "service error"
// @Router /process/list [put]
func SetList(c *gin.Context) interface{} {
	//获取环节信息
	var processes []*database.Process
	if err := c.ShouldBindJSON(&processes); err != nil {
		panic(error2.NewHttpError(400, "40001", err.Error()))
	}
	//获取演出详情
	performance := new(database.Performance)
	for k, process := range processes {
		process.Order = k
		process.ProcessID = uuid.NewV4().String()
		performance.PerformanceID = process.PerformanceID
	}
	if err := s.DB.First(performance).Error; err != nil {
		panic(err.Error())
	}
	performance.Processes = processes

	if err := s.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(performance).Error; err != nil {
		panic(err.Error())
	}

	return "success"
}

// 获取演出详情
// @Tags process
// @Summary 获取全部环节信息
// @Description get process list
// @ID get-Process-List
// @Produce  json
// @Param Authorization header string true "格式为: token OPENID 这里替换成使用者的openID" default(token OPENID)
// @Param performanceID query string true "performanceID 必填"
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=[]database.Process} "环节列表"
// @Failure 400 {object} utils.FailureResponse "40001 param error"
// @Failure 500 {object} utils.FailureResponse "service error"
// @Router /process/list [get]
func GetList(c *gin.Context) interface{} {
	PerformanceID := c.Query("performanceID")
	if PerformanceID == "" {
		panic(error2.NewHttpError(400, "40001", "performanceID null"))
	}
	performance := new(database.Performance)
	performance.PerformanceID = PerformanceID
	if err := s.DB.Preload("Processes").First(performance).Error; err != nil {
		panic(err.Error())
	}
	return performance.Processes
}
