package fetch

import (
	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func encodeResponse[T any](info js.Scope[T], res *fetch.Response) (js.Value[T], error) {
	return info.Constructor("Response").NewInstance(res)
}

func WindowOrWorkerGlobalScope_fetch[T any](info js.CallbackContext[T]) (js.Value[T], error) {
	win, err := codec.GetWindow(info)
	if err != nil {
		return nil, err
	}
	url, err := js.ConsumeArgument(info, "url", nil, codec.DecodeString)
	if err != nil {
		return nil, err
	}
	opts, err := js.ConsumeArgument(
		info,
		"options",
		codec.ZeroValue,
		decodeRequestInit,
	)
	if err != nil {
		return nil, err
	}
	f := fetch.New(win)
	req := f.NewRequest(url, opts...)
	return codec.EncodePromise(info, f.FetchAsync(req), encodeResponse)
}

func decodeRequestInit[T any](
	scope js.Scope[T], val js.Value[T],
) ([]fetch.RequestOption, error) {
	if js.IsUndefined(val) {
		return nil, nil
	}
	options := codec.Options[T, fetch.RequestOption]{
		"signal":  codec.OptDecoder[T](codec.DecodeNativeValue, fetch.WithSignal),
		"method":  codec.OptDecoder[T](codec.DecodeString, fetch.WithMethod),
		"body":    codec.OptDecoder[T](codec.DecodeRequestBody, fetch.WithBody),
		"headers": codec.OptDecoder[T](decodeHeadersInit, fetch.WithHeaders),
	}
	// These RequestInit options are accepted and parsed but otherwise ignored,
	// since this in-process browser does not perform real network requests
	// where they would matter (CORS, caching, redirects, etc.). Without this,
	// clients like htmx that always set mode/credentials would fail.
	for _, optName := range ignoredRequestOptions {
		options[optName] = func(js.Scope[T], js.Value[T]) (fetch.RequestOption, error) {
			return func(*fetch.Request) {}, nil
		}
	}
	return codec.DecodeOptions(scope, val, options)
}

var ignoredRequestOptions = []string{
	"referrer", "referrerPolicy",
	"mode", "credentials", "cache", "redirect", "integrity",
	"keepalive", "duplex", "priority", "window",
}
