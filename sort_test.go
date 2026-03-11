package sortx_test

import (
	"sort"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"github.com/yylego/sortx"
)

// TestNewSortByIndex tests NewSortByIndex with index-based comparison
// TestNewSortByIndex 测试使用索引比较的 NewSortByIndex
func TestNewSortByIndex(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	sort.Sort(sortx.NewSortByIndex(a, func(i, j int) bool {
		return a[i] < a[j]
	}))
	t.Log(spew.Sdump(a)) // Use spew to dump slice state // 使用 spew 输出切片状态
	require.Equal(t, a, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

// TestNewSortByValue tests NewSortByValue with value-based comparison
// TestNewSortByValue 测试使用值比较的 NewSortByValue
func TestNewSortByValue(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	sort.Sort(sortx.NewSortByValue(a, func(a, b int) bool {
		return a < b
	}))
	t.Log(spew.Sdump(a)) // Use spew to dump slice state // 使用 spew 输出切片状态
	require.Equal(t, a, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

// TestSliceMethods tests sort.Interface methods: Len, Less, Swap
// TestSliceMethods 测试 sort.Interface 方法：Len、Less、Swap
func TestSliceMethods(t *testing.T) {
	// Test sort.Interface methods via public API
	// 通过公开 API 测试 sort.Interface 方法
	data := []int{3, 1, 4, 1, 5}

	// Test with index-based comparison
	// 测试索引比较
	sortable1 := sortx.NewSortByIndex(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	require.Equal(t, 5, sortable1.Len())
	require.True(t, sortable1.Less(1, 0))  // data[1]=1 < data[0]=3
	require.False(t, sortable1.Less(0, 1)) // data[0]=3 > data[1]=1

	// Test Swap method
	// 测试 Swap 方法
	sortable1.Swap(0, 1)
	require.Equal(t, []int{1, 3, 4, 1, 5}, data)

	// Test with value-based comparison
	// 测试值比较
	data2 := []string{"banana", "apple", "orange"}
	sortable2 := sortx.NewSortByValue(data2, func(a, b string) bool {
		return a < b
	})

	require.Equal(t, 3, sortable2.Len())
	require.True(t, sortable2.Less(1, 0))  // "apple" < "banana"
	require.False(t, sortable2.Less(0, 1)) // "banana" > "apple"

	sortable2.Swap(0, 1)
	require.Equal(t, []string{"apple", "banana", "orange"}, data2)
}

// TestSortInterface tests sort.Interface implementation
// TestSortInterface 测试 sort.Interface 实现
func TestSortInterface(t *testing.T) {
	// Check sort.Interface methods work as expected
	// 检查 sort.Interface 方法按预期工作
	data := []float64{3.14, 2.71, 1.41, 1.73}

	sortable := sortx.NewSortByValue(data, func(a, b float64) bool {
		return a < b
	})

	// Test the interface methods
	// 测试接口方法
	require.Equal(t, 4, sortable.Len())
	require.True(t, sortable.Less(2, 0)) // 1.41 < 3.14

	// Test complete sorting
	// 测试完整排序
	sort.Sort(sortable)
	t.Log(spew.Sdump(data)) // Use spew to dump float64 slice // 使用 spew 输出浮点数切片
	expected := []float64{1.41, 1.73, 2.71, 3.14}
	require.Equal(t, expected, data)
}

// TestNewSortByIndexPanicOnNilILess tests panic when iLess is nil
// TestNewSortByIndexPanicOnNilILess 测试 iLess 为 nil 时的 panic
func TestNewSortByIndexPanicOnNilILess(t *testing.T) {
	require.Panics(t, func() {
		sortx.NewSortByIndex([]int{1, 2, 3}, nil)
	})
}

// TestNewSortByValuePanicOnNilVLess tests panic when vLess is nil
// TestNewSortByValuePanicOnNilVLess 测试 vLess 为 nil 时的 panic
func TestNewSortByValuePanicOnNilVLess(t *testing.T) {
	require.Panics(t, func() {
		sortx.NewSortByValue([]string{"a", "b"}, nil)
	})
}
