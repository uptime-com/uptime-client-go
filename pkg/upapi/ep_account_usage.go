package upapi

import (
	"context"
	"encoding/json"
	"net/http"
)

// AccountUsage represents account usage and plan limits as returned by the API.
// The API returns a single-element array with one object containing heterogeneous
// fields (strings, numbers, bools, arrays), so we use a map to preserve all data.
type AccountUsage map[string]any

// AccountUsageEndpoint provides access to account usage and plan limits.
type AccountUsageEndpoint interface {
	Get(ctx context.Context) (*AccountUsage, error)
}

func NewAccountUsageEndpoint(cbd CBD) AccountUsageEndpoint {
	return &accountUsageEndpointImpl{cbd: cbd}
}

type accountUsageEndpointImpl struct {
	cbd CBD
}

func (e *accountUsageEndpointImpl) Get(ctx context.Context) (*AccountUsage, error) {
	req, err := e.cbd.BuildRequest(ctx, "GET", "auth/account-usage/", nil, nil)
	if err != nil {
		return nil, err
	}
	rs, err := e.cbd.Do(req)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != http.StatusOK {
		return nil, ErrorFromResponse(rs)
	}
	var items []AccountUsage
	if err := json.NewDecoder(rs.Body).Decode(&items); err != nil {
		return nil, err
	}
	if len(items) == 0 {
		empty := AccountUsage{}
		return &empty, nil
	}
	return &items[0], nil
}
