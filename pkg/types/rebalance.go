package types

type RebalanceRequest struct {
	DryRun                                bool                    `param:"dryrun"`
	DataFrom                              DataSource              `param:"data_from"`
	Goals                                 []string                `param:"goals"`
	KafkaAssigner                         bool                    `param:"kafka_assigner"`
	AllowCapacityEstimation               bool                    `param:"allow_capacity_estimation"`
	ConcurrentPartitionMovementsPerBroker int                     `param:"concurrent_partition_movements_per_broker"`
	ConcurrentIntraPartitionMovements     int                     `param:"concurrent_intra_partition_movements"`
	ConcurrentLeaderMovements             int                     `param:"concurrent_leader_movements"`
	SkipHardGoalCheck                     bool                    `param:"skip_hard_goal_check"`
	ExcludedTopicsRegex                   string                  `param:"excluded_topics"`
	UseReadyDefaultGoals                  bool                    `param:"use_ready_default_goals"`
	ExcludeRecentlyDemotedBrokers         bool                    `param:"exclude_recently_demoted_brokers"`
	ExcludeRecentlyRemovedBrokers         bool                    `param:"exclude_recently_removed_brokers"`
	ReplicaMovementStrategies             ReplicaMovementStrategy `param:"replica_movement_strategies"`
	IgnoreProposalCache                   bool                    `param:"ignore_proposal_cache"`
	ReplicationThrottle                   int                     `param:"replication_throttle"`
	DestinationBrokerIDs                  []int                   `param:"destination_broker_ids"`
	RebalanceDisk                         bool                    `param:"rebalance_disk"`
	Verbose                               bool                    `param:"verbose"`
}

func RebalanceDefaults() *RebalanceRequest {
	return &RebalanceRequest{
		DryRun:                  true,
		AllowCapacityEstimation: true,
	}
}
