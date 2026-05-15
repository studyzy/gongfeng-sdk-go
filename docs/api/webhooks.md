# 项目回调钩子


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子：/api/v3/projects/diaspora%2Fdiaspora (/ 编码 %2F)

### 给项目增加回调钩子

增加项目回调钩子


```
POST /api/v3/projects/:id/hooks
```


**参数**



| 参数                      | 类型                  | 描述                                        |
| ----------------------- | ------------------- | ----------------------------------------- |
| id                      | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| url                     | string              | 回调钩子 地址                                   |
| push_events             | boolean             | 有推送事件触发回调钩子，默认 true                       |
| issues_events           | boolean             | 有缺陷事件触发回调钩子，默认 false                      |
| merge_requests_events   | boolean             | 有合并请求事件触发回调钩子，默认 false                    |
| tag_push_events         | boolean             | 有 Tag 推送事件触发回调钩子，默认 false                 |
| note_events             | boolean             | 有评论事件触发回调钩子，默认 false                      |
| review_events           | boolean             | 有评审事件触发回调钩子，默认 false                      |
| token                   | string(可选)          | 用以校验收到的负载；此token不会在包含在返回值中                |


**返回值**


```json
{
    "id": 1,
    "url": "https://git.code.tencent.com/hook/push",
    "created_at": "2015-03-29T04:48:31+0000",
    "project_id": 2365,
    "push_events": true,
    "issues_events": false,
    "merge_requests_events": false,
    "tag_push_events": false,
    "note_events": false,
    "review_events": false
}
```

### 编辑项目回调钩子

修改项目中的回调钩子


```
PUT /api/v3/projects/:id/hooks/:hook_id
```


**参数**



| 参数                      | 类型                  | 描述                                        |
| ----------------------- | ------------------- | ----------------------------------------- |
| id                      | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| url                     | string              | 回调钩子 地址                                   |
| push_events             | boolean（可选）         | 有推送事件触发回调钩子                               |
| issues_events           | boolean（可选）         | 有缺陷事件触发回调钩子                               |
| merge_requests_events   | boolean（可选）         | 有合并请求事件触发回调钩子                             |
| tag_push_events         | boolean（可选）         | 有 Tag 推送事件触发回调钩子                          |
| note_events             | boolean（可选）         | 有评论事件触发回调钩子                               |
| review_events           | boolean（可选）         | 有评审事件触发回调钩子                               |
| token                   | string(可选)          | 用以校验收到的负载；此token不会在包含在返回值中                |


**返回值**


```json
{
    "id": 1,
    "url": "https://git.code.tencent.com/hook/push",
    "created_at": "2015-03-29T04:48:31+0000",
    "project_id": 2365,
    "push_events": false,
    "issues_events": false,
    "merge_requests_events": false,
    "tag_push_events": false,
    "note_events": false,
    "review_events": false
}
```

查询项目单个回调钩子

### 根据回调钩子 id 查询项目回调钩子


```
GET /api/v3/projects/:id/hooks/:hook_id
```


**参数**



| 参数        | 类型                  | 描述                                        |
| --------- | ------------------- | ----------------------------------------- |
| id        | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| hook_id   | integer             | 回调钩子 id                                   |


**返回值**


```json
{
    "id": 2,
    "url": "https://git.code.tencent.com/hook/push",
    "created_at": "2015-03-29T04:48:31+0000",
    "project_id": 2,
    "push_events": true,
    "issues_events": false,
    "merge_requests_events": false,
    "tag_push_events": false,
    "note_events": false,
    "review_events": false
}
```

### 查询项目回调钩子列表

在项目中查询项目回调钩子列表


```
GET /api/v3/projects/:id/hooks
```


**参数**



| 参数         | 类型                  | 描述                                        |
| ---------- | ------------------- | ----------------------------------------- |
| id         | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| page       | integer             | 分页（默认值：1）                                 |
| per_page   | integer             | 默认页面大小（默认值： 20，最大值： 100）                  |


**返回值**


```json
[
    {
    "id": 1,
    "url": "https://git.code.tencent.com/hook/push",
    "created_at": "2015-03-29T04:48:31+0000",
    "project_id": 2,
    "push_events": false,
    "issues_events": false,
    "merge_requests_events": false,
    "tag_push_events": false,
    "note_events": false,
    "review_events": false
},
    {
    "id": 1,
    "url": "https://git.code.tencent.com/hook/mr",
    "created_at": "2015-03-15T04:48:31+0000",
    "project_id": 2,
    "push_events": false,
    "issues_events": false,
    "merge_requests_events": false,
    "tag_push_events": false,
    "note_events": false,
    "review_events": false
}
]
```

### 删除项目回调钩子

在项目中删除某个指定的回调钩子


```
DELETE /api/v3/projects/:id/hooks/:hook_id
```


**参数**



| 参数        | 类型                  | 描述                                        |
| --------- | ------------------- | ----------------------------------------- |
| id        | integer or string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| hook_id   | integer             | 回调钩子 id                                   |


**返回值**


### 返回状态码
