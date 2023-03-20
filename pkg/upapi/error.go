package upapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var DecodeError = errors.New("error response decode error")

type Error struct {
	Response *http.Response
	Code     string              `json:"error_code"`
	Message  string              `json:"error_message"`
	Fields   map[string][]string `json:"error_fields,omitempty"`
}

func NewError() *Error {
	return &Error{
		Fields: make(map[string][]string),
	}
}

func (e Error) Error() string {
	if e.Fields == nil {
		return fmt.Sprintf("%s %s failed: Code=%v Message=%s",
			e.Response.Request.Method,
			e.Response.Request.URL.String(),
			e.Code,
			e.Message,
		)
	} else {
		return fmt.Sprintf("%s %s failed: Code=%v Message=%s, Fields=%v",
			e.Response.Request.Method,
			e.Response.Request.URL.String(),
			e.Code,
			e.Message,
			e.Fields,
		)
	}
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
