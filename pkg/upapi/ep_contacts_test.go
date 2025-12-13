package upapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContactCreate_BodyPreserved(t *testing.T) {
	var capturedBody []byte
	var capturedContentType string
	var capturedContentLength int64

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedContentType = r.Header.Get("Content-Type")
		capturedContentLength = r.ContentLength
		var err error
		capturedBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}

		// Check if body is valid JSON with expected fields
		var data map[string]any
		if err := json.Unmarshal(capturedBody, &data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"messages":{"errors":true,"error_code":"VALIDATION_ERROR","error_message":"Invalid JSON."}}`))
			return
		}

		if _, ok := data["name"]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"messages":{"errors":true,"error_code":"VALIDATION_ERROR","error_message":"One or more fields failed validation.","error_fields":{"name":["This field is required."]}}}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"results":{"pk":123,"name":"test-contact","email_list":["test@example.com"]}}`))
	}))
	defer ts.Close()

	// Test WITHOUT trace option
	t.Run("without trace", func(t *testing.T) {
		capturedBody = nil
		capturedContentLength = 0

		client, err := New(
			WithBaseURL(ts.URL+"/"),
			WithToken("test-token"),
		)
		require.NoError(t, err)

		contact := Contact{
			Name:      "test-contact",
			EmailList: []string{"test@example.com"},
		}

		created, err := client.Contacts().Create(context.Background(), contact)
		require.NoError(t, err)
		require.NotNil(t, created)

		// Verify Content-Length was set (this was the bug - it was 0 before the fix)
		require.Greater(t, capturedContentLength, int64(0), "Content-Length should be set")

		// Verify the body was received correctly
		var receivedData map[string]any
		err = json.Unmarshal(capturedBody, &receivedData)
		require.NoError(t, err, "body should be valid JSON: %s", string(capturedBody))
		require.Equal(t, "test-contact", receivedData["name"])
		require.Equal(t, "application/json", capturedContentType)
	})

	// Test WITH trace option
	t.Run("with trace", func(t *testing.T) {
		capturedBody = nil
		capturedContentLength = 0

		client, err := New(
			WithBaseURL(ts.URL+"/"),
			WithToken("test-token"),
			WithTrace(io.Discard),
		)
		require.NoError(t, err)

		contact := Contact{
			Name:      "test-contact",
			EmailList: []string{"test@example.com"},
		}

		created, err := client.Contacts().Create(context.Background(), contact)
		require.NoError(t, err)
		require.NotNil(t, created)

		// Verify Content-Length was set
		require.Greater(t, capturedContentLength, int64(0), "Content-Length should be set")

		// Verify the body was received correctly
		var receivedData map[string]any
		err = json.Unmarshal(capturedBody, &receivedData)
		require.NoError(t, err, "body should be valid JSON: %s", string(capturedBody))
		require.Equal(t, "test-contact", receivedData["name"])
		require.Equal(t, "application/json", capturedContentType)
	})

	// Test WITH trace AND retry options
	t.Run("with trace and retry", func(t *testing.T) {
		capturedBody = nil
		capturedContentLength = 0

		client, err := New(
			WithBaseURL(ts.URL+"/"),
			WithToken("test-token"),
			WithTrace(io.Discard),
			WithRetry(3, 0, nil),
		)
		require.NoError(t, err)

		contact := Contact{
			Name:      "test-contact",
			EmailList: []string{"test@example.com"},
		}

		created, err := client.Contacts().Create(context.Background(), contact)
		require.NoError(t, err)
		require.NotNil(t, created)

		// Verify Content-Length was set
		require.Greater(t, capturedContentLength, int64(0), "Content-Length should be set")

		// Verify the body was received correctly
		var receivedData map[string]any
		err = json.Unmarshal(capturedBody, &receivedData)
		require.NoError(t, err, "body should be valid JSON: %s", string(capturedBody))
		require.Equal(t, "test-contact", receivedData["name"])
		require.Equal(t, "application/json", capturedContentType)
	})
}
