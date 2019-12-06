package types

import "errors"

type AddBrokersRequest struct {
	BrokerIDs                             []int                   `param:"brokerid"`
	DryRun                                bool                    `param:"dryrun"`
	DataFrom                              DataSource              `param:"data_from"`
	Goals                                 []string                `param:"goals"`
	KafkaAssigner                         bool                    `param:"kafka_assigner"`
	AllowCapacityEstimation               bool                    `param:"allow_capacity_estimation"`
	ConcurrentPartitionMovementsPerBroker int                     `param:"concurrent_partition_movements_per_broker"`
	ConcurrentLeaderMovements             int                     `param:"concurrent_leader_movements"`
	SkipHardGoalCheck                     bool                    `param:"skip_hard_goal_check"`
	ExcludedTopicsRegex                   string                  `param:"excluded_topics"`
	UseReadyDefaultGoals                  bool                    `param:"use_ready_default_goals"`
	ExcludeRecentlyDemotedBrokers         bool                    `param:"exclude_recently_demoted_brokers"`
	ExcludeRecentlyRemovedBrokers         bool                    `param:"exclude_recently_removed_brokers"`
	ReplicaMovementStrategies             ReplicaMovementStrategy `param:"replica_movement_strategies"`
	ReplicationThrottle                   int                     `param:"replication_throttle"`
	ThrottleAddedBroker                   bool                    `param:"throttle_added_broker"`
	Verbose                               bool                    `param:"verbose"`
}

func AddBrokersDefaults(brokerIDs []int) *AddBrokersRequest {
	return &AddBrokersRequest{
		BrokerIDs:               brokerIDs,
		DataFrom:                ValidWindows,
		DryRun:                  true,
		AllowCapacityEstimation: true,
	}
}

func (req *AddBrokersRequest) Validate() error {
	if req.BrokerIDs == nil || len(req.BrokerIDs) == 0 {
		return errors.New("BrokerID(s) cannot be empty")
	}
	return nil
}
