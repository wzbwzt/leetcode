package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestVerticalTraversal(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want [][]int
	}{{
		root: &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}},
		want: [][]int{{2}, {1}, {0, 3}},
	}}

	for i, param := range cases {
		t.Run("case:"+strconv.Itoa(i), func(t *testing.T) {
			out := verticalTraversal(param.root)
			if !reflect.DeepEqual(out, param.want) {
				t.Errorf("want:%v,but get:%v;", param.want, out)
			}
		})
	}

}
