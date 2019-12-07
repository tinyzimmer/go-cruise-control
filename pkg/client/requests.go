package client

import (
	"net/http"

	"github.com/tinyzimmer/go-cruise-control/pkg/types"
)

const (
	endpointKafkaClusterState = "kafka_cluster_state"
	endpointLoad              = "load"
	endpointPartitionLoad     = "partition_load"
	endpointState             = "state"
	endpointUserTasks         = "user_tasks"
	endpointProposals         = "proposals"

	endpointRebalance             = "rebalance"
	endpointAddBroker             = "add_broker"
	endpointRemoveBroker          = "remove_broker"
	endpointDemoteBroker          = "demote_broker"
	endpointFixOfflineReplicas    = "fix_offline_replicas"
	endpointStopProposalExecution = "stop_proposal_execution"
	endpointPauseSampling         = "pause_sampling"
	endpointResumeSampling        = "resume_sampling"
)

// GET Methods

// GetKafkaClusterState queries the health of partitions on the cluster.
func (c *Client) GetKafkaClusterState(opts *types.GetKafkaClusterStateRequest) (*types.GetKafkaClusterStateResponse, error) {
	res := &types.GetKafkaClusterStateResponse{}
	return res, c.basicRequest(http.MethodGet, endpointKafkaClusterState, opts, res)
}

// GetLoad queries the kafka cluster load. Monitor state must be RUNNING for this
// method to succeed.
func (c *Client) GetLoad(opts *types.GetLoadRequest) (*types.GetLoadResponse, error) {
	res := &types.GetLoadResponse{}
	return res, c.basicRequest(http.MethodGet, endpointLoad, opts, res)
}

// GetPartitionLoad queries the partition load sorted by the resource specified
// in the request parameters.
func (c *Client) GetPartitionLoad(opts *types.GetPartitionLoadRequest) (*types.GetPartitionLoadResponse, error) {
	res := &types.GetPartitionLoadResponse{}
	return res, c.basicRequest(http.MethodGet, endpointPartitionLoad, opts, res)
}

// GetState queries the state of Cruise Control.
func (c *Client) GetState(opts *types.GetStateRequest) (*types.GetStateResponse, error) {
	res := &types.GetStateResponse{}
	return res, c.basicRequest(http.MethodGet, endpointState, opts, res)
}

// GetUserTasks returns a list of all active and completed tasks in cruise control.
func (c *Client) GetUserTasks(opts *types.GetUserTasksRequest) (*types.GetUserTasksResponse, error) {
	res := &types.GetUserTasksResponse{}
	return res, c.basicRequest(http.MethodGet, endpointUserTasks, opts, res)
}

// GetProposals returns the optimization proposals generated based on workload
// models for the provided time range.
func (c *Client) GetProposals(opts *types.GetProposalsRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest(http.MethodGet, endpointProposals, opts, res)
}

// POST Methods

// Rebalance requests cruise control to rebalance the kafka cluster.
func (c *Client) Rebalance(opts *types.RebalanceRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest(http.MethodPost, endpointRebalance, opts, res)
}

// AddBrokers will add the given brokers to the kafka cluster.
func (c *Client) AddBrokers(opts *types.AddBrokersRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest(http.MethodPost, endpointAddBroker, opts, res)
}

// RemoveBrokers will remove the given brokers from the kafka cluster.
func (c *Client) RemoveBrokers(opts *types.RemoveBrokersRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest(http.MethodPost, endpointRemoveBroker, opts, res)
}

// DemoteBrokers will remove any partition leadership from the given brokers.
func (c *Client) DemoteBrokers(opts *types.DemoteBrokersRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest(http.MethodPost, endpointDemoteBroker, opts, res)
}

// FixOfflineReplicas will move all offline replicas from dead brokers/disks.
// This method requires Kafka > 2.x.
func (c *Client) FixOfflineReplicas(opts *types.FixOfflineReplicasRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest(http.MethodPost, endpointFixOfflineReplicas, opts, res)
}

// StopProposalExecution will tell cruise control to stop an ongoing POST
// operation.
func (c *Client) StopProposalExecution() (*types.GenericResponse, error) {
	res := &types.GenericResponse{}
	return res, c.basicRequest(http.MethodPost, endpointStopProposalExecution, nil, res)
}

// PauseSampling will pause an ongoing metrics sampling process
func (c *Client) PauseSampling(opts *types.TriggerSamplingRequest) (*types.GenericResponse, error) {
	res := &types.GenericResponse{}
	return res, c.basicRequest(http.MethodPost, endpointPauseSampling, opts, res)
}

// ResumeSampling will resume a paused metrics sampling process.
func (c *Client) ResumeSampling(opts *types.TriggerSamplingRequest) (*types.GenericResponse, error) {
	res := &types.GenericResponse{}
	return res, c.basicRequest(http.MethodPost, endpointResumeSampling, opts, res)
}
