package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// поддерживаемые валюты и курсы к "базовой" валюте (например, RUB)
var rates = map[string]float64{
	"RUB": 1,
	"USD": 76,
	"EUR": 90,
}

// Функция для чтения строки из stdin
func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Функция ввода и проверки валюты
func readCurrency(prompt string) string {
	for {
		fmt.Println("Доступные валюты:", strings.Join(getCurrencies(), ", "))
		input := strings.ToUpper(readInput(prompt))

		if _, ok := rates[input]; ok {
			return input
		}

		fmt.Println("Ошибка: такой валюты нет. Попробуйте ещё раз.")
	}
}

// Возвращает список доступных валют для подсказки
func getCurrencies() []string {
	result := make([]string, 0, len(rates))
	for k := range rates {
		result = append(result, k)
	}
	return result
}

// Функция ввода и проверки числа
func readAmount(prompt string) float64 {
	for {
		input := readInput(prompt)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка: введите число, например 100.5")
			continue
		}
		if value < 0 {
			fmt.Println("Ошибка: сумма должна быть неотрицательной")
			continue
		}
		return value
	}
}

// Функция расчёта: конвертация через базовую валюту (RUB)
func calculate(amount float64, from string, to string) float64 {
	fromRate := rates[from]
	toRate := rates[to]

	// сначала переводим в базовую (RUB), потом в целевую
	amountInBase := amount * fromRate
	result := amountInBase / toRate
	return result
}

func main() {
	fmt.Println("___Калькулятор валют___")

	// Шаг 1: ввод исходной валюты
	from := readCurrency("Введите исходную валюту: ")

	// Шаг 2: ввод суммы
	amount := readAmount("Введите сумму: ")

	// Шаг 3: ввод целевой валюты
	to := readCurrency("Введите целевую валюту: ")

	// Расчёт и вывод результата
	result := calculate(amount, from, to)
	fmt.Printf("Результат: %.2f %s из %.2f %s\n", result, to, amount, from)
}
