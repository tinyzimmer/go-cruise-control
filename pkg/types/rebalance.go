package types

// RebalanceRequest represents the parameters for a rebalance request.
type RebalanceRequest struct {
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

	// ConcurrentIntraPartitionMovements specifies the upper bound of ongoing
	// replica movements going between disks of a single broker.
	ConcurrentIntraPartitionMovements int `param:"concurrent_intra_partition_movements"`

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

	// IgnoreProposalCache specifies whether to ignore the proposal cache.
	IgnoreProposalCache bool `param:"ignore_proposal_cache"`

	// ReplicationThrottle is the upper bound on the bandwidth used to move replicas.
	ReplicationThrottle int `param:"replication_throttle"`

	// DestinationBrokerIDs specifies brokers to consider moving replicas to.
	DestinationBrokerIDs []int `param:"destination_broker_ids"`

	// RebalanceDisk specifies whether to balance load between disks within each
	// broker or between brokers in cluster
	RebalanceDisk bool `param:"rebalance_disk"`

	// Verbose returns detailed state information.
	Verbose bool `param:"verbose"`
}

// RebalanceDefaults returns a request with the default parameters for a rebalance
// request.
func RebalanceDefaults() *RebalanceRequest {
	return &RebalanceRequest{
		DryRun:                  true,
		AllowCapacityEstimation: true,
	}
}
