package sortx_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"github.com/yylego/sortx"
)

// TestSortByIndex tests SortByIndex function
// TestSortByIndex 测试 SortByIndex 函数
func TestSortByIndex(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	sortx.SortByIndex(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	t.Log(spew.Sdump(a)) // Use spew to dump slice state // 使用 spew 输出切片状态
	require.Equal(t, a, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

// TestSortByValue tests SortByValue function
// TestSortByValue 测试 SortByValue 函数
func TestSortByValue(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	sortx.SortByValue(a, func(a, b int) bool {
		return a < b
	})
	t.Log(spew.Sdump(a)) // Use spew to dump slice state // 使用 spew 输出切片状态
	require.Equal(t, a, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

// TestSortIStable tests stable sorting with index comparison
// TestSortIStable 测试使用索引比较的稳定排序
func TestSortIStable(t *testing.T) {
	// Stable sort keeps same elements in input sequence
	// 稳定排序保持相同元素的输入顺序
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
		{"David", 25},
	}

	// When ages match, keep input sequence
	// 当年龄相同时，保持输入顺序
	sortx.SortIStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	// Use spew to show struct details with field names
	// 使用 spew 显示结构体详情和字段名
	t.Log(spew.Sdump(people))

	expected := []Person{
		{"Bob", 25},
		{"David", 25},
		{"Alice", 30},
		{"Charlie", 30},
	}

	require.Equal(t, expected, people)
}

// TestSortVStable tests stable sorting with value comparison
// TestSortVStable 测试使用值比较的稳定排序
func TestSortVStable(t *testing.T) {
	// Stable sort keeps same elements in input sequence
	// 稳定排序保持相同元素的输入顺序
	type Item struct {
		Value    int
		Priority int
	}

	items := []Item{
		{1, 5},
		{2, 3},
		{3, 5},
		{4, 3},
	}

	// When priorities match, keep input sequence
	// 当优先级相同时，保持输入顺序
	sortx.SortVStable(items, func(a, b Item) bool {
		return a.Priority < b.Priority
	})

	// Use spew to show struct details with field names
	// 使用 spew 显示结构体详情和字段名
	t.Log(spew.Sdump(items))

	expected := []Item{
		{2, 3},
		{4, 3},
		{1, 5},
		{3, 5},
	}

	require.Equal(t, expected, items)
}

// TestSortEmptySlice tests sorting on slices with no elements
// TestSortEmptySlice 测试无元素切片的排序
func TestSortEmptySlice(t *testing.T) {
	// Slices with no elements should remain unchanged
	// 无元素切片应保持不变
	var noElements []int

	sortx.SortByIndex(noElements, func(i, j int) bool {
		return noElements[i] < noElements[j]
	})
	require.Empty(t, noElements)

	sortx.SortByValue(noElements, func(a, b int) bool {
		return a < b
	})
	require.Empty(t, noElements)

	sortx.SortIStable(noElements, func(i, j int) bool {
		return noElements[i] < noElements[j]
	})
	require.Empty(t, noElements)

	sortx.SortVStable(noElements, func(a, b int) bool {
		return a < b
	})
	require.Empty(t, noElements)
}

// TestSortSingleElement tests sorting on slices with one element
// TestSortSingleElement 测试单元素切片的排序
func TestSortSingleElement(t *testing.T) {
	// Slices with one element should remain unchanged
	// 单元素切片应保持不变
	single := []string{"hello"}

	sortx.SortByIndex(single, func(i, j int) bool {
		return single[i] < single[j]
	})
	require.Equal(t, []string{"hello"}, single)

	sortx.SortByValue(single, func(a, b string) bool {
		return a < b
	})
	require.Equal(t, []string{"hello"}, single)
}

// TestIsSortedByValue tests IsSortedByValue function
// TestIsSortedByValue 测试 IsSortedByValue 函数
func TestIsSortedByValue(t *testing.T) {
	require.True(t, sortx.IsSortedByValue([]int{1, 2, 3, 4, 5}, func(a, b int) bool {
		return a < b
	}))
	require.False(t, sortx.IsSortedByValue([]int{3, 1, 4, 1, 5}, func(a, b int) bool {
		return a < b
	}))
}

// TestIsSortedByIndex tests IsSortedByIndex function
// TestIsSortedByIndex 测试 IsSortedByIndex 函数
func TestIsSortedByIndex(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	require.True(t, sortx.IsSortedByIndex(a, func(i, j int) bool {
		return a[i] < a[j]
	}))
	b := []int{3, 1, 4, 1, 5}
	require.False(t, sortx.IsSortedByIndex(b, func(i, j int) bool {
		return b[i] < b[j]
	}))
}
