package utils

import "os"

func ExistData(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
