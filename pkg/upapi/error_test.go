package upapi

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrorFromResponse(t *testing.T) {
	t.Run("internal server error", func(t *testing.T) {
		rs := http.Response{
			Status:     http.StatusText(http.StatusInternalServerError),
			StatusCode: http.StatusInternalServerError,
			Body: io.NopCloser(strings.NewReader(`{
				"messages": {
					"error_code":    "500",
					"error_message": "Internal Server Error",
					"errors":        true
				}
			}`)),
			Request: &http.Request{
				URL: &url.URL{},
			},
		}
		err := NewError()
		require.ErrorAs(t, ErrorFromResponse(&rs), &err)
		require.Equal(t, "500", err.Code)
		require.Equal(t, "Internal Server Error", err.Message)
	})
	t.Run("decode error", func(t *testing.T) {
		rs := http.Response{
			Body: io.NopCloser(strings.NewReader(`invalid-json`)),
		}
		err := ErrorFromResponse(&rs)
		require.ErrorIs(t, err, DecodeError)
	})
}
