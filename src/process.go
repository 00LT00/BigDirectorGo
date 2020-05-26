package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//添加环节
func (s *Service) AddProcess(c *gin.Context) (int, interface{}) {
	userid := c.Param("userid")
	process := new(Process)
	//参数绑定
	if err := c.ShouldBindJSON(process); err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	role, err := s.checkProject(process.ProjectID, userid)
	//校验权限
	if role < 1 || role > 2 {
		return s.makeErrJSON(403, 40302, "limited access")
	}
	if err != nil {
		return s.makeErrJSON(403, 40302, err.Error())
	}
	uuid := uuid.New().String()
	process.ProcessID = uuid
	tx := s.DB.Begin()
	if err := tx.Create(process).Error; err != nil {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(process)
}
