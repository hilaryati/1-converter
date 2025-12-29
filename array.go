package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите операцию (AVG, SUM, MED):")
	opLine, _ := reader.ReadString('\n')
	op := strings.TrimSpace(strings.ToUpper(opLine))

	fmt.Println("Введите числа через запятую (например: 2, 10, 9):")
	numLine, _ := reader.ReadString('\n')
	numLine = strings.TrimSpace(numLine)

	// Разбиваем строку по запятым → получаем slice строк
	strNumbers := strings.Split(numLine, ",")

	// Создаём пустой slice чисел с начальной ёмкостью
	numbers := make([]float64, 0, len(strNumbers))

	// Парсим каждую строку в число и добавляем в slice через append
	for _, strNum := range strNumbers {
		numStr := strings.TrimSpace(strNum)
		if num, err := strconv.ParseFloat(numStr, 64); err == nil {
			numbers = append(numbers, num)
		}
	}

	if len(numbers) == 0 {
		fmt.Println("Нет чисел для расчёта")
		return
	}

	var result float64
	switch op {
	case "AVG":
		result = avg(numbers)
	case "SUM":
		result = sum(numbers)
	case "MED":
		result = med(numbers)
	default:
		fmt.Println("Неизвестная операция. Используйте AVG, SUM или MED")
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}

func sum(nums []float64) float64 {
	s := 0.0
	for _, v := range nums {
		s += v
	}
	return s
}

func avg(nums []float64) float64 {
	return sum(nums) / float64(len(nums))
}

func med(nums []float64) float64 {
	// Создаём копию slice для сортировки (не меняем оригинал)
	cp := make([]float64, len(nums))
	copy(cp, nums)
	sort.Float64s(cp)

	n := len(cp)
	if n%2 == 1 {
		return cp[n/2]
	}
	return (cp[n/2-1] + cp[n/2]) / 2
}
