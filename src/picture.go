package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

func (s *Service) GetPicture(c *gin.Context) {
	filename_:=c.Param("filename")
	filename := path.Join("..","picture",filename_)
	fmt.Println(filename)
	file,err:=os.Open(filename) // file文件实现了io.reader
	if err != nil {
		c.JSON(s.makeErrJSON(404,40400,err.Error()))
		return
	}
	fileinfo,err:=os.Stat(filename) //仅仅为了长度
	if err != nil {
		c.JSON(s.makeErrJSON(500,50000,err.Error()))
	}
	var contenttype string
	if path.Ext(filename) == ".jpg" {
		contenttype = "image/jpeg"
	}else if path.Ext(filename) == ".png" {
		contenttype = "image/png"
	}
	extraHeaders := map[string]string{ // 返回头，根据返回头浏览器将选择是下载页面还是直接展示图片
		//"Content-Disposition": `attachment; filename="` + fileName + `"`,
	}
	c.DataFromReader(200,fileinfo.Size(),contenttype,file,extraHeaders)
}

