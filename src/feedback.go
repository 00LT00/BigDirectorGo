package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
	"path"
)

func (s *Service)GetFeedBack(c *gin.Context)(int,interface{}){
	msg:=c.PostForm("msg")
	picture,err:=c.FormFile("picture")
	if err != nil {
		return s.makeErrJSON(403,40301,err.Error())
	}
	if msg== ""{
		return s.makeErrJSON(403,40302,"msg null")
	}
	feedback:=new(FeedBack)
	feedback.FeedBackID = uuid.New().String()
	feedback.Msg = msg
	tx :=s.DB.Begin()
	if err:=tx.Create(feedback).Error;err != nil {
		tx.Rollback()
		return s.makeErrJSON(500,50000,err.Error())
	}
	tx.Commit()

	//保存图片
	if picture!=nil {
		picPath :=path.Join("..","feedback")
		_,err:=os.Stat(picPath)
		if err!=nil {
			err =os.Mkdir(picPath,0700)
			if err !=nil {
				return s.makeErrJSON(500,50001,err.Error())
			}
		}
		suffix:=path.Ext(picture.Filename)
		filename:=path.Join(picPath,feedback.FeedBackID+"."+suffix)
		err =c.SaveUploadedFile(picture,filename)
		if err != nil {
			return s.makeErrJSON(500,50010,err.Error())
		}
	}
	return s.makeSuccessJSON("upload feedback success")
}
