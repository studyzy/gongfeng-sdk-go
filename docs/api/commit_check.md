# 提交检测


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子：/api/v3/projects/diaspora%2Fdiaspora (/ 编码 %2F)

### 新建检测结果

检测结果创建


```
POST /api/v3/projects/:id/commit/:sha/statuses
```


**参数**



| 参数            | 类型                  | 描述                                          |
| ------------- | ------------------- | ------------------------------------------- |
| id            | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH     |
| sha           | string              | commit hash值、分支名或 tag                       |
| state         | string              | 检测状态，可选值：pending, success, error, failure   |
| target_url    | string（可选）          | 检测路径，最大字节长度：255                             |
| description   | string（可选）          | 检测结果描述，最大字节长度：255                           |
| context       | string（可选）          | 区别于其他检测系统的标签，默认：default                     |
| block         | boolean（可选）         | 是否锁住提交和合并请求，默认：false                        |


**返回值**



```json
{
    "id": 365,
    "state": "success",
    "target_url": "https://git.code.tencent.com/mr/check",
    "description": "check success",
    "context": "jenkins/mr",
    "created_at": "2014-03-05T07:56:56+0000",
    "updated_at": "2014-03-05T07:56:56+0000",
    "block": true
}
```

### 通过 Ref 查询提交检测的组合结果

通过项目中的 Ref 查询提交检测的组合结果

注意：用户必须具有拉取版本库代码的权限


```
GET /api/v3/projects/:id/commits/:ref/status
```


**参数**



| 参数         | 类型                  | 描述                                        |
| ---------- | ------------------- | ----------------------------------------- |
| id         | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| ref        | string              | ref可以是 commit hash值、分支名或 tag              |
| page       | integer             | 分页（默认：1）                                  |
| per_page   | integer             | 默认页面大小（默认： 20，最大： 100）                    |


**返回值**


```json
{
    "state": "pending",
    "block": false,
    "statuses": [
        {
            "id": 21,
            "state": "pending",
            "target_url": "https://git.code.tencent.com/push/check",
            "description": "check pending",
            "context": "jenkins/push",
            "created_at": "2015-02-10T15:02:38+0000",
            "updated_at": "2015-02-10T15:02:38+0000",
            "block": true
        }
    ],
    "sha": "b5e3f65af2fd6d2895414a679290cad7664217b3",
    "total_count": 2
}
```

### 通过 Ref 查询提交检测结果

通过项目中的 Ref 查询提交检测的结果，最终的提交检测结果将会以时间的倒序排序返回给用户

注意：用户必须具有拉取版本库代码的权限


```
GET /api/v3/projects/:id/commits/:ref/statuses
```


**参数**



| 参数         | 类型                  | 描述                                        |
| ---------- | ------------------- | ----------------------------------------- |
| id         | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| ref        | string              | ref可以是 commit hash值、分支名或 tag              |
| page       | integer             | 分页（默认：1）                                  |
| per_page   | integer             | 默认页面大小（默认： 20，最大： 100）                    |


**返回值**


```json
[
    {
        "id": 21,
        "state": "success",
        "target_url": "https://git.code.tencent.com/mr/check",
        "description": "check success",
        "context": "jenkins/mr",
        "created_at": "2015-02-05T05:32:48+0000",
        "updated_at": "2015-02-05T05:32:48+0000",
        "block": true
    },
    {
        "id": 15,
        "state": "pending",
        "target_url": "https://git.code.tencent.com/push/check",
        "description": "check pending",
        "context": "jenkins/push",
        "created_at": "2015-02-05T05:32:48+0000",
        "updated_at": "2015-02-05T05:32:48+0000",
        "block": false
    }
]
```
