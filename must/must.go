package must

import (
	"io/ioutil"
	"log"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Check(err error) {
	check(err)
}

func ParseInt(s string, base, bits int) int64 {
	res, err := strconv.ParseInt(s, base, bits)
	check(err)
	return res
}

func Atoi(s string) int {
	res, err := strconv.Atoi(s)
	check(err)
	return res
}

func ReadAllFile(input string) string {
	b, err := ioutil.ReadFile(input)
	check(err)
	return string(b)
}
