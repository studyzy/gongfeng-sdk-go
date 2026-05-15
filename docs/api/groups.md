# 项目组


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

组权限

组/项目的权限access_level包括：

GUEST = 10
FOLLOWER = 15
REPORTER = 20
DEVELOPER = 30
MASTER = 40
OWNER = 50

### 新建项目组

创建一个新的项目组

注意: 仅对有权限创建项目组的用户可用


```
POST /api/v3/groups
```


**参数:**



| 参数            | 类型       | 描述           |
| ------------- | -------- | ------------ |
| name          | string   | 项目组的名字       |
| path          | string   | 项目组的路径       |
| description   | string   | 关于这个项目组的描述   |


**返回值：**


```json
[
  {
    "id": 2053,
    "name": "t-git",
    "path": "t-git",
    "web_url": "http://git.code.tencent.com/groups/t-git",
    "description": ""
    "avatar_url": null
  }
]
```

### 编辑项目组

更新项目组的名称和描述(暂不支持更新项目组路径)


```
PUT /api/v3/groups
```


**参数:**



| 参数            | 类型        | 描述               |
| ------------- | --------- | ---------------- |
| id            | integer   | 用户的项目组 ID 或者路径   |
| name          | string    | 项目组的名字           |
| description   | string    | 关于这个项目组的描述       |


**返回值：**


```json
[
  {
    "id": 2053,
    "name": "wevas2",
    "path": "wevas",
    "web_url": "http://git.code.tencent.com/groups/wevas",
    "description": "wevas2",
    "avatar_url": "http://git.code.tencent.com/assets/images/avatar/no_group_avatar.png"
  }
]
```
### 删除项目组

删除一个指定的项目组


```
DELETE /api/v3/groups/:id
```


**参数:**



| 参数   | 类型        | 描述                |
| ---- | --------- | ----------------- |
| id   | integer   | id = 项目组唯一标识或路径   |


**返回值**


### 200 或相关状态码
获取项目组列表

获取项目组列表

注意： 作为用户获取的是我的项目组，作为管理者获取的是所有的项目组。


```
GET /api/v3/groups
```


**参数:**



| 参数         | 类型            | 描述                         |
| ---------- | ------------- | -------------------------- |
| page       | integer（可选）   | 分页 (默认值:1)                 |
| per_page   | integer（可选）   | 默认页面大小 (默认值 20，最大值： 100)   |


**返回值：**


```json
[
  {
    "id": 2053,
    "name": "t-git",
    "path": "t-git",
    "web_url": "http://git.code.tencent.com/groups/t-git",
    "description": ""
    "avatar_url": null
  }
]
```


### 你能够通过名称或者路径查询项目组，如下。

搜索项目组

### 通过名称或者路径查询所有匹配的项目组


```
GET /api/v3/groups?search=git
```


**参数:**



| 参数         | 类型            | 描述                         |
| ---------- | ------------- | -------------------------- |
| page       | integer（可选）   | 分页 (默认值:1)                 |
| per_page   | integer（可选）   | 默认页面大小 (默认值 20，最大值： 100)   |


**返回值：**


```json
[
  {
    "id": 3157,
    "name": "git",
    "path": "git",
    "web_url": "http://git.code.tencent.com/groups/git",
    "description": "",
    "avatar_url": null
  }
]
```

项目组成员

项目组的访问级别定义如下：

GUEST     = 10
REPORTER  = 20
FOLLOWER  = 25
DEVELOPER = 30
MASTER    = 40
### OWNER     = 50

获取项目组成员列表

### 获取认证用户可见的项目组成员列表


```
GET /api/v3/groups/:id/members
```


**参数:**



| 参数         | 类型            | 描述                       |
| ---------- | ------------- | ------------------------ |
| id         | integer       | id = 项目组唯一标识或路径          |
| page       | integer（可选）   | 分页 (默认:1)                |
| per_page   | integer（可选）   | 默认页面大小 (默认 20，最大： 100)   |


**返回值：**


```json
[
  {
    "id": 1,
    "username": "xiaoshu",
    "web_url": "http://git.code.tencent.com/u/xiaoshu",
    "name": "xiaoshu",
    "state": "active",
    "avatar_url": "http://git.code.tencent.com/uploads/user/avatar/1/2c37c36551d84f9185a3d8cf04620981.gif",
    "access_level": 50
  },
  {
    "id": 2,
    "username": "long",
    "web_url": "http://git.code.tencent.com/u/long",
    "name": "long",
    "state": "active",
    "avatar_url": "http://git.code.tencent.com/uploads/user/avatar/2/2c32c36551d84f9185a3d8cf04620981.gif",
    "access_level": 30
  }
]
```

### 增加项目组成员

增加一个项目组成员


```
POST /api/v3/groups/:id/members
```


**参数:**



| 参数             | 类型        | 描述                |
| -------------- | --------- | ----------------- |
| id             | integer   | id = 项目组唯一标识或路径   |
| user_id        | integer   | 添加用户的 ID          |
| access_level   | integer   | 项目访问级别            |


**返回值：**


```json
[
  {
    "id": 2526,
    "username": "bili",
    "web_url": "http://git.code.tencent.com/u/bili",
    "name": "bili",
    "state": "active",
    "avatar_url": "http://git.code.tencent.com/uploads/user/avatar/2526/2c37c36551d84f9185a3d8cf04620981.gif",
    "access_level": 30
  }
]
```

### 修改项目组成员

更新一个项目组成员的访问权限


```
PUT /api/v3/groups/:id/members/:user_id
```


**参数:**



| 参数             | 类型        | 描述                |
| -------------- | --------- | ----------------- |
| id             | integer   | id = 项目组唯一标识或路径   |
| user_id        | integer   | 项目组成员的 ID         |
| access_level   | integer   | 项目访问级别            |


**返回值：**


```json
[
  {
    "id": 2526,
    "username": "bili",
    "web_url": "http://git.code.tencent.com/u/bili",
    "name": "bili",
    "state": "active",
    "avatar_url": "http://git.code.tencent.com/uploads/user/avatar/2526/2c37c36551d84f9185a3d8cf04620981.gif",
    "access_level": 40
  }
]
```

### 移除一个项目组成员

移除一个指定的项目组成员


```
DELETE /api/v3/groups/:id/members/:user_id
```


**参数:**



| 参数        | 类型        | 描述                |
| --------- | --------- | ----------------- |
| id        | integer   | id = 项目组唯一标识或路径   |
| user_id   | integer   | 项目组成员的 ID         |


**返回值**


### 200 或相关状态码
获取项目组的详细信息以及项目组下所有项目

### 获取对指定项目组的详细信息，以及项目组下所有项目


```
GET /api/v3/groups/:id
```


**参数:**


```json
{
    "id": 25123,
    "name": "test00",
    "path": "test00",
    "web_url": "http://git.code.tencent.com/groups/test00",
    "description": "xxx",
    "avatar_url": null,
    "projects": [
        {
            "id": 64612,
            "description": "0711",
            "public": false,
            "archived": false,
            "visibility_level": 0,
            "namespace": {
                "id": 25123,
                "name": "test00",
                "path": "test00",
                "web_url": "http://git.code.tencent.com/groups/test00",
                "description": "xxx",
                "avatar_url": null
            },
            "name": "test1",
            "name_with_namespace": "test00/test1",
            "path": "test1",
            "path_with_namespace": "test00/test1",
            "default_branch": "master",
            "ssh_url_to_repo": "git@git.code.tencent.com:test00/test1.git",
            "http_url_to_repo": "http://git.code.tencent.com/test00/test1.git",
            "https_url_to_repo": "https://git.code.tencent.com/test00/test1.git",
            "web_url": "http://git.code.tencent.com/test00/test1",
            "tag_list": [],
            "issues_enabled": true,
            "merge_requests_enabled": true,
            "wiki_enabled": true,
            "snippets_enabled": true,
            "review_enabled": true,
            "fork_enabled": false,
            "tag_name_regex": null,
            "tag_create_push_level": 30,
            "created_at": "2017-07-11T06:29:24+0000",
            "last_activity_at": "2017-07-11T06:46:33+0000",
            "creator_id": 18123,
            "avatar_url": "http://git.code.tencent.com/uploads/project/avatar/123",
            "watchs_count": 2,
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
                "commit_count": 3,
                "repository_size": 0.004
            }
        },
        {
            "id": 66062,
            "description": "gggg",
            "public": false,
            "archived": false,
            "visibility_level": 0,
            "namespace": {
                "id": 25123,
                "name": "test00",
                "path": "test00",
                "web_url": "http://git.code.tencent.com/groups/test00",
                "description": "xxx",
                "avatar_url": null
            },
            "name": "test02",
            "name_with_namespace": "test00/test02",
            "path": "test02",
            "path_with_namespace": "test00/test02",
            "default_branch": "master",
            "ssh_url_to_repo": "git@git.code.tencent.com:test00/test02.git",
            "http_url_to_repo": "http://git.code.tencent.com/test00/test02.git",
            "https_url_to_repo": "https://git.code.tencent.com/test00/test02.git",
            "web_url": "http://git.code.tencent.com/test00/test02",
            "tag_list": [],
            "issues_enabled": true,
            "merge_requests_enabled": true,
            "wiki_enabled": true,
            "snippets_enabled": true,
            "review_enabled": true,
            "fork_enabled": false,
            "tag_name_regex": null,
            "tag_create_push_level": 30,
            "created_at": "2018-07-19T07:13:22+0000",
            "last_activity_at": "2018-07-19T07:13:23+0000",
            "creator_id": 18604,
            "avatar_url": "http://git.code.tencent.com/uploads/project/avatar/66062",
            "watchs_count": 2,
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
}
```

| 参数   | 类型        | 描述                |
| ---- | --------- | ----------------- |
| id   | integer   | id = 项目组唯一标识或路径   |
