// Package sortx: Flexible slice sorting with custom comparison functions
// Provides sort.Interface implementation supporting both index-based and value-based comparisons
// Enables type-safe generic sorting without repeating boilerplate code
//
// sortx: 灵活的切片排序工具，支持自定义比较函数
// 提供 sort.Interface 实现，支持基于索引和基于值的比较
// 实现类型安全的泛型排序，避免重复样板代码
package sortx

import (
	"sort"

	"github.com/pkg/errors"
)

// Slice represents a sortable slice with custom comparison functions
// Supports both index-based and value-based comparison strategies
//
// Slice 表示一个可排序的切片，支持自定义比较函数
// 支持基于索引和基于值的比较策略
type Slice[V any] struct {
	slice []V                 // The slice to be sorted // 要排序的切片
	iLess func(i, j int) bool // Index comparison function // 索引比较函数
	vLess func(a, b V) bool   // Value comparison function // 值比较函数
}

// NewSortByIndex creates a sort.Interface using index-based comparison
// Panics when iLess is nil to detect invalid input at construction time
//
// NewSortByIndex 使用索引比较函数创建 sort.Interface
// 当 iLess 为 nil 时触发 panic，在构造时检测无效输入
func NewSortByIndex[V any](a []V, iLess func(i, j int) bool) sort.Interface {
	if iLess == nil {
		panic(errors.New("sortx: need iLess"))
	}
	return &Slice[V]{slice: a, iLess: iLess}
}

// NewSortByValue creates a sort.Interface using value-based comparison
// Panics when vLess is nil to detect invalid input at construction time
//
// NewSortByValue 使用值比较函数创建 sort.Interface
// 当 vLess 为 nil 时触发 panic，在构造时检测无效输入
func NewSortByValue[V any](a []V, vLess func(a, b V) bool) sort.Interface {
	if vLess == nil {
		panic(errors.New("sortx: need vLess"))
	}
	return &Slice[V]{slice: a, vLess: vLess}
}

// Len returns the slice length
//
// Len 返回切片长度
func (s *Slice[V]) Len() int {
	return len(s.slice)
}

// Less compares elements at indexes i and j
// Uses iLess when set, otherwise uses vLess
//
// Less 比较索引 i 和 j 处的元素
// 优先使用 iLess，否则使用 vLess
func (s *Slice[V]) Less(i, j int) bool {
	switch {
	// Use index-based comparison when iLess is set
	// 当设置了 iLess 时使用索引比较
	case s.iLess != nil:
		return s.iLess(i, j)
	// Use value-based comparison when vLess is set
	// 当设置了 vLess 时使用值比较
	case s.vLess != nil:
		return s.vLess(s.slice[i], s.slice[j])
	// Panic when no comparison function is set
	// 当没有设置比较函数时触发 panic
	default:
		panic(errors.New("sortx: less function not set"))
	}
}

// Swap exchanges elements at indexes i and j
//
// Swap 交换索引 i 和 j 处的元素
func (s *Slice[V]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}
