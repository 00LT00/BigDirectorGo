package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//验证项目是否存在，或者返回权限
func (s *Service) checkProject(projectid string, userid ...string) (int, error) {
	if s.DB.Find(&Project{}, Project{ProjectID: projectid}).RowsAffected == 0 {
		//项目不存在
		return -1, errors.New("none project")
	}
	//没有userid的参数
	if userid == nil {
		return -1, nil
	}
	//关系表中查找用户和项目的关系
	pju := new(Project_User)
	if s.DB.Where(&Project_User{ProjectID: projectid, UserID: userid[0]}).
		Find(pju).RowsAffected == 0 {
		return -1, errors.New("limited access")
	}
	return pju.Role, nil
}

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
	if err := tx.Create(&Project_User{ProjectID: project.ProjectID, UserID: project.DirectorUserID, Role: RoleTable["director"]}).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50001, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(project)
}

func (s *Service) GetProject(c *gin.Context) (int, interface{}) {
	//路由形式 /:ProjectID/*UserID
	ProjectID := c.Param("ProjectID")
	UserID := c.Param("UserID")[1:]
	result := struct {
		Project
		Role int
	}{}
	var err error
	if result.Role, err = s.checkProject(ProjectID, UserID); err != nil {
		return s.makeErrJSON(403, 40302, err.Error())
	}
	if s.DB.Where(&Project{ProjectID: ProjectID}).Find(&result.Project).RowsAffected != 1 {
		return s.makeErrJSON(404, 40401, "get error")
	}
	return s.makeSuccessJSON(result)
}
