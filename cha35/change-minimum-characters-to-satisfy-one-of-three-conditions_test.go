package main

import (
	"strconv"
	"testing"
)

func TestMin(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{
			"a",
			"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			2,
		},
	}

	for i, param := range cases {
		t.Run("case+"+strconv.Itoa(i), func(t *testing.T) {
			ans := minCharacters(param.a, param.b)
			if ans != param.want {
				t.Errorf("want:%v;but got:%v", param.want, ans)
			}
		})
	}
}
