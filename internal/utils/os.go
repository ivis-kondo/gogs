package utils

import "os"

func ExistFile(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
