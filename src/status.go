package main

import "github.com/gin-gonic/gin"

//时间太紧，只好弄一个假的websocket胡弄下，正式上线后会考虑将其更换为真正的websocket
func (s *Service) SetProjectStatus(c *gin.Context) (int, interface{}) {
	projectstatus_json := new(struct {
		ProjectStatus
		UserID string `json:"user_id" gorm:"not null" binding:"required"`
	})
	err := c.ShouldBindJSON(projectstatus_json)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	//权限校验
	role, err := s.checkProject(projectstatus_json.ProjectID, projectstatus_json.UserID)
	if err != nil {
		return s.makeErrJSON(403, 40302, err.Error())
	}
	if role != RoleTable["director"] {
		return s.makeErrJSON(403, 40301, "dont director")
	}
	//开启事务
	tx := s.DB.Begin()
	if tx.Create(&projectstatus_json.ProjectStatus).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, "insert error")
	}
	tx.Commit()
	return s.makeSuccessJSON(projectstatus_json.ProjectStatus)
}

func (s *Service) GetProjectStatus(c *gin.Context) (int, interface{}) {
	projectid := c.Query("projectid")
	userid := c.Query("userid")
	if projectid == "" || userid == "" {
		return s.makeErrJSON(403, 40301, "query error")
	}
	//权限验证
	role, err := s.checkProject(projectid, userid)
	if role < 1 || role > 7 || err != nil {
		return s.makeErrJSON(403, 40302, "limited access")
	}
	projectstatus := new(ProjectStatus)
	s.DB.Where(ProjectStatus{ProjectID: projectid}).Last(projectstatus)
	return s.makeSuccessJSON(projectstatus)
}
