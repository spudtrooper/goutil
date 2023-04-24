// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package request

import (
	"fmt"
	"time"
)

//go:generate genopts --prefix=Request --outfile=requestoptions.go "extraHeaders:map[string]string" "host:string" "customPayload:interface{}" "proxyURL:string" "timeout:time.Duration" "noRedirects" "allowedStatusCodes:[]int"

type RequestOption struct {
	f func(*requestOptionImpl)
	s string
}

func (o RequestOption) String() string { return o.s }

type RequestOptions interface {
	AllowedStatusCodes() []int
	HasAllowedStatusCodes() bool
	CustomPayload() interface{}
	HasCustomPayload() bool
	ExtraHeaders() map[string]string
	HasExtraHeaders() bool
	Host() string
	HasHost() bool
	NoRedirects() bool
	HasNoRedirects() bool
	ProxyURL() string
	HasProxyURL() bool
	Timeout() time.Duration
	HasTimeout() bool
}

func RequestAllowedStatusCodes(allowedStatusCodes []int) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		opts.has_allowedStatusCodes = true
		opts.allowedStatusCodes = allowedStatusCodes
	}, fmt.Sprintf("request.RequestAllowedStatusCodes([]int %+v)", allowedStatusCodes)}
}
func RequestAllowedStatusCodesFlag(allowedStatusCodes *[]int) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		if allowedStatusCodes == nil {
			return
		}
		opts.has_allowedStatusCodes = true
		opts.allowedStatusCodes = *allowedStatusCodes
	}, fmt.Sprintf("request.RequestAllowedStatusCodes([]int %+v)", allowedStatusCodes)}
}

func RequestCustomPayload(customPayload interface{}) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		opts.has_customPayload = true
		opts.customPayload = customPayload
	}, fmt.Sprintf("request.RequestCustomPayload(interface{} %+v)", customPayload)}
}
func RequestCustomPayloadFlag(customPayload *interface{}) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		if customPayload == nil {
			return
		}
		opts.has_customPayload = true
		opts.customPayload = *customPayload
	}, fmt.Sprintf("request.RequestCustomPayload(interface{} %+v)", customPayload)}
}

func RequestExtraHeaders(extraHeaders map[string]string) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		opts.has_extraHeaders = true
		opts.extraHeaders = extraHeaders
	}, fmt.Sprintf("request.RequestExtraHeaders(map[string]string %+v)", extraHeaders)}
}
func RequestExtraHeadersFlag(extraHeaders *map[string]string) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		if extraHeaders == nil {
			return
		}
		opts.has_extraHeaders = true
		opts.extraHeaders = *extraHeaders
	}, fmt.Sprintf("request.RequestExtraHeaders(map[string]string %+v)", extraHeaders)}
}

func RequestHost(host string) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		opts.has_host = true
		opts.host = host
	}, fmt.Sprintf("request.RequestHost(string %+v)", host)}
}
func RequestHostFlag(host *string) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		if host == nil {
			return
		}
		opts.has_host = true
		opts.host = *host
	}, fmt.Sprintf("request.RequestHost(string %+v)", host)}
}

func RequestNoRedirects(noRedirects bool) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		opts.has_noRedirects = true
		opts.noRedirects = noRedirects
	}, fmt.Sprintf("request.RequestNoRedirects(bool %+v)", noRedirects)}
}
func RequestNoRedirectsFlag(noRedirects *bool) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		if noRedirects == nil {
			return
		}
		opts.has_noRedirects = true
		opts.noRedirects = *noRedirects
	}, fmt.Sprintf("request.RequestNoRedirects(bool %+v)", noRedirects)}
}

func RequestProxyURL(proxyURL string) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		opts.has_proxyURL = true
		opts.proxyURL = proxyURL
	}, fmt.Sprintf("request.RequestProxyURL(string %+v)", proxyURL)}
}
func RequestProxyURLFlag(proxyURL *string) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		if proxyURL == nil {
			return
		}
		opts.has_proxyURL = true
		opts.proxyURL = *proxyURL
	}, fmt.Sprintf("request.RequestProxyURL(string %+v)", proxyURL)}
}

func RequestTimeout(timeout time.Duration) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		opts.has_timeout = true
		opts.timeout = timeout
	}, fmt.Sprintf("request.RequestTimeout(time.Duration %+v)", timeout)}
}
func RequestTimeoutFlag(timeout *time.Duration) RequestOption {
	return RequestOption{func(opts *requestOptionImpl) {
		if timeout == nil {
			return
		}
		opts.has_timeout = true
		opts.timeout = *timeout
	}, fmt.Sprintf("request.RequestTimeout(time.Duration %+v)", timeout)}
}

type requestOptionImpl struct {
	allowedStatusCodes     []int
	has_allowedStatusCodes bool
	customPayload          interface{}
	has_customPayload      bool
	extraHeaders           map[string]string
	has_extraHeaders       bool
	host                   string
	has_host               bool
	noRedirects            bool
	has_noRedirects        bool
	proxyURL               string
	has_proxyURL           bool
	timeout                time.Duration
	has_timeout            bool
}

func (r *requestOptionImpl) AllowedStatusCodes() []int       { return r.allowedStatusCodes }
func (r *requestOptionImpl) HasAllowedStatusCodes() bool     { return r.has_allowedStatusCodes }
func (r *requestOptionImpl) CustomPayload() interface{}      { return r.customPayload }
func (r *requestOptionImpl) HasCustomPayload() bool          { return r.has_customPayload }
func (r *requestOptionImpl) ExtraHeaders() map[string]string { return r.extraHeaders }
func (r *requestOptionImpl) HasExtraHeaders() bool           { return r.has_extraHeaders }
func (r *requestOptionImpl) Host() string                    { return r.host }
func (r *requestOptionImpl) HasHost() bool                   { return r.has_host }
func (r *requestOptionImpl) NoRedirects() bool               { return r.noRedirects }
func (r *requestOptionImpl) HasNoRedirects() bool            { return r.has_noRedirects }
func (r *requestOptionImpl) ProxyURL() string                { return r.proxyURL }
func (r *requestOptionImpl) HasProxyURL() bool               { return r.has_proxyURL }
func (r *requestOptionImpl) Timeout() time.Duration          { return r.timeout }
func (r *requestOptionImpl) HasTimeout() bool                { return r.has_timeout }

func makeRequestOptionImpl(opts ...RequestOption) *requestOptionImpl {
	res := &requestOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeRequestOptions(opts ...RequestOption) RequestOptions {
	return makeRequestOptionImpl(opts...)
}
