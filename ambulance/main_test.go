package main

import (
	"fmt"
	"testing"
)

//func TestRandomEvaluateFloorCapacity(t *testing.T) {
//	const casesAmount = 10
//	for i := 0; i < casesAmount; i++ {
//		c := HouseConfig{
//			floors:   rand.Intn(100) + 1,
//			flat:     rand.Intn(100) + 1,
//			entrance: rand.Intn(100) + 1,
//			floor:    rand.Intn(100) + 1,
//		}
//		x, ok := EvaluateFloorCapacity(c)
//		log.Printf("config: %v;\n ok: %5.5v; x: %5.5v\n\n", c, ok, x)
//	}
//}

func TestEvaluateFloorCapacity(t *testing.T) {
	for k, v := range map[
	struct {
		c                     HouseConfig
		flat, entrance, floor int
	}]bool{
		{
			c:        HouseConfig{floors: 4},
			flat:     11,
			entrance: 2,
			floor:    2,
		}: true,
		{
			c:        HouseConfig{floors: 4},
			flat:     8,
			entrance: 1,
			floor:    4,
		}: true,
		{
			c:        HouseConfig{floors: 4},
			flat:     9,
			entrance: 2,
			floor:    1,
		}: true,
		{
			c:        HouseConfig{floors: 4},
			flat:     1,
			entrance: 1,
			floor:    1,
		}: false,
	} {
		ok := k.c.EvaluateFloorCapacity(k.floor, k.entrance, k.flat)
		fmt.Println(ok, k.c.floorCapacity, v)
	}
}
