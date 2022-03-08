// Binary to write MDB CSS and JS assets
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

type writer struct{}

func (w *writer) renderTemplate(t string, name string, data interface{}) ([]byte, error) {
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

func (w *writer) writeCSS() error {
	const (
		outfile = "html/cssassets.go"
	)
	var css bytes.Buffer
	stylesheets := []string{
		"third_party/mdb/css/mdb.lite.min.css",
		"third_party/mdb/css/bootstrap.min.css",
		"third_party/mdb/css/addons/datatables.min.css",
	}
	for _, f := range stylesheets {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}
		css.WriteString(string(b))
		css.WriteString("\n")
	}
	t := `
/*
DO NOT EDIT - Generated from:
{{range .Inputs}}
- {{.}}
{{end}}
*/
package html

// Base64 encoded JS assets
const cssAssets = ` + "`" + `{{.Encoded}}` + "`"
	enc := base64.StdEncoding.EncodeToString(css.Bytes())
	data := struct {
		Encoded string
		Inputs  []string
	}{
		Encoded: enc,
		Inputs:  stylesheets,
	}
	c, err := w.renderTemplate(t, "seleniumserver.go", data)
	if err != nil {
		return errors.Errorf("renderTemplate: %v", err)
	}
	if err := ioutil.WriteFile(outfile, c, 0755); err != nil {
		return errors.Errorf("ioutil.WriteFile(%q): %v", outfile, err)
	}
	return nil
}

func (w *writer) writeJS() error {
	const (
		outfile = "html/jsassets.go"
	)
	var js bytes.Buffer
	javascripts := []string{
		"third_party/mdb/js/jquery.min.js",
		"third_party/mdb/js/mdb.min.js",
		"third_party/mdb/js/bootstrap.min.js",
		"third_party/mdb/js/addons/datatables.min.js",
	}
	for _, f := range javascripts {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}
		js.WriteString(string(b))
		js.WriteString("\n")
	}
	t := `
/*
DO NOT EDIT - Generated from:
{{range .Inputs}}
- {{.}}
{{end}}
*/
package html

// Base64 encoded JS assets
const jsAssets = ` + "`" + `{{.Encoded}}` + "`"
	enc := base64.StdEncoding.EncodeToString(js.Bytes())
	data := struct {
		Encoded string
		Inputs  []string
	}{
		Encoded: enc,
		Inputs:  javascripts,
	}
	c, err := w.renderTemplate(t, "seleniumserver.go", data)
	if err != nil {
		return errors.Errorf("renderTemplate: %v", err)
	}
	if err := ioutil.WriteFile(outfile, c, 0755); err != nil {
		return errors.Errorf("ioutil.WriteFile(%q): %v", outfile, err)
	}
	return nil
}

func writeAssets() error {
	w := writer{}
	if err := w.writeCSS(); err != nil {
		return err
	}
	if err := w.writeJS(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := writeAssets(); err != nil {
		log.Fatalf("realMain: %v", err)
	}
}
