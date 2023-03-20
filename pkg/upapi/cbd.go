package upapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type RequestBuilder interface {
	BuildRequest(ctx context.Context, method string, endpoint string, args any, data any) (*http.Request, error)
}

type ResponseDecoder interface {
	DecodeResponse(rs *http.Response, data any) error
}

type Doer interface {
	Do(rq *http.Request) (*http.Response, error)
}

// CBD is http Client, request Build and response Decoder bundle.
type CBD interface {
	Doer
	RequestBuilder
	ResponseDecoder
}

type requestBuilderImpl struct{}

func (r *requestBuilderImpl) BuildRequest(ctx context.Context, method string, endpoint string, args any, data any) (*http.Request, error) {
	var body io.ReadCloser
	if data != nil {
		buf := bytes.NewBuffer(nil)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(data)
		if err != nil {
			return nil, err
		}
		body = io.NopCloser(buf)
	}
	rq, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return nil, err
	}
	if args != nil {
		qs, err := query.Values(args)
		if err != nil {
			return nil, err
		}
		rq.URL.RawQuery = qs.Encode()
	}
	rq.Header.Set("Content-Type", "application/json")
	return rq, nil
}

type responseDecoderImpl struct{}

func (r *responseDecoderImpl) DecodeResponse(rs *http.Response, data any) error {
	if data == nil {
		return nil
	}
	dec := json.NewDecoder(rs.Body)
	return dec.Decode(data)
}

func applyOptions(cbd CBD, opts ...Option) (_ CBD, err error) {
	for i := range opts {
		cbd, err = opts[i](cbd)
		if err != nil {
			return nil, err
		}
	}
	return cbd, nil
}
