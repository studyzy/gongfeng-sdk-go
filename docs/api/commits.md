# 提交相关

提交

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

### 获取一个提交

返回项目版本库某个特定的提交


```
GET /api/v3/projects/:id/repository/commits/:sha
```


**参数**



| 参数    | 类型                 | 描述                                        |
| ----- | ------------------ | ----------------------------------------- |
| id    | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| sha   | string             | commit hash值、分支名或 tag                     |


**返回值**


```json
{
  "id": "6269835cafdd89e6d8b8d4a4e3738c79206c1d06",
  "message": "推送新内容",
  "parent_ids": [
    "11a63fa022906668338166ce7f2e7bf35d502285"
  ],
  "authored_date": "2012-06-07T07:38:33+0000",
  "author_name": "git_user1",
  "author_email": "git_user1@tencent.com",
  "committed_date": "2012-06-07T07:38:33+0000",
  "committer_name": "git_user1",
  "committer_email": "git_user1@tencent.com",
  "title": "推送新内容",
  "created_at": "2012-06-07T07:38:33+0000"
  "short_id": "6269835c"
}
```

### 获得提交的差异

返回项目版本库提交的差异。


```
GET /api/v3/projects/:id/repository/commits/:sha/diff
```


**参数**



| 参数                   | 类型                 | 描述                                        |
| -------------------- | ------------------ | ----------------------------------------- |
| id                   | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| sha                  | string             | commit hash值、分支名或 tag                     |
| path                 | string（可选）         | 文件路径                                      |
| ignore_white_space   | boolean（可选）        | 有差异的内容是否忽略空白符，默认不忽略                       |


**返回值**


```json
[
  {
    "old_path": "git-01",
    "new_path": "git-01",
    "a_mode": 33188,
    "b_mode": 33188,
    "diff": "@@ -16,7 +16,7 @@\
 \
 \	\	\	<c:if test=\\"${ not empty project.description }\\">\
 \	\	\	\	<div class=\\"str-truncated project-home-desc text-left\\" style=\\"width:250px;\\">\
-\	\	\	\	\	\	${ tgfn:markdown( project.description , markupContext) }\
+\	\	\	\	\	\	${ tgfn:renderDescription( project.description , markupContext) }\
 \	\	\	\	</div>\
 \	\	\	</c:if>\
 \	\	</div>\
",
    "new_file": false,
    "renamed_file": false,
    "deleted_file": false,
    "is_too_large": false,
    "is_collapse": false
  }
]
```

### 下载 MergeRequesr 差异文件集

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

对一个提交发表评论

向一个提交添加评论

注意： 如果选择在一个特定的提交行上添加评论，其中参数path、line和line_type是必需的。


```
POST /api/v3/projects/:id/repository/commits/:sha/comments
```


**参数**



| 参数          | 类型                 | 描述                                        |
| ----------- | ------------------ | ----------------------------------------- |
| id          | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| sha         | string             | commit hash值、分支名或 tag                     |
| note        | string             | 评论内容                                      |
| path        | string（可选）         | 文件路径                                      |
| line        | integer（可选）        | 行号                                        |
| line_type   | string（可选）         | 变更类型，可选old、new                            |


**返回值**


```json
{
    "note": "api test",
    "author": {
        "id": 2274,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/2274/a75ba2727c7a409cab1d15dd993149aa.jpg"
    },
    "created_at": "2017-08-14T03:33:55+0000"
}
```

### 获取一个提交的评论

返回项目版本库某个提交的评论


```
GET /api/v3/projects/:id/repository/commits/:sha/comments
```


**参数**



| 参数    | 类型                 | 描述                                        |
| ----- | ------------------ | ----------------------------------------- |
| id    | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| sha   | string             | commit hash值、分支名或 tag                     |


**返回值**


```json
[
  {
    "note": "api test",
    "author": {
        "id": 2274,
        "username": "git_user1",
        "web_url": "http://git.code.tencent.com/u/git_user1",
        "name": "git_user1",
        "state": "active",
        "avatar_url": "git.code.tencent.com/uploads/user/avatar/2274/a75ba2727c7a409cab1d15dd993149aa.jpg"
    },
    "created_at": "2017-08-14T03:33:55+0000"
}
]
```

### 获取项目版本库所有的提交

返回项目版本库（版本库指定文件或者目录）提交记录


```
GET /api/v3/projects/:id/repository/commits
```


**参数**



| 参数         | 类型                           | 描述                                                                                                         |
| ---------- | ---------------------------- | ---------------------------------------------------------------------------------------------------------- |
| id         | integer 或 string             | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH                                                                    |
| ref_name   | string（可选）                   | 版本库分支 或 tag，默认：默认分支                                                                                        |
| path       | string（可选）                   | 文件路径                                                                                                       |
| page       | integer（可选）                  | 分页（默认值：1）                                                                                                  |
| per_page   | integer（可选）                  | 默认页面大小（默认值： 20，最大值： 100）                                                                                   |
| since      | yyyy-MM-dd'T'HH:mm:ssZ（可选）   | 此日期及之后的提交： 例如2019-03-25T00:10:19+0000 或 2019-03-25T00:10:19+0800，时间参数必须转码，如"2019-03-25T00:10:19%2B0800"    |
| until      | yyyy-MM-dd'T'HH:mm:ssZ（可选）   | 此日期及之前的提交；例如2019-03-25T00:10:19+0000 或 2019-03-25T00:10:19+0800，时间参数必须转码，如"2019-03-25T00:10:19%2B0800" ）   |


**返回值**


```json
[
  {
    "id": "e9c5aa72c9e22592ee2611d896b730c1a95cc339",
    "short_id": "e9c5aa72"
    "title": "asd",
    "author_name": "git_user1",
    "author_email": "git_user1@tencent.com",
    "created_at": "2012-09-19T09:29:19+0000",
    "message": "asd"
  },
  {
    "id": "11a63fa022906668338166ce7f2e7bf35d502285",
    "short_id": "11a63fa0"
    "title": "dd",
    "author_name": "git_user1",
    "author_email": "git_user1@tencent.com",
    "created_at": "2012-09-19T09:29:19+0000",
    "message": "dd"
  }
]
```

### 获取某个提交对应的分支和tag

获取这个提交被推送到项目的哪个分支或tag


```
GET /api/v3/projects/:id/repository/commits/:sha/refs
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| sha        | string             | commit hash值、分支名或 tag                     |
| type       | string（可选）         | branch、tag、或all （默认：all）                  |
| page       | integer（可选）        | 分页（默认值：1）                                 |
| per_page   | integer（可选）        | 默认页面大小（默认值： 20，最大值： 100）                  |


**返回值**


```json
[
    {
        "type": "branch",
        "name": "copy"
    },
    {
        "type": "branch",
        "name": "master"
    },
    {
        "type": "tag",
        "name": "v1.1.8"
    }
]
```
