package request

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// set-cookie: at=%7B%22accountId%22%3A14723639%2C%22token%22%3A%22k19lvqorrtof1gfb1ujebcsteqatrlr6pvh3arip3815q9e0ihpj%22%7D;
var setCookieRE = regexp.MustCompile(`set-cookie: ([^=]+)=([^;]+);`)

func CurlCookies(uri string, headers map[string]string, body string) (map[string]string, error) {
	var args []string
	args = append(args, "-i")
	args = append(args, uri)
	for k, v := range headers {
		args = append(args, "-H", fmt.Sprintf("%s: %s", k, v))
	}
	args = append(args, "--data-raw", body)
	args = append(args, "--compressed")
	cmd := exec.Command("curl", args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	cookies := map[string]string{}
	for _, line := range strings.Split(stdout.String(), "\n") {
		if m := setCookieRE.FindStringSubmatch(line); len(m) == 3 {
			cookies[m[1]] = m[2]
		}
	}
	return cookies, nil
}
