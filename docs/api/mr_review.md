# MR 评审

MR评审

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子：/api/v3/projects/diaspora%2Fdiaspora (/ 编码 %2F)

### 邀请评审人

邀请某个用户去评审指定的合并请求

### reviewer_id和necessary_reviewer_id必须填写其中一个


```
POST /api/v3/projects/:id/merge_request/:merge_request_id/review/invite
```


**参数**



| 参数                      | 类型                  | 描述                                        |
| ----------------------- | ------------------- | ----------------------------------------- |
| id                      | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id        | integer             | 合并请求的 id                                  |
| reviewer_id             | integer（可选）         | 评审人的 id                                   |
| necessary_reviewer_id   | integer（可选）         | 必要评审人的 id                                 |


**返回值**


```json
{
    "author": {
        "id": 223,
        "username": "git_user01",
        "web_url": "http://git.code.tencent.com/u/git_user01",
        "name": "git_user01",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/223/3301e9f926s54482802104bb08cf7150.gif"
    },
    "reviewers": [
        {
            "id": 223,
            "username": "git_user01",
            "web_url": "http://git.code.tencent.com/u/git_user01",
            "name": "git_user01",
            "state": "active",
            "avatar_url": "git.code.tencent.com/assets/images/avatar/no_user_avatar.png",
            "type": "invite",
            "review_state": "approving",
            "created_at": "2017-07-30T07:59:08+0000",
            "updated_at": "2017-07-30T07:59:08+0000"
        },
        {
            "id": 45,
            "username": "bobi",
            "web_url": "http://git.code.tencent.com/u/bobi",
            "name": "bobi",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/45/a75ba2727c7a409cab1d15dd993149aa.jpg",
            "type": "invite",
            "review_state": "approving",
            "created_at": "2017-08-09T02:01:56+0000",
            "updated_at": "2017-08-09T02:01:56+0000"
        }
    ],
    "id": 351,
    "project_id": 14539,
    "reviewable_id": 1344,
    "reviewable_type": "merge_request",
    "commit_id": null,
    "state": "approving",
    "restrict_type": "single_approve",
    "push_reset_enabled": true,
    "created_at": "2017-06-20T01:48:30+0000",
    "updated_at": "2017-06-20T01:48:30+0000"
}
```

移除评审人

### 在指定的合并请求中移除某位指定的评审


```
DELETE /api/v3/projects/:id/merge_request/:merge_request_id/review/dismissals
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求 id                                   |
| reviewer_id        | integer             | 评审人 id                                    |


**返回值**


### 返回状态码

取消评审

### 在项目中移除某个指定合并请求的评审人


```
DELETE /api/v3/projects/:id/merge_request/:merge_request_id/review/cancel
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求 id                                   |


**返回值**


### 返回状态码

查询评审信息

### 在项目中查询某个指定的合并请求评审信息


```
GET /api/v3/projects/:id/merge_request/:merge_request_id/review
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求的 id                                  |


**返回值**


```json
{
    "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/3301e9f989b94482802104bb08cf7150.gif"
    },
    "reviewers": [
        {
            "id": 11323,
            "username": "git_user1",
            "web_url": "http://git.code.tencent.com/u/git_helper_01",
            "name": "git_user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/3301e9f989b94482802104bb08cf7150.gif",
            "type": "invite",
            "review_state": "approving",
            "created_at": "2017-07-30T08:24:51+0000",
            "updated_at": "2017-07-30T08:24:51+0000"
        }
    ],
    "id": 351,
    "project_id": 14539,
    "reviewable_id": 1344,
    "reviewable_type": "merge_request",
    "commit_id": null,
    "state": "approving",
    "restrict_type": "single_approve",
    "push_reset_enabled": true,
    "created_at": "2017-06-20T01:48:30+0000",
    "updated_at": "2017-06-20T01:48:30+0000"
}
```

### 发表评审意见

邀请评审人发表评审意见


```
PUT /api/v3/projects/:id/merge_request/:merge_request_id/reviewer/summary
```


**参数**



| 参数                 | 类型                  | 描述                                                     |
| ------------------ | ------------------- | ------------------------------------------------------ |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                |
| merge_request_id   | integer             | 合并请求的 id                                               |
| reviewer_event     | string              | 评审人事件，可选：（comment | approve | require_change | deny）   |
| summary            | string              | 评审信息摘要                                                 |


**返回值**


```json
{
    "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_helper_01",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/3301e9f989b94482802104bb08cf7150.gif"
    },
    "reviewers": [
        {
            "id": 11323,
            "username": "git_user1",
            "web_url": "http://git.code.tencent.com/u/git_user1",
            "name": "git_user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/assets/images/avatar/no_user_avatar.png",
            "type": "invite",
            "review_state": "approving",
            "created_at": "2017-07-30T07:59:08+0000",
            "updated_at": "2017-07-30T07:59:08+0000"
        },
    ],
    "id": 351,
    "project_id": 14539,
    "reviewable_id": 1344,
    "reviewable_type": "merge_request",
    "commit_id": null,
    "state": "approving",
    "restrict_type": "single_approve",
    "push_reset_enabled": true,
    "created_at": "2017-06-20T01:48:30+0000",
    "updated_at": "2017-06-20T01:48:30+0000"
}
```

重置评审状态

### 重置项目中某个指定的合并请求评审状态，这个合并请求必须处于拒绝或反馈修改的状态


```
PUT /api/v3/projects/:id/merge_request/:merge_request_id/review/reopen
```


**参数**



| 参数                 | 类型                  | 描述                                        |
| ------------------ | ------------------- | ----------------------------------------- |
| id                 | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| merge_request_id   | integer             | 合并请求 id                                   |


**返回值**


```json
{
    "author": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2763c7a409cab1d15dd993149aa.jpg"
    },
    "reviewers": [
        {
            "id": 11323,
            "username": "git_user1",
            "web_url": "http://git.code.tencent.com/u/git_user1",
            "name": "git_user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/18604/a75ba2763c7a409cab1d15dd993149aa.jpg",
            "type": "suggestion",
            "review_state": "approving",
            "created_at": "2017-06-08T07:37:52+0000",
            "updated_at": "2017-08-09T06:42:37+0000"
        },
    ],
    "id": 29856,
    "project_id": 14539,
    "reviewable_id": 1345,
    "reviewable_type": "merge_request",
    "commit_id": null,
    "state": "approving",
    "restrict_type": "single_approve",
    "push_reset_enabled": true,
    "created_at": "2017-06-08T07:37:52+0000",
    "updated_at": "2017-08-09T06:42:37+0000"
}
```
