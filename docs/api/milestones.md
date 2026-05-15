# 里程碑


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

新增里程碑

### 在指定项目内新增里程碑


```
POST /api/v3/projects/:id/milestones
```


**参数**



| 参数            | 类型                 | 描述                                        |
| ------------- | ------------------ | ----------------------------------------- |
| id            | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| title         | string             | 里程碑的标题                                    |
| description   | string（可选）         | 里程碑的描述                                    |
| due_date      | datetime（可选）       | 到期时间                                      |


**返回值**


```json
{
  "id": 1125,
  "project_id": 12321,
  "title": "年前遗留bug",
  "state": "closed",
  "iid": 3,
  "due_date": null,
  "created_at": "2013-03-22T02:12:49+0000",
  "updated_at": "2013-03-22T02:12:49+0000",
  "description": null
}
```

编辑里程碑

### 编辑项目内指定的某个里程碑


```
PUT /api/v3/projects/:id/milestones/:milestone_id
```


**参数**



| 参数             | 类型                 | 描述                                        |
| -------------- | ------------------ | ----------------------------------------- |
| id             | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| milestone_id   | integer            | 里程碑的唯一标识                                  |
| title          | string             | 里程碑的标题                                    |
| description    | string（可选）         | 里程碑的描述                                    |
| due_date       | datetime（可选）       | 到期时间                                      |
| state_event    | string （可选）        | 里程碑的事件，可选active、close                     |


**返回值**


```json
{
  "id": 1125,
  "project_id": 12321,
  "title": "年前遗留待验证bug",
  "state": "closed",
  "iid": 3,
  "due_date": null,
  "created_at": "2015-03-22T02:12:49+000",
  "updated_at": "2015-03-22T02:32:23+0000",
  "description": null
}
```

### 返回指定里程碑

返回指定项目的某个里程碑，需要有该项目的访问权限


```
GET /api/v3/projects/:id/milestones/:milestone_id
```


**参数**



| 参数             | 类型                 | 描述                                        |
| -------------- | ------------------ | ----------------------------------------- |
| id             | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| milestone_id   | integer            | 里程碑的唯一标识                                  |


**返回值**


```json
{
  "id": 1125,
  "project_id": 12321,
  "title": "年前遗留待验证bug",
  "state": "closed",
  "iid": 3,
  "due_date": null,
  "created_at": "2015-03-22T02:12:49+0000",
  "updated_at": "2015-03-22T02:32:23+0000",
  "description": null
}
```

### 删除里程碑

删除指定项目的某个里程碑


```
DELETE /api/v3/projects/:id/milestones/:milestone_id
```


**参数**



| 参数             | 类型                 | 描述                                        |
| -------------- | ------------------ | ----------------------------------------- |
| id             | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| milestone_id   | integer            | 里程碑的唯一标识                                  |


**返回值**


### 200 或相关状态码
返回某个里程碑下的所有缺陷

### 返回项目内某个指定里程碑下面的所有缺陷


```
GET /api/v3/projects/:id/milestones/:milestone_id/issues
```


**参数**



| 参数             | 类型                 | 描述                                        |
| -------------- | ------------------ | ----------------------------------------- |
| id             | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| milestone_id   | integer            | 里程碑的唯一标识                                  |


**返回值**


```json
[
  {
    "labels": [
      "ccc"
    ],
    "milestone": {
      "id": 107,
      "project_id": 5364,
      "title": "test",
      "state": "active",
      "iid": 1,
      "due_date": null,
      "created_at": "2015-03-06T05:35:35+0000",
      "updated_at": "2015-03-06T05:35:35+0000",
      "description": ""
    },
    "id": 3056,
    "project_id": 5364,
    "title": "test2",
    "state": "opened",
    "iid": 2,
    "description": "test2",
    "created_at": "2012-11-15T03:10:21+0000",
    "updated_at": "2012-11-15T03:10:21+0000",
    "author": {
      "id": 11323,
      "username": "git_user1",
      "web_url": "http://git.code.tencent.com/u/git_user1",
      "name": "git_user1",
      "state": "active",
      "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    }
    "assignees": [],
    "assignee": null,
  },
  {
    "labels": [
      "dd",
      "ccc"
    ],
    "milestone": {
      "id": 662,
      "project_id": 5364,
      "title": "test",
      "state": "active",
      "iid": 1,
      "due_date": null,
      "created_at": "2012-10-12T09:25:05+0000",
      "updated_at": "2012-10-12T09:25:05+0000",
      "description": ""
    },
    "id": 1400,
    "project_id": 5364,
    "title": "testaa",
    "state": "opened",
    "iid": 1,
    "description": "testaa",
    "created_at": "2013-06-15T09:22:09+0000",
    "updated_at": "2015-03-21T03:09:55+0000",
    "author": {
      "id": 11323,
      "username": "git_user1",
      "web_url": "http://git.code.tencent.com/u/git_user1",
      "name": "git_user1",
      "state": "active",
      "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    }
    "assignees": [{
        "id": 8771,
        "username": "git_user2",
        "web_url": "http://git.code.tencent.com/u/git_user2",
        "name": "git_user2",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/8771/a75ba2638c7a409cab1d15dd993149aa.jpg"
      }
     ],
      "assignee": [{
        "id": 8771,
        "username": "git_user2",
        "web_url": "http://git.code.tencent.com/u/git_user2",
        "name": "git_user2",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/8771/a75ba2638c7a409cab1d15dd993149aa.jpg"
        }
      ]
    }
  ]
```


### 返回里程碑列表

返回某个指定项目的里程碑列表


```
GET /api/v3/projects/:id/milestones
```


```
GET /api/v3/projects/:id/milestones?iid=42
```


**参数**



| 参数         | 类型                 | 描述                                                           |
| ---------- | ------------------ | ------------------------------------------------------------ |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                      |
| iid        | integer （可选）       | iid                                                          |
| state      | string （可选）        | 里程碑的状态，可选active, closed                                      |
| order_by   | string （可选）        | 排序字段， 允许按id,due_date,created_at,updated_at排序（默认created_at）   |
| sort       | string （可选）        | 排序方式， 允许 asc,desc（默认desc）                                    |
| page       | integer（可选）        | 分页（default：1）                                                |
| per_page   | integer（可选）        | 默认页面大小（default：20，max：100）                                   |


注意：iid != null 时state、order_by、sort、page、per_page参数不生效

**返回值**


```json
[
  {
    "id": 1125,
    "project_id": 12321,
    "title": "年前遗留待修复bug",
    "state": "active",
    "iid": 5,
    "due_date": null,
    "created_at": "2015-03-22T02:12:49+0000",
    "updated_at": "2015-03-22T02:32:23+0000",
    "description": "jjj"
  },
  {
    "id": 1104,
    "project_id": 12321,
    "title": "待解决bug",
    "state": "active",
    "iid": 13,
    "due_date": null,
    "created_at": "2014-11-15T03:20:33+0000",
    "updated_at": "2015-03-15T11:25:05+0000",
    "description": "hh"
  }
]
```
