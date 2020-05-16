package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) AddProject(c *gin.Context) (int, interface{}) {
	project := new(Project)
	//接收json
	if err := c.ShouldBindJSON(project); err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	//验证导演id
	if err := s.checkUser(project.DirectorUserID); err != nil {
		return s.makeErrJSON(404, 40400, err.Error())
	}
	//生成uuid
	project.ProjectID = uuid.New().String()
	//开启事务
	tx := s.DB.Begin()
	//新建项目
	if err := tx.Create(project).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50000, err.Error())
	}
	//把导演和项目绑定
	if err := tx.Create(&Project_User{ProjectID: project.ProjectID, UserID: project.DirectorUserID}).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50001, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(project)
}
