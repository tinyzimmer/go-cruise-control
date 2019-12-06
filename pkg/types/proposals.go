package types

type GetProposalsRequest struct {
	IgnoreProposalCache           bool       `param:"ignore_proposal_cache"`
	DataFrom                      DataSource `param:"data_from"`
	Goals                         []string   `param:"goals"`
	KafkaAssigner                 bool       `param:"kafka_assigner"`
	AllowCapacityEstimation       bool       `param:"allow_capacity_estimation"`
	ExcludedTopicsRegex           string     `param:"excluded_topics"`
	UseReadyDefaultGoals          bool       `param:"use_ready_default_goals"`
	ExcludeRecentlyDemotedBrokers bool       `param:"exclude_recently_demoted_brokers"`
	ExcludeRecentlyRemovedBrokers bool       `param:"exclude_recently_removed_brokers"`
	DestinationBrokerIDs          []string   `param:"destination_broker_ids"`
	RebalanceDisk                 bool       `param:"rebalance_disk"`
	Verbose                       bool       `param:"verbose"`
}

func GetProposalsDefaults() *GetProposalsRequest {
	return &GetProposalsRequest{
		DataFrom:                ValidWindows,
		AllowCapacityEstimation: true,
	}
}

type ProposalsResponse struct {
	Summary     ProposalsSummary `json:"summary"`
	GoalSummary []Goal           `json:"goalSummary"`
	// only populated on verbose
	LoadBeforeOptimization LoadOptimization `json:"loadBeforeOptimization"`
	LoadAfterOptimization  LoadOptimization `json:"loadAfterOptimization"`
	// only populated on verbose
	Proposals []Proposal `json:"proposals"`
	Version   int        `json:"version"`
}

type ProposalsSummary struct {
	NumIntraBrokerReplicaMovements int      `json:"numIntraBrokerReplicaMovements"`
	ExcludedBrokersForLeadership   []int    `json:"excludedBrokersForLeadership"`
	ExcludedBrokersForReplicaMove  []int    `json:"excludedBrokersForReplicaMove"`
	NumReplicaMovements            int      `json:"numReplicaMovements"`
	IntraBrokerDataToMoveMB        int      `json:"intraBrokerDataToMoveMB"`
	RecentWindows                  int      `json:"recentWindows"`
	DataToMoveMB                   int      `json:"dataToMoveMB"`
	MonitoredPartitionsPercentage  float64  `json:"monitoredPartitionsPercentage"`
	ExcludedTopics                 []string `json:"excludedTopics"`
	NumLeaderMovements             int      `json:"numLeaderMovements"`
}

type Goal struct {
	Name              string     `json:"goal"`
	Status            string     `json:"status"`
	ClusterModelStats ModelStats `json:"clusterModelStats"`
}

type ModelStats struct {
	Metadata   ModelMetadata   `json:"metadata"`
	Statistics ModelStatistics `json:"statistics"`
}

type ModelMetadata struct {
	Brokers  int `json:"brokers"`
	Replicas int `json:"replicas"`
	Topics   int `json:"topics"`
}

type ModelStatistics struct {
	AVG ModelStatsValues `json:"AVG"`
	STD ModelStatsValues `json:"STD"`
	MIN ModelStatsValues `json:"MIN"`
	MAX ModelStatsValues `json:"MAX"`
}

type ModelStatsValues struct {
	Disk            float64 `json:"disk"`
	Replicas        float64 `json:"replicas"`
	LeaderReplicas  float64 `json:"leaderReplicas"`
	CPU             float64 `json:"cpu"`
	NetworkOutbound float64 `json:"networkOutbound"`
	NetworkInbound  float64 `json:"networkInbound"`
	TopicReplicas   float64 `json:"topicReplicas"`
	PotentialNwOut  float64 `json:"potentialNwOut"`
}

type LoadOptimization struct {
	Hosts   []HostLoad   `json:"hosts"`
	Brokers []BrokerLoad `json:"brokers"`
}

type Proposal struct {
	OldLeader      int            `json:"oldLeader"`
	OldReplicas    []int          `json:"oldReplicas"`
	NewReplicas    []int          `json:"newReplicas"`
	TopicPartition TopicPartition `json:"topicPartition"`
}

type TopicPartition struct {
	Hash      int    `json:"hash"`
	Partition int    `json:"partition"`
	Topic     string `json:"topic"`
}
