package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"sync"
)

const SendUrl string = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"

type startmap struct {
	Userid    string `json:"userid"`
	Projectid string `json:"projectid"`
	CurProc   struct {
		Type  int64  `json:"type"`
		Name  string `json:"name"`
		Index int64  `json:"index"`
	}
}

func (s *Service) ActionStart(c *gin.Context) (int, interface{}) {
	template := "tuLTzAObhYDJYKxuowlx1pt5SyGLq2cCiZKqcZXZ2pw"
	token, err := s.GetToken()
	if err != nil {
		return s.makeErrJSON(500, 50003, err.Error())
	}

	start_map := new(startmap)
	err = c.ShouldBindJSON(start_map)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}

	userid := start_map.Userid
	projectid := start_map.Projectid

	//权限检查
	role, err := s.checkProject(projectid, userid)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	if role != 1 {
		return s.makeErrJSON(403, 40301, "limited access")
	}

	//获取项目所有人员
	users := make([]Project_User, 50, 100) // 其实写不写大小，在做完find后自动就更改了，作者懒的改了
	s.DB.Where(&Project_User{ProjectID: projectid}).Find(&users)

	//通过这两个方法从连接池中获取一个空的实例，可以实现连接复用，提高性能
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	//构造请求地址
	url, _ := s.makeUrl(SendUrl, Query{"access_token", token})
	req.SetRequestURI(url)
	//设置请求头
	req.Header.SetContentType("application/json")
	// 重复设置方法可以通过set函数内部的append方法去截断前一次设置的方法
	req.Header.SetMethod("POST")

	//响应体
	type resultstruct struct {
		Userid  string
		Errcode int64  `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}
	failed := make([]*resultstruct, 0, 20)

	//阻塞进程
	wait := sync.WaitGroup{}
	wait.Add(1)
	for _, v := range users {
		/*开始goroutine*/
		wait.Add(1)
		go func(userid string) {
			defer func() {
				//执行结束时减少阻塞变量
				wait.Done()
			}()
			//活动标题
			arr := [6]string{"节目", "互动", "颁奖", "致辞", "开场", "结束"}
			var name string
			if start_map.CurProc.Type == 0 {
				name = fmt.Sprintf("环节%d-节目-%s", start_map.CurProc.Index, start_map.CurProc.Name)
			} else {
				name = fmt.Sprintf("环节%d-%s", start_map.CurProc.Index, arr[start_map.CurProc.Type])
			}
			//构造发送推送的map // 不要问我为什么这么写，我也不知道我为什么要一次性套娃全套掉。。。
			sendmap := struct {
				Touser      string `json:"touser"`
				Template_id string `json:"template_id"`
				Data        struct {
					Thing1 struct {
						Value string `json:"value"`
					} `json:"thing1"`
					Thing5 struct {
						Value string `json:"value"`
					} `json:"thing5"`
				} `json:"data"`
			}{
				Template_id: template,
				Data: struct {
					Thing1 struct {
						Value string `json:"value"`
					} `json:"thing1"`
					Thing5 struct {
						Value string `json:"value"`
					} `json:"thing5"`
				}{
					Thing1: struct {
						Value string `json:"value"`
					}{
						Value: name,
					},
					Thing5: struct {
						Value string `json:"value"`
					}{
						Value: "该环节即将开始，请相关人员做好准备！",
					},
				},
			}
			result := new(resultstruct)
			result.Errcode = 0
			result.Userid = userid

			//将map转成json
			sendmap.Touser = userid
			sendjson, err := json.Marshal(sendmap)
			//fmt.Println(sendmap)
			//fmt.Println(string(sendjson))
			if err != nil {
				result := &resultstruct{Errcode: 50006, Errmsg: err.Error(), Userid: userid}
				failed = append(failed, result)
				return
			}
			//设置请求体
			req.SetBody(sendjson)
			/*向端口请求*/
			//请求出错的时候重复10次
			if err := fasthttp.Do(req, resp); err != nil {
				i := 0
				for err != nil || i <= 10 {
					err = fasthttp.Do(req, resp)
					i++
				}
				if err != nil {
					result := &resultstruct{Errcode: 50005, Errmsg: err.Error(), Userid: userid}
					failed = append(failed, result)
					return
				}
			}
			//获取响应体
			b := resp.Body()
			_ = json.Unmarshal(b, &result)
			if result.Errcode != 0 {
				//添加到失败名单里
				failed = append(failed, result)
			}
			return
		}(v.UserID)
	}
	wait.Done()
	wait.Wait()
	/*结束goroutine*/
	return s.makeSuccessJSON(failed)
}
