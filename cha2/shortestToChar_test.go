package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShorttestToChar(t *testing.T) {
	ts := []struct {
		s    string
		c    byte
		want []int
	}{
		{
			s:    "loveleetcode",
			c:    'e',
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
	}
	for i, param := range ts {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := shortestToChar(param.s, param.c)
			if !reflect.DeepEqual(res, param.want) {
				t.Errorf("got:%#v;but want:%#v", res, param.want)
			}
		})
	}

}
