package main

import (
	"testing"
)

func TestRegimeBuilder(t *testing.T) {
	inputs := map[struct {
		from   int8
		to     int8
		regime string
	}]int8{
		{10, 20, "heat"}: 20,
		{10, 20, "freeze"}: 10,
		{10, 20, "auto"}: 20,
		{10, 20, "fan"}: 10,
	}

	for k, v := range inputs {
		regime, err := RegimeBuilder(k.regime)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		result := regime.ChangeTemp(k.from, k.to)
		if result != v {
			t.Error(k, "->", result, ";", v)
		}
	}
}
