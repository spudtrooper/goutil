package io

import (
	"io/ioutil"
	"os"
	"path"
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

func TestMkdirAll(t *testing.T) {
	dirs := "path/with/dirs/"
	dirs, err := MkdirAll("path", "with", "dirs")
	if err != nil {
		t.Fatalf("MkdirAll(%q): %v", dirs, err)
	}
	defer os.RemoveAll(dirs)

	if want, got := "path/with/dirs", dirs; want != got {
		t.Errorf("MkdirAll: want: %s got: %s", want, got)
	}

	f := path.Join(dirs, "f")
	if err := ioutil.WriteFile(f, []byte{}, 0755); err != nil {
		t.Errorf("ioutil.WriteFile: want: nil got: %v", err)
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

func TestWriteFile(t *testing.T) {
	dst := "path/with/dirs/dst"
	s := "some string"
	defer os.Remove(dst)
	if err := WriteFile(dst, []byte(s)); err != nil {
		t.Fatalf("WriteFile(%q,%q): %v", dst, s, err)
	}

	b, err := ioutil.ReadFile(dst)
	if err != nil {
		t.Fatalf("ReadFile(%q): %v", dst, err)
	}

	if want, got := s, string(b); want != got {
		t.Errorf("WriteFile: want: %s got: %s", want, got)
	}
}

func TestIsDir(t *testing.T) {
	dirs := "path/with/dirs/"
	if err := os.MkdirAll(dirs, 0755); err != nil {
		t.Fatalf("MkdirAll(%q): %v", dirs, err)
	}
	defer os.RemoveAll(dirs)

	if ok := IsDir(dirs); !ok {
		t.Errorf("IsDir(%q): want: true got false", dirs)
	}
}
