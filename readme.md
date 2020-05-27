# 我是大导演

项目 -> 环节

project -> process





## 用户 `/user`

注册用户  PUT `/user/`

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
        "UserID": "12111",
        "ProjectList": [
            {
                "ID": 16,
                "CreatedAt": "2020-05-20T22:17:44+08:00",
                "UpdatedAt": "2020-05-24T21:34:58+08:00",
                "DeletedAt": null,
                "userid": "123456",
                "name": "f",
                "ProjectID": "c9c6ce1b-b581-4815-9788-5cc413640ac8",
                "DirectorName": "asdf",
                "Role": 3,
                "MemberNum": 3
            },
            {
                "ID": 18,
                "CreatedAt": "2020-05-20T22:18:24+08:00",
                "UpdatedAt": "2020-05-24T20:35:17+08:00",
                "DeletedAt": null,
                "userid": "12111",
                "name": "f",
                "ProjectID": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
                "DirectorName": "wzl",
                "Role": 1,
                "MemberNum": 1
            },
            {
                "ID": 17,
                "CreatedAt": "2020-05-20T22:18:20+08:00",
                "UpdatedAt": "2020-05-24T20:31:45+08:00",
                "DeletedAt": null,
                "userid": "12111",
                "name": "f",
                "ProjectID": "8fb06a78-99cf-4c29-8f2d-76c256ba9fb2",
                "DirectorName": "wzl",
                "Role": 1,
                "MemberNum": 1
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

为项目增加成员 POST `/member/`
示例 `/project/member/` 默认权限为3
```json
{
    "userid": "111111",
    "projectid": "c9c6ce1b-b581-4815-9788-5cc413640ac8"
}
```
返回结果：
```json
{
    "data": {
        "ID": 22,
        "CreatedAt": "2020-05-20T23:47:53.1962117+08:00",
        "UpdatedAt": "2020-05-20T23:47:53.1962117+08:00",
        "DeletedAt": null,
        "userid": "111111",
        "projectid": "c9c6ce1b-b581-4815-9788-5cc413640ac8"
    },
    "error": 0,
    "msg": "success"
}
```



获取项目的所有成员 GET `/{{userid=?}}&{{projectid=}}`
示例  `/project/?userid=111111&projectid=c9c6ce1b-b581-4815-9788-5cc413640ac8` //这里userid指的是执行此操作的用户id
返回结果：
```json
{
    "data": [
        {
            "UserID": "12111",
            "UserName": "wzl",
            "Role": 3
        },
        {
            "UserID": "123456",
            "UserName": "asdf",
            "Role": 1
        }
    ],
    "error": 0,
    "msg": "success"
}
```

- 标识用户的权限







## 环节 `/process`

- 新建环节 PUT `/process/{{userid}}`

示例 `/process/123456`

```json
{
    "process_id": "", //无效字段，填不填任意，id自动生成，以返回值为主
    "process_name": "结束",
    "process_type": 6,
    "order": 10,
    "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
    "mic_hand": 20, // 可选
    "mic_ear": 3, // 可选
    "remark": "asdfasgsgageaasdfasd" //可选
}
```

返回值

```json
{
    "data": {
        "order": 10,
        "process_id": "9a821cbe-2715-4bed-a883-789ff45b6b5d",
        "process_name": "结束",
        "process_type": 6,
        "mic_hand": 20,
        "mic_ear": 3,
        "remark": "asdfasgsgageaasdfasd",
        "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a"
    },
    "error": 0,
    "msg": "success"
}
```



- 获取环节详情 GET `/process/{{userid = }}&{{processid = }}`

示例 `/process/?processid=e2e6f9a9-6506-45cb-af0a-2eee881ebfef&userid=12111`

返回

```json
{
    "data": {
        "order": 5,
        "process_id": "e2e6f9a9-6506-45cb-af0a-2eee881ebfef",
        "process_name": "开始",
        "process_type": 1,
        "mic_hand": 20,
        "mic_ear": 0,
        "remark": "",
        "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a"
    },
    "error": 0,
    "msg": "success"
}
```



修改环节信息 POST `/process/{{userid}}`

示例 `/process/12111`

注：如果和数据库中一样，也会返回500的错误，所以如果实际未发生更改则不要请求，或者进行判断，或者写issues我改。。。

```json
{//是否可选和增加环节完全一样
    "process_id": "81da80b8-7f4e-43b8-9d8d-2fa5bb15c7a1",//必须要有
    "process_name": "环节",
    "process_type": 3,
    "order": 5,
    "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
    "mic_hand": 20
}
```

返回

```json
{
    "data": {
        "order": 5,
        "process_id": "81da80b8-7f4e-43b8-9d8d-2fa5bb15c7a1",
        "process_name": "环节",
        "process_type": 3,
        "mic_hand": 20,
        "mic_ear": 0,
        "remark": "",
        "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a"
    },
    "error": 0,
    "msg": "success"
}
```



新建环节负责人 PUT