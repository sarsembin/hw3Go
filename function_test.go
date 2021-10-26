package main

import (
	"github.com/sarsembin/hw3Go/atoi"
	"testing"
)

func Test_atoi(t *testing.T) {
	table := []struct {
		input string
		want  int
		errExpected bool
	}{
		{"-123567", -123567, false},
		{"みくだよ", 0, true},
		{"123-4213", 0, true},
		{"cake is a lie", 0, true},

	}

	for _, v := range table {
		got, err := atoi.Atoi(v.input)
		if err != nil && !v.errExpected {
			t.Errorf("Unexpected error: %v\n", err)
		}
		if got != v.want {
			t.Errorf("Result of atoi was incorrect, got: %v, want: %v\n", got, v.want)
		}
	}
}
