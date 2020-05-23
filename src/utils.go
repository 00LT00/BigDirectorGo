package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

const OpenidUrl string = "https://api.weixin.qq.com/sns/jscode2session"
const TokenUrl string = "https://api.weixin.qq.com/cgi-bin/token"

type Query struct {
	Name  string
	Value string
}

type openjson struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Openid  string `json:"openid"`
}

type tokenjson struct {
	Errcode      int64  `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Access_token string `json:"access_token"`
}

func (s *Service) makeGetUrl(baseUrl string, queries ...Query) (string, error) {
	Url := baseUrl + "?"
	for _, v := range queries {
		Url += fmt.Sprintf("%s=%s&", v.Name, v.Value)
	}
	if len(Url) <= len(baseUrl) {
		return "", errors.New("error")
	}
	return Url[:len(Url)-1], nil
}

func (s *Service) GetOpenID(code string) (string, error) {

	//通过这两个方法从连接池中获取一个空的实例，可以实现连接复用，提高性能
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 重复设置方法可以通过set函数内部的append方法去截断前一次设置的方法
	req.Header.SetMethod("GET")
	//构造get请求地址
	url, err := s.makeGetUrl(OpenidUrl,
		Query{"appid", s.Conf.Wx.AppID},
		Query{"secret", s.Conf.Wx.AppSecret},
		Query{"js_code", code},
		Query{"grant_type", "authorization_code"})
	if err != nil {
		return "", err
	}
	req.SetRequestURI(url)

	if err := fasthttp.Do(req, resp); err != nil {
		return "", errors.New("get openid err: " + err.Error())
	}

	//获取响应体
	b := resp.Body()
	respjson := openjson{}
	err = json.Unmarshal(b, &respjson)
	if err != nil {
		return "", err
	}
	if respjson.Openid == "" {
		return "", errors.New("get openid errcode: " + fmt.Sprintf("%d", respjson.Errcode) + " errmsg: " + respjson.Errmsg)
	}
	return respjson.Openid, nil

}
