package main

/*
You are given a list of positive integers nums and an integer target. Consider an operation where we remove a number v
from either the front or the back of nums and decrement target by v.

Return the minimum number of operations required to decrement target to zero. If it's not possible, return -1.

Constraints
n ≤ 100,000 where n is the length of nums

Example 1 Input nums = [3, 1, 1, 2, 5, 1, 1] target = 7 Output 3 Explanation
We can remove 1, 1 and 5 from the back to decrement target to zero.

Example 2 Input nums = [2, 4] target = 7 Output -1 Explanation There's no way to decrement target = 7 to zero.

*/

/*
思路：
`滑动窗口`

问题可以转换为求nums中最长的连续字串和=sum(nums)-target


复杂度：
时间复杂度： O(n)
空间复杂度： 1
*/

func solve(nums []int, target int) int {
	length := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}

	target = sum - target

	ans := length + 1
	i := 0
	tmp_sum := 0
	for j := 0; j < length; j++ {
		tmp_sum += nums[j]
		for tmp_sum > sum {
			tmp_sum -= nums[i]
			i++
		}

		if tmp_sum == target {
			ans = min(ans, length-(j-i+1))
		}
	}
	if ans == length+1 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
