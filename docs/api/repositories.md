# 版本库与文件操作

版本库

如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

### 下载存档的项目版本库

下载项目版本库的存档


```
GET /api/v3/projects/:id/repository/archive
```


**参数**



| 参数    | 类型                 | 描述                                        |
| ----- | ------------------ | ----------------------------------------- |
| id    | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| sha   | string（可选）         | commit hash值、分支名或 tag                     |


**返回值**


### zip下载流

获取原始文件内容

### 通过提交的hash和路径获取文件的原始内容


```
GET /api/v3/projects/:id/repository/blobs/:sha
```


```
GET /api/v3/projects/:id/repository/commits/:sha/blob
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| sha        | string             | commit hash值、分支名或 tag                     |
| filepath   | string             | 文件路径                                      |


**返回值**


返回原始内容，示例：

if (rowBounds[0] == 0 && rowBounds[1] == 2147483647) {
    this.pageSizeZero = true;
    this.pageSize = 0;
} else {
    this.pageSize = rowBounds[1];
    this.pageNum = rowBounds[1] != 0 ? (int)Math.ceil(((double)rowBounds[0] + (double)rowBounds[1]) / (double)rowBounds[1]) : 0;
### }

获取 blob 原始内容

### 通过文件 blob sha获取文件的原始内容


```
GET /api/v3/projects/:id/repository/tags
```


**参数**



| 参数    | 类型                 | 描述                                        |
| ----- | ------------------ | ----------------------------------------- |
| id    | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| sha   | string             | commit hash值、分支名或 tag                     |


**返回值**


返回原始内容，示例：

if (rowBounds[0] == 0 && rowBounds[1] == 2147483647) {
    this.pageSizeZero = true;
    this.pageSize = 0;
} else {
    this.pageSize = rowBounds[1];
    this.pageNum = rowBounds[1] != 0 ? (int)Math.ceil(((double)rowBounds[0] + (double)rowBounds[1]) / (double)rowBounds[1]) : 0;
### }

获取差异内容

### 通过对比分支，提交或者tags获取差异内容


```
GET /api/v3/projects/:id/repository/compare
```


**参数**



| 参数     | 类型                 | 描述                                        |
| ------ | ------------------ | ----------------------------------------- |
| id     | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| from   | string             | 提交的 hash、分支 或 tag                         |
| to     | string             | 提交的 hash、分支 或 tag                         |


**返回值**


```json
{
  "commit": {
    "id": "1480a4610ca01dd10cb5bc30d8a1b1eab98c09a1",
    "short_id": "455da4dc"
    "title": "xxx",
    "author_name": "git_user1",
    "author_email": "git_user1@tencent.com",
    "created_at": "2015-03-13T09:51:06+0000",
    "message": "xxx"
  },
  "commits": [
    {
      "id": "1480a4610ca01dd10cb5bc30d8a1b1eab98c09a1",
      "short_id": "6ed307f8"
      "title": "xxx",
      "author_name": "git_user1",
      "author_email": "git_user1@tencent.com",
      "created_at": "2015-03-13T09:51:06+0000",
      "message": "xxx"
    }
  ],
  "diffs": [
    {
      "old_path": "/dev/null",
      "new_path": "xxxxxxxx",
      "a_mode": 0,
      "b_mode": 33188,
      "diff": "@@ -0,0 +1 @@\
+xxx\
\\\\ No newline at end of file\
",
      "new_file": true,
      "renamed_file": false,
      "deleted_file": false,
      "is_too_large": false,
      "is_collapse": false,
      "additions": 1,
      "deletions": 2
    }
  ],
  "compare_timeout": false,
  "compare_same_ref": false,
  "over_flow": false,
  "files_total": 5,
  "commits_total": 18
}
```

### 下载Compare差异文件集

下载通过对比分支，提交或者tags获取差异内容文件集


```
GET /api/v3/projects/:id/repository/compare/changed_files
```


**参数：**



| 参数         | 类型                  | 描述                                        |
| ---------- | ------------------- | ----------------------------------------- |
| id         | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| from       | string              | 提交的 hash、分支 或 tag                         |
| to         | string              | 提交的 hash、分支 或 tag                         |
| straight   | boolean             | true：两个点比较差异，false：三个点比较差异。默认是 false      |


返回值：文件流

### 获取版本库文件和目录列表

返回项目版本库文件和目录列表


```
GET /api/v3/projects/:id/repository/tree
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| ref_name   | string（可选）         | 提交的 hash、分支 或 tag，默认：默认分支                 |
| path       | string（可选）         | 文件路径                                      |


**返回值**


```json
[
  {
    "id": "23c22934e8b4f901dd264cdd100d2f7c339014f5",
    "name": "GroupMilestoneStateStatistics.java",
    "type": "blob",
    "mode": "100644"
  },
  {
    "id": "30a6e86afd0d641b8e9715a6eab5165a71e56e80",
    "name": "GroupMilestoneStatistics.java",
    "type": "blob",
    "mode": "100644"
  },
  {
    "id": "69756319f6abcc246c11c932a3c07c923800544c",
    "name": "GroupStatisticsById.java",
    "type": "blob",
    "mode": "100644"
  },
  {
    "id": "28bdc11874295c5722276a63502010acb92cd9b6",
    "name": "GroupTypeStatistics.java",
    "type": "blob",
    "mode": "100644"
  },
  {
    "id": "b7899b52f700434959e228532171eebdc8a91bd5",
    "name": "TGitErrorStatistics.java",
    "type": "blob",
    "mode": "100644"
  }
]
```

获取单个文件内容和信息

取得一个特定文件的文件内容和一些文件信息，如名称、大小、所在提交等，文件内容content用Base64编码，解析时请注意解码


```
GET /api/v3/projects/:id/repository/files
```


**参数**



| 参数          | 类型                 | 描述                                        |
| ----------- | ------------------ | ----------------------------------------- |
| id          | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| ref         | string             | 描述信息、分支名或 tag，默认：默认分支                     |
| file_path   | string             | 文件路径（文件名）                                 |


**返回值**


```json
{
  "file_name": "foo.h",
  "file_path": "src/controller/",
  "size": 15,
  "ref": "master",
  "blob_id": "37c36524713aa8083f787066a9ed0c0d2f82bbb4",
  "commit_id": "b5e3f65af2fd6d2895414a679290cad7664217b3",
  "content": "I2lmbmRlZiBXT1JLVFJFRV9ICiNkZWZpbmUgV09SS1RSRUVfSAoKI2luY2x1ZGUgInJlZnMuaCIKCnN0cnVjdCBzdHJidWY7CgpzdHJ1Y3Qgd29ya3RyZWUgewoJY2hhciAqcGF0aDsKCWNoYXIgKmlkOwoJY2hhciAqaGVhZF9yZWY7CQkvKiBOVUxMIGlmIEhFQUQgaXMgYnJva2VuIG9yIGRldGFjaGVkICovCgljaGFyICpsb2NrX3JlYXNvbjsJLyogaW50ZXJuYWwgdXNlICovCglzdHJ1Y3Qgb2JqZWN0X2lkIGhlYWRfb2lkOwoJaW50IGlzX2RldGFjaGVkOwoJaW50IGlzX2JhcmU7CglpbnQgaXNfY3VycmVudDsKCWludCBsb2NrX3JlYXNvbl92YWxpZDsKfTs=",
  "encoding": "base64"
}
```

新增文件

### 允许您向项目版本库创建文件


```
POST /api/v3/projects/:id/repository/files
```


**参数**



| 参数               | 类型                 | 描述                                        |
| ---------------- | ------------------ | ----------------------------------------- |
| id               | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| file_path        | string             | 文件路径(文件名)                                 |
| branch_name      | string             | 分支 名                                      |
| encoding         | string（可选）         | 内容编码，可选：text、base64，默认：text               |
| content          | string             | 内容                                        |
| commit_message   | string             | 描述信息                                      |


**返回值**


```json
{
  "file_path": "controller",
  "file_name": "controller",
  "branch_name": "master"
}
```

### 删除文件

删除项目版本库文件


```
DELETE /api/v3/projects/:id/repository/files
```


**参数**



| 参数               | 类型                 | 描述                                        |
| ---------------- | ------------------ | ----------------------------------------- |
| id               | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| file_path        | string             | 文件路径（文件名）                                 |
| branch_name      | string             | 分支名                                       |
| commit_message   | string             | 描述                                        |


**返回值**


```json
{
  "file_path": "controller",
  "file_name": "controller",
  "branch_name": "master"
}
```

### 编辑文件

更新项目版本库文件


```
PUT /api/v3/projects/:id/repository/files
```


**参数**



| 参数               | 类型                 | 描述                                        |
| ---------------- | ------------------ | ----------------------------------------- |
| id               | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| file_path        | string             | 文件路径（文件名）                                 |
| branch_name      | string             | 分支名                                       |
| encoding         | string（可选）         | 内容编码，可选：text、base64，默认：text               |
| content          | string             | 内容                                        |
| commit_message   | string             | 注释                                        |


**返回值**


```json
{
  "file_path": "controller",
  "file_name": "controller",
  "branch_name": "master"
}
```
