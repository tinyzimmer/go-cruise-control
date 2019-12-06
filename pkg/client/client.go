package client

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/tinyzimmer/go-cruise-control/pkg/types"
)

// CruiseControlClient interface represents an object that carries out
// Cruise Control operations
type CruiseControlClient interface {
	// READ Methods
	GetState(*types.GetStateRequest) (*types.GetStateResponse, error)
	GetLoad(*types.GetLoadRequest) (*types.GetLoadResponse, error)
	GetPartitionLoad(*types.GetPartitionLoadRequest) (*types.GetPartitionLoadResponse, error)
	GetKafkaClusterState(*types.GetKafkaClusterStateRequest) (*types.GetKafkaClusterStateResponse, error)
	GetProposals(*types.GetProposalsRequest) (*types.ProposalsResponse, error)
	GetUserTasks(*types.GetUserTasksRequest) (*types.GetUserTasksResponse, error)

	// WRITE/UPDATE Methods
	Rebalance(*types.RebalanceRequest) (*types.ProposalsResponse, error)
	AddBrokers(*types.AddBrokersRequest) (*types.ProposalsResponse, error)
	RemoveBrokers(*types.RemoveBrokersRequest) (*types.ProposalsResponse, error)
	DemoteBrokers(*types.DemoteBrokersRequest) (*types.ProposalsResponse, error)
	FixOfflineReplicas(*types.FixOfflineReplicasRequest) (*types.ProposalsResponse, error)
	StopProposalExecution() (*types.GenericResponse, error)
	PauseSampling(*types.TriggerSamplingRequest) (*types.GenericResponse, error)
	ResumeSampling(*types.TriggerSamplingRequest) (*types.GenericResponse, error)
}

// Client implements CruiseControlClient
type Client struct {
	CruiseControlClient

	// The base URL for client operations in the format of <user-provided host>/kafkacruisecontrol
	baseURL *url.URL
	// The http client used for requests. It should include a cookiejar for session tracking.
	httpClient *http.Client
}

// New creates a new CruiseControlClient with the given options
func New(opts *types.ClientOptions) (CruiseControlClient, error) {
	// Cruise control REST docs recommend the use of cookies or recording request
	// UUIDS to avoid overloading the server with client sessions
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	url, err := url.Parse(fmt.Sprintf("%s/kafkacruisecontrol", strings.TrimSuffix(opts.URL, "/")))
	if err != nil {
		return nil, err
	}
	// Return a new client with the configured API path and cookie jar
	return &Client{
		baseURL:    url,
		httpClient: &http.Client{Jar: cookieJar},
	}, nil
}
