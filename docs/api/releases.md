# 版本发布

Releases

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子：/api/v3/projects/diaspora%2Fdiaspora (/ 编码 %2F)

### 获取releases列表

返回项目版本库 releases 列表


```
GET /api/v3/projects/:id/releases
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
    "ref": {
        "name": "01api",
        "message": "01api",
        "commit": {
            "id": "44500ce35ff269d36925aab3a53dac296a91b19f",
            "message": "Merge branch 'dev' into 'master'See merge request !3",
            "parent_ids": [
                "6145857bfd5858b3ed1a5bff4dc4cfc4b62e29ff",
                "3dad4be822a900f5b8ab7a0112da0be5ed76deba"
            ],
            "authored_date": "2016-12-20T02:14:01+0000",
            "author_name": "bili",
            "author_email": "bili@tencent.com",
            "committed_date": "2016-12-20T02:14:01+0000",
            "committer_name": "bili",
            "committer_email": "bili@tencent.com",
            "title": "Merge branch '33' into 'master'",
            "created_at": "2016-12-20T02:14:01+0000",
            "short_id": "44500ce3"
        }
    },
    "id": 2761,
    "project_id": 85436,
    "tag": "01api",
    "title": "01api",
    "type": "release",
    "attachments": null,
    "description": null,
    "created_at": "2016-12-20T03:21:31+0000",
    "updated_at": null,
    "author": null
},
{
        "ref": {
            "name": "02api",
            "message": "merged",
            "commit": {
                "id": "44200ce35ff269d36925aab3a53dac296a91b19p",
                "message": "Merge branch '33' into 'master'\\r\
\\r\
保护分支master\\r\
\\r\
33\\r\
\\r\
See merge request !3",
                "parent_ids": [
                    "6146757bfd5858b3ed1a5bff4dc4cfc4b62e29aa",
                    "3dad4af822a900f5b8ab7a0112da0be5ed76debb"
                ],
                "authored_date": "2016-12-20T02:14:01+0000",
                "author_name": "",
                "author_email": "bili@tencent.com",
                "committed_date": "2016-12-20T02:14:01+0000",
                "committer_name": "bili",
                "committer_email": "bili@tencent.com",
                "title": "Merge branch '33' into 'master'",
                "created_at": "2016-12-20T02:14:01+0000",
                "short_id": "44200ce3"
            }
        },
        "id": 2795,
        "project_id": 85436,
        "tag": "02api",
        "title": "test",
        "type": "release",
        "attachments": null,
        "description": "test",
        "created_at": "2016-12-20T06:13:06+0000",
        "updated_at": null,
        "author": null
    }
]
```
### 获取某个指定的releases

在项目版本库中获取某个指定的releases

GRT /api/v3/projects/:id/releases/{release_id}

**参数**



| 参数           | 类型                 | 描述                                        |
| ------------ | ------------------ | ----------------------------------------- |
| id           | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| release_id   | Integer            | release的id                                |


**返回值**


```json
{
    "ref": {
        "name": "01api",
        "message": "01api",
        "commit": {
            "id": "44500ce35ff269d36925aab3a53dac296a91b19f",
            "message": "Merge branch 'dev' into 'master'See merge request !3",
            "parent_ids": [
                "6145857bfd5858b3ed1a5bff4dc4cfc4b62e29ff",
                "3dad4be822a900f5b8ab7a0112da0be5ed76deba"
            ],
            "authored_date": "2016-12-20T02:14:01+0000",
            "author_name": "bili",
            "author_email": "bili@tencent.com",
            "committed_date": "2016-12-20T02:14:01+0000",
            "committer_name": "bili",
            "committer_email": "bili@tencent.com",
            "title": "Merge branch '33' into 'master'",
            "created_at": "2016-12-20T02:14:01+0000",
            "short_id": "44500ce3"
        }
    },
    "id": 2761,
    "project_id": 85436,
    "tag": "01api",
    "title": "01api",
    "type": "release",
    "attachments": null,
    "description": null,
    "created_at": "2016-12-20T03:21:31+0000",
    "updated_at": null,
    "author": null
}
```
### 新增一个releases

在项目版本库中新增一个releases


```
POST /api/v3/projects/:id/releases
```


**参数**



| 参数            | 类型                 | 描述                                           |
| ------------- | ------------------ | -------------------------------------------- |
| id            | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH      |
| tag           | string             | tag 名                                        |
| start_point   | string             | 从 commit hash、存在的 branch 或 tag 创建 releases   |
| title         | string（可选）         | 对releases的注释                                 |
| type          | string（可选）         | 默认release                                    |
| description   | string（可选）         | 对release的描述                                  |


**返回值**


```json
{
    "ref": {
        "name": "01api",
        "message": "01api",
        "commit": {
            "id": "44500ce35ff269d36925aab3a53dac296a91b19f",
            "message": "Merge branch 'dev' into 'master'See merge request !3",
            "parent_ids": [
                "6145857bfd5858b3ed1a5bff4dc4cfc4b62e29ff",
                "3dad4be822a900f5b8ab7a0112da0be5ed76deba"
            ],
            "authored_date": "2016-12-20T02:14:01+0000",
            "author_name": "bili",
            "author_email": "bili@tencent.com",
            "committed_date": "2016-12-20T02:14:01+0000",
            "committer_name": "bili",
            "committer_email": "bili@tencent.com",
            "title": "Merge branch '33' into 'master'",
            "created_at": "2016-12-20T02:14:01+0000",
            "short_id": "44500ce3"
        }
    },
    "id": 2761,
    "project_id": 85436,
    "tag": "01api",
    "title": "01api",
    "type": "release",
    "attachments": null,
    "description": null,
    "created_at": "2016-12-20T03:21:31+0000",
    "updated_at": null,
    "author": null
}
```
### 更新一个releases (TODO)

在项目版本库中更新指定的releases


```
PUT /api/v3/projects/:id/releases/{release_id}
```


**参数**



| 参数            | 类型                 | 描述                                        |
| ------------- | ------------------ | ----------------------------------------- |
| id            | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| release_id    | Integer            | release的id                                |
| description   | string（可选）         | 对release的描述                               |


**返回值**


```json
{
    "ref": {
        "name": "01api",
        "message": "01api",
        "commit": {
            "id": "44500ce35ff269d36925aab3a53dac296a91b19f",
            "message": "Merge branch 'dev' into 'master'See merge request !3",
            "parent_ids": [
                "6145857bfd5858b3ed1a5bff4dc4cfc4b62e29ff",
                "3dad4be822a900f5b8ab7a0112da0be5ed76deba"
            ],
            "authored_date": "2016-12-20T02:14:01+0000",
            "author_name": "bili",
            "author_email": "bili@tencent.com",
            "committed_date": "2016-12-20T02:14:01+0000",
            "committer_name": "bili",
            "committer_email": "bili@tencent.com",
            "title": "Merge branch '33' into 'master'",
            "created_at": "2016-12-20T02:14:01+0000",
            "short_id": "44500ce3"
        }
    },
    "id": 2761,
    "project_id": 85436,
    "tag": "01api",
    "title": "01api",
    "type": "release",
    "attachments": null,
    "description": fix bug,
    "created_at": "2016-12-20T03:21:31+0000",
    "updated_at": null,
    "author": null
}
```
### 删除releases

删除项目版本库releases，需要有项目的 admin 权限


```
DELETE /api/v3/projects/:id/releases/{release_id}
```


**参数**



| 参数           | 类型                 | 描述                                        |
| ------------ | ------------------ | ----------------------------------------- |
| id           | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| release_id   | integer            | release的id                                |


**返回值**


### 200或相关状态码
