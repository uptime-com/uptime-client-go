package uptime

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL   = "https://uptime.com/api/v1/"
	defaultUserAgent = "go-uptime"
)

// Config represents the configuration for an Client.com client
type Config struct {
	BaseURL          string
	HTTPClient       *http.Client
	Token            string
	UserAgent        string
	RateMilliseconds int
}

// Client manages communication with the Uptime.com API.
type Client struct {
	client  *http.Client // HTTP client used to communicate with the API.
	common  service
	limiter <-chan time.Time

	// Base URL for API requests. BaseURL should always be specified
	// with a trailing slash.
	BaseURL   *url.URL
	UserAgent string
	Config    *Config

	Checks       *CheckService
	Outages      *OutageService
	Tags         *TagService
	Integrations *IntegrationService
}

type service struct {
	client *Client
}

// addOptions adds the parameters in options as URL query parameters to s.
// options must be a struct whose fields may contain "url" tags.
func addOptions(urlStr string, options interface{}) (string, error) {
	v := reflect.ValueOf(options)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return urlStr, nil
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr, err
	}

	queryString, err := query.Values(options)
	if err != nil {
		return urlStr, err
	}

	u.RawQuery = queryString.Encode()
	return u.String(), nil
}

// NewClient returns a new client for the Client.com API.
func NewClient(config *Config) (*Client, error) {
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}

	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}

	if config.Token == "" {
		err := errors.New("API token is required in Config.Token field")
		return nil, err
	}

	agent := defaultUserAgent
	if config.UserAgent != "" {
		agent = config.UserAgent
	}

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:    config.HTTPClient,
		BaseURL:   baseURL,
		Config:    config,
		UserAgent: agent,
	}
	c.common.client = c
	if config.RateMilliseconds == 0 {
		config.RateMilliseconds = 1000
	}
	c.limiter = time.Tick(time.Duration(config.RateMilliseconds) * time.Millisecond)

	c.Outages = (*OutageService)(&c.common)
	c.Checks = (*CheckService)(&c.common)
	c.Tags = (*TagService)(&c.common)
	c.Integrations = (*IntegrationService)(&c.common)

	return c, nil
}

// NewRequest creates an API request for Uptime.com. urlStr may be a relative URL (with
// NO preceding slash), which will be resolved relative to the BaseURL of the Client.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL %q must have a trailing slash", c.BaseURL)
	}
	targetURL, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err = enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, targetURL.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.Config.Token))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.Config.UserAgent != "" {
		req.Header.Add("User-Agent", c.Config.UserAgent)
	}
	return req, nil
}

// Do sends an API request and returns the response from Uptime.com. The API
// response is JSON decoded and stored in the value pointed to by v, or returned
// as an error if an API error has occurred.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	// Observe rate limit
	<-c.limiter

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}
	defer resp.Body.Close()

	if err := verifyResponse(resp); err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

func verifyResponse(r *http.Response) error {
	if 200 <= r.StatusCode && r.StatusCode < 300 {
		return nil
	}
	return decodeErrorResponse(r)
}

type errorResponse struct {
	Error *Error `json:"messages"`
}

type Error struct {
	Response    *http.Response
	ErrorCode   string              `json:"error_code"`
	Message     string              `json:"error_message"`
	ErrorFields map[string][]string `json:"error_fields,omitempty"`
	errors      bool                `json:"errors"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %s failed: Code=%v Message=%s",
		e.Response.Request.Method,
		e.Response.Request.URL.String(),
		e.ErrorCode,
		e.Message,
	)
}

func decodeErrorResponse(r *http.Response) error {
	e := &errorResponse{}
	err := json.NewDecoder(r.Body).Decode(e)
	if err != nil {
		fmt.Printf("Error decoding JSON response: %v\n", err)
		return fmt.Errorf("%s %s failed: %v", r.Request.Method, r.Request.URL.String(), r.Status)
	}
	e.Error.Response = r

	if e.Error.ErrorFields != nil {
		return fmt.Errorf("%s %s failed: Code=%v, Message=%s, Fields=%v",
			r.Request.Method,
			r.Request.URL.String(),
			e.Error.ErrorCode,
			e.Error.Message,
			e.Error.ErrorFields,
		)
	}

	return e.Error
}

// Int is a helper function that allocations a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper function that allocations a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }
