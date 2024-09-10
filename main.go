package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	result, err := calculate(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func calculate(input string) (string, error) {
	input = strings.ToUpper(strings.TrimSpace(input)) // "3 + 1\n"
	parts := strings.Split(input, " ")                // "3 + 1"
	// parts = ["3","+","1"]

	if len(parts) != 3 {
		return "", errors.New("Выдача паники, так как формат математической операции не удовлетворяет заданию.")
	}

	a, b := parts[0], parts[2]
	operator := parts[1]

	isRoman := isRomanNumeral(a) && isRomanNumeral(b)    //false
	isArabic := isArabicNumeral(a) && isArabicNumeral(b) //true

	if isRoman && isArabic || (!isRoman && !isArabic) {
		return "", errors.New("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	if isRoman { //false
		return calculateRoman(a, b, operator)
	}
	return calculateArabic(a, b, operator)
}

func calculateRoman(a, b, operator string) (string, error) {
	num1, err := romanToInt(a)
	if err != nil {
		return "", err
	}
	num2, err := romanToInt(b)
	if err != nil {
		return "", err
	}

	result, err := performOperation(num1, num2, operator)
	if err != nil {
		return "", err
	}
	if result < 1 {
		return "", errors.New("Выдача паники, результат меньше единицы в римской системе")
	}
	return intToRoman(result), nil
}

func calculateArabic(a, b, operator string) (string, error) {
	num1, err := strconv.Atoi(a)
	if err != nil {
		return "", err
	}

	num2, err := strconv.Atoi(b)
	if err != nil {
		return "", err
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		return "", errors.New("Выдача паники, числа должны быть от 1 до 10 включительно.")
	}

	result, err := performOperation(num1, num2, operator)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(result), nil
}

func performOperation(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Выдача паники, деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("Выдача паники, неподдерживаемая операция.")
	}
}

func isRomanNumeral(s string) bool {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, roman := range romanNumerals {
		if s == roman {
			return true
		}
	}
	return false
}

func isArabicNumeral(number string) bool {
	_, err := strconv.Atoi(number)
	return err == nil
}

func romanToInt(s string) (int, error) {
	romanMap := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	if val, ok := romanMap[s]; ok {
		return val, nil
	}
	return 0, errors.New("Выдача паники, неправильный формат римского числа")
}

func intToRoman(num int) string {
	val := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var roman strings.Builder

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman.WriteString(syb[i])
		}
	}
	return roman.String()
}
