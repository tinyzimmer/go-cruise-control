package types

import "errors"

// AddBrokersRequest represents the parameters for an add_broker request.
type AddBrokersRequest struct {
	// BrokerIDs is a list of ids of new broker(s) to add to the cluster.
	BrokerIDs []int `param:"brokerid"`
	// DryRun will just report what would have happened without doing anything.
	DryRun bool `param:"dryrun"`
	// DataFrom specifies whether to generate proposal from available valid partitions
	// or valid windows.
	DataFrom DataSource `param:"data_from"`
	// Goals is a list of goals used to generate the proposal.
	Goals []string `param:"goals"`
	// KafkaAssigner specifies whether to use Kafka assigner mode to generate proposal.
	KafkaAssigner bool `param:"kafka_assigner"`
	// AllowCapacityEstimation specifies whether to allow broker capacity to be
	// estimated from other brokers in the cluster.
	AllowCapacityEstimation bool `param:"allow_capacity_estimation"`
	// ConcurrentPartitionMovementsPerBroker specifies the upper bound of ongoing
	// replica movements going in and out of each broker.
	ConcurrentPartitionMovementsPerBroker int `param:"concurrent_partition_movements_per_broker"`
	// ConcurrentLeaderMovements specifies the upper bound of ongoing leadership
	// movements.
	ConcurrentLeaderMovements int `param:"concurrent_leader_movements"`
	// SkipHardGoalCheck specifies whether to allow hard goal to be skipped in
	// proposal generation.
	SkipHardGoalCheck bool `param:"skip_hard_goal_check"`
	// ExcludedTopicsRegex is a regular expression to specify topic(s) not to be
	// considered for replica movement.
	ExcludedTopicsRegex string `param:"excluded_topics"`
	// UseReadyDefaultGoals specifies whether to only use ready goals to generate
	// proposal.
	UseReadyDefaultGoals bool `param:"use_ready_default_goals"`
	// ExcludeRecentlyDemotedBrokers specifies whether to allow leader replicas
	// to be moved to recently demoted broker.
	ExcludeRecentlyDemotedBrokers bool `param:"exclude_recently_demoted_brokers"`
	// ExcludeRecentlyRemovedBrokers specifies whether to allow replicas to be
	// moved to recently removed broker.
	ExcludeRecentlyRemovedBrokers bool `param:"exclude_recently_removed_brokers"`
	// ReplicaMovementStrategies specifies the replica movement strategy to use.
	ReplicaMovementStrategies ReplicaMovementStrategy `param:"replica_movement_strategies"`
	// ReplicationThrottle is the upper bound on the bandwidth used to move replicas.
	ReplicationThrottle int `param:"replication_throttle"`
	// ThrottleAddedBroker specifies whether to throttle replica movement to new
	// broker(s).
	ThrottleAddedBroker bool `param:"throttle_added_broker"`
	// Verbose returns detailed state information.
	Verbose bool `param:"verbose"`
}

// AddBrokersDefaults returns a request with the default add_broker parameters.
func AddBrokersDefaults(brokerIDs []int) *AddBrokersRequest {
	return &AddBrokersRequest{
		BrokerIDs:               brokerIDs,
		DataFrom:                ValidWindows,
		DryRun:                  true,
		AllowCapacityEstimation: true,
	}
}

// Validate checks that an AddBrokersRequest is valid.
func (req *AddBrokersRequest) Validate() error {
	if req.BrokerIDs == nil || len(req.BrokerIDs) == 0 {
		return errors.New("BrokerID(s) cannot be empty")
	}
	return nil
}
