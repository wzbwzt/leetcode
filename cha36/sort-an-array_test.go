package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSort(t *testing.T) {
	cases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 3, 2, 3, 1},
			want: []int{1, 1, 2, 3, 3},
		},
	}

	for i, v := range cases {
		t.Run("case:"+strconv.Itoa(i), func(t *testing.T) {
			ans := sortArray(v.nums)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("want:%v;but got:%v", v.want, ans)
			}
		})
	}

}
