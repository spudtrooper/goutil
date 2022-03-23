package request

import "time"

//go:generate genopts --prefix=Request --outfile=requestoptions.go "extraHeaders:map[string]string" "host:string" "customPayload:interface{}" "proxyURL:string" "timeout:time.Duration"

type RequestOption func(*requestOptionImpl)

type RequestOptions interface {
	ExtraHeaders() map[string]string
	Host() string
	CustomPayload() interface{}
	ProxyURL() string
	Timeout() time.Duration
}

func RequestExtraHeaders(extraHeaders map[string]string) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.extraHeaders = extraHeaders
	}
}
func RequestExtraHeadersFlag(extraHeaders *map[string]string) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.extraHeaders = *extraHeaders
	}
}

func RequestHost(host string) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.host = host
	}
}
func RequestHostFlag(host *string) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.host = *host
	}
}

func RequestCustomPayload(customPayload interface{}) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.customPayload = customPayload
	}
}
func RequestCustomPayloadFlag(customPayload *interface{}) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.customPayload = *customPayload
	}
}

func RequestProxyURL(proxyURL string) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.proxyURL = proxyURL
	}
}
func RequestProxyURLFlag(proxyURL *string) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.proxyURL = *proxyURL
	}
}

func RequestTimeout(timeout time.Duration) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.timeout = timeout
	}
}
func RequestTimeoutFlag(timeout *time.Duration) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.timeout = *timeout
	}
}

type requestOptionImpl struct {
	extraHeaders  map[string]string
	host          string
	customPayload interface{}
	proxyURL      string
	timeout       time.Duration
}

func (r *requestOptionImpl) ExtraHeaders() map[string]string { return r.extraHeaders }
func (r *requestOptionImpl) Host() string                    { return r.host }
func (r *requestOptionImpl) CustomPayload() interface{}      { return r.customPayload }
func (r *requestOptionImpl) ProxyURL() string                { return r.proxyURL }
func (r *requestOptionImpl) Timeout() time.Duration          { return r.timeout }

func makeRequestOptionImpl(opts ...RequestOption) *requestOptionImpl {
	res := &requestOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRequestOptions(opts ...RequestOption) RequestOptions {
	return makeRequestOptionImpl(opts...)
}
