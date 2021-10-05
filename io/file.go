// Package io contains functions dealing with IO.
package io

import (
	"io/ioutil"
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

func Copy(src string, dst string) error {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	// Write data to dst
	if err = ioutil.WriteFile(dst, data, 0644); err != nil {
		return err
	}
	return nil
}
