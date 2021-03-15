package group

import (
	"BigDirector/app/database"
	error2 "BigDirector/error"
	"BigDirector/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
)

var s = service.Service

// 新建或更新工作组
// @Tags group
// @Summary 创建或更新工作组
// @Description create or update group information
// @ID set-Group-Info
// @Accept json
// @Produce  json
// @Param process body database.Group true "组信息 performanceID必填, GroupID空则为新建, leaderID选填（组长的openID）"
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=string} "GroupID"
// @Failure 400 {object} utils.FailureResponse "40001 param error"
// @Failure 500 {object} utils.FailureResponse "service error"
// @Router /group/info [put]
func SetInfo(c *gin.Context) interface{} {
	g := new(database.Group)
	err := c.ShouldBindJSON(g)
	if err != nil {
		panic(error2.NewHttpError(400, "40001", err.Error()))
	}
	u := new(database.User)
	u.OpenID = g.LeaderID
	if err := s.DB.First(u).Error; err != nil {
		panic(err.Error())
	}
	g.Leader = u
	g.Users = append(g.Users, u)
	if g.GroupID == "" {
		g.GroupID = uuid.NewV4().String()
	}
	tx := s.DB.Begin()
	if err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(g).Error; err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	if err := updatePolicy(g); err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	_, err = s.Casbin.AddRoleForUserInDomain(g.LeaderID, g.GroupID, g.PerformanceID)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	tx.Commit()
	return g.GroupID
}

func updatePolicy(g *database.Group) error {
	var policyRules [][]string
	var oldPolicyRules [][]string
	for _, role := range g.Roles {
		policyRules = append(policyRules, []string{
			g.GroupID, g.PerformanceID, role, "set",
		})
	}
	oldPolicyRules = s.Casbin.GetPermissionsForUserInDomain(g.GroupID, g.PerformanceID)
	if _, err := s.Casbin.RemovePolicies(oldPolicyRules); err != nil {
		return err
	}
	if _, err := s.Casbin.AddPolicies(policyRules); err != nil {
		return err
	}
	return nil
}

// 获取小组详情
// @Tags performance
// @Summary 获取小组信息
// @Description get group information
// @ID get-Group-Info
// @Produce  json
// @Param groupID query string true "groupID 必填"
// @Param sign header string true "check header" default(spppk)
// @Success 200 {object} utils.SuccessResponse{data=database.Group{Leader=database.User}} "小组详情，其中roles是权限"
// @Failure 400 {object} utils.FailureResponse "40001 param error"
// @Failure 500 {object} utils.FailureResponse "service error"
// @Router /group/info [get]
func GetInfo(c *gin.Context) interface{} {
	groupID := c.Query("groupID")
	if groupID == "" {
		panic(error2.NewHttpError(400, "40001", "groupID null"))
	}
	g := new(database.Group)
	g.GroupID = groupID
	if err := s.DB.Preload("Leader").First(g).Error; err != nil {
		panic(err.Error())
	}
	var roles []string
	policies := s.Casbin.GetPermissionsForUserInDomain(g.GroupID, g.PerformanceID)
	for _, policy := range policies {
		roles = append(roles, policy[2])
	}
	g.Roles = roles
	return g
}
