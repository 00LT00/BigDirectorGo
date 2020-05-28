package main

import "github.com/gin-gonic/gin"

//设置灯光组等人
func (s *Service) SetWorker(c *gin.Context) (int, interface{}) {
	userid := c.Param("userid")
	worker := new(Worker)
	err := c.ShouldBindJSON(worker)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
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
	if tx.Where(Worker{ProjectID: worker.ProjectID, Type: worker.Type}).
		Assign(Worker{ProjectID: worker.ProjectID, WorkerID: worker.WorkerID, Type: worker.Type}).
		FirstOrCreate(&worker).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, "update error")
	}
	tx.Commit()
	return s.makeSuccessJSON(worker)
}

//type result struct {
//	workers  []Worker
//	managers []Manager
//}

//获取项目的所有管理人员
//func (s *Service) GetWorker(c *gin.Context) (int, interface{}) {
//	userid := c.Query("userid")
//	projectid := c.Query("projectid")
//	role, err := s.checkProject(projectid, userid)
//	if err != nil {
//		return s.makeErrJSON(403, 40301, err.Error())
//	}
//	if role < 1 || role > 7 {
//		return s.makeErrJSON(403, 40301, "limited access")
//	}
//
//	//先获取workers
//	role = 4
//	for ; role <= 7; role++ {
//
//	}
//
//}
