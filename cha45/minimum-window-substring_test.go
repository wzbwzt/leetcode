package main

import (
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	cases := []struct {
		s string
		t string

		want string
	}{
		{s: "ADOBECODEBANC", t: "ABC", want: "BANC"},
	}

	for i, v := range cases {
		t.Run("case:"+strconv.Itoa(i), func(t *testing.T) {
			ans := minWindow(v.s, v.t)
			if ans != v.want {
				t.Errorf("want:%v;but got:%v", v.want, ans)
			}
		})
	}

}
