# 缺陷单

缺陷

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

### 新建缺陷

给指定项目创建缺陷


```
POST /api/v3/projects/:id/issues
```


**参数**



| 参数             | 类型                 | 描述                                        |
| -------------- | ------------------ | ----------------------------------------- |
| id             | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| title          | string             | 缺陷标题                                      |
| grade          | Integer（可选）        | 权重：取值范围 1~10                              |
| description    | string（可选）         | 缺陷描述                                      |
| assignee_ids   | string（可选）         | 处理人唯一标识，允许多个，以,分隔，最多：10                   |
| milestone_id   | integer（可选）        | 里程碑 唯一标识                                  |
| labels         | string （可选）        | 缺陷标签，允许多个，以,分隔，最多：10                      |


**返回值**


```json
{
    "labels": [],
    "milestone": null,
    "id": 3253,
    "project_id": 6368,
    "title": "test-02",
    "state": "opened",
    "resolve_state": "accepted",
    "grade": null,
    "iid": 8,
    "description": null,
    "created_at": "2017-08-13T08:33:31+0000",
    "updated_at": "2017-08-13T08:33:31+0000",
    "author": {
        "id": 1055,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/1055/a75ba2727c7a409cab1d15dd993149aa.jpg"
    },
    "assignees": [],
    "assignee": null
}
```

修改缺陷

### 编辑给定项目的某个缺陷


```
PUT /api/v3/projects/:id/issues/:issue_id
```


**参数**



| 参数              | 类型                 | 描述                                        |
| --------------- | ------------------ | ----------------------------------------- |
| id              | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| issue_id        | integer            | 缺陷唯一标识                                    |
| title           | string（可选）         | 缺陷标题                                      |
| resolve_state   | string（可选）         | 解决状态：可选 resolved,accepteddenied           |
| grade           | Integer（可选）        | 权重：取值范围 1~10                              |
| description     | string（可选）         | 缺陷描述                                      |
| assignee_ids    | string（可选）         | 处理人唯一标识，允许多个，以,分隔，最大值：10                  |
| milestone_id    | integer（可选）        | 里程碑的唯一标识                                  |
| labels          | string （可选）        | 缺陷标签，允许多个，以,分隔，最大值：10                     |
| state_event     | string （可选）        | 缺陷事件，可选reopen、close                       |


**返回值**


```json
{
    "labels": [],
    "milestone": null,
    "id": 3253,
    "project_id": 6368,
    "title": "test-0202",
    "state": "opened",
    "resolve_state": "accepted",
    "grade": null,
    "iid": 8,
    "description": null,
    "created_at": "2017-08-13T08:33:31+0000",
    "updated_at": "2017-08-13T08:33:31+0000",
    "author": {
        "id": 1055,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/1055/a75ba2727c7a409cab1d15dd993149aa.jpg"
    },
    "assignees": [],
    "assignee": null
}
```

### 获取用户创建缺陷列表

返回用户创建的缺陷列表，支持搜索、分页


```
GET /api/v3/issues
```


```
GET /api/v3/issues?state=opened
```


```
GET /api/v3/issues?state=closed
```


```
GET /api/v3/issues?labels=foo
```


```
GET /api/v3/issues?labels=foo,bar
```


```
GET /api/v3/issues?labels=foo,bar&state=opened
```


**参数**



| 参数              | 类型            | 描述                                               |
| --------------- | ------------- | ------------------------------------------------ |
| resolve_state   | string（可选）    | 解决状态：可选 resolved,accepteddenied                  |
| grade           | Integer（可选）   | 权重：取值范围 1~10                                     |
| state           | string （可选）   | 缺陷状态，可选opened, closed                            |
| labels          | string （可选）   | 标签，允许多个，以,分隔，最大值：10                              |
| page            | integer（可选）   | 分页（默认值：1）                                        |
| per_page        | integer（可选）   | 默认页面大小（默认值：20，最大值：100）                           |
| order_by        | string （可选）   | 排序字段，允许按 created_at,updated_at排序（默认created_at）   |
| sort            | string （可选）   | 排序方式，允许 asc,desc（默认desc）                         |


**返回值**


```json
[
  {
        "labels": [],
        "milestone": null,
        "id": 3253,
        "project_id": 6368,
        "title": "test-0202",
        "state": "opened",
        "resolve_state": "accepted",
        "grade": null,
        "iid": 8,
        "description": null,
        "created_at": "2017-08-13T08:33:31+0000",
        "updated_at": "2017-08-13T08:37:16+0000",
        "author": {
            "id": 1055,
            "username": "git_user1",
            "web_url": "http://git.code.tencent.com/u/git_user1",
            "name": "git_user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/1055/a75ba2738c7a409cab1d15dd993149aa.jpg"
        },
        "assignees": [],
        "assignee": null
    },
   {
        "labels": [
            "bug"
        ],
        "milestone": null,
        "id": 29360,
        "project_id": 45663,
        "title": "test失败",
        "state": "opened",
        "resolve_state": "accepted",
        "grade": null,
        "iid": 206,
        "description": "null",
        "created_at": "2017-07-02T01:52:48+0000",
        "updated_at": "2017-07-02T01:52:48+0000",
        "author": {
            "id": 1055,
            "username": "git_user1",
            "web_url": "http://git.code.tencent.com/u/git_user1",
            "name": "git_user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/18604/a75ba2727c7a409cab1d15dd993149aa.jpg"
        },
        "assignees": [],
        "assignee": null
    },
]
```

### 获取项目缺陷列表

返回项目的缺陷列表，支持搜索、分页


```
GET /api/v3/projects/:id/issues
```


```
GET /api/v3/projects/:id/issues?state=opened
```


```
GET /api/v3/projects/:id/issues?state=closed
```


```
GET /api/v3/projects/:id/issues?labels=foo
```


```
GET /api/v3/projects/:id/issues?labels=foo,bar
```


```
GET /api/v3/projects/:id/issues?labels=foo,bar&state=opened
```


```
GET /api/v3/projects/:id/issues?milestone=1.0.0
```


```
GET /api/v3/projects/:id/issues?milestone=1.0.0&state=opened
```


```
GET /api/v3/projects/:id/issues?iid=42
```


**参数**



| 参数               | 类型                            | 描述                                                                                                             |
| ---------------- | ----------------------------- | -------------------------------------------------------------------------------------------------------------- |
| id               | integer 或 string              | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                                                                        |
| iid              | integer （可选）                  | 缺陷的项目编号                                                                                                        |
| state            | string （可选）                   | 缺陷状态，可选opened, closed                                                                                          |
| labels           | string （可选）                   | 标签，允许多个，以,分隔，最多：10                                                                                             |
| milestone        | string （可选）                   | 里程碑标题                                                                                                          |
| order_by         | string （可选）                   | 排序字段，允许按created_at,updated_at排序（默认created_at）                                                                  |
| sort             | string （可选）                   | 排序方式，允许asc,desc（默认desc）                                                                                        |
| page             | integer（可选）                   | 分页（默认值：1）                                                                                                      |
| per_page         | yyyy-MM-dd'T'HH:mm:ssZ （可选）   | 默认页面大小（默认值：20，最大值：100）                                                                                         |
| created_after    | integer（可选）                   | 返回给定时间及之后创建的问题；例如2019-03-25T00:10:19+0000 或 2019-03-25T00:10:19+0800，时间参数必须转码，如"2019-03-25T00:10:19%2B0800"）   |
| created_before   | yyyy-MM-dd'T'HH:mm:ssZ （可选）   | 返回给定时间及之前创建的问题；例如2019-03-25T00:10:19+0000 或 2019-03-25T00:10:19+0800，时间参数必须转码，如"2019-03-25T00:10:19%2B0800"）   |


**返回值**


```json
[
  {
    "labels": [],
    "milestone": null,
    "id": 12305,
    "project_id": 6368,
    "title": "test 消息提醒",
    "state": "closed",
    "resolve_state": "accepted",
    "grade": null,
    "iid": 1,
    "description":: "扫二维码切换账号失败，手动输入密码切换失败",
    "created_at": "2017-01-21T16:16:02+0000",
    "updated_at": "2017-01-21T16:16:02+0000",
    "author": {
      "id": 1105,
            "username": "git_user1",
            "web_url": "http://git.code.tencent.com/u/git_user1",
            "name": "git_user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/1105/a75ba2738c7a409cab1d15dd993149aa.jpg"
        },
    "assignees": [],
    "assignee": null,
  },
  {
    "labels": [
      "bug"
    ],
    "milestone": null,
        "id": 23415,
        "project_id": 6368,
        "title": "ok",
        "state": "opened",
        "resolve_state": "accepted",
        "grade": null,
        "iid": 7,
        "description": "null",
        "created_at": "2017-01-23T01:38:10+0000",
        "updated_at": "2017-01-24T06:15:53+0000",
        "author": {
            "id": 1105,
            "username": "git_user1",
            "web_url": "http://git.code.tencent.com/u/git_user1",
            "name": "git_user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/1106/a75ba2727c7a409cab1d15dd993149aa.jpg"
        },
        "assignees": [],
        "assignee": null
    },
]
```

### 查看指定缺陷

返回指定项目的某个缺陷，需要有该项目 guest 权限


```
GET /api/v3/projects/:id/issues/:issue_id
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| issue_id   | integer            | 缺陷唯一标识                                    |


**返回值**


```json
{
  "labels": [],
  "milestone": null,
  "id": 12305,
  "project_id": 6368,
  "title": "切换失败",
  "state": "closed",
  "resolve_state": "accepted",
  "grade": null,
  "iid": 180,
  "description": "扫二维码切换账号失败，手动输入密码切换失败",
  "created_at": "2017-01-21T16:16:02+0000",
  "updated_at": "2017-01-21T16:16:02+0000",
  "author": {
      "id": 1105,
      "username": "git_user1",
      "web_url": "http://git.code.tencent.com/u/git_user1",
      "name": "git_user1",
      "state": "active",
      "avatar_url": "git.code.tencent.com/uploads/user/avatar/1105/a75ba2737c7a409cab1d15dd993149aa.jpg"
  },
  "assignees": [],
  "assignee": null,
}
```

### 判断是否订阅指定项目的某个缺陷

在项目内，判断对某个指定的缺陷是否订阅了


```
GET /api/v3/projects/:id/issues/{issue_id}/subscribe
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| issue_id   | integer            | 缺陷 唯一标识                                   |


注意：有查看缺陷的权限，默认返回 true，除非取消订阅过

**返回值**


true 或 false
### 订阅指定项目的某个缺陷

在项目中订阅某个指定的缺陷


```
PUT /api/v3/projects/:id/issues/{issue_id}/subscribe
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| issue_id   | integer            | 缺陷唯一标识                                    |


**返回值**


### 200 或相关状态码
取消订阅给定项目的某个缺陷

### 在项目里，取消订阅某个指定的缺陷


```
PUT /api/v3/projects/:id/issues/{issue_id}/unsubscribe
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| issue_id   | integer            | 缺陷唯一标识                                    |


**返回值**


### 200 或相关状态码
