package main

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/go-errors/errors"
)

func renderTemplate(t string, name string, data interface{}) ([]byte, error) {
	tmpl, err := template.New(name).Parse(strings.TrimSpace(t))
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func writeJar() error {
	const (
		jar = "third_party/selenium/vendor/selenium-server.jar"
		out = "selenium/seleniumserver.go"
	)
	b, err := ioutil.ReadFile(jar)
	if err != nil {
		return errors.Errorf("ioutil.ReadFile(%q): %v", jar, err)
	}
	enc := base64.StdEncoding.EncodeToString(b)
	t := `
// DO NOT EDIT AUTOMATICALLY GENERATED: I should learn how to share data from a go module
package selenium

const seleniumServerJar = ` + "`" + `{{.Encoded}}` + "`"
	data := struct {
		Encoded string
	}{
		Encoded: string(enc),
	}
	c, err := renderTemplate(t, "seleniumserver.go", data)
	if err != nil {
		return errors.Errorf("renderTemplate: %v", err)
	}
	if err := ioutil.WriteFile(out, c, 0755); err != nil {
		return errors.Errorf("ioutil.WriteFile(%q): %v", out, err)
	}
	return nil
}

func main() {
	if err := writeJar(); err != nil {
		log.Fatalf("realMain: %v", err)
	}
}
