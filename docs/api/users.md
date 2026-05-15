# 用户


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

### 获取用户信息列表

获取用户信息列表，接口使用page 和 per_page作为分页参数限制返回列表的大小

作为普通用户

**参数:**



| 参数         | 类型            | 描述                         |
| ---------- | ------------- | -------------------------- |
| page       | integer（可选）   | 页数 (默认值:1)                 |
| per_page   | integer（可选）   | 默认页面大小 (默认值 20，最大值： 100)   |


```
GET /api/v3/users
```


**返回值：**


```json
[
    {
        "id": 2235,
        "username": "luci",
        "web_url": "http://git.code.tencent.com/u/luci",
        "name": "luci",
        "state": "active",
        "avatar_url": "git.code.tencent.com/assets/images/avatar/no_user_avatar.png"
    }
]
```


作为管理者

**参数:**



| 参数         | 类型            | 描述                         |
| ---------- | ------------- | -------------------------- |
| page       | integer（可选）   | 分页 (默认值:1)                 |
| per_page   | integer（可选）   | 默认页面大小 (默认值 20，最大值： 100)   |


```
GET /api/v3/users
```


**返回值：**


```json
[
   {
        "id": 1,
        "email": "xiaoshu@tencent.com",
        "name": "xiaoshu",
        "username": "xiaoshu",
        "web_url": "http://git.code.tencent.com/u/xiaoshu",
        "current_sign_in_at": "2016-01-11T01:55:12+0000",
        "is_admin": false,
        "projects_limit": 9999999,
        "skype": "",
        "linkedin": "",
        "twitter": "",
        "theme_id": 2,
        "bio": null,
        "state": "blocked",
        "color_scheme_id": 1,
        "website_url": "",
        "created_at": "2015-04-27T02:44:00+0000",
        "identities": [
            {
                "extern_uid": "xiaoshu",
                "provider": "ldapmain"
            },
            {
                "extern_uid": "xiaoshu",
                "provider": "tof"
            }
        ],
        "can_create_project": true
    },
   {
        "id": 657,
        "email": "long@tencent.com",
        "name": "long",
        "username": "long",
        "web_url": "http://git.code.tencent.com/u/long",
        "current_sign_in_at": "2015-04-29T09:17:49+0000",
        "is_admin": false,
        "projects_limit": 9999999,
        "skype": "",
        "linkedin": "",
        "twitter": "",
        "theme_id": 2,
        "bio": null,
        "state": "blocked",
        "color_scheme_id": 1,
        "website_url": "",
        "created_at": "2015-04-29T09:17:49+0000",
        "identities": [
            {
                "extern_uid": "long",
                "provider": "ldapmain"
            },
            {
                "extern_uid": "long",
                "provider": "tof"
            }
        ],
        "can_create_project": true
    },
]
```

### 获用户关注项目列表

获取当前用户关注的项目列表


```
GET /api/v3/user/watched
```


**返回值：**


```json
[
    {
        "id": 4484,
        "description": "",
        "public": false,
        "archived": false,
        "visibility_level": 0,
        "namespace": {
            "created_at": "2016-06-06T08:02:39+0000",
            "description": "",
            "id": 6854,
            "name": "git_user",
            "owner_id": 3654,
            "path": "git_user",
            "updated_at": "2016-06-06T08:02:39+0000"
        },
        "owner": {
            "id": 627,
            "username": "git_user",
            "web_url": "http://git.code.tencent.com/u/git_user",
            "name": "git_user",
            "state": "blocked",
            "avatar_url": "git.code.tencent.com/uploads/user/avatar/627/ae4f4ab5f5b6445bb97404a4638b68d7.gif"
        },
        "name": "pro-111",
        "name_with_namespace": "git_user/pro-111",
        "path": "pro-111",
        "path_with_namespace": "git_user/pro-111",
        "default_branch": "master",
        "ssh_url_to_repo": "git@git.code.tencent.com:git_user/pro-111.git",
        "http_url_to_repo": "http://git.code.tencent.com/git_user/pro-111.git",
        "https_url_to_repo": "https://git.code.tencent.com/git_user/pro-111.git",
        "web_url": "http://git.code.tencent.com/git_user/pro-111",
        "tag_list": [],
        "issues_enabled": true,
        "merge_requests_enabled": true,
        "wiki_enabled": true,
        "snippets_enabled": true,
        "review_enabled": true,
        "fork_enabled": false,
        "tag_name_regex": null,
        "tag_create_push_level": 30,
        "created_at": "2018-01-11T02:26:04+0000",
        "last_activity_at": "2018-03-13T07:21:28+0000",
        "creator_id": 6227,
        "avatar_url": "http://git.code.tencent.com/uploads/project/avatar/4484",
        "watchs_count": 5,
        "stars_count": 0,
        "forks_count": 2,
        "config_storage": {
            "limit_lfs_file_size": 500,
            "limit_size": 100000,
            "limit_file_size": 100000,
            "limit_lfs_size": 100000
        },
        "forked_from_project": "Forked Project not found",
        "statistics": {
            "commit_count": 8,
            "repository_size": 0.013
        }
    },
]
```

### 获取单个用户信息

获取一个单独用户的信息

### 获取某个用户的账号信息

获取某个指定用户的信息


```
GET /api/v3/users/:id
```


**参数:**



| 参数   | 类型                 | 描述               |
| ---- | ------------------ | ---------------- |
| id   | integer 或 string   | id=用户唯一标识或用户名称   |


**返回值：**


```json
{
  "id": 45,
    "email": "bobi@tencent.com",
    "name": "bobi",
    "username": "bobi",
    "web_url": "http://git.code.tencent.com/u/bobi",
    "current_sign_in_at": "2015-07-06T04:12:40+0000",
    "is_admin": false,
    "projects_limit": 9999999,
    "skype": null,
    "linkedin": null,
    "twitter": null,
    "theme_id": 1,
    "bio": null,
    "state": "active",
    "color_scheme_id": 1,
    "website_url": null,
    "created_at": "2015-01-25T08:12:01+0000",
    "identities": [
        {
            "extern_uid": "bobi",
            "provider": "ldapmain"
        },
        {
            "extern_uid": "bobi",
            "provider": "tof"
        }
    ],
    "can_create_project": true
}
```

### 当前认证用户

获取当前认证用户信息


```
GET /api/v3/user
```


**返回值：**


```json
{
  "id": 18612,
    "email": "bob@tencent.com",
    "name": "bob",
    "username": "bob",
    "web_url": "http://git.code.tencent.com/u/bob",
    "current_sign_in_at": "2017-07-19T06:51:23+0000",
    "is_admin": true,
    "projects_limit": 9999999,
    "skype": "",
    "linkedin": "",
    "twitter": "",
    "theme_id": 1,
    "bio": "",
    "state": "active",
    "color_scheme_id": 1,
    "website_url": "",
    "created_at": "2017-01-16T08:29:42+0000",
    "identities": [
        {
            "extern_uid": "bob",
            "provider": "ldapmain"
        },
        {
            "extern_uid": "bob",
            "provider": "tof"
        }
    ],
    "private_token": "QJR_AAVFp7nfoAWSJgz"
}
```

### 给当前用户创建一个SSH key

创建一个新的SSH key给当前认证的用户


```
POST /api/v3/user/keys
```


**参数:**



| 参数      | 类型       | 描述            |
| ------- | -------- | ------------- |
| title   | string   | SSH Key 的标题   |
| key     | string   | 新的 SSH key    |


**返回值：**


```json
{
  "id": 10127,
  "title": "test",
  "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDXawA7+8VH0FmLbJ0m0hMfXjmQD/DXSsbL8DicG9/jk34BLr8mQ23EXr38te5wvRejhYV4ov3/V1HGB6M2o/cDVUc1ODQpcEBa6jZ1Q/VoqVO+49+UHN/FeOgQk60bF7Nu4hXhn/e/H6Tw2fTvUW1LPpaPtqNiqcylqtZ2qfLHfUXXXvOmRmM/4sTBZnjMoK61moZgxtrYaDtZOVxVVqF+AVpnbfHUDdGZofNQe2AW2g1PXpu3ikgeyxvXGoBLKXo6r1KVYD+uSPf4OouHnUsBpYJKy6PUpjxs13yzpk65TDy1xX4xylbO0TVEhNYXax2K6ih2RPPmSpq8juZzr7RZ",
  "created_at": "2017-08-13T03:19:54+0000"
}
```


创建成功将返回状态201 Created。如果返回400 Bad Request错误将附带一个说明的错误信息:

```json
{
  "message": {
    "fingerprint": [
      "has already been taken"
    ],
    "key": [
      "has already been taken"
    ]
  }
}
```

### 获取当前用户的SSH key

获取当前认证用户的SSH keys列表


```
GET /api/v3/user/keys
```


**返回值：**


```json
[
  {
    "id": 2356,
    "title": "git-key1",
    "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDXawA7+8VH0FmLbJ0m0hMfXjmQD/DXSsbL8DicG9/jk34BLr8mQ23EXr38te5wvRejhYV4ov3/V1HGB6M2o/cDVUc1ODQpcEBa6jZ1Q/VoqVO+49+UHN/FeOgQk60bF7Nu4hXhn/e/H6Tw2fTvSSSSPpaPtqNiqcylqtZ2qfLHfUXv5vOmRmM/4sTBZnjMoK61moZgxtrYaDtZOVxVVqF+AVpnbfHUDdGZofNQe2AW2g1PXpu3ikgeyxvXGoBLKXo6r1KVYD+uSPf4OouHnUsBpYJKy6PUpjxs13yzpk65TDy1xX4xylbO0TVEhNYXax2K6ih2RPPmSpq8juZzr7RZ",
    "created_at": "2012-06-28T02:33:02+0000"
  }
]
```

### 获取某个指定的 SSH key

获取一个特定的SSH key


```
GET /api/v3/user/keys/:id
```


**参数:**



| 参数   | 类型        | 描述             |
| ---- | --------- | -------------- |
| id   | integer   | SSH key 的 id   |


**返回值：**


```json
{
  "id": 215,
  "title": "git-key1",
  "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDXawA7+8VH0FmLbJ0m0hMfXjmQD/DXSsbL8DicG9/jk34BLr8mQ23EXr38te5wvRejhYV4ov3/V1HGB6M2o/cDVUc1ODQpcEBa6jZ1Q/VoqVO+49+UHN/FeOgQk60bF7Nu4hXhn/e/H6Tw2fTvUW1LPpaPtqNiqcylqtZ2qfLHfUXv5vOmRmM/4sTBZnjMoK61moZgxtrYaDtZOVxVVqF+AVpnbfHUDdGZofNQe2AW2g1PXpu3ikgeyxvXGoBLKXo6r1KVYD+uSPf4OouHnUsBpYJKy6PUpjxs13yzpk65TDy1xX4xylbO0TVEhNYXax2K6ih2RPPmSpq8juZzr7RZ",
  "created_at": "2012-06-28T02:33:02+0000"
}
```

### 删除当前用户的 SSH key

删除当前认证用户拥有的SSH key。这是一个幂等的操作，调用它删除一个存在的或已经删除的key返回的结果都是200 OK。


```
DELETE /api/v3/user/keys/:id
```


**参数:**



| 参数   | 类型        | 描述             |
| ---- | --------- | -------------- |
| id   | integer   | SSH key 的 ID   |

### 添加邮箱

给当前认证用户增加邮箱


```
POST /api/v3/user/emails
```


**参数:**



| 参数      | 类型       | 描述     |
| ------- | -------- | ------ |
| email   | string   | 邮箱地址   |


**返回值：**


```json
{
  "id": 98,
  "email": "git02@tencent.com"
}
```


调用成功返回状态201 Created。如果发生400 Bad Request错误的话附带相关的错误信息：

```json
{
  "message": {
    "email": [
      "has not given"
    ]
  }
}
```

### 通过邮箱获取用户信息

通过邮箱获取一个用户的基本信息


```
GET /api/v3/user/email
```


**参数:**



| 参数      | 类型       | 描述     |
| ------- | -------- | ------ |
| email   | string   | 邮箱地址   |


**返回值：**


```json
{
    "name": "long",
    "username": "long",
    "web_url": "http://git.code.tencent.com/u/long"
}
```

### 获取用户邮箱列表

获取当前用户的邮箱列表


```
GET /api/v3/user/emails
```


**返回值：**


```json
[
  {
    "id": 56,
    "email": "git_user1@tencent.com"
  },
  {
    "id": 98,
    "email": "git_user1@qq.com"
  }
]
```

### 获取邮箱信息

获取某个特定邮箱信息


```
GET /api/v3/user/emails/:id
```


**参数:**



| 参数   | 类型        | 描述       |
| ---- | --------- | -------- |
| id   | integer   | 邮箱的 ID   |


**返回值：**


```json
{
  "id": 56,
  "email": "git_user1@tencent.com"
}
```

### 删除当前用户的邮箱

删除当前认证用户拥有的邮箱。这是一个幂等的操作，用户调用它删除一个存在的或已经删除的邮箱返回结果都是200 OK。


```
DELETE /api/v3/user/emails/:id
```


**参数:**



| 参数   | 类型        | 描述       |
| ---- | --------- | -------- |
| id   | integer   | 邮箱的 ID   |
