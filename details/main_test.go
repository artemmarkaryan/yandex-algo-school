package main

import "testing"

func TestEvaluate(t *testing.T) {
	for c, r := range map[[3]int]int{
		{10, 5, 2}: 4,
		{13, 5, 3}: 3,
		{14, 5, 3}: 4,
		{200, 1, 1}: 200,
		{200, 200, 33}: 6,
	} {
		if Evaluate(c[0], c[1], c[2]) != r {
			t.Logf("%v -> %v: %v", c, Evaluate(c[0], c[1], c[2]), r)
			t.Fail()
		}
	}
}
