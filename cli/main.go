// Package cli is a helper for using github.com/spudtrooper/goutil.
package cli

import (
	"context"
	"flag"
	"io/ioutil"

	"github.com/go-errors/errors"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/gitversion"
	minimalcli "github.com/spudtrooper/minimalcli/app"
)

var (
	curlFile       = flags.String("curl_file", "file containing curl command")
	curlString     = flags.String("curl_string", "string containing curl command")
	curlOutfile    = flags.String("curl_outfile", "file to which we dump the output of importing a curl command")
	curlRun        = flags.Bool("curl_run", "run generated file")
	curlBodyStruct = flags.Bool("curl_body_struct", "create the body string by generating a struct, creating an instance of this struct, and serializing this object to a string. The is flaky, so we'll always generate the plain string, and you can just comment out the struct or turn this flag off.")
	curlUnescape   = flag.Bool("curl_unescape", true, "unescape strings and place them in url.QueryEscape() calls, when possible")
)

func Main(ctx context.Context) error {
	app := minimalcli.Make()
	app.Init()

	if gitversion.CheckVersionFlag() {
		return nil
	}

	app.Register("CurlImport", func(context.Context) error {
		if *curlRun && *curlOutfile == "" {
			return errors.Errorf("doesn't make sense to set --curl_run without an outfile")
		}
		var s string
		if *curlString != "" {
			s = *curlString
		} else if *curlFile != "" {
			b, err := ioutil.ReadFile(*curlFile)
			if err != nil {
				return errors.Errorf("reading %s: %v", *curlFile, err)
			}
			s = string(b)
		}
		if s == "" {
			return errors.Errorf("required either --curl_file or --curl_string")
		}
		return curlImport(s, *curlOutfile, *curlRun, *curlBodyStruct, *curlUnescape)
	})

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}
