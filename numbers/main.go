/*
Телефонные номера в адресной книге мобильного телефона имеют один из следующих
форматов: +7<код><номер> 8<код><номер> <номер> где <номер> — это семь цифр, а
<код> — это три цифры или три цифры в круглых скобках. Если код не указан, то
считается, что он равен 495. Кроме того, в записи телефонного номера может
стоять знак “-” между любыми двумя цифрами (см. пример). На данный момент в
адресной книге телефона Васи записано всего три телефонных номера, и он хочет
записать туда еще один. Но он не может понять, не записан ли уже такой номер в
телефонной книге. Помогите ему! Два телефонных номера совпадают, если у них
равны коды и равны номера. Например, +7(916)0123456 и 89160123456 — это один и
тот же номер.

Формат ввода
В первой строке входных данных записан номер телефона, который Вася хочет
добавить в адресную книгу своего телефона. В следующих трех строках записаны три
номера телефонов, которые уже находятся в адресной книге телефона Васи.
Гарантируется, что каждая из записей соответствует одному из трех приведенных в
условии форматов.

Формат вывода
Для каждого телефонного номера в адресной книге выведите YES (заглавными
буквами), если он совпадает с тем телефонным номером, который Вася хочет
добавить в адресную книгу или NO (заглавными буквами) в противном случае.
*/
package main

import (
	"fmt"
	"strings"
)

type PhoneModifier interface {
	Apply(s string) (r string)
}

type ReduceDashPhoneModifier struct{}

func (m ReduceDashPhoneModifier) Apply(s string) (r string) {
	return strings.Join(strings.Split(s, "-"), "")
}

type ReduceBracketsPhoneModifier struct{}

func (m ReduceBracketsPhoneModifier) Apply(s string) (r string) {
	r = strings.ReplaceAll(
		strings.ReplaceAll(
			s,
			"(",
			"",
		),
		")",
		"",
	)
	return
}

type TrimCountryCodePhoneModifier struct{}

func (m TrimCountryCodePhoneModifier) Apply(s string) (r string) {
	if s[0] == '8' {
		r = s[1:]
	} else if s[:2] == "+7" {
		r = s[2:]
	} else {
		r = s
	}
	return r
}

type AddCityCodePhoneModifier struct{}

func (m AddCityCodePhoneModifier) Apply(s string) (r string) {
	if len(s) == 7 {
		r = "495" + s
	} else {
		r = s
	}

	return r
}

func main() {
	var phones [4]string
	modifiers := []PhoneModifier{
		ReduceDashPhoneModifier{},
		ReduceBracketsPhoneModifier{},
		TrimCountryCodePhoneModifier{},
		AddCityCodePhoneModifier{},
	}

	for i := 0; i < len(phones); i++ {
		_, _ = fmt.Scan(&phones[i])
		for _, m := range modifiers {
			phones[i] = m.Apply(phones[i])
		}
	}
	for i := 1; i < len(phones); i++ {
		if phones[i] == phones[0] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}


}
