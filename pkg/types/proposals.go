package types

// GetProposalsRequest represents the parameters for a proposals request.
type GetProposalsRequest struct {
	// IgnoreProposalCache specifies whether to ignore the proposal cache.
	IgnoreProposalCache bool `param:"ignore_proposal_cache"`

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

	// DestinationBrokerIDs specifies brokers to consider moving replicas to.
	DestinationBrokerIDs []string `param:"destination_broker_ids"`

	// RebalanceDisk specifies whether to balance load between brokers or between
	// disks within a broker.
	RebalanceDisk bool `param:"rebalance_disk"`

	// Verbose returns detailed state information.
	Verbose bool `param:"verbose"`
}

// GetProposalsDefaults returns a request with the default proposals parameters.
func GetProposalsDefaults() *GetProposalsRequest {
	return &GetProposalsRequest{
		DataFrom:                ValidWindows,
		AllowCapacityEstimation: true,
	}
}

// ProposalsResponse represents a response to a proposals request.
type ProposalsResponse struct {
	Summary                ProposalsSummary `json:"summary"`
	GoalSummary            []Goal           `json:"goalSummary"`
	LoadBeforeOptimization ClusterLoad      `json:"loadBeforeOptimization"`
	LoadAfterOptimization  ClusterLoad      `json:"loadAfterOptimization"`
	Proposals              []Proposal       `json:"proposals"`
	Version                int              `json:"version"`
}

// ProposalsSummary represents the summary of the proposals.
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

// Goal represents a goal to be satisfied by a proposal.
type Goal struct {
	Name              string     `json:"goal"`
	Status            string     `json:"status"`
	ClusterModelStats ModelStats `json:"clusterModelStats"`
}

// ModelStats represents the statistics for a training model.
type ModelStats struct {
	Metadata   ModelMetadata   `json:"metadata"`
	Statistics ModelStatistics `json:"statistics"`
}

// ModelMetadata represents the metadata for a given model.
type ModelMetadata struct {
	Brokers  int `json:"brokers"`
	Replicas int `json:"replicas"`
	Topics   int `json:"topics"`
}

// ModelStatistics represents groups of statistical data for a model.
type ModelStatistics struct {
	AVG ModelStatsValues `json:"AVG"`
	STD ModelStatsValues `json:"STD"`
	MIN ModelStatsValues `json:"MIN"`
	MAX ModelStatsValues `json:"MAX"`
}

// ModelStatsValues represents the actual values for a given model.
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

// Proposal represents the actions taken for a single proposal.
type Proposal struct {
	OldLeader      int            `json:"oldLeader"`
	OldReplicas    []int          `json:"oldReplicas"`
	NewReplicas    []int          `json:"newReplicas"`
	TopicPartition TopicPartition `json:"topicPartition"`
}

// TopicPartition represents a single partition for a topic.
type TopicPartition struct {
	Hash      int    `json:"hash"`
	Partition int    `json:"partition"`
	Topic     string `json:"topic"`
}
