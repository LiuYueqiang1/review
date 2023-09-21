package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	count := len(nums)
	if count < 3 {
		return [][]int{}
	}

	sort.Ints(nums) // 排序后简化的问题，可以使用滑动窗口
	ans := make([][]int, 0)
	// 寻找 -a=b+c
	for a := 0; a < count; a++ {
		if a > 0 && (nums[a] == nums[a-1]) { //去掉重复a，因为排过序，所以只跟上一个比
			continue
		}
		target := (-nums[a])
		c := count - 1                   //如果c=b+1可能会越界
		for b := a + 1; b < count; b++ { // 固定b找c
			if (b > a+1) && nums[b] == nums[b-1] { //去掉重复b
				continue
			}
			for b < c && (nums[b]+nums[c] > target) {
				c--
			}
			if b == c {
				break
			}
			if nums[b]+nums[c] == target {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}

		}
	}
	return ans
}
func main() {
	var a int
	fmt.Scan(&a)
	num := make([]int, 0)
	nums := append(num, a)
	result := threeSum(nums)
	fmt.Println(result)
}
