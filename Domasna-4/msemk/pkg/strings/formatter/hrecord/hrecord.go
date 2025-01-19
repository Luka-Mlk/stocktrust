package hrecfmt

import (
	"fmt"
	"strconv"
	"strings"
)

func EUDecimalToUSFromStr(str string) (float32, error) {
	// if str == "0" {
	// 	return 0, nil
	// }
	if str == "" {
		return 0, nil
	}
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, ",", ".")
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(value), nil
}

func FloatToStr(f float32) string {
	str := strconv.FormatFloat(float64(f), 'f', -1, 64)
	decimalIndex := strings.Index(str, ".")
	if decimalIndex == -1 {
		return formatInteger(str)
	}
	integerPart := str[:decimalIndex]
	fractionalPart := str[decimalIndex+1:]
	formattedInteger := formatInteger(integerPart)
	if len(fractionalPart) == 1 {
		fractionalPart += "0"
	} else if len(fractionalPart) > 2 {
		fractionalPart = fractionalPart[:2]
	}
	return formattedInteger + "," + fractionalPart
}

func formatInteger(s string) string {
	var formatted string
	for i := len(s); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		formatted = s[start:i] + "." + formatted
	}
	if len(formatted) > 0 && formatted[len(formatted)-1] == '.' {
		formatted = formatted[:len(formatted)-1]
	}
	return formatted
}

func FormatDate(date string) string {
	sepstr := strings.Split(date, ".")
	return fmt.Sprintf("%s-%s-%s", sepstr[2], sepstr[1], sepstr[0])
}
