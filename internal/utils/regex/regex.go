package regex

import (
	"regexp"
	"strconv"
	"strings"
)

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
