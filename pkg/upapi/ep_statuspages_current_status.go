package upapi

import (
	"context"
	"fmt"
	"net/http"
)

// StatusPageCurrentStatus represents the current operational status of a status page
type StatusPageCurrentStatus struct {
	StatusPage                                             // Embed StatusPage to inherit all its fields
	GlobalIsOperational bool                               `json:"global_is_operational"`
	ActiveIncidents     []interface{}                      `json:"active_incidents,omitempty"`
	UpcomingMaintenance []interface{}                      `json:"upcoming_maintenance,omitempty"`
	Components          []StatusPageCurrentStatusComponent `json:"components,omitempty"`
	Metrics             []interface{}                      `json:"metrics,omitempty"`
}

// StatusPageCurrentStatusComponent represents the current status of a component
type StatusPageCurrentStatusComponent struct {
	PK          int64  `json:"pk,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
}

// StatusPageCurrentStatusResponse wraps the current status response (API returns directly without wrapper)
type StatusPageCurrentStatusResponse StatusPageCurrentStatus

func (r StatusPageCurrentStatusResponse) Item() StatusPageCurrentStatus {
	return StatusPageCurrentStatus(r)
}

// StatusPageCurrentStatusEndpoint provides access to status page current status
type StatusPageCurrentStatusEndpoint interface {
	Get(context.Context) (*StatusPageCurrentStatus, error)
}

type statusPageCurrentStatusEndpointImpl struct {
	CBD
	endpoint string
}

// NewStatusPageCurrentStatusEndpoint creates a new status page current status endpoint
func NewStatusPageCurrentStatusEndpoint(cbd CBD, statusPagePK PrimaryKeyable) StatusPageCurrentStatusEndpoint {
	endpoint := fmt.Sprintf("statuspages/%d/current-status", statusPagePK.PrimaryKey())
	return &statusPageCurrentStatusEndpointImpl{
		CBD:      cbd,
		endpoint: endpoint,
	}
}

func (e *statusPageCurrentStatusEndpointImpl) Get(ctx context.Context) (*StatusPageCurrentStatus, error) {
	rq, err := e.BuildRequest(ctx, "GET", e.endpoint+"/", nil, nil)
	if err != nil {
		return nil, err
	}
	rs, err := e.Do(rq)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != http.StatusOK {
		return nil, ErrorFromResponse(rs)
	}
	var resp StatusPageCurrentStatusResponse
	if err := e.DecodeResponse(rs, &resp); err != nil {
		return nil, err
	}
	result := resp.Item()
	return &result, nil
}
