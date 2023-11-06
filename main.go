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

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите выражение")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		tokens := strings.Fields(text)
		if len(tokens) != 3 {
			log.Fatal("ошибка: неверное выражение")
		}
		num1 := tokens[0]
		num2 := tokens[2]
		operand := tokens[1]
		strconv.Atoi(num1)
		var toNumber1, toNumber2 int
		var err1 error
		toNumber1, toNumber2, err1 = convert(num1, num2)
		if err1 != nil {
			result1 := convertFromRoman(num1)
			result2 := convertFromRoman(num2)
			if result1 != "0" && result2 != "0" {
				toNumber1, toNumber2, _ = convert(result1, result2)
				result := calculate(operand, toNumber1, toNumber2)
				if result <= 0 {
					log.Fatal("ошибка: результат не может быть выведен в римских числах")
				}
				fmt.Println(convertToRoman(result))
			} else {
				log.Fatal("ошибка: неправильный ввод")
			}
		} else {
			fmt.Println(calculate(operand, toNumber1, toNumber2))
		}
	}
}

func convert(num1 string, num2 string) (int, int, error) {
	toNumber1, err := strconv.Atoi(num1)
	if err != nil {
		return 0, 0, errors.New("ошибка:введенное число не арабское")
	}
	toNumber2, err := strconv.Atoi(num2)
	if err != nil {
		return 0, 0, errors.New("ошибка:введенное число не арабское")
	}
	if !((toNumber1 >= 1 && toNumber1 <= 10) && (toNumber2 >= 1 && toNumber2 <= 10)) {
		return 0, 0, errors.New("ошибка: значения не в диапазоне")
	}
	return toNumber1, toNumber2, nil
}

func convertFromRoman(num1 string) string {
	switch num1 {
	case "I":
		return "1"
	case "II":
		return "2"
	case "III":
		return "3"
	case "IV":
		return "4"
	case "V":
		return "5"
	case "VI":
		return "6"
	case "VII":
		return "7"
	case "VIII":
		return "8"
	case "IX":
		return "9"
	case "X":
		return "10"
	}
	return "0"
}

func convertToRoman(result int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var roman strings.Builder
	for _, conversion := range conversions {
		for result >= conversion.value {
			roman.WriteString(conversion.digit)
			result -= conversion.value
		}
	}
	return roman.String()
}

func calculate(operand string, toNumber1 int, toNumber2 int) int {
	switch operand {
	case "+":
		return (toNumber1 + toNumber2)
	case "-":
		return (toNumber1 - toNumber2)
	case "*":
		return (toNumber1 * toNumber2)
	case "/":
		return (toNumber1 / toNumber2)
	}
	return 101
}