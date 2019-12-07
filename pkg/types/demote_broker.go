package types

import "errors"

// DemoteBrokersRequest represents the parameters for a demote_broker request.
type DemoteBrokersRequest struct {
	// BrokerIDs is a list of ids of new broker(s) to demote.
	BrokerIDs []int `param:"brokerid"`
	// BrokerIDsAndLogDirs is a list of broker id and logdir pairs to be demoted
	// in the cluster. The format of a single entry is `id-logdir`.
	BrokerIDsAndLogDirs []string `param:"brokerid_and_logdirs"`
	// DryRun will just report what would have happened without doing anything.
	DryRun bool `param:"dryrun"`
	// AllowCapacityEstimation specifies whether to allow broker capacity to be
	// estimated from other brokers in the cluster.
	AllowCapacityEstimation bool `param:"allow_capacity_estimation"`
	// ConcurrentLeaderMovements specifies the upper bound of ongoing leadership
	// movements.
	ConcurrentLeaderMovements int `param:"concurrent_leader_movements"`
	// SkipUrpDemotion specifies whether to skip demoting leader replicas for
	// under replicated partitions.
	SkipUrpDemotion bool `param:"skip_urp_demotion"`
	// ExcludeFollowerDemotion specifies whether to skip demoting follower replicas
	// on the broker to be demoted.
	ExcludeFollowerDemotion bool `param:"exclude_follower_demotion"`
	// ExcludeRecentlyDemotedBrokers specifies whether to allow leader replicas
	// to be moved to recently demoted broker.
	ExcludeRecentlyDemotedBrokers bool `param:"exclude_recently_demoted_brokers"`
	// ReplicaMovementStrategies specifies the replica movement strategy to use.
	ReplicaMovementStrategies ReplicaMovementStrategy `param:"replica_movement_strategies"`
	// ReplicationThrottle is the upper bound on the bandwidth used to move replicas.
	ReplicationThrottle int `param:"replication_throttle"`
	// Verbose returns detailed state information.
	Verbose bool `param:"verbose"`
}

// DemoteBrokersDefaults returns a request with the default demote_broker parameters.
func DemoteBrokersDefaults(brokerIDs []int) *DemoteBrokersRequest {
	return &DemoteBrokersRequest{
		BrokerIDs:               brokerIDs,
		DryRun:                  true,
		AllowCapacityEstimation: true,
	}
}

// Validate checks that a DemoteBrokersRequest is valid.
func (req *DemoteBrokersRequest) Validate() error {
	if req.BrokerIDs == nil || len(req.BrokerIDs) == 0 {
		if req.BrokerIDsAndLogDirs == nil || len(req.BrokerIDsAndLogDirs) == 0 {
			return errors.New("One of BrokerIDs or BrokerIDsAndLogDirs must be specified")
		}
	}
	return nil
}
