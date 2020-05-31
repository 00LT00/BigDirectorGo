package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"image/jpeg"
	"io/ioutil"

	//_ "image/png"
	"os"
	"path"
)

const WxacodeUrl = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"

//暂时不用
func (s *Service) MakeWxacodePicture(projectid string) error {
	//func (s *Service)MakeWxacodePicture(c *gin.Context)(int,interface{}){
	token, err := s.GetToken()

	//通过这两个方法从连接池中获取一个空的实例，可以实现连接复用，提高性能
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 重复设置方法可以通过set函数内部的append方法去截断前一次设置的方法,同时转换成[]byte
	req.Header.SetMethod("POST")
	//构造get请求地址
	url, err := s.makeUrl(WxacodeUrl,
		Query{"access_token", token})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	req.SetRequestURI(url)
	requestbody, err := json.Marshal(map[string]string{
		"scene": projectid,
	})
	req.SetBody(requestbody)
	if err := fasthttp.Do(req, resp); err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return errors.New(string(resp.Body()))
	}

	b := resp.Body()
	var responsebody struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}
	err = nil
	err = json.Unmarshal(b, &responsebody)
	if err == nil {
		return errors.New(fmt.Sprintf("%d,%s", responsebody.Errcode, responsebody.Errmsg))
	}

	buffer := bytes.NewBuffer(b)
	img, err := jpeg.Decode(buffer)
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join("..", "wxacode", projectid+".jpg"))
	defer f.Close()
	if err != nil {
		return err
	}
	fmt.Println(f, b, img)
	err = jpeg.Encode(f, img, nil)
	if err != nil {
		return err
	}
	//返回nil
	return nil
}

func(s *Service)MakeWxacodeBuffer(projectid string)(string,error){
	//func (s *Service)MakeWxacodePicture(c *gin.Context)(int,interface{}){
	token, err := s.GetToken()

	//通过这两个方法从连接池中获取一个空的实例，可以实现连接复用，提高性能
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 重复设置方法可以通过set函数内部的append方法去截断前一次设置的方法,同时转换成[]byte
	req.Header.SetMethod("POST")
	//构造get请求地址
	url, err := s.makeUrl(WxacodeUrl,
		Query{"access_token", token})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	req.SetRequestURI(url)
	requestbody, err := json.Marshal(map[string]string{
		"scene": projectid,
	})
	req.SetBody(requestbody)
	if err := fasthttp.Do(req, resp); err != nil {
		return "",err
	}

	if resp.StatusCode() != 200 {
		return "",errors.New(string(resp.Body()))
	}

	b := resp.Body()
	var responsebody struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}
	err = nil
	err = json.Unmarshal(b, &responsebody)
	if err == nil {
		return "",errors.New(fmt.Sprintf("%d,%s", responsebody.Errcode, responsebody.Errmsg))
	}

	return string(b),nil
}

func(s *Service)GetWxacodeBuffer(c *gin.Context)(int,interface{}){
	projectid :=c.Param("projectid")
	filename :=path.Join("..","wxacode",projectid)
	buffer,err:=ioutil.ReadFile(filename)
	if err != nil {
		return s.makeErrJSON(500,50000,err.Error())
	}
	return s.makeSuccessJSON(buffer)
}