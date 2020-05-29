package main

import (
	"github.com/gin-gonic/gin"
)

//设置灯光组等人
func (s *Service) SetWorker(c *gin.Context) (int, interface{}) {
	userid := c.Param("userid")
	worker := new(Worker)
	err := c.ShouldBindJSON(worker)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	//验证要设置的权限
	if worker.Role < 4 || worker.Role > 7 {
		return s.makeErrJSON(403, 40304, "type error")
	}
	role, err := s.checkProject(worker.ProjectID, userid)
	if err != nil {
		return s.makeErrJSON(403, 40302, err.Error())
	}
	if role != 1 {
		return s.makeErrJSON(403, 40302, "limited access")
	}
	role, err = s.checkProject(worker.ProjectID, worker.WorkerID)
	if err != nil || role < 1 || role > 7 {
		return s.makeErrJSON(403, 40303, "none member")
	}
	tx := s.DB.Begin()
	if tx.Where(Worker{ProjectID: worker.ProjectID, Role: worker.Role}).
		Assign(Worker{ProjectID: worker.ProjectID, WorkerID: worker.WorkerID, Role: worker.Role}).
		FirstOrCreate(&worker).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, "update error")
	}
	tx.Commit()
	return s.makeSuccessJSON(worker)
}

type Result struct {
	Workers []struct {
		Worker
		WorkerName string
		PhoneNum   string
	}
	Managers []struct {
		Manager
		ManagerName string
		PhoneNum    string
		ProcessName string
	}
}

//获取项目的所有管理人员
func (s *Service) GetWorker(c *gin.Context) (int, interface{}) {
	userid := c.Query("userid")
	projectid := c.Query("projectid")
	result := new(Result)
	role, err := s.checkProject(projectid, userid)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	if role < 1 || role > 7 {
		return s.makeErrJSON(403, 40301, "limited access")
	}

	//先获取workers
	role = 4
	for ; role <= 7; role++ {
		worker := new(struct {
			Worker
			WorkerName string
			PhoneNum   string
		})
		if s.DB.Where(&Worker{Role: role, ProjectID: projectid}).Find(&worker.Worker).RowsAffected == 1 {
			user := new(User)
			if s.DB.Where(&User{UserId: worker.WorkerID}).Find(&user).RowsAffected != 1 {
				return s.makeErrJSON(500, 50001, "get workerinfo error")
			}
			worker.PhoneNum = user.PhoneNum
			worker.WorkerName = user.UserName
			result.Workers = append(result.Workers, *worker)
		}
	}
	// 找manager,其实这里应该使用视图或者是直接用外连接做，但这里因为使用频率不高,为了和上面写法统一，就没去写连接
	processes := make([]*Process, 10, 20)
	s.DB.Where(&Process{ProjectID: projectid}).Order("order").Find(&processes)
	for _, process := range processes {
		manager := new(struct {
			Manager
			ManagerName string
			PhoneNum    string
			ProcessName string
		})
		manager.ProcessID = process.ProcessID
		manager.ManagerID = process.ManagerID
		manager.ProcessName = process.ProcessName
		user := new(User)
		if s.DB.Table("users").Where("user_id = ?", manager.ManagerID).Find(&user).RowsAffected > 1 {
			return s.makeErrJSON(500, 50001, "get managerinfo error")
		}
		manager.ManagerName = user.UserName
		manager.PhoneNum = user.PhoneNum
		result.Managers = append(result.Managers, *manager)
	}
	return s.makeSuccessJSON(result)
}
