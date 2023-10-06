package arpscan

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var (
	arpScanRE = regexp.MustCompile(`^(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s+((?:[\da-fA-F]{2}:){5}[\da-fA-F]{2})\s+(.+)$`)
)

type Result struct {
	IP, MAC, Manufacturer string
}

func Scan() ([]Result, error) {
	var res []Result

	cmd := exec.Command("arp-scan", "--localnet")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run arp-scan: %v", err)
	}

	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		m := arpScanRE.FindStringSubmatch(line)
		if len(m) == 4 {
			if ip, mac, man := m[1], m[2], m[3]; ip != "" {
				res = append(res, Result{
					IP:           ip,
					MAC:          mac,
					Manufacturer: strings.TrimSpace(man),
				})
			}
		}
	}

	return res, nil
}
