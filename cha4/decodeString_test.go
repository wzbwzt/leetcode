package main

import (
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{"3[a2[c]]", "accaccacc"},
		{"11[ac]", "acacacacacacacacacacac"},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("case", i), func(t *testing.T) {
			res := decodeString(v.s)
			if res != v.want {
				t.Errorf("get:%s ;but want:%s", res, v.want)
			}
		})
	}
}
