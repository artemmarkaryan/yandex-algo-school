/*
Имеется N кг металлического сплава. Из него изготавливают заготовки массой K кг
каждая. После этого из каждой заготовки вытачиваются детали массой M кг каждая
(из каждой заготовки вытачивают максимально возможное количество деталей). Если
от заготовок после этого что-то остается, то этот материал возвращают к началу
производственного цикла и сплавляют с тем, что осталось при изготовлении
заготовок. Если того сплава, который получился, достаточно для изготовления хотя
бы одной заготовки, то из него снова изготавливают заготовки, из них – детали и
т.д. Напишите программу, которая вычислит, какое количество деталей может быть
получено по этой технологии из имеющихся исходно N кг сплава.

Формат ввода
Вводятся N, K, M. Все числа натуральные и не превосходят 200.

Формат вывода
Выведите одно число — количество деталей, которое может получиться
по такой технологии.
*/

package main

import "fmt"

func main() {
	var n, k, m, details int
	_, _ = fmt.Scanf("%v %v %v", &n, &k, &m)

	for n >= k {
		billets := n / k
		n = n % k
		details += (k / m) * billets
		n += (k % m) * details
	}
	fmt.Print(details)
}