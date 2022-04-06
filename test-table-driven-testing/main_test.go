package main

import "testing"

func TestPingPong(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{0, "ping-pong"},
		{1, ""},
		{2, "ping"},
		{3, "pong"},
		{4, "ping"},
		{5, ""},
		{6, "ping-pong"},
		{7, ""},
		{8, "ping"},
		{9, "pong"},
	}
	for _, test := range tests {
		result := pingPong(test.input)
		if result != test.want {
			t.Fatalf("pingPong(%d)=%s, want=%s", test.input, result, test.want)
		}
	}
}
