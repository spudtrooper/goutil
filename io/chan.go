package io

import (
	"bufio"
	"os"
	"strings"
	"sync"
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

func StringsFromFiles(filepaths []string, sOpts ...StringsFromFileOption) (chan string, chan error, error) {
	opts := MakeStringsFromFileOptions(sOpts...)

	res := make(chan string)
	errs := make(chan error)

	fs := make(chan string)
	go func() {
		for _, f := range filepaths {
			fs <- f
		}
		close(fs)
	}()

	go func() {
		var wg sync.WaitGroup
		for filepath := range fs {

			f, err := os.Open(filepath)
			if err != nil {
				errs <- err
				continue
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
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
			}()
		}

		wg.Wait()
		close(res)
		close(errs)
	}()

	return res, errs, nil
}
