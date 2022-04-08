package must

import (
	"encoding/json"
	"io/fs"
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

func WriteFile(f string, b []byte) {
	chk(io.WriteFile(f, b))
}

func Marshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	chk(err)
	return b
}

func Unmarshal(b []byte, v interface{}) {
	err := json.Unmarshal(b, v)
	chk(err)
}

func ReadDir(dirname string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dirname)
	chk(err)
	return files
}
