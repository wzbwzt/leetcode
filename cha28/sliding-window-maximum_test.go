package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSliding(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		// {
		// 	nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
		// 	k:    3,
		// 	want: []int{3, 3, 5, 5, 6, 7},
		// },
		{
			nums: []int{1, 3, -1, 5, 5, 3, 6, 7},
			k:    4,
			want: []int{5, 5, 5, 6, 7},
		},
	}

	for i, v := range cases {
		t.Run("case:"+strconv.Itoa(i), func(t *testing.T) {
			out := maxSlidingWindow(v.nums, v.k)
			if !reflect.DeepEqual(out, v.want) {
				t.Errorf("want:%#v;but got:%#v", v.want, out)
			}
		})

	}

}
