package paths

import "os"

func IsExist(path string) bool {
	if _, err := os.Lstat(path); err == nil {
		return true
	}
	return false
}

func IsNotExist(path string) bool {
	if _, err := os.Lstat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func IsLinkTargetExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

