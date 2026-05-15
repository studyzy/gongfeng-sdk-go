# 项目相关

项目成员

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

项目权限

组/项目的权限access_level包括：

GUEST = 10
FOLLOWER = 15
REPORTER = 20
DEVELOPER = 30
MASTER = 40
### OWNER = 50

获取项目成员列表

### 获取项目所有的成员列表


```
GET /api/v3/projects/:id/members
```


**参数：**



| 参数         | 类型                  | 描述                                        |
| ---------- | ------------------- | ----------------------------------------- |
| page       | integer             | 页数 (默认值:1)                                |
| per_page   | integer             | 每页列出成员数 (默认值 20)                          |
| id         | integer 或者 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| query      | string (可选)         | 搜索成员的字符串                                  |


**返回值：**


```json
[ {
        "id": 11323,
        "username": "git-user1",
        "web_url": "http://git.code.tencent.com/u/git-user1",
        "name": "git-user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/111323/a75ba2727c7a409cab1d15dd993149aa.jpg",
        "access_level": 30
    } ]
```

### 增加项目成员

增加一个项目成员，这个方法是幂等的，即多次增加同一个成员结果是一致的


```
POST /api/v3/projects/:id/members
```


**参数：**



| 参数             | 类型        | 描述                                        |
| -------------- | --------- | ----------------------------------------- |
| id             | integer   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| user_id        | integer   | 增加的用户的 ID                                 |
| access_level   | integer   | 项目访问级别                                    |


```
POST /api/v3/projects/:id/members
```


**返回值：**


```json
[ {
    "id": 5520,
    "username": "git-zhou",
    "web_url": "http://git.code.tencent.com/u/git-zhou",
    "name": "git-zhou",
    "state": "active",
    "avatar_url": "git.code.tencent.com/uploads/user/avatar/5520/a75ba2727c7a409cab1d15dd993149aa.jpg",
    "access_level": 30
} ]
```

### 修改项目成员

更新一个指定项目成员的访问级别


```
PUT /api/v3/projects/:id/members/:user_id
```


**参数：**



| 参数             | 类型        | 描述                                        |
| -------------- | --------- | ----------------------------------------- |
| id             | integer   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| user_id        | integer   | 用户的 ID                                    |
| access_level   | integer   | 项目访问级别                                    |


**返回值：**


```json
[ {
    "id": 5520,
    "username": "git-zhou",
    "web_url": "http://git.code.tencent.com/u/git-zhou",
    "name": "git-zhou",
    "state": "active",
    "avatar_url": "git.code.tencent.com/uploads/user/avatar/5520/a75ba2727c7a409cab1d15dd993149aa.jpg",
    "access_level": 40
} ]
```

### 删除项目成员

在项目内删除一个指定的成员


```
DELETE /api/v3/projects/:id/members/:user_id
```


**参数：**



| 参数        | 类型        | 描述                                        |
| --------- | --------- | ----------------------------------------- |
| id        | integer   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| user_id   | integer   | 用户的 ID                                    |


**返回值：**


### 200或相关状态码
获取项目内的某个指定成员信息

### 在项目内获取某个指定的成员


```
GET /api/v3/projects/:id/members/:user_id
```


**参数：**



| 参数        | 类型        | 描述                                        |
| --------- | --------- | ----------------------------------------- |
| id        | integer   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| user_id   | integer   | 成员的 ID                                    |


**返回值：**


```json
[ {
    "id": 5520,
    "username": "git-zhou",
    "web_url": "http://git.code.tencent.com/u/git-zhou",
    "name": "git-zhou",
    "state": "active",
    "avatar_url": "git.code.tencent.com/uploads/user/avatar/5520/a75ba2727c7a409cab1d15dd993149aa.jpg",
    "access_level": 40
} ]
```

项目

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子：/api/v3/projects/diaspora%2Fdiaspora(/ 编码 %2F)

项目可见度级别

工蜂Git系统的项目级别有私有、内部、公共三种。你能够通过定义项目visibility_level属性来决定项目级别。

内部版本
私有项目，visibility_level = 0，项目的访问必须显式授予每个用户
公共项目，visibility_level = 10，任何登录用户都可以克隆该项目

非内部版本
私有项目，visibility_level = 0，项目的访问必须显式授予每个用户
内部项目，visibility_level = 10，任何登录用户都可以克隆该项目
### 公共项目，visibility_level = 20，可以不经任何身份验证克隆该项目。

创建项目

### 创建一个自己的项目


```
POST /api/v3/projects
```


**参数：**



| 参数                 | 类型                      | 描述                              |
| ------------------ | ----------------------- | ------------------------------- |
| name               | string                  | 项目名                             |
| path               | string（可选）              | 项目版本库路径，默认：path = name          |
| fork_enabled       | boolean（可选）             | 项目是否可以被fork，默认：false            |
| namespace_id       | integer 或者 string（可选）   | 项目所属命名空间 ，默认用户的命名空间             |
| description        | string（可选）              | 项目描述                            |
| visibility_level   | integer（可选）             | 项目可视范围，默认visibility_level = 0   |


注意：namespace_id != null 时需要用户拥有在指定命名空间中创建项目的权限

**返回值：**


```json
{
    "id": 32833,
    "description": null,
    "public": false,
    "archived": false,
    "visibility_level": 0,
    "namespace": {
        "created_at": "2017-01-29T07:49:48+0000",
        "description": "git_user1",
        "id": 15302,
        "name": "git_user1",
        "owner_id": 13525,
        "path": "git_user1",
        "updated_at": "2017-01-29T07:49:48+0000"
    },
    "owner": {
        "id": 13525,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/13525/fec6c477729445df9a6e0a4a05b4a86c.png"
    },
    "name": "testapi",
    "name_with_namespace": "git_user1/testapi",
    "path": "testapi",
    "path_with_namespace": "git_user1/testapi",
    "default_branch": "master",
    "ssh_url_to_repo": "git@git.code.tencent.com:git_user1/testapi.git",
    "http_url_to_repo": "http://git.code.tencent.com/git/git_user1/testapi.git",
    "https_url_to_repo": "https://git.code.tencent.com/git/git_user1/testapi.git",
    "web_url": "http://git.code.tencent.com/git_user1/testapi",
    "tag_list": [],
    "issues_enabled": true,
    "merge_requests_enabled": true,
    "wiki_enabled": true,
    "snippets_enabled": true,
    "review_enabled": true,
    "fork_enabled": false,
    "tag_name_regex": null,
    "tag_create_push_level": 30,
    "created_at": "2017-08-13T07:25:53+0000",
    "last_activity_at": "201-08-13T07:25:53+0000",
    "creator_id": 13525,
    "avatar_url": "http://git.code.tencent.com/uploads/project/avatar/32833",
    "watchs_count": 0,
    "stars_count": 0,
    "forks_count": 0,
    "config_storage": {
        "limit_lfs_file_size": 500,
        "limit_size": 100000,
        "limit_file_size": 100000,
        "limit_lfs_size": 100000
    },
    "forked_from_project": "Forked Project not found",
    "statistics": {
        "commit_count": 0,
        "repository_size": 0
    }
}
```

### 编辑项目

修改项目设置


```
PUT /api/v3/projects/:id
```


修改成功，返回修改后的项目信息；参数错误，返回400

**参数：**



| 参数                       | 类型            | 描述                   |
| ------------------------ | ------------- | -------------------- |
| name                     | string（可选）    | 项目名                  |
| description              | string（可选）    | 项目描述                 |
| default_branch           | string（可选）    | 项目默认分支               |
| limit_file_size          | float（可选）     | 文件大小限制，单位:MB         |
| limit_lfs_file_size      | float（可选）     | LFS 文件大小限制，单位:MB     |
| issues_enabled           | boolean（可选）   | 缺陷配置                 |
| merge_requests_enabled   | boolean（可选）   | 合并请求配置               |
| wiki_enabled             | boolean（可选）   | 维基配置                 |
| review_enabled           | boolean（可选）   | 评审配置                 |
| fork_enabled             | boolean（可选）   | 是否可以被fork，默认:false   |
| tag_name_regex           | string（可选）    | 推送或创建 tag 规则         |
| tag_create_push_level    | integer（可选）   | 推送或创建 tag 权限         |
| visibility_level         | integer（可选）   | 项目可视范围               |


注意：
tag_create_push_level = 0 任何不能推送或创建 tag
tag_create_push_level = 30 DEVELOPER 以上角色才能推送或创建 tag>
### tag_create_push_level = 40 MASTER 以上角色才能推送或创建 tag

**返回值：**


```json
{
    "id": 32833,
    "description": null,
    "public": false,
    "archived": false,
    "visibility_level": 0,
    "namespace": {
        "created_at": "2017-01-29T07:49:48+0000",
        "description": "git_user1",
        "id": 15302,
        "name": "git_user1",
        "owner_id": 13525,
        "path": "git_user1",
        "updated_at": "2017-01-29T07:49:48+0000"
    },
    "owner": {
        "id": 13525,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/13525/fec6c477729445df9a6e0a4a05b4a86c.png"
    },
    "name": "testapp",
    "name_with_namespace": "git_user1/testapp",
    "path": "testapi",
    "path_with_namespace": "git_user1/testapi",
    "default_branch": "master",
    "ssh_url_to_repo": "git@git.code.tencent.com:git_user1/testapi.git",
    "http_url_to_repo": "http://git.code.tencent.com/git/git_user1/testapi.git",
    "https_url_to_repo": "https://git.code.tencent.com/git/git_user1/testapi.git",
    "web_url": "http://git.code.tencent.com/git_user1/testapi",
    "tag_list": [],
    "issues_enabled": true,
    "merge_requests_enabled": true,
    "wiki_enabled": true,
    "snippets_enabled": true,
    "review_enabled": true,
    "fork_enabled": false,
    "tag_name_regex": null,
    "tag_create_push_level": 30,
    "created_at": "2017-08-13T07:25:53+0000",
    "last_activity_at": "201-08-13T07:25:53+0000",
    "creator_id": 13525,
    "avatar_url": "http://git.code.tencent.com/uploads/project/avatar/32833",
    "watchs_count": 0,
    "stars_count": 0,
    "forks_count": 0,
    "config_storage": {
        "limit_lfs_file_size": 500,
        "limit_size": 100000,
        "limit_file_size": 100000,
        "limit_lfs_size": 100000
    },
    "forked_from_project": "Forked Project not found",
    "statistics": {
        "commit_count": 0,
        "repository_size": 0
    }
}
```

### 获取用户授权的项目列表

返回用户的授权项目列表，支持搜索、分页


```
GET /api/v3/projects
```


**参数：**



| 参数         | 类型            | 描述                                                                             |
| ---------- | ------------- | ------------------------------------------------------------------------------ |
| search     | string （可选）   | 搜索条件，模糊匹配path,name                                                             |
| archived   | boolean（可选）   | 归档状态，archived = true限制为查询归档项目，默认不区分归档状态                                        |
| order_by   | string （可选）   | 排序字段，允许按 id,name,path,created_at,updated_at,last_activity_at排序（默认created_at）   |
| sort       | string （可选）   | 排序方式，允许asc,desc（默认 desc）                                                       |
| page       | integer（可选）   | 页数（默认值：1）                                                                      |
| per_page   | integer（可选）   | 默认页面大小（默认值：20，最大值：100）                                                         |


**返回值：**


```json
[
  {
    "id": 8922,
    "description": null,
    "public": false,
    "archived": false,
    "visibility_level": 0,
    "namespace": {
        "created_at": "2017-01-16T08:29:43+0000",
        "description": "git_user1",
        "id": 11323,
        "name": "git_user1",
        "owner_id": 11323,
        "path": "git_user1",
        "updated_at": "2017-01-16T08:29:43+0000"
    },
    "owner": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2727c7a409cab1d15dd993149aa.jpg"
    },
    "name": "test-01",
    "name_with_namespace": "git_user1/test-01",
    "path": "test-01",
    "path_with_namespace": "git_user1/test-01",
    "default_branch": "master",
    "ssh_url_to_repo": "git@git.code.tencent.com:git_user1/test-01.git",
    "http_url_to_repo": "http://git.code.tencent.com/git_user1/test-01.git",
    "https_url_to_repo": "https://git.code.tencent.com/git_user1/test-01.git",
    "web_url": "http://git.code.tencent.com/git_user1/test-01",
    "tag_list": [],
    "issues_enabled": true,
    "merge_requests_enabled": true,
    "wiki_enabled": true,
    "snippets_enabled": true,
    "review_enabled": true,
    "fork_enabled": false,
    "tag_name_regex": null,
    "tag_create_push_level": 30,
    "created_at": "2017-08-13T07:37:14+0000",
    "last_activity_at": "2017-08-13T07:37:14+0000",
    "creator_id": 11323,
    "avatar_url": "http://git.code.tencent.com/uploads/project/avatar/70703",
    "watchs_count": 0,
    "stars_count": 0,
    "forks_count": 0,
    "config_storage": {
        "limit_lfs_file_size": 500,
        "limit_size": 100000,
        "limit_file_size": 100000,
        "limit_lfs_size": 100000
    },
    "forked_from_project": "Forked Project not found",
    "statistics": {
        "commit_count": 0,
        "repository_size": 0
    },
    "permissions": {
        "project_access": {
          "access_level": null
        },
         "share_group_access": {
          "access_level": null
        },
        "group_access": {
          "access_level": 50
        }
    }
}
]
```

### 获取用户拥有的项目列表

返回用户拥有的项目列表，支持搜索、分页


```
GET /api/v3/projects/owned
```


**参数：**



| 参数         | 类型            | 描述                                                                             |
| ---------- | ------------- | ------------------------------------------------------------------------------ |
| search     | string（可选）    | 搜索条件，模糊匹配path, name                                                            |
| archived   | boolean（可选）   | 归档状态，archived = true限制为查询归档项目，默认不区分归档状态                                        |
| order_by   | string（可选）    | 排序字段，允许按 id,name,path,created_at,updated_at,last_activity_at排序（默认created_at）   |
| sort       | string（可选）    | 排序方式，允许 asc or desc（默认 desc）                                                   |
| page       | integer（可选）   | 页数（默认值：1）                                                                      |
| per_page   | integer（可选）   | 默认页面大小（默认值：20，最大值：100）                                                         |


**返回值：**


```json
[
  {
    "id": 8922,
    "description": null,
    "public": false,
    "archived": false,
    "visibility_level": 0,
    "namespace": {
        "created_at": "2017-01-16T08:29:43+0000",
        "description": "git_user1",
        "id": 11323,
        "name": "git_user1",
        "owner_id": 11323,
        "path": "git_user1",
        "updated_at": "2017-01-16T08:29:43+0000"
    },
    "owner": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2727c7a409cab1d15dd993149aa.jpg"
    },
    "name": "test-01",
    "name_with_namespace": "git_user1/test-01",
    "path": "test-01",
    "path_with_namespace": "git_user1/test-01",
    "default_branch": "master",
    "ssh_url_to_repo": "git@git.code.tencent.com:git_user1/test-01.git",
    "http_url_to_repo": "http://git.code.tencent.com/git_user1/test-01.git",
    "https_url_to_repo": "https://git.code.tencent.com/git_user1/test-01.git",
    "web_url": "http://git.code.tencent.com/git_user1/test-01",
    "tag_list": [],
    "issues_enabled": true,
    "merge_requests_enabled": true,
    "wiki_enabled": true,
    "snippets_enabled": true,
    "review_enabled": true,
    "fork_enabled": false,
    "tag_name_regex": null,
    "tag_create_push_level": 30,
    "created_at": "2017-08-13T07:37:14+0000",
    "last_activity_at": "2017-08-13T07:37:14+0000",
    "creator_id": 11323,
    "avatar_url": "http://git.code.tencent.com/uploads/project/avatar/70703",
    "watchs_count": 0,
    "stars_count": 0,
    "forks_count": 0,
    "config_storage": {
        "limit_lfs_file_size": 500,
        "limit_size": 100000,
        "limit_file_size": 100000,
        "limit_lfs_size": 100000
    },
    "forked_from_project": "Forked Project not found",
    "statistics": {
        "commit_count": 0,
        "repository_size": 0
    }
}
]
```

### 删除项目

删除项目，会自动删除关联的资源（如缺陷、合并请求、标签等等）


```
DELETE /api/v3/projects/:id
```


**参数：**



| 参数   | 类型                 | 描述                                        |
| ---- | ------------------ | ----------------------------------------- |
| id   | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |


**返回值：**


### 200 或相关状态码
获取项目详细信息

根据 ID 或 NAMESPACE_PATH/PROJECT_PATH返回项目的详细信息，至少需要有该项目的guest权限


```
GET /api/v3/projects/:id
```


**参数：**



| 参数   | 类型                 | 描述                                        |
| ---- | ------------------ | ----------------------------------------- |
| id   | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |


**返回值：**


```json
{
    "id": 8922,
    "description": null,
    "public": false,
    "archived": false,
    "visibility_level": 0,
    "namespace": {
        "created_at": "2017-01-16T08:29:43+0000",
        "description": "git_user1",
        "id": 11323,
        "name": "git_user1",
        "owner_id": 11323,
        "path": "git_user1",
        "updated_at": "2017-01-16T08:29:43+0000"
    },
    "owner": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2727c7a409cab1d15dd993149aa.jpg"
    },
    "name": "test-01",
    "name_with_namespace": "git_user1/test-01",
    "path": "test-01",
    "path_with_namespace": "git_user1/test-01",
    "default_branch": "master",
    "ssh_url_to_repo": "git@git.code.tencent.com:git_user1/test-01.git",
    "http_url_to_repo": "http://git.code.tencent.com/git_user1/test-01.git",
    "https_url_to_repo": "https://git.code.tencent.com/git_user1/test-01.git",
    "web_url": "http://git.code.tencent.com/git_user1/test-01",
    "tag_list": [],
    "issues_enabled": true,
    "merge_requests_enabled": true,
    "wiki_enabled": true,
    "snippets_enabled": true,
    "review_enabled": true,
    "fork_enabled": false,
    "tag_name_regex": null,
    "tag_create_push_level": 30,
    "created_at": "2017-08-13T07:37:14+0000",
    "last_activity_at": "2017-08-13T07:37:14+0000",
    "creator_id": 11323,
    "avatar_url": "http://git.code.tencent.com/uploads/project/avatar/70703",
    "watchs_count": 0,
    "stars_count": 0,
    "forks_count": 0,
    "config_storage": {
        "limit_lfs_file_size": 500,
        "limit_size": 100000,
        "limit_file_size": 100000,
        "limit_lfs_size": 100000
    },
    "forked_from_project": "Forked Project not found",
    "statistics": {
        "commit_count": 0,
        "repository_size": 0
    }
    "suggestion_reviewers": [],
    "necessary_reviewers": [],
    "path_reviewer_rules": "",
    "approver_rule": 1,
    "necessary_approver_rule": 0,
    "can_approve_by_creator": true,
    "auto_create_review_after_push": true,
    "push_reset_enabled": false,
    "merge_request_template": null,
}
```

与组共享项目

### 允许与组共享项目


```
POST  /api/v3/projects/:id/share
```


**参数：**



| 参数             | 类型                 | 描述                                        |
| -------------- | ------------------ | ----------------------------------------- |
| id             | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| group_id       | integer            | 要与之共享的组的id                                |
| group_access   | integer            | 授予组的权限级别                                  |


**返回值：**


### 200或相关状态码
获取项目的share group列表

### 获取项目中所有的sharegroup


```
GET  /api/v3/projects/:id/shares
```


**参数：**



| 参数   | 类型                 | 描述                                        |
| ---- | ------------------ | ----------------------------------------- |
| id   | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |


**返回值：**


```json
[
   {
        "project_id": 29799,
        "group_id": 74393,
        "group_access": 20,
        "created_at": "2017-03-11T07:32:17+0000",
        "updated_at": "2017-03-11T07:32:17+0000",
    },
    {
      "project_id": 29799,
      "group_id": 220053,
      "group_access": 30,
      "created_at": "2018-01-04T03:52:19+0000",
      "updated_at": "2018-01-04T03:52:19+0000",
    }
]
```

删除组中共享项目链接

### 从组中取消共享项目。


```
DELETE  /api/v3/projects/:id/share/:group_id
```


**参数：**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| group_id   | integer            | 组的id                                      |


**返回值：**


### 200或相关状态码
通过名称搜索项目

### 过滤用户授权的项目列表，支持分页，如果是管理员，支持过滤所有项目


```
GET /api/v3/projects
```


**参数：**



| 参数                  | 类型            | 描述                                                                            |
| ------------------- | ------------- | ----------------------------------------------------------------------------- |
| search              | string        | 搜索条件，模糊匹配path, name.                                                          |
| with_archived       | boolean（可选）   | 归档状态，with_archived = true限制为查询归档项目，默认不区分                                      |
| with_push           | boolean（可选）   | 推送状态，with_push = true限制为查询推送过的项目，默认不区分                                        |
| abandoned           | boolean（可选）   | 活跃状态，abandoned = true限制为查询最近半年更新过的项目，默认全部                                     |
| visibility_levels   | string（可选）    | 项目可视范围，默认visibility_levels = "0, 10, 20"                                      |
| order_by            | string（可选）    | 排序字段，允许按id,name,path,created_at,updated_at,last_activity_at排序（默认created_at）   |
| sort                | string（可选）    | 排序方式，允许asc,desc（默认 desc）                                                      |
| page                | integer（可选）   | 页数（默认:1）                                                                      |
| per_page            | integer（可选）   | 默认页面大小（默认值： 20，最大值： 100）                                                      |


注意：当用户为非管理员时， with_archived、with_push、abandoned、visibility_levels过滤参数不生效

**返回值：**


```json
[
  {
    "id": 8922,
    "description": null,
    "public": false,
    "archived": false,
    "visibility_level": 0,
    "namespace": {
        "created_at": "2017-01-16T08:29:43+0000",
        "description": "git_user1",
        "id": 11323,
        "name": "git_user1",
        "owner_id": 11323,
        "path": "git_user1",
        "updated_at": "2017-01-16T08:29:43+0000"
    },
    "owner": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2727c7a409cab1d15dd993149aa.jpg"
    },
    "name": "test-01",
    "name_with_namespace": "git_user1/test-01",
    "path": "test-01",
    "path_with_namespace": "git_user1/test-01",
    "default_branch": "master",
    "ssh_url_to_repo": "git@git.code.tencent.com:git_user1/test-01.git",
    "http_url_to_repo": "http://git.code.tencent.com/git_user1/test-01.git",
    "https_url_to_repo": "https://git.code.tencent.com/git_user1/test-01.git",
    "web_url": "http://git.code.tencent.com/git_user1/test-01",
    "tag_list": [],
    "issues_enabled": true,
    "merge_requests_enabled": true,
    "wiki_enabled": true,
    "snippets_enabled": true,
    "review_enabled": true,
    "fork_enabled": false,
    "tag_name_regex": null,
    "tag_create_push_level": 30,
    "created_at": "2017-08-13T07:37:14+0000",
    "last_activity_at": "2017-08-13T07:37:14+0000",
    "creator_id": 11323,
    "avatar_url": "http://git.code.tencent.com/uploads/project/avatar/70703",
    "watchs_count": 0,
    "stars_count": 0,
    "forks_count": 0,
    "config_storage": {
        "limit_lfs_file_size": 500,
        "limit_size": 100000,
        "limit_file_size": 100000,
        "limit_lfs_size": 100000
    },
    "forked_from_project": "Forked Project not found",
    "statistics": {
        "commit_count": 0,
        "repository_size": 0
    }
}
]
```

### 查询项目的事件列表

查询项目的事件列表


```
GET /api/v3/projects/:id/events
```


**参数：**



| 参数                | 类型                     | 描述                                        |
| ----------------- | ---------------------- | ----------------------------------------- |
| id                | integer or string      | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| user_id_or_name   | integer 或 string（可选）   | 用户的 id或用户名                                |
| page              | integer                | 页数（默认值：1）                                 |
| per_page          | integer                | 默认页面大小（默认值： 20，最大值： 100）                  |


**返回值：**


```json
[
    {
        "title": null,
        "project_id": 28,
        "action_name": "CREATED",
        "target_id": 387,
        "target_type": "MergeRequest",
        "author_id": 11323,
        "author_username": "git_user1",
        "created_at": "2016-09-25T08:46:28+0000",
        "target_title": null,
        "data": {}
    },
    {
        "title": null,
        "project_id": 28,
        "action_name": "COMMENTED",
        "target_id": 732,
        "target_type": "note",
        "author_id": 11323,
        "author_username": "git_user1",
        "created_at": "2016-09-25T08:46:28+0000",
        "target_title": null,
        "data": {}
      },
]
```

标星

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子：/api/v3/projects/diaspora%2Fdiaspora (/ 编码 %2F)

对指定项目标星

### 对指定项目进行标星


```
PUT /api/v3/projects/:id/star
```


**参数：**



| 参数   | 类型                 | 描述                                        |
| ---- | ------------------ | ----------------------------------------- |
| id   | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |


**返回值：**


### true 或 flase

取消对指定项目标星

### 取消当前用户对指定项目标星


```
DELETE /api/v3/projects/:id/star
```


**参数：**



| 参数   | 类型                 | 描述                                        |
| ---- | ------------------ | ----------------------------------------- |
| id   | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |


**返回值：**


200 或相关状态码
查看对指定项目是否标星

### 查看当前用户对指定项目是否标星


```
GET /api/v3/projects/:id/star
```


**参数：**



| 参数   | 类型                 | 描述                                        |
| ---- | ------------------ | ----------------------------------------- |
| id   | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |


**返回值：**


### true 或 false

获取标星项目列表

### 获取对指定项目的标星用户列表


```
GET /api/v3/projects/:id/stars
```


**参数：**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| page       | integer（可选）        | 页数（默认值：1）                                 |
| per_page   | integer（可选）        | 默认页面大小（默认值： 20，最大值： 100）                  |


**返回值：**


```json
[
   {
        "project_id": 2654,
        "user": {
            "id": 11323,
            "username": "git_user1",
            "web_url": "http://git.code.tencent.com/u/git_user1",
            "name": "git_user1",
            "state": "active",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/a75ba2727c7a409cab1d15dd993149aa.jpg"
        }
    }
]
```
