package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	cases := []struct {
		nums []int
		want float64
	}{
		{
			nums: []int{3, 4, 5, 6},
			want: 0.5,
		},
	}

	for i, v := range cases {
		t.Run("case:"+strconv.Itoa(i), func(t *testing.T) {
			ans := solve(v.nums)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("want:%v;but got %v", v.want, ans)
			}
		})
	}
}
