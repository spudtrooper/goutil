// curlimport will imort a curl command into the goutil/request framework.
package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/slice"
)

var (
	// curl 'https://api.seatgeek.com/2/list?is_group=false&lat=40.7662&lon=-73.9862&tag=trending_events_list
	curlCmdRE = regexp.MustCompile(`^curl([^']*)'([^']+)'`)
	//  -H 'origin: https://seatgeek.com' \
	headerRE = regexp.MustCompile(`\s*-H '([^:]+): ([^']+)'`)
)

type uriParam struct{ key, val string }
type header struct{ key, val string }
type curlCmd struct {
	opts      []string
	uri       string
	uriParams []uriParam
	headers   []header
}

type renderedParam struct {
	Key, Val string
}

type rawParam struct {
	Key  string
	Val  interface{}
	Type rawParamType
}
type rawParamType string

const (
	rawParamTypeBool    rawParamType = "bool"
	rawParamTypeInt     rawParamType = "int"
	rawParamTypeFloat   rawParamType = "float"
	rawParamTypeComplex rawParamType = "complex"
	rawParamTypeString  rawParamType = "string"
)

func needsQueryEscape(s string) bool {
	return strings.Contains(s, "%2")
}

func createCurlCode(c curlCmd) (string, error) {
	if len(c.opts) > 0 {
		log.Printf("OOOOPS: can't support any options yet. You tried to specify the following options: %v", c.opts)
	}
	t := `
	uri := request.CreateRoute("{{.URI}}",
		{{range .QuotedURLParams}}request.Param{"{{.Key}}", ` + "`" + `{{.Val}}` + "`" + `},
		{{end}}{{range .QueryEscapedURLParams}}request.Param{"{{.Key}}", ` + "url.QueryEscape(`" + `{{.Val}}` + "`" + `)},
		{{end}}{{range .URLParams}}request.Param{"{{.Key}}", {{.Val}}},
		{{end}}
	)
	type cookiePart struct{ key, val string }
	var cookieParts = []cookiePart{
		{{range .Cookies}} { "{{.Key}}", ` + "`" + `{{.Val}}` + "`" + `},
		{{end}}{{range .QueryEscapedCookies}} { "{{.Key}}", ` + "url.QueryEscape(`" + `{{.Val}}` + "`)" + `},
		{{end}}
	}
	headers := map[string]string{
		{{range .Headers}}"{{.Key}}": ` + "`" + `{{.Val}}` + "`" + `,
		{{end}}
	}

	serializeCookie := func() string {
		if len(cookieParts) > 0 {
			var cs []string
			for _, c := range cookieParts {
				cs=append(cs, fmt.Sprintf("%s=%s", c.key, c.val))
			}
			return strings.Join(cs, "; ")
		}		
		return ""
	}
	if c := serializeCookie(); c != "" {
		headers["cookie"] = c
	}
	var payload interface{}
	res, err := request.Get(uri, &payload, request.RequestExtraHeaders(headers))
	check.Err(err)
	log.Printf("result: %+v", res)
	log.Printf("payload: %s", request.MustFormatString(payload))
	`
	var headers []renderedParam
	var cookies, queryEscapedCookies []renderedParam
	var urlParams, quotedUrlParams, queryEscapedUrlParams []rawParam
	const ()
	for _, x := range c.uriParams {
		v := x.val
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			urlParams = append(urlParams, rawParam{Key: x.key, Val: n, Type: rawParamTypeInt})
			continue
		}
		// Parse bools after ints, since 1 parses to true, probably 0 to false.
		if n, err := strconv.ParseBool(v); err == nil {
			urlParams = append(urlParams, rawParam{Key: x.key, Val: n, Type: rawParamTypeBool})
			continue
		}
		if n, err := strconv.ParseFloat(v, 64); err == nil {
			urlParams = append(urlParams, rawParam{Key: x.key, Val: n, Type: rawParamTypeFloat})
			continue
		}
		if n, err := strconv.ParseComplex(v, 64); err == nil {
			urlParams = append(urlParams, rawParam{Key: x.key, Val: n, Type: rawParamTypeComplex})
			continue
		}
		if needsQueryEscape(v) {
			unescaped, err := url.QueryUnescape(v)
			if err != nil {
				return "", errors.Errorf("url.QueryUnescape(%q): %v", v, err)
			}
			queryEscapedUrlParams = append(queryEscapedUrlParams, rawParam{Key: x.key, Val: unescaped, Type: rawParamTypeString})
		} else {
			quotedUrlParams = append(quotedUrlParams, rawParam{Key: x.key, Val: v, Type: rawParamTypeString})
		}
	}
	for _, h := range c.headers {
		if strings.ToLower(h.key) == "cookie" {
			for _, c := range strings.Split(h.val, "; ") {
				parts := strings.SplitN(c, "=", 2)
				var key, val string
				if len(parts) == 0 {
					continue
				}
				if len(parts) == 1 {
					key = parts[0]
				} else if len(parts) == 2 {
					key, val = parts[0], parts[1]
				} else {
					return "", errors.Errorf("unexpected cookie parts: %+v", parts)
				}
				if needsQueryEscape(val) {
					unescaped, err := url.QueryUnescape(val)
					if err != nil {
						return "", errors.Errorf("url.QueryUnescape(%q): %v", val, err)
					}
					queryEscapedCookies = append(queryEscapedCookies, renderedParam{key, unescaped})
				} else {
					cookies = append(cookies, renderedParam{key, val})
				}
			}
		} else {
			headers = append(headers, renderedParam{h.key, h.val})
		}
	}
	var data = struct {
		URI                   string
		Headers               []renderedParam
		Cookies               []renderedParam
		QueryEscapedCookies   []renderedParam
		URLParams             []rawParam
		QuotedURLParams       []rawParam
		QueryEscapedURLParams []rawParam
	}{
		URI:                   c.uri,
		Headers:               headers,
		Cookies:               cookies,
		QueryEscapedCookies:   queryEscapedCookies,
		URLParams:             urlParams,
		QuotedURLParams:       quotedUrlParams,
		QueryEscapedURLParams: queryEscapedUrlParams,
	}
	res, err := renderTemplate(t, "curl-code", data)
	if err != nil {
		return "", err
	}
	return res, nil
}

func createCurlMainCode(code string) (string, error) {
	t := `
	package main

	func main() {
		flag.Parse()
		{{.Code}}
	}
	`
	var data = struct {
		Code string
	}{
		Code: code,
	}
	res, err := renderTemplate(t, "curl-code", data)
	if err != nil {
		return "", err
	}
	return res, nil
}

func renderTemplate(t string, name string, data interface{}) (string, error) {
	tmpl, err := template.New(name).Parse(strings.TrimSpace(t))
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func curlImport(content, outfile string, run bool) error {
	c := curlCmd{}

	for _, line := range strings.Split(content, "\n") {
		if c.uri == "" {
			m := curlCmdRE.FindStringSubmatch(line)
			check.Check(len(m) == 3, check.CheckMessage(fmt.Sprintf("expecting match of length 3 and got: %+v for line: %s", m, line)))
			opts, uri := m[1], m[2]
			rawUri := strings.SplitN(uri, "?", 2)[0]
			c.uri = rawUri
			u, err := url.Parse(uri)
			if err != nil {
				return errors.Errorf("url.Parse(%q): %v", uri, err)
			}
			for _, p := range strings.Split(u.RawQuery, "&") {
				parts := strings.SplitN(p, "=", 2)
				var key, val string
				if len(parts) == 0 {
					continue
				}
				if len(parts) == 1 {
					key = parts[0]
				} else if len(parts) == 2 {
					key, val = parts[0], parts[1]
				} else {
					return errors.Errorf("unexpected parts: %+v", parts)
				}
				c.uriParams = append(c.uriParams, uriParam{key, val})
			}
			c.opts = slice.Strings(strings.TrimSpace(opts), " ")
			continue
		}
		if m := headerRE.FindStringSubmatch(line); len(m) == 3 {
			key, val := m[1], m[2]
			c.headers = append(c.headers, header{key, val})
		}
	}

	code, err := createCurlCode(c)
	if err != nil {
		return errors.Errorf("createCurlCode(%+v): %v", c, err)
	}
	if outfile != "" {
		mainCode, err := createCurlMainCode(code)
		if err != nil {
			return errors.Errorf("createCurlMainCode(%q): %v", code, err)
		}
		if err := ioutil.WriteFile(outfile, []byte(mainCode), 0755); err != nil {
			return errors.Errorf("writing main code to %s: %v", outfile, err)
		}
		if err := exec.Command("go", "fmt", outfile).Run(); err != nil {
			return errors.Errorf("formatting main code to %s: %v", outfile, err)
		}
		if err := exec.Command("goimports", "-w", outfile).Run(); err != nil {
			return errors.Errorf("fixing imports for main code in %s: %v", outfile, err)
		}
		log.Printf("wrote code to %s", outfile)

		if run {
			log.Printf("running %s", outfile)
			cmd := exec.Command("go", "run", outfile)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return errors.Errorf("running main code: %v", err)
			}

		}
	} else {
		fmt.Println(code)
	}
	return nil
}
