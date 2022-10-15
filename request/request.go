package request

import (
	"context"
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
	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/flags"
	goutiljson "github.com/spudtrooper/goutil/json"
	"github.com/spudtrooper/goutil/or"
)

var (
	requestStats        = flags.Bool("request_stats", "print verbose debugging of request timing")
	requestDebug        = flags.Bool("request_debug", "print verbose debugging of requests")
	requestDebugPayload = flags.Bool("request_debug_payload", "print verbose debugging of request payloads")
	readFromURLCache    = flags.Bool("request_read_from_url_cache", "read from the url cache")
	writeToURLCache     = flags.Bool("request_write_to_url_cache", "write to the url cache after every request")
	urlCachePort        = flags.Int("request_url_cache_port", "port for the url cache")
	urlCacheDBName      = flags.String("request_url_cache_db_name", "DB name for the url cache")
)

var globalURLCache *urlCache

type Param struct {
	Key string
	Val interface{}
}

// Params is an array of Param
type Params []Param

// ParamsBuilder allows building a Params
type ParamsBuilder interface {
	Add(key string, val interface{}) ParamsBuilder
	AddIfNotDefault(key string, val interface{}) ParamsBuilder
	Build() Params
}

// MakeParamsBuilder creates a new ParamsBuilder
func MakeParamsBuilder() ParamsBuilder {
	return &paramsBuilder{
		params: Params{},
	}
}

type paramsBuilder struct {
	params Params
}

func (b *paramsBuilder) Add(key string, val interface{}) ParamsBuilder {
	b.params = append(b.params, MakeParam(key, val))
	return b
}

func (b *paramsBuilder) AddIfNotDefault(key string, val interface{}) ParamsBuilder {
	b.params = b.params.AddIfNotDefault(key, val)
	return b
}

func (b *paramsBuilder) Build() Params {
	return b.params
}

// MakeParam creates a Param
func MakeParam(key string, val interface{}) Param {
	return Param{Key: key, Val: val}
}

// AddStringIfNotEmpty adds another params if val is not empty
func (p Params) AddStringIfNotEmpty(key, val string) Params {
	if val != "" {
		return append(p, Param{key, val})
	}
	return p
}

func isDefault(val interface{}) bool {
	switch v := val.(type) {
	case bool:
		return !v
	case int:
		return v == 0
	case uint8:
		return v == uint8(0)
	case uint16:
		return v == uint16(0)
	case uint32:
		return v == uint32(0)
	case uint64:
		return v == uint64(0)
	case int8:
		return v == int8(0)
	case int16:
		return v == int16(0)
	case int32:
		return v == int32(0)
	case int64:
		return v == int64(0)
	case float32:
		return v == float32(0)
	case float64:
		return v == float64(0)
	case string:
		return v == ""
	}
	return false
}

// AddIfNotDefault adds another params if val is not a default value
func (p Params) AddIfNotDefault(key string, val interface{}) Params {
	if !isDefault(val) {
		return append(p, Param{key, val})
	}
	return p
}

// Don't use, use MakeURL, this has a stupid name.
func CreateRoute(base string, ps ...Param) string {
	if len(ps) == 0 {
		return base
	}
	params := MakeRequestParams(ps...)
	return fmt.Sprintf("%s?%s", base, params)
}

func MakeURL(base string, ps ...Param) string {
	if len(ps) == 0 {
		return base
	}
	params := MakeRequestParams(ps...)
	return fmt.Sprintf("%s?%s", base, params)
}

func MakeRequestParams(ps ...Param) string {
	if len(ps) == 0 {
		return ""
	}
	var ss []string
	for _, p := range ps {
		s := fmt.Sprintf("%s=%s", p.Key, url.QueryEscape(fmt.Sprintf("%v", p.Val)))
		ss = append(ss, s)
	}
	return strings.Join(ss, "&")
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

// TODO: Move body to a RequestOption
func Put(url string, result interface{}, body io.Reader, rOpts ...RequestOption) (*Response, error) {
	return request("PUT", url, result, body, rOpts...)
}

func Delete(url string, result interface{}, rOpts ...RequestOption) (*Response, error) {
	return request("DELETE", url, result, nil, rOpts...)
}

func request(method, uri string, result interface{}, body io.Reader, rOpts ...RequestOption) (*Response, error) {
	if *readFromURLCache {
		resp, err := mustGetGlobalURLCache().FindRequest(context.Background(), uri)
		if err != nil {
			return nil, err
		}
		if resp != nil {
			log.Printf("responding with cached response")
			return resp, nil
		}
		log.Printf("no cached response")
	}

	opts := MakeRequestOptions(rOpts...)

	timeout := or.Duration(opts.Timeout(), 10*time.Second)
	start := time.Now()

	var client *http.Client
	if opts.ProxyURL() != "" {
		proxyUrl, err := url.Parse(opts.ProxyURL())
		if err != nil {
			return nil, errors.Errorf("parsing proxy: %s: %v", opts.ProxyURL(), err)
		}
		client = &http.Client{
			Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
			Timeout:   timeout,
		}
	} else {
		client = &http.Client{
			Timeout: timeout,
		}
	}
	if opts.NoRedirects() {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	for k, v := range opts.ExtraHeaders() {
		req.Header.Set(k, v)
	}
	if *requestDebug {
		log.Printf("requesting %s %s", method, uri)
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

	isAllowedStatusCode := func() bool {
		for _, code := range opts.AllowedStatusCodes() {
			if code == resp.StatusCode {
				return true
			}
		}
		if resp.StatusCode == http.StatusOK {
			return true
		}
		if resp.StatusCode == 302 /* TODO: use code */ {
			return true
		}
		return false
	}

	if !isAllowedStatusCode() {
		if *requestDebug {
			log.Printf("got non-OK status code: %d", resp.StatusCode)
		}
		return nil, errors.Errorf("request status code: %d", resp.StatusCode)
	}

	var data []byte

	if resp.StatusCode != 302 {
		data, err = ioutil.ReadAll(resp.Body)
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
	}

	var isHTML bool
	accepts := strings.Split(req.Header.Get("Accept"), ",")
	for _, a := range accepts {
		if a == "text/html" {
			isHTML = true
			break
		}
	}

	isJSON := true
	if isHTML {
		isJSON = false
	}

	if *requestDebug {
		if isJSON {
			prettyJSON, err := PrettyPrintJSON(data)
			if err != nil {
				log.Printf("ignoring prettyPrintJSON error: %v", err)
				prettyJSON = string(data)
			}
			log.Printf("from url %q have JSON response %s", uri, prettyJSON)
			for k, vs := range resp.Header {
				for _, v := range vs {
					log.Printf("%s: %s", color.HiWhiteString(k), v)
				}
			}
		}
		if isHTML {
			log.Printf("from url %q have HTML response %s", uri, string(data))
		}
	}

	if len(data) > 0 {
		if opts.CustomPayload() != nil {
			if isJSON {
				if err := json.Unmarshal(data, opts.CustomPayload()); err != nil {
					return nil, err
				}
			}
		} else if result != nil {
			if isJSON {
				if err := json.Unmarshal(data, &result); err != nil {
					return nil, err
				}
				if *requestDebug {
					log.Printf("got response: %+v", result)
				}
			}
		}
	}

	var cookies []Cookie
	for k, v := range resp.Header {
		if strings.ToLower(k) != "set-cookie" {
			continue
		}
		for _, cookie := range v {
			c := strings.SplitN(cookie, "; ", 2)[0]
			parts := strings.SplitN(c, "=", 2)
			k, v := parts[0], ""
			if len(parts) == 2 {
				v = parts[1]
			}
			cookies = append(cookies, Cookie{k, v})
		}
	}

	res := Response{
		Resp:    resp,
		Data:    data,
		Cookies: cookies,
	}

	if *writeToURLCache {
		if err := mustGetGlobalURLCache().SaveRequest(context.Background(), uri, res); err != nil {
			return nil, err
		}
	}

	if result != nil && *requestDebugPayload {
		s, err := goutiljson.ColorMarshal(result)
		if err != nil {
			return nil, err
		}
		log.Printf("payload: %+v", s)
	}

	return &res, nil
}

func mustGetGlobalURLCache() *urlCache {
	if globalURLCache == nil {
		log.Printf("connecting to global url cache")
		var opts []ConnectToURLCacheOption
		if *urlCachePort != 0 {
			opts = append(opts, ConnectToURLCachePort(*urlCachePort))
		}
		if *urlCacheDBName != "" {
			opts = append(opts, ConnectToURLCacheDbName(*urlCacheDBName))
		}
		cache, err := ConnectToURLCache(context.Background(), opts...)
		if err != nil {
			panic(err.Error())
		}
		globalURLCache = cache
		log.Printf("globalURLCache %+v", globalURLCache)
	}
	return globalURLCache
}

func CreateCookie(cookie [][2]string) string {
	var cs []string
	for _, c := range cookie {
		cs = append(cs, fmt.Sprintf("%s=%s", c[0], c[1]))
	}
	return strings.Join(cs, "; ")
}
