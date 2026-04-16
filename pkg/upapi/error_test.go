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
	t.Run("flat field errors", func(t *testing.T) {
		rs := http.Response{
			StatusCode: http.StatusBadRequest,
			Body: io.NopCloser(strings.NewReader(`{
				"messages": {
					"error_code":    "VALIDATION_ERROR",
					"error_message": "One or more fields failed validation.",
					"error_fields":  {"locations": ["Please select at least one location."]}
				}
			}`)),
			Request: &http.Request{URL: &url.URL{}},
		}
		err := NewError()
		require.ErrorAs(t, ErrorFromResponse(&rs), &err)
		require.Equal(t, []any{"Please select at least one location."}, err.Fields["locations"])
	})
	t.Run("nested field errors are flattened", func(t *testing.T) {
		rs := http.Response{
			StatusCode: http.StatusBadRequest,
			Body: io.NopCloser(strings.NewReader(`{
				"messages": {
					"error_code":    "VALIDATION_ERROR",
					"error_message": "One or more fields failed validation.",
					"error_fields":  {"cloudstatusconfig": {"service_name": ["Object with name=aws-ec2-us-east-1 does not exist."]}}
				}
			}`)),
			Request: &http.Request{URL: &url.URL{}},
		}
		err := NewError()
		require.ErrorAs(t, ErrorFromResponse(&rs), &err)
		require.Equal(t,
			[]any{"Object with name=aws-ec2-us-east-1 does not exist."},
			err.Fields["cloudstatusconfig.service_name"],
		)
	})
}
