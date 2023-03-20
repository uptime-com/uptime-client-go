package upapi

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type cbdMock struct {
	mock.Mock
	Doer
	RequestBuilder
	ResponseDecoder
}

func (m *cbdMock) Do(rq *http.Request) (*http.Response, error) {
	args := m.Called(rq)
	if err := args.Error(1); err != nil {
		return nil, err
	}
	return args.Get(0).(*http.Response), nil
}

func (m *cbdMock) BuildRequest(ctx context.Context, method, path string, arg, data any) (*http.Request, error) {
	args := m.Called(ctx, method, path, arg, data)
	if err := args.Error(1); err != nil {
		return nil, err
	}
	return args.Get(0).(*http.Request), nil
}

func (m *cbdMock) DecodeResponse(r *http.Response, v interface{}) error {
	args := m.Called(r, v)
	return args.Error(0)
}
