package regex

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func CheckAlphabet(value string) bool {
	re := regexp.MustCompile("[A-Za-z]+")
	return re.MatchString(value)
}

func CheckTelephoneFormat(tel string) bool {
	re := regexp.MustCompile(`(^0(\d{1}\-?\d{4}|\d{2}\-?\d{3}|\d{3}\-?\d{2}|\d{4}\-?\d{1})\-?\d{4}$|^0[5789]0\-?\d{4}\-?\d{4}$)`)
	return re.MatchString(tel)
}

// Check ORCID Format
// This code is based on [https://github.com/NII-DG/nii-dg/blob/main/nii_dg/utils.py check_orcid_id()]
func CheckORCIDFormat(value string) bool {

	re := regexp.MustCompile(`^(\d{4}-){3}\d{3}[\dX]$`)
	if !re.MatchString(value) {
		return false
	}

	var checksum int
	if value[len(value)-1:] == "X" {
		checksum = 10
	} else {
		num, _ := strconv.Atoi(value[len(value)-1:])
		checksum = int(num)
	}

	sum_val := 0

	v := strings.Replace(value, "-", "", -1)
	for _, num := range strings.Split(v[:len(v)-1], "") {
		n, _ := strconv.Atoi(num)
		sum_val = (sum_val + n) * 2
	}

	if (12-(sum_val%11))%11 != checksum {
		return false
	} else {
		return true
	}
}

// Check e-Rad researcher number Format
// This code is based on [https://github.com/NII-DG/nii-dg/blob/main/nii_dg/utils.py check_erad_researcher_number()]
func CheckERadRearcherNumberFormat(value string) bool {

	if len(value) != 8 {
		return false
	}

	check_digit, _ := strconv.Atoi(strings.Split(value, "")[0])
	sum_val := 0

	for i, num := range strings.Split(value, "") {
		if i == 0 {
			continue
		} else if i%2 == 0 {
			number, _ := strconv.Atoi(num)
			sum_val = sum_val + (number * 2)
		} else {
			number, _ := strconv.Atoi(num)
			sum_val = sum_val + number
		}
	}
	println(fmt.Sprintf("sum_val : %d, value : %s, check_digit : %d, remain : %d", sum_val, value, check_digit, (sum_val % 10)))
	if (sum_val % 10) != check_digit {
		return false
	}
	return true
}
