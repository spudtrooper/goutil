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
	"github.com/spudtrooper/jsontogo/jsontogo"
)

var (
	// curl 'https://api.seatgeek.com/2/list?is_group=false&lat=40.7662&lon=-73.9862&tag=trending_events_list
	curlCmdRE = regexp.MustCompile(`^curl([^']*)'([^']+)'`)
	//  -H 'origin: https://seatgeek.com' \
	headerRE = regexp.MustCompile(`\s*-H '([^:]+): ([^']+)'`)
	//   --data-raw '{...}' \
	dataRawRE = regexp.MustCompile(`\s*--data-raw '([^']+)'`)
)

type uriParam struct{ key, val string }
type header struct{ key, val string }
type curlCmd struct {
	opts      []string
	uri       string
	uriParams []uriParam
	headers   []header
	data      string
}

type renderedParam struct {
	Key, Val string
}

type rawParamType string
type rawParam struct {
	Key  string
	Val  interface{}
	Type rawParamType
}
type rawParams struct {
	URLParams             []rawParam
	QueryEscapedURLParams []rawParam
	QuotedURLParams       []rawParam
}

const (
	rawParamTypeBool    rawParamType = "bool"
	rawParamTypeInt     rawParamType = "int"
	rawParamTypeFloat   rawParamType = "float"
	rawParamTypeComplex rawParamType = "complex"
	rawParamTypeString  rawParamType = "string"
)

func fillRawParams(k, v string, p *rawParams, unescape bool) error {
	if n, err := strconv.ParseInt(v, 10, 64); err == nil {
		p.URLParams = append(p.URLParams, rawParam{Key: k, Val: n, Type: rawParamTypeInt})
		return nil
	}
	// Parse bools after ints, since 1 parses to true, probably 0 to false.
	if n, err := strconv.ParseBool(v); err == nil {
		p.URLParams = append(p.URLParams, rawParam{Key: k, Val: n, Type: rawParamTypeBool})
		return nil
	}
	if n, err := strconv.ParseFloat(v, 64); err == nil {
		p.URLParams = append(p.URLParams, rawParam{Key: k, Val: n, Type: rawParamTypeFloat})
		return nil
	}
	if n, err := strconv.ParseComplex(v, 64); err == nil {
		p.URLParams = append(p.URLParams, rawParam{Key: k, Val: n, Type: rawParamTypeComplex})
		return nil
	}
	if unescape && needsQueryEscape(v) {
		unescaped, err := url.QueryUnescape(v)
		if err != nil {
			return err
		}
		p.QueryEscapedURLParams = append(p.QueryEscapedURLParams, rawParam{Key: k, Val: unescaped, Type: rawParamTypeString})
		return nil
	}
	p.QuotedURLParams = append(p.QuotedURLParams, rawParam{Key: k, Val: v, Type: rawParamTypeString})
	return nil
}

func needsQueryEscape(s string) bool {
	return strings.Contains(s, "%2")
}

func createCurlCode(c curlCmd, unescape bool) (string, error) {
	if len(c.opts) > 0 {
		log.Printf("OOOOPS: can't support any options yet. You tried to specify the following options: %v", c.opts)
	}
	t := `
	// Options
	printData := true
	printCookies := true
	printPayload := true

	// Data
	uri := request.MakeURL("{{.URI}}",
		{{range .URLParams.QuotedURLParams}}request.Param{"{{.Key}}", ` + "`" + `{{.Val}}` + "`" + `},
		{{end}}{{range .URLParams.QueryEscapedURLParams}}request.Param{"{{.Key}}", ` + "url.QueryEscape(`" + `{{.Val}}` + "`" + `)},
		{{end}}{{range .URLParams.URLParams}}request.Param{"{{.Key}}", {{.Val}}},
		{{end}}
	)
	cookie := [][2]string{
		{{range .Cookies}} { "{{.Key}}", ` + "`" + `{{.Val}}` + "`" + `},
		{{end}}{{range .QueryEscapedCookies}} { "{{.Key}}", ` + "url.QueryEscape(`" + `{{.Val}}` + "`)" + `},
		{{end}}
	}
	headers := map[string]string{
		{{range .Headers}}"{{.Key}}": ` + "`" + `{{.Val}}` + "`" + `,
		{{end}}
	}{{ if  and (not .DataParams.QuotedURLParams) (not .DataParams.QueryEscapedURLParams) (not .DataParams.URLParams) }}
	body := ` + "`" + `{{.Data}}` + "`" + `
	{{ if .SerializeBodyOject }}
	{
		{{.DataStruct}}
		bodyObject := {{.BodyObject}}
		body = string(request.MustJSONMarshal(bodyObject))
	}
	{{ end }}

	{{ else }}
	body := request.MakeRequestParams(
		{{range .DataParams.QuotedURLParams}}request.Param{"{{.Key}}", ` + "`" + `{{.Val}}` + "`" + `},
		{{end}}{{range .DataParams.QueryEscapedURLParams}}request.Param{"{{.Key}}", ` + "url.QueryEscape(`" + `{{.Val}}` + "`" + `)},
		{{end}}{{range .DataParams.URLParams}}request.Param{"{{.Key}}", {{.Val}}},
		{{end}}
	)
	{{ end }}

	// Make the request
	if len(cookie) > 0 {
		var cs []string
		for _, c := range cookie {
			cs=append(cs, fmt.Sprintf("%s=%s", c[0], c[1]))
		}
		if c:= strings.Join(cs, "; "); c != "" {
			headers["cookie"] = c
		}
	}
	
	var payload interface{}
	var res *request.Response
	var err error
	if body == "" {
		res, err = request.Get(uri, &payload, request.RequestExtraHeaders(headers))
	} else {
		res, err = request.Post(uri, &payload, strings.NewReader(body), request.RequestExtraHeaders(headers))
	}
	if printData {
		log.Printf("data: %s", string(res.Data))
	}
	if printCookies {
		log.Printf("cookies: %v", res.Cookies)
	}
	if printPayload {
		log.Printf("payload: %s", request.MustFormatString(payload))
	}
	check.Err(err)
	`

	var headers []renderedParam
	var cookies, queryEscapedCookies []renderedParam
	var urlParams rawParams

	for _, x := range c.uriParams {
		fillRawParams(x.key, x.val, &urlParams, unescape)
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

	isRawData := func(s string) bool {
		if s == "" {
			return true
		}
		if string(s[0]) == "{" {
			return true
		}
		return false
	}

	// Types could be laid out with declarations following uses. Since we're
	// adding this to funftion scope, we need to make sure declarations come
	// first.
	//
	// e.g.
	//
	// type BodyT struct {
	// 	 Variables Variables `json:"variables"`
	// }
	// type Variables struct {
	// 	 ID string `json:"id"`
	// }
	reverseTypeOrder := func(s string) string {
		types := []string{""}
		for _, line := range strings.Split(s, "\n") {
			types[len(types)-1] += line + "\n"
			if line == "}" {
				types = append(types, "")
			}
		}
		slice.Reverse(types)
		return strings.Join(types, "")
	}

	createBodyObject := func(dataStruct, data string) string {
		replace := func(word string) string {
			var parts []string
			for _, part := range strings.Split(word, "_") {
				parts = append(parts, strings.Title(part))
			}
			res := strings.Join(parts, "")
			res = strings.ReplaceAll(res, `"`, "")
			return res
		}
		keyRE := regexp.MustCompile(`"[^"]+":`)
		return keyRE.ReplaceAllStringFunc(data, replace)
	}

	var rawData string
	var dataParams rawParams
	var dataStruct string
	var bodyObject string

	if isRawData(c.data) {
		rawData = c.data
		ds, err := jsontogo.JSONToGo(c.data, "Body")
		if err != nil {
			return "", errors.Errorf("converting %q to a struct: %v", c.data, err)
		}
		dataStruct = reverseTypeOrder(ds)
		bodyObject = "Body" + createBodyObject(dataStruct, c.data)
	} else {
		for _, p := range strings.Split(c.data, "&") {
			parts := strings.SplitN(p, "=", 2)
			k, v := parts[0], ""
			if len(parts) == 2 {
				v = parts[1]
			}
			fillRawParams(k, v, &dataParams, unescape)
		}
	}

	var data = struct {
		URI                 string
		Headers             []renderedParam
		Cookies             []renderedParam
		QueryEscapedCookies []renderedParam
		URLParams           rawParams
		Data                string
		DataStruct          string
		DataParams          rawParams
		BodyObject          string
		SerializeBodyOject  bool
	}{
		URI:                 c.uri,
		Headers:             headers,
		Cookies:             cookies,
		QueryEscapedCookies: queryEscapedCookies,
		URLParams:           urlParams,
		Data:                rawData,
		DataStruct:          dataStruct,
		DataParams:          dataParams,
		BodyObject:          bodyObject,
		SerializeBodyOject:  *curlBodyStruct,
	}
	res, err := renderTemplate(t, "curl-code", data)
	if err != nil {
		return "", errors.Errorf("rendering template: %v", err)
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

func curlImport(content, outfile string, run, createBodyStruct, unescape bool) error {
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
				if key != "" {
					c.uriParams = append(c.uriParams, uriParam{key, val})
				}
			}
			c.opts = slice.Strings(strings.TrimSpace(opts), " ")
			continue
		}
		if m := headerRE.FindStringSubmatch(line); len(m) == 3 {
			key, val := m[1], m[2]
			c.headers = append(c.headers, header{key, val})
		}
		if m := dataRawRE.FindStringSubmatch(line); len(m) == 2 {
			data := m[1]
			data = strings.ReplaceAll(data, "\\\\n", "\\n")
			c.data = data
		}
	}

	code, err := createCurlCode(c, unescape)
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
