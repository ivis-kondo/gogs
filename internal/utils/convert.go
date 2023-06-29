package utils

import (
	"strconv"
	"strings"
	"unsafe"
)

func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func StringToBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{data, len(data)},
	))
}

func NumericStringToInt(data string) (int, error) {
	return strconv.Atoi(data)
}

func NumericIntToString(data int) string {
	return strconv.Itoa(data)
}

func RemoveAllHalfWidthSpace(data string) string {
	return strings.ReplaceAll(data, " ", "")
}

func RemoveAllFullWidthSpace(data string) string {
	return strings.ReplaceAll(data, "　", "")
}

func RemoveAllSpace(data string) string {
	return RemoveAllFullWidthSpace(RemoveAllHalfWidthSpace(data))
}
