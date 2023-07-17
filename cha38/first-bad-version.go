package main

/*
你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。

假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。

你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

示例:

给定 n = 5，并且 version = 4 是第一个错误的版本。

调用 isBadVersion(3) -> false
调用 isBadVersion(5) -> true
调用 isBadVersion(4) -> true

所以，4 是第一个错误的版本。
*/

/*
思路：
`二分法`
典型的二分寻`找最左边的满足条件的值`。
一句话概括就是：寻找最左边和寻找指定值的差别就是碰到等于号的处理情况。 如果是寻找最左边那么碰到等于继续收缩右边界（寻找最右边就是收缩左边界），查找指定值则是直接返回。

复杂度：
时间复杂度： O(logN)
空间复杂度： O(1)
*/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 */
func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	l, r := 1, n
	for l <= r {
		mid := (l + r) / 2
		if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
