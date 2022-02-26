package request

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spudtrooper/goutil/flags"
)

var (
	requestStats = flags.Bool("request_stats", "print verbose debugging of request timing")
	requestDebug = flags.Bool("request_debug", "print verbose debugging of requests")
)

type Param struct {
	Key string
	Val interface{}
}

type Params []Param

func (p Params) AddStringIfNotEmpty(key, val string) Params {
	if val != "" {
		return append(p, Param{key, val})
	}
	return p
}

func CreateRoute(base string, ps ...Param) string {
	if len(ps) == 0 {
		return base
	}
	var ss []string
	for _, p := range ps {
		s := fmt.Sprintf("%s=%s", p.Key, url.QueryEscape(fmt.Sprintf("%v", p.Val)))
		ss = append(ss, s)
	}
	return fmt.Sprintf("%s?%s", base, strings.Join(ss, "&"))
}

type Cookie struct {
	Key, Val string
}

type Response struct {
	Resp    *http.Response
	Data    []byte
	Cookies []Cookie
}

func Get(route string, result interface{}, rOpts ...RequestOption) (*Response, error) {
	return request("GET", route, result, nil, rOpts...)
}

// TODO: Move body to a RequestOption
func Post(url string, result interface{}, body io.Reader, rOpts ...RequestOption) (*Response, error) {
	return request("POST", url, result, body, rOpts...)
}

// TODO: Move body to a RequestOption
func Patch(url string, result interface{}, body io.Reader, rOpts ...RequestOption) (*Response, error) {
	return request("PATCH", url, result, body, rOpts...)
}

func Delete(url string, result interface{}, rOpts ...RequestOption) (*Response, error) {
	return request("DELETE", url, result, nil, rOpts...)
}

func request(method, url string, result interface{}, body io.Reader, rOpts ...RequestOption) (*Response, error) {
	opts := MakeRequestOptions(rOpts...)

	start := time.Now()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range opts.ExtraHeaders() {
		req.Header.Set(k, v)
	}
	if *requestDebug {
		log.Printf("requesting %s %s", method, url)
		if len(opts.ExtraHeaders()) > 0 {
			log.Printf("  with extra headers:")
			for k, v := range opts.ExtraHeaders() {
				log.Printf("    %s: %s", color.HiWhiteString(k), v)
			}
		}
		log.Printf("  headers:")
		for k, v := range req.Header {
			log.Printf("    %s: %s", color.HiWhiteString(k), v)
		}
		log.Printf("  body: %v", body)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	reqStop := time.Now()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	readStop := time.Now()

	if *requestStats {
		reqDur := reqStop.Sub(start)
		readDur := readStop.Sub(reqStop)
		totalDur := readStop.Sub(start)
		log.Printf("request stats: total:%v request:%v read:%v", totalDur, reqDur, readDur)
	}

	if *requestDebug {
		prettyJSON, err := PrettyPrintJSON(data)
		if err != nil {
			log.Printf("ignoring prettyPrintJSON error: %v", err)
			prettyJSON = string(data)
		}
		log.Printf("from url %q have response %s", url, prettyJSON)
		for k, vs := range resp.Header {
			for _, v := range vs {
				log.Printf("%s: %s", color.HiWhiteString(k), v)
			}
		}
	}

	if len(data) > 0 {
		if opts.CustomPayload() != nil {
			if err := json.Unmarshal(data, opts.CustomPayload()); err != nil {
				return nil, err
			}
		} else if result != nil {
			if err := json.Unmarshal(data, &result); err != nil {
				return nil, err
			}
			if *requestDebug {
				log.Printf("got response: %+v", result)
			}
		}
	}

	var cookies []Cookie
	for k, v := range resp.Header {
		if strings.ToLower(k) != "set-cookie" {
			continue
		}
		for _, cookie := range v {
			c := strings.SplitN(cookie, "; ", 1)[0]
			parts := strings.SplitN(c, "=", 2)
			k, v := parts[0], ""
			if len(parts) == 2 {
				v = parts[1]
			}
			cookies = append(cookies, Cookie{k, v})
		}
	}

	res := &Response{
		Resp:    resp,
		Data:    data,
		Cookies: cookies,
	}
	return res, nil
}
