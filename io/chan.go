package io

import (
	"bufio"
	"os"
	"strings"
)

func StringsFromFile(filepath string, sOpts ...StringsFromFileOption) (chan string, error) {
	opts := MakeStringsFromFileOptions(sOpts...)

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	res := make(chan string)
	go func() {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			s := scanner.Text()
			if opts.SkipEmpty() && s == "" {
				continue
			}
			if opts.CommentStart() != "" && strings.HasPrefix(s, opts.CommentStart()) {
				continue
			}
			res <- s
		}
		close(res)
	}()

	return res, nil
}
