# 标签


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

### 新增标签

创建一个标签给指定项目，每个项目标签最多100个


```
POST /api/v3/projects/:id/labels
```


**参数**



| 参数      | 类型                 | 描述                                        |
| ------- | ------------------ | ----------------------------------------- |
| id      | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| name    | string             | 标签名                                       |
| color   | string             | 标签颜色，举例：#428bca                           |


**返回值**


```json
{
  "color": "#428bca",
  "name": "界面"
}
```

### 修改标签

修改指定项目的某个标签


```
PUT /api/v3/projects/:id/labels
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| name       | string             | 旧标签名                                      |
| new_name   | string             | 新标签名                                      |
| color      | string             | 标签颜色，举例：#428bca                           |


**返回值**


```json
{
  "color": "#428bca",
  "name": "新界面"
}
```

### 获取标签列表

返回给定项目的所有标签


```
GET /api/v3/projects/:id/labels
```


**参数**



| 参数         | 类型                 | 描述                                        |
| ---------- | ------------------ | ----------------------------------------- |
| id         | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| order_by   | string（可选）         | 排序字段，允许按 name,created_at排序（默认name）        |
| sort       | string（可选）         | 排序方式，允许 asc,desc（默认asc）                   |
| page       | integer（可选）        | 分页（默认值：1）                                 |
| per_page   | integer（可选）        | 默认页面大小（默认值：20，最大值：100）                    |


**返回值**


```json
[
  {
    "color": "#d9534f",
    "name": "bug"
  },
  {
    "color": "#d9534f",
    "name": "critical"
  },
  {
    "color": "#d9534f",
    "name": "confirmed"
  }
]
```

### 删除标签

删除指定项目某个标签


```
DELETE /api/v3/projects/:id/labels
```


**参数**



| 参数     | 类型                 | 描述                                        |
| ------ | ------------------ | ----------------------------------------- |
| id     | integer 或 string   | id = 项目唯一标识或NAMESPACE_PATH/PROJECT_PATH   |
| name   | string             | 标签名                                       |


**返回值**


### 删除成功返回 200，参数错误返回 400，标签不存在返回 404
