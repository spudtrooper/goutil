package net

import (
	"bufio"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// https://golangcode.com/download-a-file-from-a-url/
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func ReadURL(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	return ioutil.ReadAll(resp.Body)
}

//go:generate genopts --function=ReadURLChan "skipEmpty:bool" "commentStart:string"
func ReadURLChan(url string, optss ...ReadURLChanOption) (chan string, error) {
	opts := MakeReadURLChanOptions(optss...)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Body == nil {
		return nil, errors.Errorf("nil body")
	}

	res := make(chan string)
	go func() {
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
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
