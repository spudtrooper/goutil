package io

import (
	"fmt"
	"strings"
)

type stringOutput struct {
	buf []string
}

func NewStringOutput() *stringOutput { return &stringOutput{} }

func (s *stringOutput) Println(args ...any) {
	var c string
	for _, a := range args {
		c += fmt.Sprintf("%v", a)
	}
	s.buf = append(s.buf, c+"\n")
}

func (s *stringOutput) Printf(format string, args ...any) {
	s.buf = append(s.buf, fmt.Sprintf(format, args...))
}

func (s *stringOutput) String() string {
	return strings.Join(s.buf, "")
}

func (s *stringOutput) Clear() {
	s.buf = []string{}
}
