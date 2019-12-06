package types

import "errors"

type DemoteBrokersRequest struct {
	BrokerIDs                     []int                   `param:"brokerid"`
	BrokerIDsAndLogDirs           []string                `param:"brokerid_and_logdirs"`
	DryRun                        bool                    `param:"dryrun"`
	AllowCapacityEstimation       bool                    `param:"allow_capacity_estimation"`
	ConcurrentLeaderMovements     int                     `param:"concurrent_leader_movements"`
	SkipUrpDemotion               bool                    `param:"skip_urp_demotion"`
	ExcludeFollowerDemotion       bool                    `param:"exclude_follower_demotion"`
	ExcludeRecentlyDemotedBrokers bool                    `param:"exclude_recently_demoted_brokers"`
	ReplicaMovementStrategies     ReplicaMovementStrategy `param:"replica_movement_strategies"`
	ReplicationThrottle           int                     `param:"replication_throttle"`
	Verbose                       bool                    `param:"verbose"`
}

func DemoteBrokersDefaults(brokerIDs []int) *DemoteBrokersRequest {
	return &DemoteBrokersRequest{
		BrokerIDs:               brokerIDs,
		DryRun:                  true,
		AllowCapacityEstimation: true,
	}
}

func (req *DemoteBrokersRequest) Validate() error {
	if req.BrokerIDs == nil || len(req.BrokerIDs) == 0 {
		if req.BrokerIDsAndLogDirs == nil || len(req.BrokerIDsAndLogDirs) == 0 {
			return errors.New("One of BrokerIDs or BrokerIDsAndLogDirs must be specified")
		}
	}
	return nil
}
