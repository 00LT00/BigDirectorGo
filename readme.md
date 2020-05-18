# 我是大导演

项目 -> 环节

project -> process





## 用户 `/user`

注册用户  PUT

```json
{
    "openid": "123456",
    "username": "asdf",
    "qqnum": "150521321", // 可选
    "avatar":"asdfasdfsa",
   	"phonenum":"13456"
}
```



获取用户信息  GET

示例：GET `/user/123456`

结果：

```json
{
    "data": {
        "openid": "123456",
        "username": "asdf",
        "phonenum": "13456",
        "avatar": "asdfdlfajsldfjsaldkf",
        "qqnum": "150521321",
        "CreatedAt": "2020-05-16T21:30:44+08:00",
        "UpdatedAt": "2020-05-16T21:55:52+08:00",
        "DeletedAt": null
    },
    "error": 0,
    "msg": "success"
}
```



修改用户信息 PUT(小程序不支持patch)

示例：PUT`/user/123456`

```json
{
    "openid": "123456", //必选
    "username": "asdf",
    "qqnum": "150521321",
    "phonenum": "13456",
    "avatar": "asdfdlfajsldfjsaldkf"
}
```

结果：

```json
{
    "data": {
        "openid": "123456",
        "username": "asdf",
        "phonenum": "13456",
        "avatar": "11111",
        "qqnum": "150521321",
        /*无效数据
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null
        */
    },
    "error": 0,
    "msg": "success"
}
```



获取用户参与的项目 GET `/user/{{userid}}/project`

示例：GET `/user/123456/project`

```json
{
    "data": {
        "UserID": "123456",
        "ProjectList": [
            {
                "ID": 6,
                "CreatedAt": "2020-05-18T01:39:45+08:00",
                "UpdatedAt": "2020-05-18T01:39:45+08:00",
                "DeletedAt": null,
                "userid": "123456",
                "name": "asdfadfasf",
                "ProjectID": "af35bc80-88ca-400f-9633-74b2c1cecb62",
                "DirectorName": "asdf"
            },
            {
                "ID": 5,
                "CreatedAt": "2020-05-18T01:39:36+08:00",
                "UpdatedAt": "2020-05-18T01:39:36+08:00",
                "DeletedAt": null,
                "userid": "123456",
                "name": "ffffffff",
                "ProjectID": "dae5b3ac-06ea-471c-8a4d-b48e6b8addfb",
                "DirectorName": "asdf"
            },
            {
                "ID": 9,
                "CreatedAt": "2020-05-18T01:40:39+08:00",
                "UpdatedAt": "2020-05-18T01:40:39+08:00",
                "DeletedAt": null,
                "userid": "123456",
                "name": "00000asf",
                "ProjectID": "ffc4285d-eb7d-4b9f-9725-9978b053af5d",
                "DirectorName": "asdf"
            },
            {
                "ID": 10,
                "CreatedAt": "2020-05-18T01:43:13+08:00",
                "UpdatedAt": "2020-05-18T01:43:13+08:00",
                "DeletedAt": null,
                "userid": "123456",
                "name": "00000asf",
                "ProjectID": "129f0a6d-39eb-4a68-a93c-992d23e27a36",
                "DirectorName": "asdf"
            }
        ]
    },
    "error": 0,
    "msg": "success"
}
```







| error |    msg     |                 报错原因                 |
| :---: | :--------: | :--------------------------------------: |
| 40301 |  json错误  | json的格式或者数据不对，比如说必选项为空 |
| 40400 | none user  |          用户id没找到对应的账户          |
| 50000 |  插入错误  |      数据库问题或者有主键重复等错误      |
| 40302 |  json错误  |                 同40301                  |
| 40303 | openid错误 |         地址栏id和json id不一致          |
| 40401 | none user  |                 同40400                  |
| 50001 |  更新错误  |                 详见msg                  |



## 项目 `/project`

新建项目 PUT

示例: PUT `/project/`

```json
{
    "userid": "123456",
    "name": "ffffffff"
}
```

返回结果:

```json
{
    "data": {
        "ID": 4,
        "CreatedAt": "2020-05-17T05:13:17.3518602+08:00",
        "UpdatedAt": "2020-05-17T05:13:17.3518602+08:00",
        "DeletedAt": null,
        "userid": "123456",
        "name": "ffffffff",
        "ProjectID": "e6903829-9c63-4323-93cc-4ed055d526de"
    },
    "error": 0,
    "msg": "success"
}
```



获取项目详情 GET`/project/{{projectid}}/{{userid}}`

示例：GET`/project/9dc83213-7a2e-42f6-8337-dfcb07b9062a/123456`

```json
{
    "data": {
        "ID": 1,
        "CreatedAt": "2020-05-17T14:42:55+08:00",
        "UpdatedAt": "2020-05-17T14:42:55+08:00",
        "DeletedAt": null,
        "userid": "123456",
        "name": "ffffffff", //项目名称
        "ProjectID": "9dc83213-7a2e-42f6-8337-dfcb07b9062a",
        "Role": 0
    },
    "error": 0,
    "msg": "success"
}
```

- 项目的环节 还没做

修改项目信息 PUT `/project/{{projectid}}/{{userid}}`   //userid是当前操作的用户的id

示例：PUT `/project/4d8629ef-e7d2-422a-b7af-3dd85437f7cc/111111` 

111111用户将导演转给了123456用户，前提是111111是项目的原导演，123456是成员

```json
{
    "userid": "123456",
    "name": "00000"
}
```



获取项目的所有成员 GET `/user`

- 标识用户的权限

新建负责人 PUT







## 环节 `/process`

新建环节 PUT

获取环节详情 GET

修改环节信息 PATCH

新建环节负责人 PUT