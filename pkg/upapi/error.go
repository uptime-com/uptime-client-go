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

type Error struct {
	Response *http.Response
	Code     string           `json:"error_code"`
	Message  string           `json:"error_message"`
	Fields   map[string][]any `json:"error_fields,omitempty"`
}

func NewError() *Error {
	return &Error{
		Fields: make(map[string][]any),
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
