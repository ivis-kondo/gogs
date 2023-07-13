package utils

import "strings"

func ContainsSpace(data string) bool {
	return ContainsFullWidthSpace(data) || ContainsHalfWidthSpace(data)
}

func ContainsFullWidthSpace(data string) bool {
	return strings.Contains(data, "　")
}

func ContainsHalfWidthSpace(data string) bool {
	return strings.Contains(data, " ")
}
