# 复刻


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子：/api/v3/projects/diaspora%2Fdiaspora (/ 编码 %2F)

### Fork 项目

Fork项目到用户的命名空间


```
POST /api/v3/projects/fork/:id
```


**参数**



| 参数   | 类型                  | 描述                                        |
| ---- | ------------------- | ----------------------------------------- |
| id   | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |


**返回值**


```json
{
    "id": 561,
    "description": null,
    "public": false,
    "archived": false,
    "visibility_level": 0,
    "namespace": {
        "created_at": "2017-01-29T07:49:48+0000",
        "description": "git_user1",
        "id": 2513,
        "name": "git_user1",
        "owner_id": 11323,
        "path": "git_user1",
        "updated_at": "2017-01-29T07:49:48+0000"
    },
    "owner": {
        "id": 11323,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/11323/fec6c477729445df9a6e0a4a05b4a86c.png"
    },
    "name": "fork_project",
    "name_with_namespace": "git_user1/fork_project",
    "path": "fork_project",
    "path_with_namespace": "git_user1/fork_project",
    "default_branch": "master",
    "ssh_url_to_repo": "git@git.code.tencent.com:git_user1/fork_project.git",
    "http_url_to_repo": "http://git.code.tencent.com/git/git_user1/fork_project.git",
    "https_url_to_repo": "https://git.code.tencent.com/git/git_user1/fork_project.git",
    "web_url": "http://git.code.tencent.com/git_user1/fork_project",
    "tag_list": [],
    "issues_enabled": true,
    "merge_requests_enabled": true,
    "wiki_enabled": true,
    "snippets_enabled": true,
    "review_enabled": true,
    "fork_enabled": false,
    "tag_name_regex": null,
    "tag_create_push_level": 30,
    "created_at": "2017-08-14T06:29:25+0000",
    "last_activity_at": "2017-08-14T06:29:25+0000",
    "creator_id": 11323,
    "avatar_url": "http://git.code.tencent.com/uploads/project/avatar/561",
    "watchs_count": 1,
    "stars_count": 0,
    "forks_count": 0,
    "config_storage": {
        "limit_lfs_file_size": 500,
        "limit_size": 100000,
        "limit_file_size": 100,
        "limit_lfs_size": 100000
    },
    "forked_from_project": {
        "path": "fork_project",
        "path_with_namespace": "test-06/fork_project",
        "name": "fork_project",
        "id": 32811,
        "name_with_namespace": "test-0607/fork_project"
    },
    "statistics": {
        "commit_count": 0,
        "repository_size": 0
    }
}
```

将项目 Fork 到另外一个命名空间

### 将项目 Fork 到另外一个命名空间，创建 fork 关系


```
POST /api/v3/projects/:id/fork/:forked_from_id
```


**参数**



| 参数               | 类型                  | 描述                                        |
| ---------------- | ------------------- | ----------------------------------------- |
| id               | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| forked_from_id   | integer             | 需要 fork 的项目 id                            |


**返回值**


### 状态码

删除 Fork 关系

### 删除 fork 的项目，删除 fork 关系


```
DELETE /api/v3/projects/:id/fork
```


**参数**



| 参数   | 类型                  | 描述                                        |
| ---- | ------------------- | ----------------------------------------- |
| id   | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |


**返回值**


### 状态码
