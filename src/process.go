package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/* 用超级笨的方法实现的增删改查，全删全插入，后端界的耻辱
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
*/
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

func (s *Service) UpdateProcess(c *gin.Context) (int, interface{}) {
	userid := c.Param("userid")
	processes := make([]*Process, 10, 20)
	if err := c.ShouldBindJSON(&processes); err != nil {
		return 500, err.Error()
	}
	role, err := s.checkProject(processes[0].ProjectID, userid)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	if role < 1 || role > 2 {
		return s.makeErrJSON(403, 40301, "limited access")
	}
	nowProcessID := []string{}
	//使环节id保持不变
	for _, process := range processes {
		if s.DB.Where(Process{ProcessID: process.ProcessID, ProjectID: process.ProjectID}).Find(&Process{}).RowsAffected == 0 {
			process.ProcessID = uuid.New().String()
		}
		nowProcessID = append(nowProcessID, process.ProcessID)
	}
	tx := s.DB.Begin()
	err = tx.Where("process_id not in (?)", nowProcessID).Delete(processes).Error
	if err != nil {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, err.Error())
	}
	for _, process := range processes {
		err := tx.Where(Process{ProcessID: process.ProcessID}).
			Assign(Process{ProcessID: process.ProcessID, Order: process.Order, ProcessName: process.ProcessName,
				ProcessType: process.ProcessType, MicHand: process.MicHand, MicEar: process.MicEar, Remark: process.Remark,
				ProjectID: process.ProjectID, ManagerID: process.ManagerID}).
			FirstOrCreate(&process).Error
		if err != nil {
			tx.Rollback()
			return s.makeErrJSON(500, 50001, string(process.Order)+" update error")
		}
	}
	tx.Commit()
	return s.makeSuccessJSON(processes)
}

type Manager struct {
	ManagerID string `json:"manager_id" binding:"required"`
	ProcessID string `json:"process_id" binding:"required"`
}

func (s *Service) SetManager(c *gin.Context) (int, interface{}) {
	userid := c.Param("userid")
	manager := new(Manager)
	err := c.ShouldBindJSON(manager)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	process := new(Process)
	s.DB.Where(&Process{ProcessID: manager.ProcessID}).Find(process)
	//验证操作者身份
	role, err := s.checkProject(process.ProjectID, userid)
	if role != 1 {
		return s.makeErrJSON(403, 40302, "limited access")
	}
	if err != nil {
		return s.makeErrJSON(403, 40302, err.Error())
	}
	//验证要设置的负责人的身份
	err = s.checkUser(manager.ManagerID)
	if err != nil {
		return s.makeErrJSON(403, 40303, err.Error())
	}

	process.ManagerID = manager.ManagerID
	tx := s.DB.Begin()
	if tx.Model(&Process{}).Where(&Process{ProcessID: manager.ProcessID}).Updates(process).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, "update error")
	}
	tx.Commit()
	return s.makeSuccessJSON(process)
}
