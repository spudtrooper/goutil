package io

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestStringsFromFile(t *testing.T) {
	f := "TestStringsFromFile.txt"
	content := `one

two

three`
	if err := ioutil.WriteFile(f, []byte(content), 0755); err != nil {
		t.Fatalf("WriteFile(%q): %v", f, err)
	}
	defer os.Remove(f)

	stringChan, err := StringsFromFile(f)
	if err != nil {
		t.Fatalf("StringsFromFile: %v", err)
	}

	var stringList []string
	for s := range stringChan {
		log.Println(s)
		stringList = append(stringList, s)
	}
	if want, got := []string{"one", "", "two", "", "three"}, stringList; !reflect.DeepEqual(want, got) {
		t.Errorf("want != got: %+v != %v", want, got)
	}
}

func TestStringsFromFileSkipEmpty(t *testing.T) {
	f := "TestStringsFromFileSkipEmpty.txt"
	content := `one

two

three`
	if err := ioutil.WriteFile(f, []byte(content), 0755); err != nil {
		t.Fatalf("WriteFile(%q): %v", f, err)
	}
	defer os.Remove(f)

	stringChan, err := StringsFromFile(f, StringsFromFileSkipEmpty(true))
	if err != nil {
		t.Fatalf("StringsFromFile: %v", err)
	}

	var stringList []string
	for s := range stringChan {
		log.Println(s)
		stringList = append(stringList, s)
	}
	if want, got := []string{"one", "two", "three"}, stringList; !reflect.DeepEqual(want, got) {
		t.Errorf("want != got: %+v != %v", want, got)
	}
}

func TestStringsFromFileComments(t *testing.T) {
	f := "TestStringsFromFileSkipEmpty.txt"
	content := `one
#two
three`
	if err := ioutil.WriteFile(f, []byte(content), 0755); err != nil {
		t.Fatalf("WriteFile(%q): %v", f, err)
	}
	defer os.Remove(f)

	stringChan, err := StringsFromFile(f, StringsFromFileCommentStart("#"))
	if err != nil {
		t.Fatalf("StringsFromFile: %v", err)
	}

	var stringList []string
	for s := range stringChan {
		log.Println(s)
		stringList = append(stringList, s)
	}
	if want, got := []string{"one", "three"}, stringList; !reflect.DeepEqual(want, got) {
		t.Errorf("want != got: %+v != %v", want, got)
	}
}
