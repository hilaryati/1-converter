package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Структура для информации о валюте
type CurrencyInfo struct {
	Rate float64
	Name string
}

// Все данные о валютах в одной map (ключ - валюта, значение - курс и название)
var currencies = map[string]CurrencyInfo{
	"RUB": {1.0, "Рубли"},
	"USD": {76.0, "Доллары США"},
	"EUR": {90.0, "Евро"},
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

		if _, ok := currencies[input]; ok {
			return input
		}

		fmt.Println("Ошибка: такой валюты нет. Попробуйте ещё раз.")
	}
}

// Возвращает список доступных валют для подсказки
func getCurrencies() []string {
	result := make([]string, 0, len(currencies))
	for k := range currencies {
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
	fromRate := currencies[from].Rate
	toRate := currencies[to].Rate

	// Сначала переводим в базовую (RUB), потом в целевую
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

	// Расчёт и вывод результата с названиями валют
	result := calculate(amount, from, to)
	fmt.Printf("%.2f %s (%s) = %.2f %s (%s)\n",
		amount, from, currencies[from].Name,
		result, to, currencies[to].Name)
}
