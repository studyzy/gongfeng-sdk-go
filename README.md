# gongfeng-sdk-go

腾讯工蜂 https://code.tencent.com/help/api/prepare 封装的 SDK。

## 特性

- 基于 `go-gitlab`，覆盖工蜂兼容的全部 GitLab REST API
- 提供统一 `Call` 方法，支持直接访问任意 REST endpoint
- 支持自定义 BaseURL、API 版本、HTTP 客户端

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
"net/http"

gongfeng "github.com/studyzy/gongfeng-sdk-go"
)

func main() {
client, err := gongfeng.NewClient("your-private-token", nil)
if err != nil {
panic(err)
}

var result map[string]any
_, err = client.Call(context.Background(), http.MethodGet, "/projects", nil, &result)
if err != nil {
panic(err)
}

fmt.Println(result)
}
```

## 自定义配置

```go
client, err := gongfeng.NewClient("token", &gongfeng.Options{
BaseURL:    "https://code.tencent.com/",
APIVersion: "v3",
})
```
