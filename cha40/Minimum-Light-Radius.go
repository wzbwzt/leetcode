package main

import "sort"

/*
You are given a list of integers nums representing coordinates of houses on a 1-dimensional line.
You have 3 street lights that you can put anywhere on the coordinate line and a light at coordinate
x lights up houses in [x - r, x + r], inclusive. Return the smallest r required such that we can place the 3 lights and all the houses are lit up.

Constraints

n ≤ 100,000 where n is the length of nums
Example 1
Input
nums = [3, 4, 5, 6]
Output
0.5
Explanation
If we place the lamps on 3.5, 4.5 and 5.5 then with r = 0.5 we can light up all 4 houses.

*/

/*
思路：
`能力检测` `最左二分`



复杂度：
时间复杂度： O(nlogn+logn∗log(MAX−MIN))
空间复杂度： O(1)
*/

func solve(nums []int) float64 {
	if len(nums) <= 3 {
		return 0
	}
	sort.Ints(nums)
	l := 0
	r := nums[len(nums)-1] - nums[0]
	for l <= r {
		mid := l + (r-l)>>1
		if isPossible(nums, mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return float64(l) / 2
}

func isPossible(nums []int, mid int) bool {
	start := nums[0]
	end := start + mid
	for i := 0; i < 3; i++ {
		index := 0
		for index < len(nums) && nums[index] <= end {
			index++
		}
		if index >= len(nums) {
			return true
		}
		start = nums[index]
		end = start + mid
	}
	return false
}
