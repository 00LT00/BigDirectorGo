package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//验证项目是否存在，或者返回权限
// -1 只检测项目是否存在的返回值，无报错就是项目存在
//  0 无关人员

func (s *Service) checkProject(projectid string, userid ...string) (int, error) {
	if s.DB.Find(&Project{}, Project{ProjectID: projectid}).RowsAffected == 0 {
		//项目不存在
		return -1, errors.New("none project")
	}
	//没有userid的参数,表示项目存在
	if userid == nil {
		return -1, nil
	}
	//关系表中查找用户和项目的关系
	pju := new(Project_User)
	if s.DB.Where(&Project_User{ProjectID: projectid, UserID: userid[0]}).
		Find(pju).RowsAffected == 0 {
		return 0, errors.New("limited access")
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
	uuid4 := uuid.New().String()
	project.ProjectID = uuid4[:len(uuid4)-4]
	//开启事务
	tx := s.DB.Begin()
	//新建项目
	if err := tx.Create(project).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50000, err.Error())
	}
	//把导演和项目绑定
	if err := tx.Create(&Project_User{ProjectID: project.ProjectID, UserID: project.DirectorUserID, Role: RoleTable["director"].(int)}).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50001, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(project)
}

func (s *Service) AddMember(c *gin.Context) (int, interface{}) {
	pju := new(Project_User)
	if err := c.ShouldBindJSON(pju); err != nil {
		return s.makeErrJSON(403, 40307, err.Error())
	}
	role, err := s.checkProject(pju.ProjectID, pju.UserID)
	if role == -1 && err != nil {
		return s.makeErrJSON(404, 40402, err.Error())
	}
	if 1 <= role && role <= 6 {
		return s.makeErrJSON(403, 40308, errors.New("Already bound"))
	}

	pju.Role = RoleTable["member"].(int)
	tx := s.DB.Begin()
	if err := tx.Create(pju).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50006, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(pju)
}

func (s *Service) GetProject(c *gin.Context) (int, interface{}) {
	//路由形式 /:ProjectID/*UserID
	ProjectID := c.Param("projectid")
	UserID := c.Param("userid")[1:]
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

//更改项目名称
func (s *Service) UpdateProjectName(c *gin.Context) (int, interface{}) {
	ProjectID := c.Param("projectid")
	project := new(Project)
	if err := c.ShouldBindJSON(project); err != nil {
		return s.makeErrJSON(403, 40303, err.Error())
	}
	role, err := s.checkProject(ProjectID, project.DirectorUserID)
	if err != nil {
		return s.makeErrJSON(403, 40304, err.Error())
	}
	if role != 1 {
		return s.makeErrJSON(403, 40305, "not director")
	}
	tx := s.DB.Begin()
	if err := tx.Model(&Project{}).Where(Project{ProjectID: ProjectID}).Updates(&project).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50002, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(project)
}

func (s *Service) UpdateProjectUserid(c *gin.Context) (int, interface{}) {
	ProjectID := c.Param("projectid")
	tempstruct := struct {
		UserID         string `json:"userid" binding:"required"`
		DirectorUserID string `json:"directoruserid" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&tempstruct); err != nil {
		return s.makeErrJSON(403, 40311, err.Error())
	}

	//验证用户和项目的关系
	if role, err := s.checkProject(ProjectID, tempstruct.UserID); err != nil {
		return s.makeErrJSON(403, 40312, err.Error())
	} else if role != 1 {
		return s.makeErrJSON(403, 40313, errors.New("limited access"))
	}

	tx := s.DB.Begin()
	if err := tx.Model(&Project{}).Where(Project{ProjectID: ProjectID}).Updates(Project{DirectorUserID: tempstruct.DirectorUserID}).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50003, err.Error())
	}
	if err := tx.Model(&Project_User{}).
		Where(Project_User{ProjectID: ProjectID, UserID: tempstruct.UserID}).     //选择正在使用的帐号（原导演）
		Update(Project_User{Role: RoleTable["member"].(int)}).Error; err != nil { // 权限降为 member
		tx.Callback()
		return s.makeErrJSON(500, 50004, err.Error())
	}
	if err := tx.Model(&Project_User{}).
		Where(Project_User{ProjectID: ProjectID, UserID: tempstruct.DirectorUserID}). //json中的新导演
		Update(Project_User{Role: RoleTable["director"].(int)}).Error; err != nil {   // 权限升为 director
		tx.Callback()
		return s.makeErrJSON(500, 50005, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(tempstruct)
}

//获取当前项目的所有成员
func (s *Service) GetProjectUser(c *gin.Context) (int, interface{}) {
	userid := c.Query("userid")
	projectid := c.Query("projectid")
	role, err := s.checkProject(projectid, userid)
	if err != nil {
		return s.makeErrJSON(403, 40309, err.Error())
	}
	if role <= 0 || role > 6 {
		return s.makeErrJSON(403, 40310, "none role")
	}
	type member struct {
		UserID   string
		UserName string
		Role     int
	}
	members := make([]*member, 50, 100)
	err = s.DB.Table("project_user").
		Select("project_user.user_id,users.user_name,project_user.role").
		Joins("left join users on project_user.user_id = users.user_id").
		Where(&Project_User{ProjectID: projectid}).Scan(&members).Error
	if err != nil {
		return s.makeErrJSON(500, 50007, err.Error())
	}
	return s.makeSuccessJSON(members)
}
