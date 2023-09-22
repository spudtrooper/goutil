package csv

import (
	"testing"
)

func TestSimpleWriter_Empty(t *testing.T) {
	w := NewSimpleWriter()
	w.Flush()
	if want, have := "", w.Done(); have != want {
		t.Errorf("Expected output:\n%s\n\nhave output:\n%s", want, have)
	}
}

func TestSimpleWriter_Write(t *testing.T) {
	w := NewSimpleWriter()
	check(t, w.Write([]string{"a", "b", "c"}))
	if want, have := "a,b,c\n", w.Done(); have != want {
		t.Errorf("Expected output:\n%s\n\nhave output:\n%s", want, have)
	}
}

func TestSimpleWriter_Write_many(t *testing.T) {
	w := NewSimpleWriter()
	check(t, w.Write([]string{"a", "b", "c"}))
	check(t, w.Write([]string{"d", "e", "f"}))
	check(t, w.Write([]string{"g", "h", "i"}))
	if want, have := "a,b,c\nd,e,f\ng,h,i\n", w.Done(); have != want {
		t.Errorf("Expected output:\n%s\n\nhave output:\n%s", want, have)
	}
}

func TestSimpleWriter_WriteAll(t *testing.T) {
	w := NewSimpleWriter()
	check(t, w.WriteAll([][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
	}))
	if want, have := "a,b,c\nd,e,f\ng,h,i\n", w.Done(); have != want {
		t.Errorf("Expected output:\n%s\n\nhave output:\n%s", want, have)
	}
}
func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
