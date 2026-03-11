package sortx

import "sort"

// SortByIndex sorts the slice using index-based comparison
// Compares elements at given indices
//
// SortByIndex 使用索引比较对切片排序
// 通过给定索引比较元素
func SortByIndex[V any](a []V, iLess func(i, j int) bool) {
	// Execute sorting with index comparison
	// 使用索引比较执行排序
	sort.Sort(NewSortByIndex(a, iLess))
}

// SortByValue sorts the slice using value-based comparison
// Compares elements using values
//
// SortByValue 使用值比较对切片排序
// 通过值比较元素
func SortByValue[V any](a []V, vLess func(a, b V) bool) {
	// Execute sorting with value comparison
	// 使用值比较执行排序
	sort.Sort(NewSortByValue(a, vLess))
}

// SortIStable sorts the slice with stable mode using index-based comparison
// Same elements remain in input sequence
//
// SortIStable 使用索引比较对切片进行稳定排序
// 相同元素保持输入时的顺序
func SortIStable[V any](a []V, iLess func(i, j int) bool) {
	// Execute stable sorting with index comparison
	// 使用索引比较执行稳定排序
	sort.Stable(NewSortByIndex(a, iLess))
}

// SortVStable sorts the slice with stable mode using value-based comparison
// Same elements remain in input sequence
//
// SortVStable 使用值比较对切片进行稳定排序
// 相同元素保持输入时的顺序
func SortVStable[V any](a []V, vLess func(a, b V) bool) {
	// Execute stable sorting with value comparison
	// 使用值比较执行稳定排序
	sort.Stable(NewSortByValue(a, vLess))
}
