package main

/*
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5

*/

/*
思路：
`二分法`
a^2<=x
以[1，x]为边界，找到符合条件的最大a值



复杂度：
时间复杂度：O(logN)
空间复杂度：O(1)
*/

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if ans != -1 {
		return ans
	}
	return l
}
