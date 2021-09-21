package io

import "os"

func FileExists(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}
	return true
}
