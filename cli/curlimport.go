// curlimport will imort a curl command into the goutil/request framework.
package cli

import (
	"bytes"
	"flag"
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
	// Flags
	debug = flag.Bool("curl_import_debug", false, "Verbose logging")

	// Other

	// curl 'https://api.seatgeek.com/2/list?is_group=false&lat=40.7662&lon=-73.9862&tag=trending_events_list
	curlCmdRE = regexp.MustCompile(`^curl([^']*)'([^']+)'`)
	//  -H 'origin: https://seatgeek.com' \
	headerRE = regexp.MustCompile(`\s*-H '([^:]+): ([^']+)'`)
	//   --data-raw '{...}' \
	dataRawRE = regexp.MustCompile(`\s*--data-raw '([^']+)'`)
	//  -X 'PUT' \
	methodRE = regexp.MustCompile(`\s*-X '([^']+)'`)
)

type uriParam struct{ key, val string }
type header struct{ key, val string }
type curlCmd struct {
	opts      []string
	uri       string
	uriParams []uriParam
	headers   []header
	data      string
	method    string
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
	URLParams                 []rawParam
	QueryEscapedURLParamsVal  []rawParam
	QueryEscapedURLParamsKey  []rawParam
	QueryEscapedURLParamsBoth []rawParam
	QuotedURLParams           []rawParam
}

const (
	rawParamTypeBool    rawParamType = "bool"
	rawParamTypeInt     rawParamType = "int"
	rawParamTypeFloat   rawParamType = "float"
	rawParamTypeComplex rawParamType = "complex"
	rawParamTypeString  rawParamType = "string"
)

func addToRawParams(k, v string, p *rawParams, unescape bool) error {
	if n, err := strconv.ParseInt(v, 10, 64); err == nil {
		if needsQueryEscape(k) {
			unescapedKey, err := url.QueryUnescape(k)
			if err != nil {
				return err
			}
			p.QueryEscapedURLParamsKey = append(p.QueryEscapedURLParamsKey, rawParam{Key: unescapedKey, Val: v, Type: rawParamTypeInt})
			return nil
		}
		p.URLParams = append(p.URLParams, rawParam{Key: k, Val: n, Type: rawParamTypeInt})
		return nil
	}
	// Parse bools after ints, since 1 parses to true, probably 0 to false.
	if n, err := strconv.ParseBool(v); err == nil {
		if needsQueryEscape(k) {
			unescapedKey, err := url.QueryUnescape(k)
			if err != nil {
				return err
			}
			p.QueryEscapedURLParamsKey = append(p.QueryEscapedURLParamsKey, rawParam{Key: unescapedKey, Val: v, Type: rawParamTypeBool})
			return nil
		}
		p.URLParams = append(p.URLParams, rawParam{Key: k, Val: n, Type: rawParamTypeBool})
		return nil
	}
	if n, err := strconv.ParseFloat(v, 64); err == nil {
		if needsQueryEscape(k) {
			unescapedKey, err := url.QueryUnescape(k)
			if err != nil {
				return err
			}
			p.QueryEscapedURLParamsKey = append(p.QueryEscapedURLParamsKey, rawParam{Key: unescapedKey, Val: v, Type: rawParamTypeFloat})
			return nil
		}
		p.URLParams = append(p.URLParams, rawParam{Key: k, Val: n, Type: rawParamTypeFloat})
		return nil
	}
	if n, err := strconv.ParseComplex(v, 64); err == nil {
		if needsQueryEscape(k) {
			unescapedKey, err := url.QueryUnescape(k)
			if err != nil {
				return err
			}
			p.QueryEscapedURLParamsKey = append(p.QueryEscapedURLParamsKey, rawParam{Key: unescapedKey, Val: v, Type: rawParamTypeComplex})
			return nil
		}
		p.URLParams = append(p.URLParams, rawParam{Key: k, Val: n, Type: rawParamTypeComplex})
		return nil
	}
	if unescape {
		if needsQueryEscape(k) && needsQueryEscape(v) {
			unescapedKey, err := url.QueryUnescape(k)
			if err != nil {
				return err
			}
			unescapedVal, err := url.QueryUnescape(v)
			if err != nil {
				return err
			}
			p.QueryEscapedURLParamsBoth = append(p.QueryEscapedURLParamsBoth, rawParam{Key: unescapedKey, Val: unescapedVal, Type: rawParamTypeString})
			return nil
		}
		if needsQueryEscape(v) {
			unescapedVal, err := url.QueryUnescape(v)
			if err != nil {
				return err
			}
			p.QueryEscapedURLParamsVal = append(p.QueryEscapedURLParamsVal, rawParam{Key: k, Val: unescapedVal, Type: rawParamTypeString})
			return nil
		}
		if needsQueryEscape(k) {
			unescapedKey, err := url.QueryUnescape(k)
			if err != nil {
				return err
			}
			p.QueryEscapedURLParamsKey = append(p.QueryEscapedURLParamsKey, rawParam{Key: unescapedKey, Val: v, Type: rawParamTypeString})
			return nil
		}
	}
	p.QuotedURLParams = append(p.QuotedURLParams, rawParam{Key: k, Val: v, Type: rawParamTypeString})
	return nil
}

func needsQueryEscape(s string) bool {
	return strings.Contains(s, "%2")
}

func debugCurlCmd(c curlCmd) {
	var buf bytes.Buffer
	out := func(t string, a ...interface{}) {
		buf.WriteString(fmt.Sprintf(t, a...))
		buf.WriteString("\n")
	}
	out("")
	out("Debugging curl command...")
	log.Println(buf.String())
	out("%d opts", len(c.opts))
	for i, o := range c.opts {
		out(" [%3d] %s", i, o)
	}
	out("%d headers", len(c.headers))
	for i, o := range c.headers {
		out(" [%3d] %q = %q", i, o.key, o.val)
	}
	out("%d uriParams", len(c.uriParams))
	for i, o := range c.uriParams {
		out(" [%3d] %q = %q", i, o.key, o.val)
	}
	fmt.Println(buf.String())
}

func createCurlCode(c curlCmd, unescape bool) (string, error) {
	if *debug {
		debugCurlCmd(c)
	}
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
		{{end}}{{range .URLParams.QueryEscapedURLParamsVal}}request.Param{"{{.Key}}", ` + "url.QueryEscape(`" + `{{.Val}}` + "`" + `)},
		{{end}}{{range .URLParams.QueryEscapedURLParamsKey}}request.Param{` + "url.QueryEscape(`" + `{{.Key}}` + "`" + `)` + `, {{.Val}}},
		{{end}}{{range .URLParams.QueryEscapedURLParamsBoth}}request.Param{` + "url.QueryEscape(`" + `{{.Key}}` + "`" + `)` + `, ` + "url.QueryEscape(`" + `{{.Val}}` + "`" + `)},
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
		"cookie": request.CreateCookie(cookie),
	}{{ if  and (not .DataParams) }}
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
		{{range .DataParams}}request.Param{"{{.Key}}", ` + "`" + `{{.Val}}` + "`" + `},
		{{end}}
	)
	{{ end }}

	// Make the request
	var payload interface{}
	var res *request.Response
	var err error
	{{ if eq .Method "GET" }}
		res, err = request.Get(uri, &payload, request.RequestExtraHeaders(headers))
	{{ else if eq .Method "POST" }}
		res, err = request.Post(uri, &payload, strings.NewReader(body), request.RequestExtraHeaders(headers))
  {{ else if eq .Method "DELETE" }}
		res, err = request.Delete(uri, &payload, request.RequestExtraHeaders(headers))
	{{ else if eq .Method "PUT" }}
		res, err = request.Put(uri, &payload, strings.NewReader(body), request.RequestExtraHeaders(headers))
	{{ else }}
		if body == "" {
			res, err = request.Get(uri, &payload, request.RequestExtraHeaders(headers))
		} else {
			res, err = request.Post(uri, &payload, strings.NewReader(body), request.RequestExtraHeaders(headers))
		}
	{{ end }}
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
		addToRawParams(x.key, x.val, &urlParams, unescape)
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
	var dataParams []rawParam
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
			val, err := url.QueryUnescape(v)
			if err != nil {
				return "", errors.Errorf("unescaping %s: %v", v, err)
			}
			// TODO: Create a struct from the body if you can
			dataParams = append(dataParams, rawParam{
				Key: k,
				Val: val,
			})
			// addToRawParams(k, v, &dataParams, unescape)
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
		DataParams          []rawParam
		BodyObject          string
		SerializeBodyOject  bool
		Method              string
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
		Method:              c.method,
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

func parseCurlCmd(content string) (*curlCmd, error) {
	c := &curlCmd{}

	for _, line := range strings.Split(content, "\n") {
		if c.uri == "" {
			m := curlCmdRE.FindStringSubmatch(line)
			check.Check(len(m) == 3, check.CheckMessage(fmt.Sprintf("expecting match of length 3 and got: %+v for line: %s", m, line)))
			opts, uri := m[1], m[2]
			rawUri := strings.SplitN(uri, "?", 2)[0]
			c.uri = rawUri
			u, err := url.Parse(uri)
			if err != nil {
				return nil, errors.Errorf("url.Parse(%q): %v", uri, err)
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
					return nil, errors.Errorf("unexpected parts: %+v", parts)
				}
				if key != "" {
					c.uriParams = append(c.uriParams, uriParam{key, val})
				}
			}
			if arr := slice.Strings(strings.TrimSpace(opts), " "); len(arr) > 0 {
				c.opts = arr
			}
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
		if m := methodRE.FindStringSubmatch(line); len(m) == 2 {
			method := m[1]
			c.method = method
		}
	}

	if c.method == "" {
		if c.data == "" {
			c.method = "GET"
		} else {
			c.method = "POST"
		}
	}

	return c, nil
}

func curlImport(content, outfile string, run, createBodyStruct, unescape bool) error {
	c, err := parseCurlCmd(content)
	if err != nil {
		return errors.Errorf("parseCurlCmd: %v", err)
	}

	code, err := createCurlCode(*c, unescape)
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
