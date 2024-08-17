package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	// Создаем регулярное выражение для удаления подстроки с пробелами
	re := regexp.MustCompile(`\s*` + regexp.QuoteMeta(stringTwo) + `\s*`)
	trimmedStr := re.ReplaceAllString(stringOne, " ")

	// Обрезаем лишние пробелы в начале и конце результата
	result := strings.TrimSpace(trimmedStr)

	// Возвращаем результат с ограничением длины до 40 символов
	return TrimStringsAfter40(result)
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

	// Используем регулярное выражение для корректного считывания строки с кавычками
	re := regexp.MustCompile(`^"(.+?)"\s*(\+|\-|\*|\/)\s*(".*?"|\d+)$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		panic("Некорректное выражение")
	}

	str1 := matches[1]
	operator := matches[2]
	str2 := matches[3]
	
	// Проверка длины первой строки
	if len(str1) > 10 {
		panic("Первая строка не может превышать 10 символов")
	}

	// Если вторая часть это строка в кавычках, удаляем кавычки
	if strings.HasPrefix(str2, "\"") && strings.HasSuffix(str2, "\"") {
		str2 = strings.Trim(str2, "\"")
	} else {
		// Если это не строка, проверяем, что это число от 1 до 10
		num, err := strconv.Atoi(str2)
		if err != nil || num < 1 || num > 10 {
			panic("Некорректное число, допустимы числа от 1 до 10")
		}
	}

	// Выполнение операции
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
		if strings.TrimSpace(input) == "" {
			fmt.Println("Ошибка! Введите валидное выражение!!!")
			continue
		}

		// Обработка паники
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Произошла ошибка:", r)
				}
			}()

			result := Calculate(input)
			fmt.Println(result)
		}()
	}
}
