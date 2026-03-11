package main

import (
	"fmt"

	"github.com/yylego/sortx"
)

func main() {
	// Sort integers using index comparison
	// 使用索引比较排序整数
	numbers := []int{5, 3, 8, 1, 4}
	sortx.SortByIndex(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println("Sorted using index:", numbers)

	// Sort strings using value comparison
	// 使用值比较排序字符串
	fruits := []string{"apple", "banana", "orange", "date"}
	sortx.SortByValue(fruits, func(a, b string) bool {
		return a < b
	})
	fmt.Println("Sorted using value:", fruits)
}
