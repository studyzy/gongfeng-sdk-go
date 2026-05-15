# 项目分支管理

分支

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

### 创建项目分支

在项目里面创建分支


```
POST /api/v3/projects/:id/repository/branches
```


**参数:**



| 参数            | 类型                 | 描述                                        |
| ------------- | ------------------ | ----------------------------------------- |
| id            | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| branch_name   | string             | 分支名                                       |
| ref           | string             | 从 提交的SHA 或存在的分支 创建新分支                     |


**返回值:**


```json
{
  "name": "123456",
  "protected": false,
  "developers_can_push": false,
  "developers_can_merge": false,
  "commit": {
    "id": "6269835cafdd2356d8b8d4a4e3738c79206c1d06",
    "message": "test",
    "parent_ids": [
      "12345fa022906668338166ce7f2e7bf35d502285"
    ],
    "authored_date": "2012-02-01T06:18:54+0000",
    "author_name": "git_user2",
    "author_email": "git_user2@tencent.com",
    "committed_date": "2012-02-01T06:18:54+0000",
    "committer_name": "git_user2",
    "committer_email": "git_user2@tencent.com",
    "title": "test",
    "created_at": "2012-02-01T06:18:54+0000"
    "short_id": "6269835c"
  }
}
```

### 删除分支

删除项目版本库分支


```
DELETE /api/v3/projects/:id/repository/branches/:branch
```


**参数:**



| 参数       | 类型                 | 描述                                        |
| -------- | ------------------ | ----------------------------------------- |
| id       | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| branch   | string             | 分支 名                                      |


**返回值:**


```json
{
  "branch_name": "fenzhi01"
}
```

### 分支列表

返回项目版本库 分支 列表，按字母顺序排序


```
GET /api/v3/projects/:id/repository/branches
```


**参数:**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| page       | integer（可选）        | 分页（默认值：1）                                 |
| per_page   | integer（可选）        | 默认页面大小（默认：20，最大：100）                      |


**返回值:**


```json
[
  {
    "name": "fenzhi01",
    "protected": false,
    "developers_can_push": false,
    "developers_can_merge": false,
    "commit": {
      "id": "e99ab09234cfa855a129305e7960261cf145950a",
      "message": "提交测试",
      "parent_ids": [
        "e9c5aa72c9e22592ee2611d896b730c1a95cc339"
      ],
      "authored_date": "2011-03-03T00:01:05+0000",
      "author_name": "git_user1",
      "author_email": "git_user1@tencent.com",
      "committed_date": "2011-03-03T00:01:05+0000",
      "committer_name": "git_user1",
      "committer_email": "git_user1@tencent.com",
      "title": "nihaoshijie",
      "created_at": "2011-03-03T00:01:05+0000"
      "short_id": "e99ab092"
    }
  },
  {
    "name": "fenzhi02",
    "protected": false,
    "developers_can_push": false,
    "developers_can_merge": false,
    "commit": {
      "id": "9bd9d12cbab4260bdd9d38fe1140a6cdac7cd2ea",
      "message": "test",
      "parent_ids": [
        "1230a4610ca01dd10cb5bc30d8a1b1eab98c09a1"
      ],
      "authored_date": "2011-09-19T09:09:14+0000",
      "author_name": "git_user1",
      "author_email": "git_user1@tencent.com",
      "committed_date": "2011-09-19T09:09:14+0000",
      "committer_name": "git_user1",
      "committer_email": "git_user1@tencent.com",
      "title": "666",
      "created_at": "2011-09-19T09:09:14+0000"
      "short_id": "9bd9d12c"
    }
  }
]
```

### 获取分支详情

返回项目版本库某个分支详情


```
GET /api/v3/projects/:id/repository/branches/:branch
```


**参数:**



| 参数       | 类型                | 描述                                        |
| -------- | ----------------- | ----------------------------------------- |
| id       | integer 或string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| branch   | string            | 分支名                                       |


**返回值:**


```json
{
  "name": "master",
  "protected": true,
  "developers_can_push": false,
  "developers_can_merge": false,
  "commit": {
    "id": "d21f748b73507a227a92c954bd3c89bf4e78e897",
    "message": "Merge branch 'fenzhi01' into 'master'\\r\
\\r\
xxx\\r\
\\r\
xxx\\r\
\\r\
See merge request !7",
    "parent_ids": [
      "9bd9d12crgd4260bdd9d38fe1140a6cdac7cd2ea",
      "e99db05604cfa385a129305e7660261cf161950a"
    ],
    "authored_date": "2012-01-01T01:11:11+0000",
    "author_name": "git_user1",
    "author_email": "git_user1@tencent.com",
    "committed_date": "2012-01-01T01:11:11+0000",
    "committer_name": "git_user1",
    "committer_email": "git_user2",
    "title": "Merge branch 'fenzhi01' into 'master'",
    "created_at": "2012-01-01T01:11:11+0000"
    "short_id": "d21f748b"
  }
}
```

### 获取保护分支详情

返回项目版本库某个 保护分支 详情


```
GET /api/v3/projects/:id/repository/branches/:branch/protect
```


**参数：**



| 参数       | 类型                | 描述                                        |
| -------- | ----------------- | ----------------------------------------- |
| id       | integer 或string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| branch   | string            | 分支名                                       |


**返回值：**


```json
{
  "protected": true,
  "developers_can_push": false,
  "developers_can_merge": false,
  "suggestion_reviewers": [],
  "necessary_reviewers": [],
  "name": "master",
  "commit": {
    "id": "d21f748b73507a227a92c954bd3c89bf4e78e897",
    "message": "Merge branch 'fenzhi01' into 'master'\\r\
\\r\
xxx\\r\
\\r\
xxx\\r\
\\r\
See merge request !7",
    "parent_ids": [
      "9bd9d12crgd4260bdd9d38fe1140a6cdac7cd2ea",
      "e99db05604cfa385a129305e7660261cf161950a"
    ],
    "authored_date": "2012-01-01T01:11:11+0000",
    "author_name": "git_user1",
    "author_email": "git_user1@tencent.com",
    "committed_date": "2012-01-01T01:11:11+0000",
    "committer_name": "git_user1",
    "committer_email": "git_user2",
    "title": "Merge branch 'fenzhi01' into 'master'",
    "created_at": "2012-01-01T01:11:11+0000",
    "short_id": "d21f748b"
  }
    "push_reset_enabled": true,
    "can_approve_by_creator": true,
    "auto_create_review_after_push": true,
    "approver_rule": 1,
    "necessary_approver_rule": 0,
    "merge_request_template": null,
    "path_reviewer_rules": ""
}
```
