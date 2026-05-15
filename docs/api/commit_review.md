# Commit 评审

日常代码评审（Commit Review）

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子：/api/v3/projects/diaspora%2Fdiaspora (/ 编码 %2F)

### 新建Commit 评审

在项目中新建一个Commit评审

审批人规则：-1：所有评审人通过，1：单评审通过，2+：多评审通过(用>=2的数字代表需要几位必要评审人通过)。

必要审批人规则：-1：所有必要评审人通过，1：单必要评审通过，2+：多必要评审通过(用>=2的数字代表需要几位必要评审人通过)，0: 不需要必要评审人即可通过。


```
POST /api/v3/projects/:id/review
```


**参数**



| 参数                        | 类型                  | 描述                                        |
| ------------------------- | ------------------- | ----------------------------------------- |
| id                        | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| title                     | string              | 标题                                        |
| source_branch             | string              | 源分支名（默认选择该分支最新的提交点）                       |
| target_branch             | string              | 目标分支名（默认选择该分支最新的提交点）                      |
| description               | string（可选）          | 描述                                        |
| source_commit             | string              | 源提交点                                      |
| target_commit             | string              | 目标提交点                                     |
| target_project_id         | integer（可选）         | 目标项目 id                                   |
| reviewer_ids              | string（可选）          | 评审人 id (只能是id。多个评审人用,隔开)                  |
| necessary_reviewer_ids    | string（可选）          | 必要评审人id (只能是id。多个评审人用,隔开)                 |
| approver_rule             | integer（可选）         | 评审人规则                                     |
| necessary_approver_rule   | integer（可选）         | 必要评审人规则                                   |


**返回值：**


```json
{
    "title": "api新建",
    "description": null,
    "author": {
        "id": 0001,
        "username": "git_user01",
        "web_url": "http://git.code.tencent.com/u/git_user01",
        "name": "git_user01",
        "state": "active",
        "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png"
    },
    "reviewers": [],
    "id": 276498,
    "project_id": 9837556,
    "reviewable_id": 105,
    "reviewable_type": "comparison",
    "state": "approving",
    "approver_rule": 1 ,
    "necessary_approver_rule": 0,
    "push_reset_enabled": true,
    "created_at": "2019-04-19T07:23:18+0000",
    "updated_at": "2019-04-19T07:23:18+0000"
}    
```
### 获取项目中的Commit评审

获取项目中所有的Commit评审


```
GET /api/v3/projects/:id/reviews
```


**参数**



| 参数          | 类型                  | 描述                                                                     |
| ----------- | ------------------- | ---------------------------------------------------------------------- |
| id          | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                                |
| author_id   | integer（可选）         | 作者的 id                                                                 |
| state       | string（可选）          | 代码状态，可选值：approving, approving change_required 或 closed，不填写返回所有的代码评审。   |
| order_by    | string（可选）          | 排序字段， 允许按 created_at, updated_at 排序（默认 created_at）                     |
| sort        | string（可选）          | 排序方式， 允许 asc or desc（默认 desc）                                          |
| page        | integer（可选）         | 分页（默认值：1）                                                              |
| per_page    | integer（可选）         | 默认页面大小（默认值： 20，最大值： 100）                                               |


**返回值：**


```json
{
        "title": "创建差异",
        "description": "创建差异",
        "author": {
            "id": 0001,
            "username": "git_user01",
            "web_url": "http://git.code.tencent.com/u/git_user01",
            "name": "git_user01",
            "state": "active",
            "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png"
        },
        "reviewers": null,
        "id": 27695,
        "project_id": 98376,
        "reviewable_id": 102,
        "reviewable_type": "comparison",
        "state": "approving",
        "approver_rule": 1,
        "necessary_approver_rule": 0,
        "push_reset_enabled": true,
        "created_at": "2019-04-18T09:23:51+0000",
        "updated_at": "2019-04-18T09:23:51+0000"
}
```
### 获取项目中某个具体的Commit评审

获取项目中某个具体的Commit评审情况


```
GET /api/v3/projects/:id/review/:review_id
```


**参数**



| 参数          | 类型                  | 描述                                        |
| ----------- | ------------------- | ----------------------------------------- |
| id          | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| review_id   | integer             | 代码评审的 id                                  |


**返回值：**


```json
{
    "title": "创建差异",
    "description": "创建差异",
    "commits": [
        {
            "id": "8a7f0548f9c18ce7e28ed8fb27ad700d44768362",
            "commit_date": "2019-04-24T08:24:43+0000"
        },
        {
            "id": "479f9d544f57fc571af9a88fb7e8dea1e56d899a",
            "commit_date": "2019-04-18T07:34:52+0000"
        }
    ],
    "author": {
        "id": 0001,
        "username": "git_user01",
        "web_url": "http://git.example.qq.com/u/git_user01",
        "name": "git_user01",
        "state": "active",
        "avatar_url": "http://git.example.qq.com/assets/images/avatar/no_user_avatar.png"
    },
    "reviewers": [],
    "id": 276495,
    "project_id": 9837556,
    "reviewable_id": 102,
    "reviewable_type": "comparison",
    "state": "approving",
    "approver_rule": 1,
    "necessary_approver_rule": 0,
    "push_reset_enabled": true,
    "created_at": "2019-04-18T09:23:51+0000",
    "updated_at": "2019-04-18T09:23:51+0000"
}
```
### 邀请评审人

给项目中的Commit评审添加评审人

### reviewer_id和necessary_reviewer_id两个必须有一个有值


```
POST /api/v3/projects/:id/review/:review_id/invite
```


**参数**



| 参数                      | 类型                  | 描述                                        |
| ----------------------- | ------------------- | ----------------------------------------- |
| id                      | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| review_id               | integer             | 代码评审的 id                                  |
| reviewer_id             | integer（可选）         | 评审人的 id (只能是id。多个评审人用,隔开)                 |
| necessary_reviewer_id   | integer（可选）         | 必要评审人的 id (只能是id。多个评审人用,隔开)               |


**返回值：**


```json
{
    "title": "api增加reviewer_id",
    "description": null,
    "author": {
        "id": 0001,
        "username": "git_user01",
        "web_url": "http://git.code.tencent.com/u/git_user01",
        "name": "git_user01",
        "state": "active",
        "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png"
    },
    "reviewers": [
        {
            "id": 0002,
            "username": "git_user02",
            "web_url": "http://git.code.tencent.com/u/git_user02",
            "name": "git_user02",
            "state": "active",
            "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png",
            "type": "invite",
            "review_state": "approving",
            "created_at": "2019-04-19T08:30:20+0000",
            "updated_at": "2019-04-19T08:30:20+0000"
        }
    ],
    "id": 276502,
    "project_id": 9837556,
    "reviewable_id": 108,
    "reviewable_type": "comparison",
    "state": "approving",
    "approver_rule": 1,
    "necessary_approver_rule": 0,
    "push_reset_enabled": true,
    "created_at": "2019-04-19T08:01:54+0000",
    "updated_at": "2019-04-19T08:30:15+0000"
}
```
### 移除评审人

在项目中的Commit评审中移除某位评审人


```
DELETE /api/v3/projects/:id/review/:review_id/dismissals
```


**参数**



| 参数            | 类型                  | 描述                                        |
| ------------- | ------------------- | ----------------------------------------- |
| id            | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| review_id     | integer             | 代码评审的 id                                  |
| reviewer_id   | integer             | 评审人的 id                                   |


**返回值：**


200等相关状态码
### 发表评审意见

在项目中的Commit评审发表评审意见


```
PUT /api/v3/projects/:id/review/:review_id/reviewer/summary
```


**参数**



| 参数               | 类型                  | 描述                                                     |
| ---------------- | ------------------- | ------------------------------------------------------ |
| id               | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                |
| review_id        | integer             | 代码评审的 id                                               |
| reviewer_event   | string              | 评审人事件，可选：（comment | approve | require_change | deny）   |
| summary          | string              | 评审信息摘要                                                 |


**返回值：**


```json
{
    "title": "创建一个评审",
    "description": "创建一个评审",
    "author": {
        "id": 0001,
        "username": "git_user01",
        "web_url": "http://git.code.tencent.com/u/git_user01",
        "name": "git_user01",
        "state": "active",
        "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png"
    },
    "reviewers": [
        {
            "id": 0002,
            "username": "git_user02",
            "web_url": "http://git.test.code.oa.com/u/git_user02",
            "name": "git_user02",
            "state": "active",
            "avatar_url":
            "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png",
            "type": "suggestion",
            "review_state": "approving",
            "created_at": "2019-04-19T09:03:50+0000",
            "updated_at": "2019-04-19T09:06:28+0000"

    ],
    "id": 276504,
    "project_id": 9837556,
    "reviewable_id": 110,
    "reviewable_type": "comparison",
    "state": "change_denied",
    "approver_rule": 1,
    "necessary_approver_rule": 0,
    "push_reset_enabled": true,
    "created_at": "2019-04-19T09:03:50+0000",
    "updated_at": "2019-04-19T09:06:47+0000"
}
重置评审状态

在项目中重置某个指定的Commit评审状态

PUT /api/v3/projects/:id/review/:review_id/reopen

参数

参数	类型	描述
id	integer or string	id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH
review_id	integer	代码评审的 id

返回值：

{
    "title": "创建一个评审",
    "description": "创建一个评审意见",
    "author": {
        "id": 0001,
        "username": "git_user01",
        "web_url": "http://git.code.tencent.com/u/git_user01",
        "name": "git_user01",
        "state": "active",
        "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png"
    },
    "reviewers": [
        {
            "id": 0002,
            "username": "git_user02",
            "web_url": "http://git.code.tencent.com/u/git_user02",
            "name": "git_user01",
            "state": "active",
            "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png",
            "type": "suggestion",
            "review_state": "approving",
            "created_at": "2019-04-19T09:03:50+0000",
            "updated_at": "2019-04-19T09:11:33+0000"

    ],
    "id": 276504,
    "project_id": 9837556,
    "reviewable_id": 110,
    "reviewable_type": "comparison",
    "state": "approving",
    "approver_rule": 1,
    "necessary_approver_rule": 0,
    "push_reset_enabled": true,
    "created_at": "2019-04-19T09:03:50+0000",
    "updated_at": "2019-04-19T09:11:33+0000"
}
更新Commit评审

在项目内更新某个指定的Commit评审

PUT /api/v3/projects/:id/review/:review_id

参数

参数	类型	描述
id	integer or string	id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH
request_id	integer	代码评审的 id
title	string	代码评审的标题
description	string（可选）	代码评审的描述

返回值：

{
    "title": "修改标题",
    "description": "创建一个评审意见",
    "author": {
        "id": 0001,
        "username": "git_user01",
        "web_url": "http://git.code.tencent.com/u/git_user01",
        "name": "git_user01",
        "state": "active",
        "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png"
    },
    "reviewers": [
        {
            "id": 0002,
            "username": "git_user02",
            "web_url": "http://git.code.tencent.com/u/git_user02",
            "name": "git_user01",
            "state": "active",
            "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_user_avatar.png",
            "type": "suggestion",
            "review_state": "approving",
            "created_at": "2019-04-19T09:03:50+0000",
            "updated_at": "2019-04-19T09:11:33+0000"

    ],
    "id": 276504,
    "project_id": 9837556,
    "reviewable_id": 110,
    "reviewable_type": "comparison",
    "state": "approving",
    "approver_rule": 1,
    "necessary_approver_rule": 0,
    "push_reset_enabled": true,
    "created_at": "2019-04-19T09:03:50+0000",
    "updated_at": "2019-04-19T09:11:33+0000"
}
下载 Commit Review差异文件集

下载指定代码评审的差异文件集

GET /api/v3/projects/:id/review/:review_id/changed_files

参数：

参数	类型	描述
id	integer or string	id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH
review_id	integer	代码评审的id

返回值：文件流
```
