package utils

import "regexp"

func CheckTelephoneFormat(tel string) bool {
	re := regexp.MustCompile(`(^0(\d{1}\-?\d{4}|\d{2}\-?\d{3}|\d{3}\-?\d{2}|\d{4}\-?\d{1})\-?\d{4}$|^0[5789]0\-?\d{4}\-?\d{4}$)`)
	return re.MatchString(tel)
}
