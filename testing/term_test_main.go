// Test for term since stty won't work in a unit test.
package main

import (
	"fmt"
	"log"

	"github.com/spudtrooper/goutil/term"
)

func TestDetectTerminalSize() error {
	s, err := term.DetectTerminalSize()
	if err != nil {
		return fmt.Errorf("DetectTerminalSize unexpected error: %v", err)
	}
	if s.Height == 0 {
		return fmt.Errorf("DetectTerminalSize: expected Height >0 and got %d", s.Height)
	}
	if s.Width == 0 {
		return fmt.Errorf("DetectTerminalSize: expected Width >0 and got %d", s.Width)
	}
	return nil
}

func main() {
	if err := TestDetectTerminalSize(); err != nil {
		log.Fatalf("TestDetectTerminalSize: %v", err)
	}
	fmt.Println("ok  	github.com/spudtrooper/goutil/term")
}
