package user

import (
	error2 "BigDirector/error"
	"BigDirector/service"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
	"strconv"
)

type wxRespStruct struct {
	Openid      string `json:"openid"`
	Session_key string `json:"session_key,-"`
	Unionid     string `json:"unionid"`
	Errcode     int    `json:"errcode,omitempty"`
	Errmsg      string `json:"errmsg,omitempty"`
}

func getOpenID(code string) interface{} {
	wx := service.Conf.Wx
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	// 设置微信接口
	base, err := url.Parse("https://api.weixin.qq.com/sns/jscode2session")
	if err != nil {
		panic(err.Error())
	}
	query := base.Query()
	query.Add("appid", wx.AppID)
	query.Add("secret", wx.AppSecret)
	query.Add("js_code", code)
	query.Add("grant_type", "authorization_code")
	base.RawQuery = query.Encode()
	req.SetRequestURI(base.String())

	err = fasthttp.Do(req, resp)
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode() != http.StatusOK {
		panic("wx request is not statusOK")
	}

	res := new(wxRespStruct)

	err = json.Unmarshal(resp.Body(), res)
	if err != nil {
		panic(err.Error())
	}
	if res.Errcode != 0 {
		panic(error2.NewHttpError(500, strconv.Itoa(res.Errcode), res.Errmsg))
	}
	return res
}
