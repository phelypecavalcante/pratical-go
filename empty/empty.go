package main

import "fmt"

func main() {
	fmt.Println(max([]int{3, 1, 2}))
	fmt.Println(max([]float64{3, 1, 2}))
}

type Number interface {
	int | float64
}

func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}
