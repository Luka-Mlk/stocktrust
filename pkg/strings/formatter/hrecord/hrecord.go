package hrecfmt

import (
	"fmt"
	"strconv"
	"strings"
)

func EUDecimalToUSFromStr(str string) (float32, error) {
	if str == "0" {
		return 0, nil
	}
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

func FormatDate(date string) string {
	sepstr := strings.Split(date, ".")
	return fmt.Sprintf("%s-%s-%s", sepstr[2], sepstr[1], sepstr[0])
}
