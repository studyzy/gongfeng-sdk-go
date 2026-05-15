# gongfeng-sdk-go

腾讯工蜂 REST API 的 Go SDK，完整覆盖工蜂 v3 API。

## 特性

- 基于标准库自建 HTTP Client，无重型第三方依赖
- 按资源拆分 Service，每个 API 资源独立文件
- 所有方法支持 `context.Context`
- 统一的分页、认证、错误处理
- 函数式选项模式（`WithBaseURL`、`WithHTTPClient`）

## 安装

```bash
go get github.com/studyzy/gongfeng-sdk-go
```

## 快速开始

```go
package main

import (
	"context"
	"fmt"
	"log"

	gongfeng "github.com/studyzy/gongfeng-sdk-go"
)

func main() {
	client, err := gongfeng.NewClient("your-private-token")
	if err != nil {
		log.Fatal(err)
	}

	// 获取项目列表
	projects, _, err := client.Projects.ListProjects(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range projects {
		fmt.Printf("%d: %s\n", p.ID, p.PathWithNamespace)
	}
}
```

## 自定义配置

```go
// 连接私有部署的工蜂实例
client, err := gongfeng.NewClient("token",
	gongfeng.WithBaseURL("https://git.example.com/"),
)
```

## 支持的 API

| Service | 说明 |
| --- | --- |
| `Projects` | 项目 CRUD、搜索、成员管理 |
| `Groups` | 项目组 CRUD、成员管理 |
| `Branches` | 分支创建、删除、保护分支 |
| `Tags` | Tag 列表、创建、删除 |
| `MergeRequests` | MR 创建、合并、评论、变更查询 |
| `Commits` | 提交查询、Diff、评论、Ref |
| `Repositories` | 存档下载、文件树、文件 CRUD、分支比较 |
| `Issues` | 缺陷 CRUD |
| `Notes` | MR/Issue 评论 CRUD |
| `Labels` | 标签 CRUD |
| `Milestones` | 里程碑 CRUD |
| `Users` | 用户查询 |
| `Namespaces` | 命名空间列表与搜索 |
| `Releases` | 版本发布 CRUD |
| `Webhooks` | 回调钩子 CRUD |
| `Forks` | Fork 项目 |
| `Watchers` | 项目关注/取消关注 |
| `Session` | 获取私有令牌 |
| `CommitStatuses` | 提交检测（CI 状态） |
| `Reviews` | MR 评审 + Commit 评审 |

## API 文档

详细的工蜂 REST API 文档见 [docs/api/](docs/api/README.md)。
