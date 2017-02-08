package enterprise_test

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/influxdata/chronograf"
	"github.com/influxdata/chronograf/enterprise"
)

type ControlClient struct {
	Cluster            *enterprise.Cluster
	ShowClustersCalled bool
}

func NewMockControlClient(addr string) *ControlClient {
	_, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}

	return &ControlClient{
		Cluster: &enterprise.Cluster{
			DataNodes: []enterprise.DataNode{
				enterprise.DataNode{
					HTTPAddr: addr,
				},
			},
		},
	}
}

func (cc *ControlClient) ShowCluster() (*enterprise.Cluster, error) {
	cc.ShowClustersCalled = true
	return cc.Cluster, nil
}

type TimeSeries struct {
	URLs     []string
	Response Response

	QueryCtr int
}

type Response struct{}

func (r *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (ts *TimeSeries) Query(ctx context.Context, q chronograf.Query) (chronograf.Response, error) {
	ts.QueryCtr++
	return &Response{}, nil
}

func (ts *TimeSeries) Connect(ctx context.Context, src *chronograf.Source) error {
	return nil
}

func NewMockTimeSeries(urls ...string) *TimeSeries {
	return &TimeSeries{
		URLs:     urls,
		Response: Response{},
	}
}
