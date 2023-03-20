package upapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptrace"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWithAPIToken(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cbdm := new(cbdMock)
	cbdm.
		On("BuildRequest", ctx, http.MethodGet, "/", nil, nil).
		Return(http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)).
		Once()

	cbd, err := WithToken("example-token")(cbdm)
	require.NoError(t, err)

	rq, err := cbd.BuildRequest(ctx, http.MethodGet, "/", nil, nil)
	require.NoError(t, err)
	require.Equal(t, "Token example-token", rq.Header.Get("Authorization"))

	cbdm.AssertExpectations(t)
}

func TestWithUserAgent(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cbdm := new(cbdMock)
	cbdm.
		On("BuildRequest", ctx, http.MethodGet, "/", nil, nil).
		Return(http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)).
		Once()

	c, err := WithUserAgent("foo bar baz")(cbdm)
	require.NoError(t, err)

	req, err := c.BuildRequest(ctx, http.MethodGet, "/", nil, nil)
	require.NoError(t, err)
	require.Equal(t, "foo bar baz", req.Header.Get("User-Agent"))

	cbdm.AssertExpectations(t)
}

func TestWithBaseURL(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const (
		endpoint = "foo-bar-baz/"
		base     = "https://example.com/api/v1/irrelevant-bit"
		expect   = "https://example.com/api/v1/foo-bar-baz/"
	)

	cbdm := new(cbdMock)
	cbdm.
		On("BuildRequest", ctx, http.MethodGet, endpoint, nil, nil).
		Return(http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)).
		Once()

	c, err := WithBaseURL(base)(cbdm)
	require.NoError(t, err)

	req, err := c.BuildRequest(ctx, http.MethodGet, endpoint, nil, nil)
	require.NoError(t, err)
	require.Equal(t, expect, req.URL.String())

	cbdm.AssertExpectations(t)
}

func TestWithRateLimit(t *testing.T) {
	cbdm := new(cbdMock)
	cbdm.On("Do", mock.Anything).Return(&http.Response{}, nil).Times(10)

	cbd, err := WithRateLimit(100)(cbdm)
	require.NoError(t, err)

	rq, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	now := time.Now()
	for i := 0; i < 10; i++ {
		_, err = cbd.Do(rq)
		require.NoError(t, err)
	}
	require.GreaterOrEqual(t, time.Since(now), time.Millisecond*90)

	cbdm.AssertExpectations(t)
}

func TestWithRateLimitEvery(t *testing.T) {
	cbdm := new(cbdMock)
	cbdm.On("Do", mock.Anything).Return(&http.Response{}, nil).Times(10)

	cbd, err := WithRateLimitEvery(time.Millisecond * 10)(cbdm)
	require.NoError(t, err)

	rq, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	now := time.Now()
	for i := 0; i < 10; i++ {
		_, err = cbd.Do(rq)
		require.NoError(t, err)
	}
	require.GreaterOrEqual(t, time.Since(now), time.Millisecond*90)

	cbdm.AssertExpectations(t)
}

func TestWithHTTPClientTrace(t *testing.T) {
	cbdm := new(cbdMock)
	cbdm.
		On("Do", mock.Anything).
		Once().
		Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("")),
		}, nil).
		Run(func(args mock.Arguments) {
			rq := args.Get(0).(*http.Request)
			tr := httptrace.ContextClientTrace(rq.Context())
			require.NotNil(t, tr)
		})

	cbd, err := WithTrace(io.Discard)(cbdm)
	require.NoError(t, err)

	rq, _ := http.NewRequest(http.MethodGet, "/", nil)
	_, err = cbd.Do(rq)
	require.NoError(t, err)

	cbdm.AssertExpectations(t)
}
