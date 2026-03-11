[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/sortx/release.yml?branch=main&label=BUILD)](https://github.com/yylego/sortx/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/sortx)](https://pkg.go.dev/github.com/yylego/sortx)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/sortx/main.svg)](https://coveralls.io/github/yylego/sortx?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yylego/sortx.svg)](https://github.com/yylego/sortx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/sortx)](https://goreportcard.com/report/github.com/yylego/sortx)

# sortx

`sortx` 是一个 Go 包，它提供了一种简单灵活的方式来使用自定义比较函数对切片进行排序。它利用了 Go 的泛型和 `sort.Interface`，避免了为不同类型重复实现排序逻辑。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 安装

要安装 `sortx` 包，可以使用以下命令：

```bash
go get github.com/yylego/sortx
```

## 使用

该包提供了几种排序函数，支持不同的比较策略。以下是可用的主要函数：

### `SortByIndex`

使用基于索引的比较函数 `iLess` 对切片 `a` 进行排序。

```go
sortx.SortByIndex(a []V, iLess func(i, j int) bool)
```

- `a`: 要排序的切片。
- `iLess`: 比较切片中两个元素索引的函数。
- 使用提供的基于索引的比较函数对切片进行排序。

### `SortByValue`

使用基于值的比较函数 `vLess` 对切片 `a` 进行排序。

```go
sortx.SortByValue(a []V, vLess func(a, b V) bool)
```

- `a`: 要排序的切片。
- `vLess`: 比较切片中两个元素值的函数。
- 使用提供的基于值的比较函数对切片进行排序。

### `SortIStable`

使用基于索引的比较函数 `iLess` 对切片 `a` 进行排序，并保持相等元素的原始顺序（稳定排序）。

```go
sortx.SortIStable(a []V, iLess func(i, j int) bool)
```

- `a`: 要排序的切片。
- `iLess`: 比较切片中两个元素索引的函数。
- 使用基于索引的比较函数对切片进行排序，同时保持相等元素的原始顺序（稳定排序）。

### `SortVStable`

使用基于值的比较函数 `vLess` 对切片 `a` 进行排序，并保持相等元素的原始顺序（稳定排序）。

```go
sortx.SortVStable(a []V, vLess func(a, b V) bool)
```

- `a`: 要排序的切片。
- `vLess`: 比较切片中两个元素值的函数。
- 使用基于值的比较函数对切片进行排序，同时保持相等元素的原始顺序（稳定排序）。

## 示例

### 基础排序

使用索引比较和值比较对整数和字符串进行排序：

```go
package main

import (
	"fmt"

	"github.com/yylego/sortx"
)

func main() {
	// 使用索引比较排序整数
	numbers := []int{5, 3, 8, 1, 4}
	sortx.SortByIndex(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println("按索引排序:", numbers) // 输出: [1 3 4 5 8]

	// 使用值比较排序字符串
	fruits := []string{"apple", "banana", "orange", "date"}
	sortx.SortByValue(fruits, func(a, b string) bool {
		return a < b
	})
	fmt.Println("按值排序:", fruits) // 输出: [apple banana date orange]
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 稳定排序

稳定排序保持相同元素的输入顺序：

```go
package main

import (
	"fmt"

	"github.com/yylego/sortx"
)

type Task struct {
	Name string
	Rank int
}

func main() {
	// 稳定排序保持相同元素的输入顺序
	tasks := []Task{
		{"Task-A", 2},
		{"Task-B", 1},
		{"Task-C", 2},
		{"Task-D", 1},
	}

	// 按 Rank 排序，相同 Rank 保持输入顺序
	sortx.SortVStable(tasks, func(a, b Task) bool {
		return a.Rank < b.Rank
	})

	fmt.Println("稳定排序结果:")
	for _, task := range tasks {
		fmt.Printf("  %s (Rank: %d)\n", task.Name, task.Rank)
	}
	// 输出:
	//   Task-B (Rank: 1)
	//   Task-D (Rank: 1)
	//   Task-A (Rank: 2)
	//   Task-C (Rank: 2)
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

## 更多示例

### 自定义结构体排序

**按结构体字段排序：**
```go
type Person struct {
	Name string
	Age  int
}
people := []Person{{"Alice", 30}, {"Bob", 25}}
sortx.SortByValue(people, func(a, b Person) bool {
	return a.Age < b.Age
})
```

**多条件排序：**
```go
sortx.SortByValue(people, func(a, b Person) bool {
	if a.Age != b.Age {
		return a.Age < b.Age
	}
	return a.Name < b.Name
})
```

### 降序排序

**按降序排列：**
```go
numbers := []int{1, 5, 3, 9, 2}
sortx.SortByValue(numbers, func(a, b int) bool {
	return a > b // 使用 > 实现降序
})
```

### 使用 sort.Interface

**创建 sort.Interface 实现高级用法：**
```go
data := []int{5, 2, 8, 1}
sortable := sortx.NewSortByValue(data, func(a, b int) bool {
	return a < b
})
sort.Sort(sortable) // 使用标准 sort 包
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2026-03-11 12:00:00.000000 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/sortx.svg?variant=adaptive)](https://starchart.cc/yylego/sortx)
