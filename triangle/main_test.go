package main

import "testing"

func TestCalculate(t *testing.T) {
	cases := map[[3]int]bool {
		{1, 2, 3}: false,
		{3, 4, 5}: true,
	}
	for k, v := range cases {
		if Calculate(k[0], k[1], k[2]) != v {
			t.Error(k, v)
		}
	}
}
