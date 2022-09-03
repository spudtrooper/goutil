package or

import (
	"testing"
)

func TestAny(t *testing.T) {
	if want, got := 0, Any(0, 0); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", 0, 0, want, got)
	}
	if want, got := 1, Any(1, 0); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", 1, 0, want, got)
	}
	if want, got := 1, Any(1, 2); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", 1, 2, want, got)
	}

	if want, got := "", Any("", ""); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", "", "", want, got)
	}
	if want, got := "1", Any("1", ""); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", "1", "", want, got)
	}
	if want, got := "1", Any("1", "2"); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", "1", "2", want, got)
	}
}
