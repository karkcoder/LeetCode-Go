package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				//Do nothing??
			} else if (nums[i] + nums[j]) == target {
				fmt.Println(nums[i])
				fmt.Println(nums[j])
				return ([]int{i, j})
			}
		}
	}
	return []int{0, 0}
}
