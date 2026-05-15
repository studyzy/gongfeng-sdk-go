# 使用前必读

## 如何调用

每个API均符合标准RESTFUL规范，每个请求的路径均包含 `/api/v3`。比如:

```
GET http://git.code.tencent.com/api/v3/projects?private_token=ddwnTYP0bBSWaDPD7_NY
```

在接口文档说明中，会注明应当使用那种HTTP方法，可以是 `GET`, `POST`, `PUT`, `DELETE`。

- 如果接口用于**获取数据**，则成功时返回 `200 OK`，返回值为 JSON 格式
- 如果接口用于**插入或创建数据**，成功时返回 `201 Created`，并会把新创建的项目用 JSON 格式返回
- 如果接口用于**修改数据**，则成功时返回 `200 OK`，同时修改了的数据会以 JSON 格式返回
- 如果接口用于**删除数据**，则成功时返回 `200 OK`

## 返回值

| 返回值 | 涵义 |
| --- | --- |
| 200 OK | 操作成功 |
| 201 Created | 创建成功 |
| 400 Bad Request | 参数错误，或是参数格式错误 |
| 401 Unauthorized | 认证失败 |
| 403 Forbidden | 账号并没有该操作的权限 |
| 404 Not Found | 资源不存在，也可能是账号没有该项目的权限（为防止黑客撞库获取库列表） |
| 405 Method Not Allowed | 没有该接口 |
| 409 Conflict | 已经有一个同名项目存在了 |
| 422 Unprocessable | 操作不能进行 |
| 500 Server Error | 服务器出错 |

## 认证方式

每个API调用都需要认证，支持 `Private_Key` 方式。如果认证失败（比如认证码错误或者过期），会返回401错误，例如：

```json
{
  "message": "401 Unauthorized"
}
```

### Private_Key 方式

账号的 Private_key 在 `http://git.code.tencent.com/profile/account` 设置。

**使用URL参数调用**

使用URL请求参数 `private_token`，例如：

```
GET http://git.code.tencent.com/api/v3/projects?private_token=ddwnTYP0bBSWaDPD7_NY
```

**使用HEADER调用**

使用 HTTP_HEADER 的 `PRIVATE-TOKEN`，例如：

```bash
curl --header "PRIVATE-TOKEN: ddwnTYP0bBSWaDPD7_NY" "http://git.code.tencent.com/api/v3/projects"
```

## 分页

所有获取列表的接口均支持分页，分页参数：

| 参数 | 含义 |
| --- | --- |
| page | 页号（从1开始），默认值为1 |
| per_page | 每页的项目数 (default: 20, max: 100) |

例如要获取项目中所有的Commit，您需要按照页号循环，逐页获取。

用户可以在Header中获取以下参数对所查询的结果进行分页：

| Header | 描述 |
| --- | --- |
| X-Total | 项目总数 |
| X-Total-Pages | 总页数 |
| X-Per-Page | 每页的项目数 |
| X-Page | 当前页面的索引（从1开始） |
| X-Next-Page | 下一页的索引 |
| X-Prev-Page | 上一页的索引 |

## iid 与 id 的区别

`iid` 与 `id` 的含义完全不同：

- **iid**：在项目中唯一的，主要用于分页或者生成网页URL。它的值会对应网页中的URL，例如项目中的第10号Issue。
- **id**：在系统中全局唯一的。
