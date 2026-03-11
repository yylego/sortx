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
	// 稳定排序保持相同元素的输入顺序
	tasks := []Task{
		{"Task-A", 2},
		{"Task-B", 1},
		{"Task-C", 2},
		{"Task-D", 1},
	}

	// Sort using Rank, same Rank keeps input sequence
	// 按 Rank 排序，相同 Rank 保持输入顺序
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
