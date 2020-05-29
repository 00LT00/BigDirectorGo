# 我是大导演

项目 -> 环节

project -> process

[TOC]

## 用户 `/user`

### 注册用户  

  格式：PUT `/`

- 示例 `/user/`
```json
{
    "openid": "123456",
    "username": "asdf",
    "qqnum": "150521321", // 可选
    "avatar":"asdfasdfsa",
    "phonenum":"13456" //可选
}
```



### 获取用户信息  

  格式：GET `/{{userid}}`

- 示例 GET `/user/123456`

- 返回

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



### 修改用户信息 

  格式：PUT`/{{userid}}`(小程序不支持patch)

- 示例 PUT`/user/123456`
  ```json
  {
      "openid": "123456", //必选
      "username": "asdf",
      "qqnum": "150521321",
      "phonenum": "13456",
      "avatar": "asdfdlfajsldfjsaldkf"
  }
  ```

- 返回

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



### 获取用户参与的项目 
  格式：GET `/user/{{userid}}/project`

- 示例 GET `/user/123456/project`

- 返回
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

### 新建项目 
  格式：PUT `/`

- 示例: PUT `/project/`
  ```json
  {
      "userid": "123456",
      "name": "ffffffff"
  }
  ```

- 返回

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



###获取项目详情 
格式：GET`/{{projectid}}/{{userid}}`

- 示例 GET`/project/9dc83213-7a2e-42f6-8337-dfcb07b9062a/123456`

- 返回
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

### 获取项目的环节 
格式：GET `/{{processid}}`

- 示例 `/project/process/f3852e84-130a-4ab8-be69-7fae2628ba3a`

- 返回

  ```json
  {
      "data": [
          {
              "order": 5,
              "process_id": "81da80b8-7f4e-43b8-9d8d-2fa5bb15c7a1",
              "process_name": "环节",
              "process_type": 3,
              "mic_hand": 20,
              "mic_ear": 0,
              "remark": "",
              "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a"
          },
          {
              "order": 5,
              "process_id": "e2e6f9a9-6506-45cb-af0a-2eee881ebfef",
              "process_name": "开始",
              "process_type": 1,
              "mic_hand": 20,
              "mic_ear": 0,
              "remark": "",
              "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a"
          },
          {
              "order": 10,
              "process_id": "9a821cbe-2715-4bed-a883-789ff45b6b5d",
              "process_name": "结束",
              "process_type": 6,
              "mic_hand": 20,
              "mic_ear": 3,
              "remark": "asdfasgsgageaasdfasd",
              "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a"
          }
      ],
      "error": 0,
      "msg": "success"
  }
  ```



### 修改项目名称 
格式：PUT `/pnm/{{projectid}}`

- 示例 PUT `/project/pnm/4d8629ef-e7d2-422a-b7af-3dd85437f7cc` 

  ```json
  {
      "userid": "123456",
      "name": "00000"
  }
  ```

- 返回
  ```json
  {
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "userid": "12111",
        "name": "f",
        "ProjectID": ""
    },
    "error": 0,
    "msg": "success"
  }
  ```

### 修改项目导演
格式：PUT `/uid/{{projectid}}`

- 示例 PUT `/project/uid/4d8629ef-e7d2-422a-b7af-3dd85437f7cc` 

  ```json
  {
    "userid": "123456", // 从userid转给directoruserid
    "directoruserid": "789456"
  }
  ```
- 返回
  ```json
  {
    "data": {
        "userid": "123456",
        "directoruserid": "789456"
    },
    "error": 0,
    "msg": "success"
  }
  ```

### 为项目增加成员 
格式：POST `/member/`
- 示例 `/project/member/` 默认权限为3

  ```json
  {
      "userid": "111111",
      "projectid": "c9c6ce1b-b581-4815-9788-5cc413640ac8"
  }
  ```
- 返回
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



### 获取项目的所有成员 
格式：GET `/user/{{userid=?&projectid=?}}`
- 示例  `/project/user/?userid=111111&projectid=c9c6ce1b-b581-4815-9788-5cc413640ac8` //这里userid指的是执行此操作的用户id

- 返回
  ```json
  {
      "data": [
          {
              "UserID": "oOMHn5QPr690ztewIhWIJkieitGo",
              "UserName": "111",
              "Role": 3,
              "Avatar": "asdfdlfajsldfjsaldkf"
          },
          {
              "UserID": "12111",
              "UserName": "111",
              "Role": 3,
              "Avatar": "asdfdlfajsldfjsaldkf"
          },
          {
              "UserID": "123456",
              "UserName": "111",
              "Role": 1,
              "Avatar": "asdfdlfajsldfjsaldkf"
          }
      ],
      "error": 0,
      "msg": "success"
  }
  ```



## 环节 `/process`



//作废！！！！ - 新建环节 PUT `/{{userid}}`

//示例 `/process/123456`

//```json
//{
//  "process_id": "", //无效字段，填不填任意，id自动生成，以返回值为主
//  "process_name": "结束",
//  "process_type": 6,
//  "order": 10,
//  "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
//  "mic_hand": 20, // 可选
//  "mic_ear": 3, // 可选
//  "remark": "asdfasgsgageaasdfasd" //可选
//}
//```

//返回值

//```json
//{
//  "data": {
//      "order": 10,
//      "process_id": "9a821cbe-2715-4bed-a883-789ff45b6b5d",
//      "process_name": "结束",
//      "process_type": 6,
//      "mic_hand": 20,
//      "mic_ear": 3,
//      "remark": "asdfasgsgageaasdfasd",
//      "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a"
//  },
//  "error": 0,
//  "msg": "success"
//}
//```



### 获取环节详情 
格式：GET `/{{userid=?&processid=?}}`

- 示例 `/process/?processid=e2e6f9a9-6506-45cb-af0a-2eee881ebfef&userid=12111`

- 返回

  ```json
  {
      "data": {
          "order": 3,
          "process_id": "ac776572-e06c-42f7-af7f-e9b513221c05",
          "process_name": "结束",
          "process_type": 3,
          "mic_hand": 215,
          "mic_ear": 0,
          "remark": "",
          "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
          "manager_id": "123456"
      },
      "error": 0,
      "msg": "success"
  }
  ```



### 修改环节信息 
格式：PUT `/{{userid}}`

- 示例 `/process/12111`

  ```json
  {
      "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
      "processes": [
          {
            "process_id":"cf94866d-abaa-4263-ac20-b2ca8cf12c8d", //新环节不用写，但是已经存在的环节一定要写
              "process_type": 6,
              "order": 19,
              "mic_hand": 20,
              "mic_ear": 3
          },
         {
              "process_name": "名字",
              "process_type": 2,
              "order": 10,
              "mic_hand": 2,
              "mic_ear": 3
          },
         {
              "process_type": 4,
              "order": 17,
              "mic_hand": 20,
              "mic_ear": 3
          }
      ]
  }
  ```

- 返回

  ```json
  {
      "data": [
          {
              "order": 1,
              "process_id": "459ccef2-5309-4f97-9bd3-97291cb76c3a",
              "process_name": "开始",
              "process_type": 1,
              "mic_hand": 20,
              "mic_ear": 0,
              "remark": "",
              "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
              "manager_id": ""
          },
          {
              "order": 2,
              "process_id": "075bc1a0-b499-4efd-ba3b-ce6f31e31c13",
              "process_name": "青花瓷",
              "process_type": 2,
              "mic_hand": 2,
              "mic_ear": 0,
              "remark": "",
              "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
              "manager_id": ""
          },
          {
              "order": 3,
              "process_id": "ac776572-e06c-42f7-af7f-e9b513221c05",
              "process_name": "结束",
              "process_type": 3,
              "mic_hand": 215,
              "mic_ear": 0,
              "remark": "",
              "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
              "manager_id": "123456"
          }
      ],
      "error": 0,
      "msg": "success"
  }
  ```

### 设置环节负责人 
格式：POST`/{{userid}}`

- 示例 `/process/12111`

  ```json
  {
      "manager_id": "123456",
      "process_id": "ac776572-e06c-42f7-af7f-e9b513221c05"
  }
  ```

- 返回

  ```json
  {
      "data": {
          "order": 3,
          "process_id": "ac776572-e06c-42f7-af7f-e9b513221c05",
          "process_name": "结束",
          "process_type": 3,
          "mic_hand": 215,
          "mic_ear": 0,
          "remark": "",
          "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
          "manager_id": "123456"
      },
      "error": 0,
      "msg": "success"
  }
  ```

## 管理人员 `/worker`

### 设置除节目以外的其他负责人 

格式：PUT `/{{userid}}`

- 示例：`PUT /worker/12111`

  ```json
  {
      "worker_id": "123456",
      "project_id": "dd406d2d-dbe2-42f6-a2b3-c5b36eb7",
      "role": 6
  }
  ```

- 返回

  ```json
  {
      "data": {
          "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
          "worker_id": "123456",
          "role": 5
      },
      "error": 0,
      "msg": "success"
  }
  ```

### 获取项目除导演以外的所有负责人

格式：GET `/{{projectid=?&userid=?}}`

- 示例 `/worker/?userid=12111&projectid=f3852e84-130a-4ab8-be69-7fae2628ba3a`

- 返回

  ```json
  {
      "data": {
          "Workers": [
              {
                  "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
                  "worker_id": "111",
                  "role": 4,
                  "WorkerName": "asdf",
                  "PhoneNum": "13456"
              },
              {
                  "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
                  "worker_id": "123456",
                  "role": 5,
                  "WorkerName": "asdf",
                  "PhoneNum": "123"
              },
              {
                  "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
                  "worker_id": "123456",
                  "role": 6,
                  "WorkerName": "asdf",
                  "PhoneNum": "123"
              },
              {
                  "project_id": "f3852e84-130a-4ab8-be69-7fae2628ba3a",
                  "worker_id": "123456",
                  "role": 7,
                  "WorkerName": "asdf",
                  "PhoneNum": "123"
              }
          ],
          "Managers": [
              {
                  "manager_id": "111",
                  "process_id": "0ea1a6d5-5c95-40ec-8366-039d271c8539",
                  "ManagerName": "asdf",
                  "PhoneNum": "13456"
              },
              {
                  "manager_id": "12111",
                  "process_id": "3ba76e12-276e-42d2-8beb-ec4b1d289d80",
                  "ManagerName": "wzl",
                  "PhoneNum": "13456"
              },
              {
                  "manager_id": "123456",
                  "process_id": "cf94866d-abaa-4263-ac20-b2ca8cf12c8d",
                  "ManagerName": "asdf",
                  "PhoneNum": "123"
              }
          ]
      },
      "error": 0,
      "msg": "success"
  }
  ```

  