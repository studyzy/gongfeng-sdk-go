# 合并请求（MR）


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子：/api/v3/projects/diaspora%2Fdiaspora (/ 编码 %2F)

### 新增合并请求

在项目中创建一个合并请求

审批人规则：-1：所有评审人通过，1：单评审通过，2+：多评审通过(用>=2的数字代表需要几位必要评审人通过)。

必要审批人规则：-1：所有必要评审人通过，1：单必要评审通过，2+：多必要评审通过(用>=2的数字代表需要几位必要评审人通过)，0: 不需要必要评审人即可通过。


```
POST /api/v3/projects/:id/merge_requests
```


**参数**



| 参数                        | 类型                  | 描述                                        |
| ------------------------- | ------------------- | ----------------------------------------- |
| id                        | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| source_branch             | string              | 源分支                                       |
| target_branch             | string              | 目标分支                                      |
| title                     | string              | 合并请求的标题                                   |
| assignee_id               | integer （可选）        | 分配人 id                                    |
| description               | string （可选）         | 合并请求的描述                                   |
| target_project_id         | integer （可选）        | 目标项目的 id                                  |
| labels                    | string （可选）         | 合并请求的标签，多个请用英文逗号分隔                        |
| reviewers                 | string （可选）         | 评审人id (只能是id。多个评审人用,隔开)                   |
| necessary_reviewers       | string （可选）         | 必要评审人id (只能是id。多个评审人用,隔开)                 |
| approver_rule             | integer（可选）         | 评审人规则                                     |
| necessary_approver_rule   | integer（可选）         | 必要评审人规则                                   |


**返回值**


```json
{
    "id": 25563,
    "title": "merge request title",
    "target_project_id": 12321,
    "target_branch": "master",
    "source_project_id": 12321,
    "source_branch": "branch1",
    "state": "opened",
    "iid": 5,
    "description": "desc",
    "created_at": "2015-03-21T06:26:38+0000",
    "updated_at": "2015-03-21T06:26:38+0000",
    "labels": [],
    "assignee": null,
    "author": {
        "id": 1055,
        "name": "jim",
        "username": "jim",
        "state": "active",
        "avatar_url": null
    },
    "milestone": {
        "id": 1103,
        "project_id": 12321,
        "title": "milestone1",
        "state": "active",
        "iid": 1,
        "due_date": null,
        "created_at": "2012-10-15T06:02:45+0000",
        "updated_at": "2012-10-15T06:02:45+0000",
        "description": "milestone1"
    },
    "project_id": 12321,
    "work_in_progress": false,
    "upvotes": 0,
    "downvotes": 0
}
```

### 合并合并请求

在项目内合并某个指定的合并请求


```
PUT /api/v3/projects/:id/merge_request/:merge_request_id/merge
```


**参数**



| 参数                     | 类型                  | 描述                                        |
| ---------------------- | ------------------- | ----------------------------------------- |
| id                     | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id       | integer             | 合并请求的 id                                  |
| merge_commit_message   | string （可选）         | 合并合并请求的描述消息                               |


**返回值**


```json
{
    "id": 3256,
    "title": "merge request title",
    "target_project_id": 12321,
    "target_branch": "master",
    "source_project_id": 12321,
    "source_branch": "branch1",
    "state": "merged",
    "iid": 26,
    "description": "desc",
    "created_at": "2013-02-03T08:04:53+0000",
    "updated_at": "2013-02-03T13:24:34+0000",
    "labels": [],
    "assignee": {
        "id": 11323,
        "name": "git_user1",
        "username": "git_user1",
        "state": "active",
        "avatar_url": null
    },
    "author": {
        "id": 15651,
        "name": "git_user3",
        "username": "git_user3",
        "state": "active",
        "avatar_url": null
    },
    "milestone": {
        "id": 1103,
        "project_id": 12321,
        "title": "milestone1",
        "state": "active",
        "iid": 1,
        "due_date": null,
        "created_at": "2013-10-15T08:33:31+0000",
        "updated_at": "2013-10-15T08:33:56+0000",
        "description": "milestone1"
    },
    "project_id": 12321,
    "work_in_progress": false,
    "upvotes": 0,
    "downvotes": 0
}
```

### 获取合并请求中的提交

获取项目中某个指定合并请求的提交列表


```
GET /api/v3/projects/:id/merge_requests/:merge_request_id/commits
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | string              | 合并请求的ID                                   |


**返回值**


```json
[
    {
        "id": "34a677f562eaaeb8a54afe08d65b74165604adda",
        "short_id": "34a677f5",
        "title": null,
        "author_name": "git_user1",
        "author_email": "git_user1@tencent.com",
        "created_at": "2015-10-17T07:24:07+0000",
        "message": "fix"
    },
    {
    "id": "1dfbb3a9d7fcb9a9709f21712f501e9d5835137a",
    "short_id": "1dfbb3a9",
    "title": "branch",
    "author_name": "git_user1",
    "author_email": "git_user1@tencent.com",
    "created_at": "2018-10-16T08:39:49+0000",
    "message": "branch"
  }

]
```

添加合并请求的评论

### 在指定的合并请求下添加评论


```
POST /api/v3/projects/:id/merge_request/:merge_request_id/comments
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求的 id                                  |
| note               | integer             | 评论内容                                      |


**返回值**


```json
{
    "id": 3564,
    "body": "note test",
    "attachment": null,
    "author": {
        "id": 11323,
        "username": "git—user1",
        "web_url": "http://git.code.tencent.com/u/git—user1",
        "name": "git—user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "created_at": "2013-03-21T07:38:27+0000",
    "system": false
}
```

### 获取合并请求评论列表

获取项目内某个指定合并请求的评论列表


```
GET /api/v3/projects/:id/merge_request/:merge_request_id/comments
```


**参数**



| 参数                 | 类型                            | 描述                                                                                                             |
| ------------------ | ----------------------------- | -------------------------------------------------------------------------------------------------------------- |
| id                 | integer or string             | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                                                                        |
| merge_request_id   | integer                       | 合并请求的 id                                                                                                       |
| page               | integer                       | 分页（默认值：1）                                                                                                      |
| per_page           | integer                       | 默认页面大小（默认值： 20，最大值： 100）                                                                                       |
| created_after      | integer（可选）                   | 返回给定时间及之后创建的问题；例如2019-03-25T00:10:19+0000 或 2019-03-25T00:10:19+0800，时间参数必须转码，如"2019-03-25T00:10:19%2B0800"）   |
| created_before     | yyyy-MM-dd'T'HH:mm:ssZ （可选）   | 返回给定时间及之前创建的问题；例如2019-03-25T00:10:19+0000 或 2019-03-25T00:10:19+0800，时间参数必须转码，如"2019-03-25T00:10:19%2B0800"）   |


**返回值**


```json
[
    {
        "id": 2356,
        "body": "note test",
        "attachment": null,
        "author": {
        "id": 11323,
        "username": "git—user1",
        "web_url": "http://git.code.tencent.com/u/git—user1",
        "name": "git—user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
        "created_at": "2013-01-31T06:12:09+0000",
        "system": false
    },
    {
        "id": 2345,
        "body": "note-test",
        "attachment": null,
        "author": {
        "id": 11323,
        "username": "git—user1",
        "web_url": "http://git.code.tencent.com/u/git—user1",
        "name": "git—user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
        "created_at": "2013-03-20T10:17:05+0000",
        "system": true
    }
]
```

### 更新合并请求

在项目内更新某个指定的合并请求


```
PUT /api/v3/projects/:id/merge_request/:merge_request_id
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求的 id                                  |
| target_branch      | string （可选）         | 目标分支                                      |
| assignee_id        | integer （可选）        | 分配人 id                                    |
| title              | string （可选）         | 合并请求的标题                                   |
| state_event        | string （可选）         | 新的状态，可选值：（close|reopen）                   |
| description        | string （可选）         | 合并请求的描述                                   |
| labels             | string （可选）         | 合并请求的标签，多个请用英文逗号分隔                        |


**返回值**


```json
{
    "id": 38546,
    "title": "merge request update",
    "target_project_id": 12321,
    "target_branch": "master",
    "source_project_id": 12321,
    "source_branch": "branch1",
    "state": "closed",
    "iid": 9,
    "description": "desc",
    "created_at": "2013-01-30T08:24:46+0000",
    "updated_at": "2013-03-21T08:34:44+0000",
    "labels": [],
    "assignee": null,
    "author": {
        "id": 11323,
        "name": "git_user1",
        "username": "git_user1",
        "state": "active",
        "avatar_url": null
    },
    "milestone": {
        "id": 1103,
        "project_id": 12321,
        "title": "milestone1",
        "state": "active",
        "iid": 1,
        "due_date": null,
        "created_at": "2012-11-13T08:06:31+0000",
        "updated_at": "2012-11-13T08:06:31+0000",
        "description": "milestone1"
    },
    "project_id": 12321,
    "work_in_progress": false,
    "upvotes": 0,
    "downvotes": 0
}
```

### 获取合并请求列表

查询项目的合并请求列表


```
GET /api/v3/projects/:id/merge_requests
```


```
GET /api/v3/projects/:id/merge_requests?state=opened
```


```
GET /api/v3/projects/:id/merge_requests?state=all
```


```
GET /api/v3/projects/:id/merge_requests?iid=42
```


**参数**



| 参数         | 类型                  | 描述                                                   |
| ---------- | ------------------- | ---------------------------------------------------- |
| id         | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH              |
| iid        | integer（可选）         | 项目里面的合并请求id编号                                        |
| state      | string （可选）         | 合并请求状态，可选值：merged, opened 或 closed，不填写返回所有的合并请求      |
| order_by   | string （可选）         | 排序字段， 允许按 created_at, updated_at 排序（默认 created_at）   |
| sort       | string （可选）         | 排序方式， 允许 asc or desc（默认 desc）                        |
| page       | integer             | 分页（默认值：1）                                            |
| per_page   | integer             | 默认页面大小（默认值： 20，最大值： 100）                             |


**返回值**


```json
[
    {
        "id": 2556,
        "title": "merge request title",
        "target_project_id": 12321,
        "target_branch": "master",
        "source_project_id": 12321,
        "source_branch": "branch1",
        "state": "opened",
        "iid": 9,
        "description": "desc",
        "created_at": "2013-01-19T07:02:50+0000",
        "updated_at": "2013-01-19T07:02:50+0000",
        "assignee": {
            "id": 333,
            "username": "git_user2",
            "web_url": "http://git.code.tencent.com/u/git—user2",
            "name": "git_user2",
            "state": "active",
            "avatar_url": null
        },
        "author": {
        "id": 11323,
        "username": "git—user1",
        "web_url": "http://git.code.tencent.com/u/git—user1",
        "name": "git—user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "milestone": null,
    "project_id": 4559,
    "work_in_progress": false,
    "upvotes": 0,
    "downvotes": 0
    },
    {
        "id": 2535,
        "title": "merge request title",
        "target_project_id": 7,
        "target_branch": "master",
        "source_project_id": 7,
        "source_branch": "branch1",
        "state": "opened",
        "iid": 10,
        "description": "desc",
        "created_at": "2013-01-30T08:24:46+0000",
        "updated_at": "2013-03-21T06:04:50+0000",
        "assignee": null,
            "author": {
            "id": 11323,
            "username": "git—user1",
            "web_url": "http://git.code.tencent.com/u/git—user1",
            "name": "git—user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "milestone": null,
    "project_id": 4559,
    "work_in_progress": false,
    "upvotes": 0,
    "downvotes": 0
    }
]
```

查询合并请求的代码变更

### 显示某个指定合并请求的详情、包含的文件及修改


```
GET /api/v3/projects/:id/merge_request/:merge_request_id/changes
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求的 id                                  |


**返回值**


```json
{
    "id": 26534,
    "title": "merge request title",
    "target_project_id": 12321,
    "target_branch": "master",
    "source_project_id": 12321,
    "source_branch": "branch1",
    "state": "opened",
    "iid": 11323,
    "description": "desc",
    "created_at": "2013-03-21T06:26:38+0000",
    "updated_at": "2013-03-21T06:26:38+0000",
    "labels": [],
    "assignee": null,
    "author": {
            "id": 11323,
            "username": "git—user1",
            "web_url": "http://git.code.tencent.com/u/git—user1",
            "name": "git—user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "milestone": {
        "id": 1103,
        "project_id": 4559,
        "title": "milestone1",
        "state": "active",
        "iid": 1,
        "due_date": null,
        "created_at": "2012-11-15T08:02:31+0000",
        "updated_at": "2012-11-15T08:02:31+0000",
        "description": "milestone1"
    },
    "files": [
        {
            "old_path": "/dev/null",
            "new_path": "test",
            "a_mode": 0,
            "b_mode": 33188,
            "diff": "@@ -0,0 +1 @@\
+hello, world\
\\\\ No newline at end of file\
",
            "new_file": true,
            "renamed_file": false,
            "deleted_file": false,
            "is_too_large": false,
            "is_collapse": false,
            "additions": 1,
            "deletions": 2
        }
    ],
    "project_id": 4559,
    "work_in_progress": false,
    "upvotes": 0,
    "downvotes": 0
}
```

### 查询项目合并请求

查询项目下某个指定合并请求


```
GET /api/v3/projects/:id/merge_request/:merge_request_id
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求的 id                                  |


**返回值**


```json
{
    "id": 235,
    "title": "merge request title",
    "target_project_id": 12321,
    "target_branch": "master",
    "source_project_id": 12321,
    "source_branch": "branch1",
    "state": "opened",
    "iid": 9,
    "description": "desc",
    "created_at": "2013-01-30T08:24:46+0000",
    "updated_at": "2013-03-21T08:34:44+0000",
    "labels": [],
    "assignee": null,
    "author": {
            "id": 11323,
            "username": "git—user1",
            "web_url": "http://git.code.tencent.com/u/git—user1",
            "name": "git—user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2738c7a409cab1d15dd993149aa.jpg"
    },
    "milestone": null,
    "project_id": 45595,
    "work_in_progress": false,
    "upvotes": 0,
    "downvotes": 0
}
```

### 查询用户是否订阅请求合并

在项目里查询用户是否订阅了某个指定合并请求


```
GET /api/v3/projects/:id/merge_request/:merge_request_id/subscribe
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 请求合并的 id                                  |


**返回值**


true or false

### 订阅请求合并

在项目内订阅某个指定的合并请求


```
PUT /api/v3/projects/:id/merge_request/:merge_request_id/subscribe
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求的 id                                  |


**返回值**


### 状态码
取消订阅合并请求

### 在项目内取消订阅某个指定的合并请求


```
PUT /api/v3/projects/:id/merge_request/:merge_request_id/unsubscribe
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求的 id                                  |


**返回值**


### 状态码
下载 MergeRequest 差异文件集

下载指定合并请求的差异文件集

注意：仅支持下载open状态的合并请求


```
GET /api/v3/projects/:id/merge_requests/:merge_request_id/changed_files
```


**参数：**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求的 id                                  |


返回值：文件流
