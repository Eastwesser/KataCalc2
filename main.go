package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TrimStringsAfter40 обрезает строку до 40 символов и добавляет '...', если эта строка длиннее 40 символов
func TrimStringsAfter40(s string) string {
	if len(s) > 40 {
		return s[:40] + "..."
	}
	return s
}

// Основная логика вычислений Add, Subtract, Multiply, Divide

func Add(stringOne, stringTwo string) string {
	return TrimStringsAfter40(stringOne + stringTwo)
}

func Subtract(stringOne, stringTwo string) string {
	if !strings.Contains(stringOne, stringTwo) {
		return TrimStringsAfter40(stringOne)
	}
	return TrimStringsAfter40(strings.ReplaceAll(stringOne, stringTwo, ""))
}

func Multiply(stringOne, stringTwo string) string {
	n, err := strconv.Atoi(stringTwo)
	if err != nil || n < 1 || n > 10 {
		panic("Некорректное число для умножения")
	}

	return TrimStringsAfter40(strings.Repeat(stringOne, n))
}

func Divide(stringOne, stringTwo string) string {
	n, err := strconv.Atoi(stringTwo)
	if err != nil || n < 1 || n > 10 {
		panic("Некорректное число для деления")
	}

	// Если длина строки меньше делителя, возвращаем пустую строку
	if len(stringOne) < n {
		return ""
	}

	// Длина одной части строки
	partLength := len(stringOne) / n

	// Возвращаем первую часть строки
	return TrimStringsAfter40(stringOne[:partLength])
}

// Calculate функция записывает и показывает результаты вычислений
func Calculate(input string) string {
	input = strings.TrimSpace(input)
	parts := strings.SplitN(input, " ", 3)

	if len(parts) != 3 {
		panic("Некорректное выражение")
	}

	str1 := parts[0]
	operator := parts[1]
	str2 := parts[2]

	// Проверяем, что строки находятся в кавычках
	if !(strings.HasPrefix(str1, "\"") && strings.HasSuffix(str1, "\"")) {
		panic("Первая строка должна быть заключена в кавычки")
	}

	str1 = strings.Trim(str1, "\"")

	// Проверка для второго аргумента, если это строка
	if strings.HasPrefix(str2, "\"") && strings.HasSuffix(str2, "\"") {
		str2 = strings.Trim(str2, "\"")
	} else {
		// Если это не строка, проверяем что это число от 1 до 10
		num, err := strconv.Atoi(str2)
		if err != nil || num < 1 || num > 10 {
			panic("Некорректное число, допустимы числа от 1 до 10")
		}
	}

	if len(str1) == 0 || len(str2) == 0 {
		panic("Некорректное выражение")
	}

	switch operator {
	case "+":
		return Add(str1, str2)
	case "-":
		return Subtract(str1, str2)
	case "*":
		return Multiply(str1, str2)
	case "/":
		return Divide(str1, str2)
	default:
		panic("Недопустимая операция")
	}
}

// Эта функция запускает наш калькулятор
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите выражение: ")

		scanner.Scan()
		input := scanner.Text()

		// Проверка на пустой ввод
		if strings.TrimSpace(input) == "" || !strings.Contains(input, " ") {
			fmt.Println("Ошибка! Введите валидное выражение!!!")
			continue
		}

		// Обработка паники
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Произошла ошибка:", r)
					os.Exit(1)
				}
			}()

			result := Calculate(input)
			fmt.Println(result)
		}()
	}
}
