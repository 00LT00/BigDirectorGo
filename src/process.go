package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//检测环节是否存在，顺便返回次序
func (s *Service) checkProcessOrder(processid string, projectid string) (int64, error) {
	process := new(Process)
	s.DB.Where(Process{ProcessID: processid, ProjectID: projectid}).Find(process)
	var err error = nil
	if process.Order == 0 || processid == "" {
		err = errors.New("none process or projectid err")
	}
	return process.Order, err
}

//添加环节
func (s *Service) AddProcess(c *gin.Context) (int, interface{}) {
	userid := c.Param("userid")
	process := new(Process)
	//参数绑定
	if err := c.ShouldBindJSON(process); err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	//校验权限
	role, err := s.checkProject(process.ProjectID, userid)
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

//修改环节
func (s *Service) UpdateProcess(c *gin.Context) (int, interface{}) {
	userid := c.Param("userid")
	process := new(Process)
	if err := c.ShouldBindJSON(process); err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	//user校验权限
	role, err := s.checkProject(process.ProjectID, userid)
	if role < 1 || role > 2 {
		return s.makeErrJSON(403, 40302, "limited access")
	}
	if err != nil {
		return s.makeErrJSON(403, 40302, err.Error())
	}
	//process 权限校验
	_, err = s.checkProcessOrder(process.ProcessID, process.ProjectID)
	if err != nil {
		return s.makeErrJSON(403, 40303, err.Error())
	}

	//开启事务
	tx := s.DB.Begin()
	if tx.Model(&Process{}).
		Where(Process{ProcessID: process.ProcessID, ProjectID: process.ProjectID}).
		Updates(&process).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, errors.New("update error or none update"))
	}
	tx.Commit()
	return s.makeSuccessJSON(process)
}

//查找环节
func (s *Service) GetProcess(c *gin.Context) (int, interface{}) {
	processid := c.Query("processid")
	userid := c.Query("userid")
	process := new(Process)
	s.DB.Where(Process{ProcessID: processid}).Find(process)
	//if process.ProjectID == "" {
	//	return s.makeErrJSON(403,40301,errors.New("none process"))
	//}
	//通过获取的projectid验证user，同时验证process是否合法
	_, err := s.checkProject(process.ProjectID, userid)
	if err != nil {
		return s.makeErrJSON(403, 40302, err.Error())
	}
	return s.makeSuccessJSON(process)
}
