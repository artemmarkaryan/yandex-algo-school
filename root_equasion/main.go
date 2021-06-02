/*


Решите в целых числах уравнение:
sqrt(ax+b) = c

a, b, c – данные целые числа: найдите все решения или сообщите, что решений в
целых числах нет

Формат ввода
Вводятся три числа a, b и c по одному в строке.

Формат вывода
Программа должна вывести все решения уравнения в порядке возрастания, либо NO
SOLUTION (заглавными буквами), если решений нет. Если решений бесконечно много,
вывести MANY SOLUTIONS.

*/
package main

import (
	"fmt"
	"strconv"
)

const (
	NoSolutions   = "NO SOLUTION"
	ManySolutions = "MANY SOLUTIONS"
)

func isInt(n float32) bool {
	return n == float32(int(n))
}

func main() {
	var a, b, c int
	var result_s string

	_, _ = fmt.Scanf("%v\n%v\n%v", &a, &b, &c)

	if c < 0 {
		result_s = NoSolutions
	} else if a == 0 {
		if c*c == b {
			result_s = ManySolutions
		} else {
			result_s = NoSolutions
		}
	} else if c == 0 {
		if a == 0 {
			if b == 0 {
				result_s = ManySolutions
			} else if isInt(float32(-b) / float32(a)){
				result_s = strconv.Itoa(int(float32(-b) / float32(a)))
			} else { // non-int solution
				result_s = NoSolutions
			}
		} else { // a != 0
			v := float32(-b) / float32(a)
			if isInt(v){
				result_s = strconv.Itoa(int(v))
			} else { // non-int solution
				result_s = NoSolutions
			}
		}
	} else { // c > 0
		if a == 0 {
			if c * c == b {
				result_s = ManySolutions
			} else {
				result_s = NoSolutions
			}
		} else { // a != 0
			v := (float32(c * c) - float32(b)) / float32(a)
			if isInt(v){
				result_s = strconv.Itoa(int(v))
			} else { // non-int solution
				result_s = NoSolutions
			}
		}
	}
	fmt.Print(result_s)
}
