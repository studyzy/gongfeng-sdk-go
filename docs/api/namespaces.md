# 命名空间


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

### 获取命名空间列表

获取一个命名空间的列表

注意：作为用户显示的是我的命名空间, 作为管理者显示所有的命名空间。


```
GET /api/v3/namespaces
```


**参数:**



| 参数         | 类型             | 描述                         |
| ---------- | -------------- | -------------------------- |
| page       | integer （可选）   | 分页 (默认值:1)                 |
| per_page   | integer （可选）   | 默认页面大小 (默认值 20，最大值： 100)   |


**返回值：**


```json
[
  {
    "id": 5,
    "path": "git_user1",
    "kind": "user"
  },
  {
    "id": 35,
    "path": "tengzz",
    "kind": "group"
  }
]
```


### 你能够通过用户名称和路径查询命名空间，如下。

搜索命名空间

### 查询所有的匹配你的用户名称或者路径的命名空间


```
GET /api/v3/namespaces?search=luci
```


**返回值：**


```json
[
  {
    "id": 5,
    "path": "luciyu",
    "kind": "user"
  }
]
```
