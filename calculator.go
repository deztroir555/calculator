package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

type Calculator struct{}

func (c *Calculator) Calculate(input string) (string, error) {
	var num1, num2, result int
	var operator string

	// Разбиваем строку на отдельные элементы
	items := strings.Split(input, " ")
	if len(items) != 3 {
		return "", errors.New("неверный формат запроса")
	}

	// Проверяем оператор
	operator = items[1]
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		return "", errors.New("неверный оператор")
	}

	// Проверяем, что числа в правильном диапазоне
	num1, err := parseNumber(items[0])
	if err != nil {
		return "", err
	}

	num2, err = parseNumber(items[2])
	if err != nil {
		return "", err
	}

	// Выполняем операции
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", errors.New("деление на ноль")
		}
		result = num1 / num2
	}

	// Возвращаем результат в нужном формате
	if isRoman(items[0]) && isRoman(items[2]) {
		if result <= 0 {
			return "", errors.New("результат меньше единицы")
		}

		return toRoman(result), nil
	} else if !isRoman(items[0]) && !isRoman(items[2]) {
		return strconv.Itoa(result), nil
	} else {
		return "", errors.New("нельзя работы с разными системами счисления")
	}
}

func parseNumber(input string) (int, error) {
	// Проверяем, что число в правильном диапазоне
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("неверное число")
	}

	if n < 1 || n > 10 {
		return 0, errors.New("число должно быть в диапазоне от 1 до 10")
	}

	return n, nil
}

func isRoman(input string) bool {
	_, ok := romanNumerals[input]
	return ok
}

func toRoman(input int) string {
	var result string

	for input > 0 {
		switch {
		case input >= 10:
			result += "X"
			input -= 10
		case input >= 9:
			result += "IX"
			input -= 9
		case input >= 5:
			result += "V"
			input -= 5
		case input >= 4:
			result += "IV"
			input -= 4
		default:
			result += "I"
			input--
		}
	}

	return result
}

func main() {
	calc := &Calculator{}

	for {
		fmt.Print("Введите выражение: ")
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		result, err := calc.Calculate(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}
}
