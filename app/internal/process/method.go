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
// @Summary 创建或更改环节信息
// @Description create or update process information
// @ID set-Process-Info
// @Accept json
// @Produce  json
// @Param process body []database.Process true "process结构体数组 performanceID必须一致且存在"
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=string} "success"
// @Failure 400 {object} utils.FailureResponse "40001 param error"
// @Failure 500 {object} utils.FailureResponse "service error"
// @Router /process/info [put]
func SetInfo(c *gin.Context) interface{} {
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
