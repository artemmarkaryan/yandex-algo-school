/*
В офисе, где работает программист Петр, установили кондиционер нового типа. Этот
кондиционер отличается особой простотой в управлении. У кондиционера есть всего
лишь два управляемых параметра: желаемая температура и режим работы.

Кондиционер может работать в следующих четырех режимах:

«freeze» — охлаждение. В этом режиме кондиционер может только уменьшать
температуру. Если температура в комнате и так не больше желаемой, то он
выключается.

«heat» — нагрев. В этом режиме кондиционер может только увеличивать температуру.
Если температура в комнате и так не меньше желаемой, то он выключается.

«auto» — автоматический режим. В этом режиме кондиционер может как увеличивать,
так и уменьшать температуру в комнате до желаемой.

«fan» — вентиляция. В этом режиме кондиционер осуществляет только вентиляцию
воздуха и не изменяет температуру в комнате.

Кондиционер достаточно мощный, поэтому при настройке на правильный режим работы
он за час доводит температуру в комнате до желаемой.

Требуется написать программу, которая по заданной температуре в комнате troom,
установленным на кондиционере желаемой температуре tcond и режиму работы
определяет температуру, которая установится в комнате через час.

Формат ввода

Первая строка входного файла содержит два целых числа troom, и tcond,
разделенных ровно одним пробелом (–50 ≤ troom ≤ 50, –50 ≤ tcond ≤ 50).

Вторая строка содержит одно слово, записанное строчными буквами латинского алфавита — режим работы кондиционера.

Формат вывода

Выходной файл должен содержать одно целое число — температуру, которая установится в комнате через час.
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Regime interface {
	ChangeTemp(from, to int8) (result int8)
}

func RegimeBuilder(s string) (Regime, error) {
	regime, ok := map[string]Regime{
		"freeze": FreezeRegime{},
		"heat":   heatRegime{},
		"auto":   AutoRegime{},
		"fan":    FanRegime{},
	}[s]
	if ok {
		return regime, nil
	} else {
		return nil, errors.New("no such regime")
	}
}

type FreezeRegime struct{}

func (f FreezeRegime) ChangeTemp(from, to int8) (result int8) {
	if from > to {
		result = to
	} else {
		result = from
	}
	return
}

type heatRegime struct{}

func (f heatRegime) ChangeTemp(from, to int8) (result int8) {
	if from < to {
		result = to
	} else {
		result = from
	}
	return
}

type AutoRegime struct{}

func (f AutoRegime) ChangeTemp(from, to int8) (result int8) {
	result = to
	return
}

type FanRegime struct{}

func (f FanRegime) ChangeTemp(from, to int8) (result int8) {
	result = from
	return
}

//goland:noinspection ALL
func main() {
	var troom, tcond int
	var regimeInput string

	scanner := bufio.NewScanner(os.Stdin)
	{
		scanner.Scan();
		troom, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
		tcond, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[1])
	}
	{
		scanner.Scan();
		regimeInput = scanner.Text()
	}
	selectedRegime, err := RegimeBuilder(regimeInput)
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Print(selectedRegime.ChangeTemp(int8(troom), int8(tcond)))
}
