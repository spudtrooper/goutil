package term

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type TermSize struct {
	Height, Width int
}

// https://stackoverflow.com/questions/16569433/get-terminal-size-in-go
func DetectTerminalSize() (*TermSize, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	output := strings.TrimSpace(string(out))
	parts := strings.Split(output, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid output from stty: %s", output)
	}
	height, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing height from %s: %v", parts[0], err)
	}
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing width from %s: %v", parts[1], err)
	}
	ts := &TermSize{Height: height, Width: width}
	return ts, nil
}
