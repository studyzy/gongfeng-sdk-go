# 关注人

项目关注

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

### 关注者列表

返回指定项目的关注人列表


```
GET /api/v3/projects/:id/watchers
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| page       | integer（可选）        | 分页（默认值：1）                                 |
| per_page   | integer（可选）        | 默认页面大小（默认值：20，最大值：100）                    |


**返回值**


```json
[
  {
        "project_id": 6563,
        "mute": true,
        "user": {
            "id": 11323,
            "username": "git_user1",
            "web_url": "http://git.code.tencent.com/u/git_user1",
            "name": "git_user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
        }
    },
  {
    "project_id": 6563,
    "mute": false,
    "user": {
      "id": 11325,
      "username": "git_user2",
            "web_url": "http://git.domain..com/u/git_user2",
            "name": "git_user2",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/11325/a75ba3738c7a409cab1d15dd993149aa.jpg"
    }
  }
]
```

是否关注给定项目

### 查看当前用户是否关注指定项目


```
GET /api/v3/projects/:id/watch
```


**参数**



| 参数   | 类型                 | 描述                                        |
| ---- | ------------------ | ----------------------------------------- |
| id   | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |


**返回值**


### true 或 false

关注项目

### 关注指定项目


```
PUT /api/v3/projects/:id/watch
```


**参数**



| 参数     | 类型                 | 描述                                               |
| ------ | ------------------ | ------------------------------------------------ |
| id     | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH          |
| mute   | boolean            | 是否静音，mute = true将不会收到除参与和订阅以外通知，默认mute = false   |


**返回值**


```json
{
  "project_id": 2564,
  "mute": false,
  "user": {
    "id": 11323,
    "username": "git_user1",
    "web_url": "http://git.code.tencent.com/u/git_user1",
    "name": "git_user1",
    "state": "active",
    "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
  }
}
```

### 取消关注项目

取消关注指定项目


```
DELETE /api/v3/projects/:id/watch
```


**参数**



| 参数   | 类型                 | 描述                                        |
| ---- | ------------------ | ----------------------------------------- |
| id   | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |


**返回值**


### 200 或 相关状态码
