package client

import (
	"github.com/tinyzimmer/go-cruise-control/pkg/types"
)

// GET Methods

func (c *Client) GetKafkaClusterState(opts *types.GetKafkaClusterStateRequest) (*types.GetKafkaClusterStateResponse, error) {
	res := &types.GetKafkaClusterStateResponse{}
	return res, c.basicRequest("GET", "kafka_cluster_state", opts, res)
}

func (c *Client) GetLoad(opts *types.GetLoadRequest) (*types.GetLoadResponse, error) {
	res := &types.GetLoadResponse{}
	return res, c.basicRequest("GET", "load", opts, res)
}

func (c *Client) GetPartitionLoad(opts *types.GetPartitionLoadRequest) (*types.GetPartitionLoadResponse, error) {
	res := &types.GetPartitionLoadResponse{}
	return res, c.basicRequest("GET", "partition_load", opts, res)
}

func (c *Client) GetState(opts *types.GetStateRequest) (*types.GetStateResponse, error) {
	res := &types.GetStateResponse{}
	return res, c.basicRequest("GET", "state", opts, res)
}

func (c *Client) GetUserTasks(opts *types.GetUserTasksRequest) (*types.GetUserTasksResponse, error) {
	res := &types.GetUserTasksResponse{}
	return res, c.basicRequest("GET", "user_tasks", opts, res)
}

func (c *Client) GetProposals(opts *types.GetProposalsRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest("GET", "proposals", opts, res)
}

// POST Methods

func (c *Client) Rebalance(opts *types.RebalanceRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest("POST", "rebalance", opts, res)
}

func (c *Client) AddBrokers(opts *types.AddBrokersRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest("POST", "add_broker", opts, res)
}

func (c *Client) RemoveBrokers(opts *types.RemoveBrokersRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest("POST", "remove_broker", opts, res)
}

func (c *Client) DemoteBrokers(opts *types.DemoteBrokersRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest("POST", "demote_broker", opts, res)
}

func (c *Client) FixOfflineReplicas(opts *types.FixOfflineReplicasRequest) (*types.ProposalsResponse, error) {
	res := &types.ProposalsResponse{}
	return res, c.basicRequest("POST", "fix_offline_replicas", opts, res)
}

func (c *Client) StopProposalExecution() (*types.GenericResponse, error) {
	res := &types.GenericResponse{}
	return res, c.basicRequest("POST", "stop_proposal_execution", nil, res)
}

func (c *Client) PauseSampling(opts *types.TriggerSamplingRequest) (*types.GenericResponse, error) {
	res := &types.GenericResponse{}
	return res, c.basicRequest("POST", "pause_sampling", opts, res)
}

func (c *Client) ResumeSampling(opts *types.TriggerSamplingRequest) (*types.GenericResponse, error) {
	res := &types.GenericResponse{}
	return res, c.basicRequest("POST", "resume_sampling", opts, res)
}
