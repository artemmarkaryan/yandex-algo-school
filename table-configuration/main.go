/*


В школе решили на один прямоугольный стол поставить два прямоугольных ноутбука.
Ноутбуки нужно поставить так, чтобы их стороны были параллельны сторонам стола.
Определите, какие размеры должен иметь стол, чтобы оба ноутбука на него
поместились, и площадь стола была минимальна.

Формат ввода
Вводится четыре натуральных числа, первые два задают размеры одного ноутбука, а
следующие два — размеры второго. Числа не превышают 1000.

Формат вывода
Выведите два числа — размеры стола. Если возможно несколько ответов, выведите
любой из них (но только один).
 */
package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, length int
}

func (r *Rectangle) Transform() {
	r.length, r.width = r.width, r.length
}

func (r *Rectangle) Square() int {
	return r.width * r.length
}

func getTable(l1, l2 Rectangle) (rectangle Rectangle, square int) {
	rectangle.length = int(math.Max(float64(l1.length), float64(l2.length)))
	rectangle.width = l1.width + l2.width
	return rectangle, rectangle.Square()
}

func main() {
	var l1, l2, t, min_t Rectangle
	min_s := math.MaxInt32
	_, _ = fmt.Scanf("%v %v %v %v", &l1.width, &l1.length, &l2.width, &l2.length)

	t, s := getTable(l1, l2)
	if s < min_s {
		min_s = s
		min_t = t
	}

	l1.Transform()
	t, s = getTable(l1, l2)
	if s < min_s {
		min_s = s
		min_t = t
	}

	l2.Transform()
	t, s = getTable(l1, l2)
	if s < min_s {
		min_s = s
		min_t = t
	}

	l1.Transform()
	t, s = getTable(l1, l2)
	if s < min_s {
		min_s = s
		min_t = t
	}

	fmt.Printf("%v %v", min_t.length, min_t.width)
}
