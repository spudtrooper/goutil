package io

import (
	"os"
	"path"
)

func FileExists(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}
	return true
}

func MkdirAll(paths ...string) (string, error) {
	outDir := path.Join(paths...)
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return "", err
	}
	return outDir, nil
}
