package main

import (
	"strconv"
	"strings"
	"unicode"
)

/*
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。



示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/

/*
思路:
栈思路,遍历s,压入栈，遇到']'弹栈
只有四种可能出现的字符: 字母\ 数字(可能对位)\ [ \ ]

复杂度：
时间复杂度为O(n)
空间复杂度为O(n)
*/
func decodeString(s string) string {
	stack := []string{}
	last := []int{}
	tmpCount := ""
	for _, v := range s {
		if unicode.IsDigit(v) {
			tmpCount += string(v)
			continue
		}

		if len(tmpCount) != 0 {
			stack = append(stack, tmpCount)
			tmpCount = ""
		}

		if v == '[' {
			last = append(last, len(stack))
		}
		if v == ']' {
			new := last[len(last)-1]
			num := stack[new-1]
			value := make([]string, len(stack)-new-1)
			copy(value, stack[new+1:])
			stack = stack[:new-1]
			num_i, _ := strconv.Atoi(string(num))
			for i := 0; i < num_i; i++ {
				stack = append(stack, strings.Join(value, ""))
			}
			last = last[:len(last)-1]
			continue
		}

		stack = append(stack, string(v))
	}
	return strings.Join(stack, "")
}
