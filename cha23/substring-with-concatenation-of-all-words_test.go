package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSubstring(t *testing.T) {
	cases := []struct {
		s     string
		words []string
		want  []int
	}{
		{
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			want:  []int{0, 9},
		},
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "word"},
			want:  nil,
		},
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "good"},
			want:  []int{8},
		},
	}

	for i, v := range cases {
		t.Run("case:"+strconv.Itoa(i), func(t *testing.T) {
			out := findSubstring(v.s, v.words)
			if !reflect.DeepEqual(out, v.want) {
				t.Errorf("want:%#v;but got:%#v", v.want, out)
			}
		})
	}
}
