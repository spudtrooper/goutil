package io

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	existsFile := "exists"
	if err := ioutil.WriteFile(existsFile, []byte{}, 0755); err != nil {
		t.Fatalf("WriteFile(%q): %v", existsFile, err)
	}
	defer os.Remove(existsFile)

	if exists := FileExists(existsFile); !exists {
		t.Errorf("FileExists(%q) want: true got: false", existsFile)
	}

	doesntExistFile := "doesntExist"
	if exists := FileExists(doesntExistFile); exists {
		t.Errorf("FileExists(%q) want: false got: true", doesntExistFile)
	}
}

func TestCopy(t *testing.T) {
	src := "src"
	s := "some string"
	if err := ioutil.WriteFile(src, []byte(s), 0755); err != nil {
		t.Fatalf("WriteFile(%q): %v", src, err)
	}
	defer os.Remove(src)

	dst := "dst"
	if err := Copy(src, dst); err != nil {
		t.Fatalf("Copy(%q,%q): %v", src, dst, err)
	}

	b, err := ioutil.ReadFile(dst)
	if err != nil {
		t.Fatalf("ReadFile(%q): %v", dst, err)
	}

	if want, got := s, string(b); want != got {
		t.Errorf("Copy: want: %s got: %s", want, got)
	}
}
