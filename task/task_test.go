package task

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeBuilderEmpty(t *testing.T) {
	b := MakeBuilder()
	tasks := b.Build()
	if err := tasks.Go(); err != nil {
		t.Fatalf("Go(): %v", err)
	}
}

func TestMakeBuilder(t *testing.T) {
	var record []string
	add := func(s string) TaskFn {
		return func() error {
			record = append(record, s)
			return nil
		}
	}

	b := MakeBuilder()
	b.Add("one", add("one"))
	b.Add("two", add("two"))
	b.Add("tre", add("tre"))
	tasks := b.Build()
	if err := tasks.Go(); err != nil {
		t.Fatalf("Go(): %v", err)
	}
	if want, got := []string{"one", "two", "tre"}, record; !reflect.DeepEqual(want, got) {
		t.Errorf("Go(): want: %v got: %v", want, got)
	}
}

func TestMakeBuilderWithError(t *testing.T) {
	b := MakeBuilder()
	b.Add("one", func() error {
		return fmt.Errorf("error")
	})
	tasks := b.Build()
	if err := tasks.Go(); err == nil {
		t.Errorf("Go(): expected non-nil, got %v", err)
	}
}
