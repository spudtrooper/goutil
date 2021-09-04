// Package errors contains utilities to accumulate and join errors.
package errors

import (
	"fmt"
	"sync"
)

// ErrorCollector accumulates errors so they can be joined.
type ErrorCollector interface {
	// Accumulates `err` if it is not nil
	Add(err error)

	// Creates an error from all accumulated non-nil errors
	Build() error

	// Returns true if this has accumulated any non-nil errors yet
	Empty() bool
}

// MakeErrorCollector creates a ErrorCollector that is not thread safe.
func MakeErrorCollector() ErrorCollector {
	return &errorCollector{}
}

type errorCollector struct {
	errs []error
}

func (e *errorCollector) Add(err error) {
	if err != nil {
		e.errs = append(e.errs, err)
	}
}

func (e *errorCollector) Build() error {
	return joinErrors(e.errs...)
}

func (e *errorCollector) Empty() bool {
	return len(e.errs) == 0
}

// MakeSyncErrorCollector creates a ErrorCollector that is thread-safe.
func MakeSyncErrorCollector() ErrorCollector {
	return &threadSafeCollector{}
}

type threadSafeCollector struct {
	errs []error
	mu   sync.Mutex
}

func (e *threadSafeCollector) Add(err error) {
	if err != nil {
		e.mu.Lock()
		defer e.mu.Unlock()
		e.errs = append(e.errs, err)
	}
}

func (e *threadSafeCollector) Build() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	return joinErrors(e.errs...)
}

func (e *threadSafeCollector) Empty() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return len(e.errs) == 0
}

// https://stackoverflow.com/questions/33470649/combine-multiple-error-strings/33514248
// Probably overkill?
func joinErrors(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}
	var joinErrsR func(string, int, ...error) error
	joinErrsR = func(soFar string, count int, errs ...error) error {
		if len(errs) == 0 {
			if count == 0 {
				return nil
			}
			return fmt.Errorf(soFar)
		}
		current := errs[0]
		next := errs[1:]
		if current == nil {
			return joinErrsR(soFar, count, next...)
		}
		count++
		if count == 1 {
			return joinErrsR(fmt.Sprintf("%s", current), count, next...)
		} else if count == 2 {
			return joinErrsR(fmt.Sprintf("1: %s\n2: %s", soFar, current), count, next...)
		}
		return joinErrsR(fmt.Sprintf("%s\n%d: %s", soFar, count, current), count, next...)
	}
	return joinErrsR("", 0, errs...)
}
