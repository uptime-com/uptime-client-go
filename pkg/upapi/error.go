package upapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var DecodeError = errors.New("error response decode error")

// FieldErrors is the decoded `error_fields` payload. The server returns a
// flat `field -> [msg, ...]` map for top-level field errors and a nested
// `field -> {subfield: [msg, ...]}` object when the offending field is
// itself a structured value (e.g. `cloudstatusconfig.service_name`).
// UnmarshalJSON flattens nested objects into dotted-path keys so callers
// always see a flat `map[string][]any`.
type FieldErrors map[string][]any

func (f *FieldErrors) UnmarshalJSON(data []byte) error {
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	out := make(FieldErrors)
	flattenFieldErrors(out, "", raw)
	*f = out
	return nil
}

func flattenFieldErrors(dst FieldErrors, prefix string, src map[string]any) {
	for k, v := range src {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}
		switch t := v.(type) {
		case []any:
			dst[key] = append(dst[key], t...)
		case map[string]any:
			flattenFieldErrors(dst, key, t)
		default:
			dst[key] = append(dst[key], t)
		}
	}
}

type Error struct {
	Response *http.Response
	Code     string      `json:"error_code"`
	Message  string      `json:"error_message"`
	Fields   FieldErrors `json:"error_fields,omitempty"`
}

func NewError() *Error {
	return &Error{
		Fields: make(FieldErrors),
	}
}

func (e Error) Error() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%s %s failed: Code=%v Message=%s",
		e.Response.Request.Method,
		e.Response.Request.URL.String(),
		e.Code,
		e.Message,
	))
	if e.Fields != nil {
		b.WriteString(fmt.Sprintf(", Fields=%v", e.Fields))
	}
	return b.String()
}

func ErrorFromResponse(r *http.Response) error {
	data := new(struct {
		Error *Error `json:"messages"`
	})
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		if errors.Is(err, io.EOF) {
			return &Error{
				Response: r,
				Code:     strconv.Itoa(r.StatusCode),
				Message:  http.StatusText(r.StatusCode),
			}
		}
		return fmt.Errorf("%w: %s", DecodeError, err.Error())
	}
	data.Error.Response = r
	return data.Error
}
