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
	if err := tx.Create(&Project_User{ProjectID: project.ProjectID, UserID: project.DirectorUserID, Role: RoleTable["director"].(int)}).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50001, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(project)
}

func (s *Service) AddMember(c *gin.Context) (int, interface{}){
	pju:=new(Project_User)
	if err:=c.ShouldBindJSON(pju);err!=nil{
		return s.makeErrJSON(403,40307,err.Error())
	}
	role,err:=s.checkProject(pju.ProjectID,pju.UserID)
	if err != nil {
		return s.makeErrJSON(404,40402,err.Error())
	}
	if  1<=role && role<=6 {
		return s.makeErrJSON(403,40308,errors.New("Already bound"))
	}

	pju.Role = RoleTable["member"].(int)
	tx:=s.DB.Begin()
	if err := tx.Create(pju).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500,50006,err.Error())
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

func (s *Service) UpdateProject(c *gin.Context) (int, interface{}) {
	ProjectID:=c.Param("projectid")
	UserID:=c.Param("userid")[1:] //路由匹配会把‘/’也放进来，必须做处理
	project:=new(Project)
	if err:=c.ShouldBindJSON(project);err!=nil{
		return s.makeErrJSON(403,40303, err.Error())
	}
	//如果不改导演
	if project.DirectorUserID == UserID {
		role,err:= s.checkProject(ProjectID,UserID)
		if err != nil {
			return s.makeErrJSON(403,40304,err.Error())
		}
		if role !=1 {
			return s.makeErrJSON(403,40305,"not director")
		}
		tx:=s.DB.Begin()
		if err:=tx.Model(&Project{}).Updates(project).Error;err!=nil {
			tx.Callback()
			return s.makeErrJSON(500,50002,err.Error())
		}
		tx.Commit()
		return s.makeSuccessJSON(project)
	}
	//如果导演改了，要判断新导演是否有权限，同时更改项目表和权限表

	//验证用户和项目的关系
	if _,err:=s.checkProject(ProjectID,project.DirectorUserID);err != nil { // 这里directoruserid指的是从json发过来的，应该是修改后的userid，作者懒得改字段了，凑活看
		return s.makeErrJSON(403,40306,err.Error())
	}
	tx:=s.DB.Begin()
	if err:=tx.Model(&Project{}).Updates(project).Error;err!=nil {
		tx.Callback()
		return s.makeErrJSON(500,50003,err.Error())
	}
	if err:=tx.Model(&Project_User{}).
		Where(Project_User{ProjectID:ProjectID,UserID:UserID}).  //选择正在使用的帐号（原导演）
		Update(Project_User{Role:RoleTable["member"].(int)}).Error;err!=nil{ // 权限降为 member
		tx.Callback()
		return s.makeErrJSON(500,50004,err.Error())
	}
	if err:=tx.Model(&Project_User{}).
		Where(Project_User{ProjectID:ProjectID,UserID:project.DirectorUserID}).  //json中的新导演
		Update(Project_User{Role:RoleTable["director"].(int)}).Error;err!=nil{ // 权限升为 director
		tx.Callback()
		return s.makeErrJSON(500,50005,err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(project)
}
