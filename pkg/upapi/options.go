package upapi

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

type Option func(cbd CBD) (CBD, error)

func WithUserAgent(ua string) Option {
	return func(cbd CBD) (CBD, error) {
		return &withUserAgentCBD{cbd, ua}, nil
	}
}

type withUserAgentCBD struct {
	CBD
	userAgent string
}

func (c *withUserAgentCBD) BuildRequest(ctx context.Context, method string, endpoint string, args any, data any) (*http.Request, error) {
	req, err := c.CBD.BuildRequest(ctx, method, endpoint, args, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	return req, nil
}

func WithBaseURL(baseurl string) Option {
	return func(cbd CBD) (CBD, error) {
		base, err := url.Parse(baseurl)
		if err != nil {
			return nil, err
		}
		base = base.ResolveReference(&url.URL{Path: "./"})
		return &withBaseURLCBD{cbd, base}, nil
	}
}

type withBaseURLCBD struct {
	CBD
	base *url.URL
}

func (s *withBaseURLCBD) BuildRequest(ctx context.Context, method string, endpoint string, opts any, data any) (*http.Request, error) {
	req, err := s.CBD.BuildRequest(ctx, method, endpoint, opts, data)
	if err != nil {
		return nil, err
	}
	req.URL = s.base.ResolveReference(req.URL)
	return req, nil
}

func WithToken(token string) Option {
	return func(cbd CBD) (CBD, error) {
		return &withTokenCBD{cbd, token}, nil
	}
}

type withTokenCBD struct {
	CBD
	token string
}

func (s *withTokenCBD) BuildRequest(ctx context.Context, method string, endpoint string, opts any, data any) (*http.Request, error) {
	req, err := s.CBD.BuildRequest(ctx, method, endpoint, opts, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Token "+s.token)
	return req, nil
}

func WithRateLimit(rateLimit float64) Option {
	return func(cbd CBD) (CBD, error) {
		return &withRateLimitCBD{cbd, rate.NewLimiter(rate.Limit(rateLimit), 1)}, nil
	}
}

func WithRateLimitEvery(every time.Duration) Option {
	return func(cbd CBD) (CBD, error) {
		return &withRateLimitCBD{cbd, rate.NewLimiter(rate.Every(every), 1)}, nil
	}
}

type withRateLimitCBD struct {
	CBD
	limiter *rate.Limiter
}

func (s *withRateLimitCBD) Do(rq *http.Request) (*http.Response, error) {
	if err := s.limiter.Wait(rq.Context()); err != nil {
		return nil, err
	}
	return s.CBD.Do(rq)
}

func WithTrace(w io.Writer) Option {
	return func(cbd CBD) (CBD, error) {
		return &withTraceCBD{cbd, log.New(w, "››› ", log.LstdFlags)}, nil
	}
}

type withTraceCBD struct {
	CBD
	log *log.Logger
}

func (t *withTraceCBD) Do(rq *http.Request) (*http.Response, error) {
	rq = rq.WithContext(httptrace.WithClientTrace(rq.Context(), t.trace()))
	buf := bytes.NewBuffer(nil)
	if rq.Body != nil {
		rq.Body = io.NopCloser(io.TeeReader(rq.Body, buf))
	}
	t.log.Println(rq.Method, rq.URL.String())
	rs, err := t.CBD.Do(rq)
	if err != nil {
		return nil, err
	}
	if buf.Len() > 0 {
		t.log.Println("WroteRequestBody")
		s := bufio.NewScanner(buf)
		for s.Scan() {
			t.log.Println(" +", s.Text())
		}
		buf.Reset()
	}
	t.log.Println("GotResponseHeader", rs.StatusCode)
	_ = rs.Header.Write(buf)
	s := bufio.NewScanner(buf)
	for s.Scan() {
		t.log.Println(" +", s.Text())
	}
	t.log.Println("GotResponseBody")
	buf.Reset()
	s = bufio.NewScanner(io.TeeReader(rs.Body, buf))
	for s.Scan() {
		t.log.Println(" +", s.Text())
	}
	rs.Body = io.NopCloser(buf)
	return rs, nil
}

func (t *withTraceCBD) trace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			t.log.Println("DNSStart", info)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			t.log.Println("DNSDone", info)
		},
		TLSHandshakeStart: func() {
			t.log.Println("TLSHandshakeStart")
		},
		TLSHandshakeDone: func(_ tls.ConnectionState, err error) {
			if err != nil {
				t.log.Println("TLSHandshakeError", err)
			} else {
				t.log.Println("TLSHandshakeDone")
			}
		},
		WroteHeaderField: func(key string, value []string) {
			if key == "authorization" {
				for i := range value {
					if strings.HasPrefix(value[i], "Token ") {
						value[i] = "Token " + strings.Repeat("*", len(value[i])-6)
					}
				}
			}
			t.log.Println("WroteHeaderField", key, value)
		},
		WroteHeaders: func() {
			t.log.Println("WroteHeaders")
		},
		WroteRequest: func(_ httptrace.WroteRequestInfo) {
			t.log.Println("WroteRequest")
		},
	}
}

func WithRetry(limit int, maxDelay time.Duration, w io.Writer) Option {
	if w == nil {
		w = io.Discard
	}
	l := log.New(w, "••• ", 0)
	return func(cbd CBD) (CBD, error) {
		return &autoRetryCBD{cbd, l, limit, maxDelay}, nil
	}
}

type autoRetryCBD struct {
	CBD
	log           *log.Logger
	countdown     int
	retryAfterCap time.Duration
}

func (r *autoRetryCBD) Do(rq *http.Request) (*http.Response, error) {
	var bodyBytes []byte
	if rq.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(rq.Body)
		if err != nil {
			return nil, err
		}
		rq.Body.Close()
		rq.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	return r.doWithRetry(rq, bodyBytes)
}

func (r *autoRetryCBD) doWithRetry(rq *http.Request, bodyBytes []byte) (*http.Response, error) {
	rs, err := r.CBD.Do(rq)
	if err != nil {
		return nil, err
	}
	if rs.StatusCode != http.StatusTooManyRequests {
		return rs, nil
	}
	if r.countdown == 0 {
		return rs, nil
	}

	retryAfterSeconds, err := strconv.Atoi(rs.Header.Get("Retry-After"))
	if err != nil {
		return rs, nil
	}
	retryAfter := time.Duration(retryAfterSeconds) * time.Second

	r.log.Printf("server instructed us to retry in %s", retryAfter)
	if r.retryAfterCap > 0 && retryAfter > r.retryAfterCap {
		r.log.Printf("...but delay is capped to %s", r.retryAfterCap)
		retryAfter = r.retryAfterCap
	}

	wait, cancel := context.WithTimeout(rq.Context(), retryAfter)
	defer cancel()
	<-wait.Done()

	r.log.Printf("proceeding, %d attempts left", r.countdown)

	if bodyBytes != nil {
		rq.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	r.countdown--
	return r.doWithRetry(rq, bodyBytes)
}

type withSubaccountCBD struct {
	CBD
	subaccount int64
}

func (s *withSubaccountCBD) BuildRequest(ctx context.Context, method string, endpoint string, opts any, data any) (*http.Request, error) {
	req, err := s.CBD.BuildRequest(ctx, method, endpoint, opts, data)
	if err != nil {
		return nil, err
	}
	if s.subaccount > 0 {
		req.Header.Set("X-Subaccount", strconv.FormatInt(s.subaccount, 10))
	}
	return req, nil
}

func WithSubaccount(subaccount int64) Option {
	return func(cbd CBD) (CBD, error) {
		return &withSubaccountCBD{cbd, subaccount}, nil
	}
}
