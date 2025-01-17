package validation

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateDate(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, fieldValue)
	if !match {
		return false
	}
	parsedDate, err := time.Parse("2006-01-02", fieldValue)
	if err != nil {
		return false
	}
	today := time.Now()
	if parsedDate.After(today) {
		return false
	}
	return true
}
