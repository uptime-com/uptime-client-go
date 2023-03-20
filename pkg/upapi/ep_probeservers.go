package upapi

import "context"

type ProbeServer struct {
	Location  string `json:"location"`
	ProbeName string `json:"probe_name"`
	IPAddress string `json:"ip_address"`
}

type ProbeServerListResponse []ProbeServer

type ProbeServerListOptions struct{}

func (r ProbeServerListResponse) List() []ProbeServer {
	return r
}

type ProbeServersEndpoint interface {
	List(ctx context.Context) ([]ProbeServer, error)
}

func NewProbeServersEndpoint(cbd CBD) ProbeServersEndpoint {
	return &probeServersEndpointImpl{
		EndpointLister: NewEndpointLister[ProbeServerListResponse, ProbeServer, ProbeServerListOptions](cbd, "probe-servers"),
	}
}

type probeServersEndpointImpl struct {
	EndpointLister[ProbeServerListResponse, ProbeServer, ProbeServerListOptions]
}

func (p probeServersEndpointImpl) List(ctx context.Context) ([]ProbeServer, error) {
	return p.EndpointLister.List(ctx, ProbeServerListOptions{})
}
