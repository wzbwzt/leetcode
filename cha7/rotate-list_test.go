package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotateList(t *testing.T) {
	cases := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			head: &ListNode{Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}, Val: 1},
			k:    2,
			want: &ListNode{Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}, Val: 4},
		},
		{
			head: &ListNode{Next: nil, Val: 1},
			k:    0,
			want: &ListNode{Next: nil, Val: 1},
		},
		{
			head: &ListNode{Next: nil, Val: 1},
			k:    1,
			want: &ListNode{Next: nil, Val: 1},
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			out := rotateRight(v.head, v.k)
			if !reflect.DeepEqual(out, v.want) {
				t.Errorf("want:%#v;but got:%#v", v.want, out)
			}
		})
	}
}
