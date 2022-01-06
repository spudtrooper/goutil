package must

import (
	"io/ioutil"
	"strconv"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/io"
)

func chk(err error) {
	check.Err(err)
}

func Check(err error) {
	chk(err)
}

func ParseInt(s string, base, bits int) int64 {
	res, err := strconv.ParseInt(s, base, bits)
	chk(err)
	return res
}

func Atoi(s string) int {
	res, err := strconv.Atoi(s)
	chk(err)
	return res
}

func ReadAllFile(input string) string {
	b, err := ioutil.ReadFile(input)
	chk(err)
	return string(b)
}

func ReadLines(input string) []string {
	lines, err := io.ReadLines(input)
	chk(err)
	return lines
}
