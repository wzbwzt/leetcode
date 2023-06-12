package main

import "strings"

/*
题目描述:
leetcode:https://leetcode.cn/problems/shortest-distance-to-a-character/

给定一个字符串 S 和一个字符 C。返回一个代表字符串 S 中每个字符到字符串 S 中的字符 C 的最短距离的数组。

示例 1:

输入: S = "loveleetcode", C = 'e'
输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
说明:

- 字符串 S 的长度范围为 [1, 10000]。
- C 是一个单字符，且保证是字符串 S 里的字符。
- S 和 C 中的所有字母均为小写字母。
*/

/*
思路：
遍历数组
left:当前遍历字符左边最近的的c下标，没有的话为len(s)
right: 当前遍历字符的右边最近的c下标


复杂度：
空间复杂度为 O(n)
时间复杂度为 O(n)
*/

func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)

	l := 0
	if s[0] != c {
		l = n
	}
	r := strings.IndexByte(s[1:], c) + 1

	for i := 0; i < n; i++ {
		res[i] = min(abs(i-l), abs(r-i))
		if i == r {
			l = r
			r = strings.IndexByte(s[l+1:], c) + l + 1
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
