[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/sortx/release.yml?branch=main&label=BUILD)](https://github.com/yylego/sortx/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/sortx)](https://pkg.go.dev/github.com/yylego/sortx)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/sortx/main.svg)](https://coveralls.io/github/yylego/sortx?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yylego/sortx.svg)](https://github.com/yylego/sortx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/sortx)](https://goreportcard.com/report/github.com/yylego/sortx)

# sortx

`sortx` is a Go package that provides a simple and flexible approach to sort slices using custom comparison functions. It leverages Go's generics and the `sort.Interface` to avoid repeating the implementation of sorting logic for different types.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Installation

To install the `sortx` package, you can use the following command:

```bash
go get github.com/yylego/sortx
```

## Usage

The package offers multiple functions for sorting slices with different comparison strategies. Below are the main functions available:

### `SortByIndex`

Sorts the slice `a` using an index-based comparison function `iLess`.

```go
sortx.SortByIndex(a []V, iLess func(i, j int) bool)
```

- `a`: The slice to be sorted.
- `iLess`: The function that compares the indices of two elements in the slice.
- Sorts the slice in place using the provided index-based comparison function.

### `SortByValue`

Sorts the slice `a` using a value-based comparison function `vLess`.

```go
sortx.SortByValue(a []V, vLess func(a, b V) bool)
```

- `a`: The slice to be sorted.
- `vLess`: The function that compares the values of two elements in the slice.
- Sorts the slice in place using the provided value-based comparison function.

### `SortIStable`

Sorts the slice `a` using an index-based comparison function `iLess` and preserves the sequence of same elements (stable sort).

```go
sortx.SortIStable(a []V, iLess func(i, j int) bool)
```

- `a`: The slice to be sorted.
- `iLess`: The function that compares the indices of two elements in the slice.
- Sorts the slice in place while maintaining the input sequence of same elements (stable sort).

### `SortVStable`

Sorts the slice `a` using a value-based comparison function `vLess` and preserves the sequence of same elements (stable sort).

```go
sortx.SortVStable(a []V, vLess func(a, b V) bool)
```

- `a`: The slice to be sorted.
- `vLess`: The function that compares the values of two elements in the slice.
- Sorts the slice in place while maintaining the input sequence of same elements (stable sort).

## Example

### Basic Sorting

Sort integers and strings using index-based and value-based comparison:

```go
package main

import (
	"fmt"

	"github.com/yylego/sortx"
)

func main() {
	// Sort integers using index comparison
	numbers := []int{5, 3, 8, 1, 4}
	sortx.SortByIndex(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println("Sorted using index:", numbers) // Output: [1 3 4 5 8]

	// Sort strings using value comparison
	fruits := []string{"apple", "banana", "orange", "date"}
	sortx.SortByValue(fruits, func(a, b string) bool {
		return a < b
	})
	fmt.Println("Sorted using value:", fruits) // Output: [apple banana date orange]
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Stable Sorting

Stable sort preserves input sequence of same elements:

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
	// Stable sort preserves input sequence of same elements
	tasks := []Task{
		{"Task-A", 2},
		{"Task-B", 1},
		{"Task-C", 2},
		{"Task-D", 1},
	}

	// Sort using Rank, same Rank keeps input sequence
	sortx.SortVStable(tasks, func(a, b Task) bool {
		return a.Rank < b.Rank
	})

	fmt.Println("Stable sorted tasks:")
	for _, task := range tasks {
		fmt.Printf("  %s (Rank: %d)\n", task.Name, task.Rank)
	}
	// Output:
	//   Task-B (Rank: 1)
	//   Task-D (Rank: 1)
	//   Task-A (Rank: 2)
	//   Task-C (Rank: 2)
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

## Examples

### Custom Struct Sorting

**Sort using struct field:**
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

**Sort with multiple conditions:**
```go
sortx.SortByValue(people, func(a, b Person) bool {
	if a.Age != b.Age {
		return a.Age < b.Age
	}
	return a.Name < b.Name
})
```

### Descending Sort

**Sort in descending sequence:**
```go
numbers := []int{1, 5, 3, 9, 2}
sortx.SortByValue(numbers, func(a, b int) bool {
	return a > b // Use > for descending
})
```

### Using sort.Interface

**Create sort.Interface for advanced usage:**
```go
data := []int{5, 2, 8, 1}
sortable := sortx.NewSortByValue(data, func(a, b int) bool {
	return a < b
})
sort.Sort(sortable) // Use standard sort package
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2026-03-11 12:00:00.000000 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/sortx.svg?variant=adaptive)](https://starchart.cc/yylego/sortx)
