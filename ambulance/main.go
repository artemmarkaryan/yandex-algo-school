/*
Бригада скорой помощи выехала по вызову в один из отделенных районов. К
сожалению, когда диспетчер получил вызов, он успел записать только адрес дома и
номер квартиры K1, а затем связь прервалась. Однако он вспомнил, что по этому же
адресу дома некоторое время назад скорая помощь выезжала в квартиру K2, которая
расположена в подъезда P2 на этаже N2. Известно, что в доме M этажей и
количество квартир на каждой лестничной площадке одинаково. Напишите программу,
которая вычилсяет номер подъезда P1 и номер этажа N1 квартиры K1.

Формат ввода
Во входном файле записаны пять положительных целых чисел K1, M, K2, P2, N2. Все
числа не превосходят 106.

Формат вывода
Выведите два числа P1 и N1. Если входные данные не позволяют однозначно
определить P1 или N1, вместо соответствующего числа напечатайте 0. Если входные
данные противоречивы, напечатайте два числа –1 (минус один).
*/

// unfinished

package main

import (
	"fmt"
)

const maxFloorCapacity = 5

type HouseConfig struct {
	floors, // всего этажей
	floorCapacity int
}

func EvaluateFloorCapacity(flat, entrance, floor, floors int) (floorCapacities []int, ok bool) {
	//if floors *
	for x := 1; x <= maxFloorCapacity; x++ {
		entranceCondition := ((flat-1)/(x*floors))+1 == entrance
		floorCondition := ((flat-1)%(x*floors))/x+1 == floor
		//fmt.Println("x: ", x)
		//fmt.Println("entranceCondition: ", entranceCondition)
		//fmt.Println("floorCondition: ", floorCondition)
		if entranceCondition && floorCondition {
			floorCapacities = append(floorCapacities, x)
		}
	}
	return
}

func (c *HouseConfig) EvaluateFloor(flat int) (floor int, ok bool) {
	if flat < 1 {
		return
	} else {
		ok = true
	}
	floor = ((flat-1)%(c.floorCapacity*c.floors))/c.floorCapacity+1
	return
}

func (c *HouseConfig) EvaluateEntrance(flat int) (entrance int, ok bool) {
	if flat < 1 {
		return
	} else {
		ok = true
	}
	entrance = ((flat-1)%(c.floorCapacity*c.floors))/c.floorCapacity+1
	return
}



func main() {
	var k1, m, k2, p2, n2 int
	_, _ = fmt.Scanf("%v %v %v %v %v", &k1, &m, &k2, &p2, &n2)
	//c := HouseConfig{m, 0}

	//capacityOk := c.EvaluateFloorCapacity(k2, p2, n2)
	//if !capacityOk {
	//	fmt.Print(-1)
	//} else {
	//	log.Print(c.floorCapacity, c.floors)
		//entrance, _ := c.EvaluateEntrance(k1)
		//floor, _ := c.EvaluateFloor(k1)
		//fmt.Printf("%v %v", entrance, floor)
	//}
}
