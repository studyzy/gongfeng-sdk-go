# Tag 相关

Tag

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

### Tag列表

返回项目版本库 tag 列表，按字母顺序排序


```
GET /api/v3/projects/:id/repository/tags
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| page       | integer（可选）        | 分页（default：1）                             |
| per_page   | integer（可选）        | 默认页面大小（default：20，max：100）                |


**返回值**


```json
[
  {
    "name": "v1.0",
    "message": "1.0版本第一次上线版本",
    "commit": {
      "id": "6be50b7afaf98f49896b45ac6c19fb19315286cc",
      "message": "Merge branch 'master'",
      "parent_ids": [
        "f888582ad6976daef5ba793d4663a60d5b031151",
        "6ad05205ec9a4681d1ce17f38ba0a64729981db9"
      ],
      "authored_date": "2016-10-23T23:10:19+0000",
      "author_name": "userExample",
      "author_email": "user_example@tencent.com",
      "committed_date": "2016-10-23T23:10:19+0000",
      "committer_name": "userExample",
      "committer_email": "user_example@tencent.com",
      "title": "Merge branch 'master'",
      "created_at": "2016-10-23T23:10:19+0000"
      "short_id": "6be50b7a"
    }
  },
  {
    "name": "v1.0.1",
    "message": "1.0.0版本 bugfixed",
    "commit": {
      "id": "7eee7c871a74f114bbc0df10c9bbacee870ae409",
      "message": "fixed：测试环境 pageContext.request.localName 处理速度慢问题处理",
      "parent_ids": [
        "6b7fa77bf10d348195e1a79312471f8b3d1aef4a"
      ],
      "authored_date": "2016-11-21T03:19:38+0000",
      "author_name": "userExample",
      "author_email": "user_example@tencent.com",
      "committed_date": "2016-11-21T03:19:38+0000",
      "committer_name": "userExample",
      "committer_email": "user_example@tencent.com",
      "title": "fixed：测试环境 pageContext.request.localName 处理速度慢问题处理",
      "created_at": "2016-11-21T03:19:38+0000"
      "short_id": "7eee7c87"
    }
  }
]
```

### 获取指定Tag

返回项目版本库某个 tag 详情


```
GET /api/v3/projects/:id/repository/tags/:tag
```


**参数**



| 参数    | 类型                 | 描述                                        |
| ----- | ------------------ | ----------------------------------------- |
| id    | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| tag   | string             | tag 名                                     |


**返回值**


```json
{
  "name": "v1.0",
  "message": "1.0版本第一次上线版本",
  "commit": {
    "id": "6be50b7afaf98f49896b45ac6c19fb19315286cc",
    "message": "Merge branch 'master'",
    "parent_ids": [
      "f888582ad6976daef5ba793d4663a60d5b031151",
      "6ad05205ec9a4681d1ce17f38ba0a64729981db9"
    ],
    "authored_date": "2016-10-23T23:10:19+0000",
    "author_name": "userExample",
    "author_email": "user_example@tencent.com",
    "committed_date": "2016-10-23T23:10:19+0000",
    "committer_name": "userExample",
    "committer_email": "user_example@tencent.com",
    "title": "Merge branch 'master'",
    "created_at": "2016-10-23T23:10:19+0000"
    "short_id": "6be50b7a"
  }
}
```

### 创建Tag

创建项目版本库 tag


```
POST /api/v3/projects/:id/repository/tags
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| tag_name   | string             | tag 名                                     |
| ref        | string             | 从 commit hash、存在的 branch 或 tag 创建 tag     |
| message    | string（可选）         | 描述、注释                                     |


**返回值**


```json
{
  "name": "ccccccccccc",
  "message": "message",
  "commit": {
    "id": "d21f748b73507a227a92c954bd3c89bf4e78e897",
    "message": "Merge branch 'b1' into 'master'\\r\
\\r\
xxx\\r\
\\r\
xxx\\r\
\\r\
See merge request !78",
    "parent_ids": [
      "78a1d588ce01b0e6aec46c7cc5aa407d6e7cfd24",
      "75029d005ac29572f04c461ab3257714245c9798"
    ],
    "authored_date": "2017-06-07T07:38:33+0000",
    "author_name": "userExample",
    "author_email": "user_example@tencent.com",
    "committed_date": "2017-06-07T07:38:33+0000",
    "committer_name": "userExample",
    "committer_email": "user_example@tencent.com",
    "title": "Merge branch 'b1' into 'master'",
    "created_at": "2017-06-07T07:38:33+0000"
    "short_id": "d21f748b"
  }
}
```

### 删除Tag

删除项目版本库 tag，需要有项目的 master 权限


```
DELETE /api/v3/projects/:id/repository/tags/:tag
```


**参数**



| 参数    | 类型                 | 描述                                        |
| ----- | ------------------ | ----------------------------------------- |
| id    | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| tag   | string             | tag 名                                     |


**返回值**


```json
{
  "tag_name": "my-removed-tag"
}
```
