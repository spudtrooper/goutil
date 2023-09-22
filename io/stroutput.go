package io

import (
	"fmt"
	"strings"
)

type stringOutput struct {
	buf []string
}

func NewStringOutput() *stringOutput { return &stringOutput{} }

func (s *stringOutput) Println() {
	s.buf = append(s.buf, "")
}

func (s *stringOutput) Printf(format string, args ...any) {
	s.buf = append(s.buf, fmt.Sprintf(format, args...))
}

func (s *stringOutput) String() string {
	return strings.Join(s.buf, "\n")
}

func (s *stringOutput) Clear() {
	s.buf = []string{}
}
