package helpers

import (
	"aprian1337/thukul-service/utilities/constants"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func IsZero(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

func IsInt(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}

func IsDate(str string) bool {
	_, err := time.Parse(constants.BirthdayFormat, str)
	if err != nil {
		return false
	}
	return true
}
