package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortedListToBST(t *testing.T) {
	cases := []struct {
		header *ListNode
		want   *TreeNode
	}{
		{
			header: &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}},
			want:   &TreeNode{Val: 0, Left: &TreeNode{Val: -3, Left: &TreeNode{Val: -10}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}}},
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			out := sortedListToBST(v.header)
			if !reflect.DeepEqual(out, v.want) {
				t.Errorf("want:%#v;but got:%#v", v.want, out)
			}
		})
	}
}
