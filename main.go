package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(8, []int{1,1}))
}
func minSubArrayLen(s int, nums []int) int {
	sum, min := 0, math.MaxInt64
	slow := 0
	flag := 0
	for i := 0; i < len(nums); {
		if sum += nums[i]; sum >= s {
			if temp := i - slow + 1; min > temp {
				min = temp
			}
			sum = 0
			slow++
			i = slow
			flag = 1
			continue
		}
		i++
	}
	if flag == 0 {
		return 0
	}
	return min
}