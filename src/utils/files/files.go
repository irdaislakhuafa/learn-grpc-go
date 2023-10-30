package files

import "os"

func IsFileExist(pathName string) bool {
	if _, err := os.Stat(pathName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
