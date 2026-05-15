# 评论


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子：/api/v3/projects/diaspora%2Fdiaspora (/ 编码 %2F)

### 创建合并请求的评论

在项目内给某个指定合并请求新增评论


```
POST /api/v3/projects/:id/merge_requests/:merge_request_id/notes
```


**参数**



| 参数                 | 类型                  | 描述                                                          |
| ------------------ | ------------------- | ----------------------------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                     |
| merge_request_id   | integer             | 合并请求 id                                                     |
| body               | string              | 评论的内容                                                       |
| path               | string（可选）          | 文件路径                                                        |
| line               | string（可选）          | 行号                                                          |
| line_type          | string（可选）          | 变更类型，可选old、new                                              |
| reviewer_state     | string（可选）          | 单文件评审的状态，可选：（approved | change_required | change_denied ）   |


**返回值**


```json
{
    "id": 265654,
    "body": "查看中",
    "attachment": null,
    "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "created_at": "2015-03-22T06:52:05+0000",
    "system": true
}
```

### 编辑合并请求的评论

在项目内编辑某个指定合并请求的指定评论


```
PUT /api/v3/projects/:id/merge_requests/:merge_request_id/notes/:note_id
```


**参数**



| 参数                 | 类型                  | 描述                                                          |
| ------------------ | ------------------- | ----------------------------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                     |
| merge_request_id   | integer             | 合并请求 id                                                     |
| note_id            | integer             | 评论 id                                                       |
| body               | string              | 评论的内容                                                       |
| reviewer_state     | string（可选）          | 单文件评审的状态，可选：（approved | change_required | change_denied ）   |


**返回值**


```json
{
    "id": 265451,
    "body": "有待修改ing",
    "attachment": null,
    "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "created_at": "2015-03-22T06:52:05+0000",
    "system": true
}
```

### 查询合并请求的单个评论

在项目内查询某个指定合并请求的指定评论


```
GET /api/v3/projects/:id/merge_requests/:merge_request_id/notes/:note_id
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求 id                                   |
| note_id            | integer             | 评论 id                                     |


**返回值**


```json
{
    "id": 32561,
    "body": "ok",
    "attachment": null,
    "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "created_at": "2015-03-22T06:52:05+0000",
    "system": true
}
```

### 查询合并请求的评论列表

在项目内查询某个指定合并请求的评论列表


```
GET /api/v3/projects/:id/merge_requests/:merge_request_id/notes
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求 id                                   |
| page               | integer             | 分页（默认值：1）                                 |
| per_page           | integer             | 默认页面大小（默认值： 20，最大值： 100）                  |


**返回值**


```json
[
    {
        "id": 16545,
        "body": "milestone removed",
        "attachment": null,
        "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
        "created_at": "2015-03-20T10:17:05+0000",
        "system": true
    },
    {
        "id": 16525,
        "body": "Assignee removed",
        "attachment": null,
        "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
        "created_at": "2018-03-20T10:17:05+0000",
        "system": false
    }
]
```

### 查询代码评审的单个评论

在项目内查看某个指定代码评审的指定评论


```
GET /api/v3/projects/:id/reviews/:review_id/notes/:note_id
```


**参数**



| 参数          | 类型                  | 描述                                        |
| ----------- | ------------------- | ----------------------------------------- |
| id          | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| review_id   | integer             | 代码评审 id                                   |
| note_id     | integer             | 评论 id                                     |


**返回值**


```json
{
    "id": 323561,
    "body": "new",
    "attachment": null,
    "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "created_at": "2015-04-25T12:07:41+0000",
    "system": true
}
```

### 查询代码评审的评论

在项目内获取某个指定代码评审的评论


```
GET /api/v3/projects/:id/reviews/:review_id/notes
```


**参数**



| 参数          | 类型                  | 描述                                        |
| ----------- | ------------------- | ----------------------------------------- |
| id          | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| review_id   | integer             | 代码评审的 id                                  |
| page        | integer             | 分页（默认值：1）                                 |
| per_page    | integer             | 默认页面大小（默认值： 20 ，最大值： 100）                 |


**返回值**


```json
[
    {
        "id": 167754,
        "body": "milestone removed",
        "attachment": null,
        "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
        "created_at": "2015-03-20T10:17:05+0000",
        "system": true
    },
    {
        "id": 167755,
        "body": "这里要改一下",
        "attachment": null,
        "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
        "created_at": "2015-03-20T10:17:05+0000",
        "system": false
    }
]
```

### 创建代码评审的评论

在项目内给某个指定代码评审新增评论


```
POST /api/v3/projects/:id/reviews/:review_id/notes
```


**参数**



| 参数               | 类型                  | 描述                                                          |
| ---------------- | ------------------- | ----------------------------------------------------------- |
| id               | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                     |
| review_id        | integer             | 代码评审 id                                                     |
| body             | string              | 评论的内容                                                       |
| path             | string（可选）          | 文件路径                                                        |
| line             | string（可选）          | 行号                                                          |
| line_type        | string（可选）          | 变更类型，可选old、new                                              |
| reviewer_state   | string（可选）          | 单文件评审的状态，可选：（approved | change_required | change_denied ）   |


**返回值**


```json
{
    "id": 1673429,
    "body": "评论",
    "attachment": null,
    "author": {
        "id": 0001,
        "username": "git_user01",
        "web_url": "http://git.code.tencent.com/u/git_user01",
        "name": "git_user01",
        "state": "active",
        "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png"
    },
    "created_at": "2019-04-19T09:58:21+0000",
    "system": false
}
```

### 编辑代码评审的评论

在项目内编辑某个指定代码评审的指定评论


```
PUT /api/v3/projects/:id/reviews/:review_id/notes/note_id
```


**参数**



| 参数               | 类型                  | 描述                                                          |
| ---------------- | ------------------- | ----------------------------------------------------------- |
| id               | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                     |
| review_id        | integer             | 代码评审 id                                                     |
| note_id          | integer             | 评论 id                                                       |
| body             | string              | 评论的内容                                                       |
| reviewer_state   | string（可选）          | 单文件评审的状态，可选：（approved | change_required | change_denied ）   |


**返回值**


```json
{
    "id": 1673431,
    "body": "修改第一次评论",
    "attachment": null,
    "author": {
        "id": 0001,
        "username": "git_user01",
        "web_url": "http://git.code.tencent.com/u/git_user01",
        "name": "git_user01",
        "state": "active",
        "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png"
    },
    "created_at": "2019-04-19T10:03:15+0000",
    "system": false
}
```

### 创建项目的缺陷评论

在项目内给某个指定的缺陷创建评论


```
POST /api/v3/projects/:id/issues/:issue_id/notes
```


**参数**



| 参数         | 类型                  | 描述                                        |
| ---------- | ------------------- | ----------------------------------------- |
| id         | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| issue_id   | integer             | 缺陷 id                                     |
| body       | string              | 评论的内容                                     |


**返回值**


```json
{
    "id": 25546,
    "body": "请尽快处理",
    "attachment": null,
    "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "created_at": "2015-03-22T06:43:36+0000",
    "system": false,
    "upvote": false,
    "downvote": false
}
```

### 修改项目的缺陷评论

修改项目内某个指定缺陷的指定评论


```
PUT /api/v3/projects/:id/issues/:issue_id/notes/:note_id
```


**参数**



| 参数         | 类型                  | 描述                                        |
| ---------- | ------------------- | ----------------------------------------- |
| id         | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| issue_id   | integer             | 缺陷 id                                     |
| note_id    | integer             | 评论 id                                     |
| body       | string              | 评论 内容                                     |


**返回值**


```json
{
    "id": 25642,
    "body": "已经修复待验证",
    "attachment": null,
    "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "created_at": "2015-03-20T07:13:22+0000",
    "system": true,
    "upvote": false,
    "downvote": false
}
```

### 查询单个缺陷的评论

在项目内查询某个指定的缺陷的指定评论


```
GET /api/v3/projects/:id/issues/:issue_id/notes/:note_id
```


**参数**



| 参数         | 类型                  | 描述                                        |
| ---------- | ------------------- | ----------------------------------------- |
| id         | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| issue_id   | integer             | 缺陷 id                                     |
| note_id    | integer             | 评论 id                                     |


**返回值**


```json
{
    "id": 25644,
    "body": "note1",
    "attachment": null,
    "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "created_at": "2015-03-20T07:13:22+0000",
    "system": false,
    "upvote": false,
    "downvote": false
}
```

### 获取缺陷评论列表

查询项目的某个缺陷的评论列表


```
GET /api/v3/projects/:id/issues/:issue_id/notes
```


**参数**



| 参数         | 类型                  | 描述                                        |
| ---------- | ------------------- | ----------------------------------------- |
| id         | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| issue_id   | integer             | 缺陷 id                                     |
| page       | integer             | 分页（默认值：1）                                 |
| per_page   | integer             | 默认页面大小（默认值： 20，最大值： 100）                  |


**返回值**


```json
[
    {
        "id": 2,
        "body": "note1",
        "attachment": null,
        "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
        "created_at": "2015-03-20T07:13:22+0000",
        "system": false,
        "upvote": false,
        "downvote": false
    },
    {
        "id": 3,
        "body": "note2 @git_user2 ",
        "attachment": null,
        "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
        "created_at": "2015-03-22T02:33:46+0000",
        "system": false,
        "upvote": false,
        "downvote": false
    }
]
```
