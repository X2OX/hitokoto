# [Hitokoto](https://github.com/X2OX/hitokoto)

[![JetBrains Open Source Licenses](https://img.shields.io/badge/-JetBrains%20Open%20Source%20License-000?style=flat-square&logo=JetBrains&logoColor=fff&labelColor=000)](https://www.jetbrains.com/?from=blackdatura)
[![GoDoc](https://pkg.go.dev/badge/pkg.go.dev/go.x2ox.com/hitokoto)](https://pkg.go.dev/go.x2ox.com/hitokoto)
[![Sourcegraph](https://sourcegraph.com/github.com/X2OX/hitokoto/-/badge.svg)](https://sourcegraph.com/github.com/X2OX/hitokoto?badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/X2OX/hitokoto)](https://goreportcard.com/report/github.com/X2OX/hitokoto)
[![Release](https://img.shields.io/github/v/release/X2OX/hitokoto.svg)](https://github.com/X2OX/hitokoto/releases)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

总有那么些句子，会在不经意间突然击中你的心。

## 项目说明

一言的 API 并不算少，但是拿着前几年的 API 进行调用测试时，发现就只有金山词霸的源还能使用。 而且不能自定义数据算是之前的使用中一直会有的遗憾。收集了一些公开的数据，然后进行了去重等等处理。

对于很多功能并不支持，但是可能未来会支持。 例如：
- 唯一 ID
- 来源
- 分类标签（动漫、游戏、电视剧等）
- 返回 JavaScript 调用函数
- 批量返回

有一些功能可能不会支持。例如：
- 返回 GBK 及其他非 UTF-8 编码的内容
- 限定句子长度

很多功能都是比较方便实现的，例如 ID，返回函数等等，只要有人有需求就可以加上。但是对于一些开倒车的功能是不会加上的，例如编码问题。
而限定句子长度， 对于 Serverless 来说，对数组进行扫描有浪费资源的嫌疑。
分类标签，可能需要考虑清楚数据结构后进行处理。

## 使用方式
- JSON `https://hitokoto.x2ox.com/api/hitokoto?type=json`
- 字符串 `https://hitokoto.x2ox.com/api/hitokoto`

## 自定义部署

那么最简单的方式是一键部署，手动部署和一键部署的区别不大。
如果需要自定义词库，自行修改词库内容即可。
如果需要自定义词库，且需要集成部署，则需要稍微改变一下。


### 一键部署
[![](https://vercel.com/button)](https://vercel.com/new/project?template=https://github.com/X2OX/hitokoto&utm_source=x2ox&utm_campaign=oss)

### 手动部署
1. Fork Repo
2. 在 Vercel 中导入部署

### 集成部署
```go
package hitokoto

import (
	"net/http"

	"go.x2ox.com/hitokoto/api/hitokoto"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	hitokoto.Handler(w, r)
}
```

### 更新

一键部署及手动部署，同步上游即可更新。 自定义部署更新 Go Mod 内的版本即可。


[![Powered by Vercel](https://hitokoto.x2ox.com/powered-by-vercel.svg)](https://vercel.com?utm_source=x2ox&utm_campaign=oss)

