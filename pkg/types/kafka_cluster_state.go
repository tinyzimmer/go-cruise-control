package types

// GetKafkaClusterStateRequest represents the parameters for a kafka_cluster_state
// request.
type GetKafkaClusterStateRequest struct {
	// TopicRegex is a regex for filtering partitions to report on.
	TopicRegex string `param:"topic"`

	// Verbose returns detailed state information.
	Verbose bool `param:"verbose"`
}

// GetKafkaClusterStateDefaults returns a request with the default kafka_cluster_state
// parameters.
func GetKafkaClusterStateDefaults() *GetKafkaClusterStateRequest {
	return &GetKafkaClusterStateRequest{}
}

// GetKafkaClusterStateResponse represents a response from the kafka_cluster_state
// endpoint.
type GetKafkaClusterStateResponse struct {
	KafkaPartitionState KafkaPartitionState `json:"KafkaPartitionState"`
	KafkaBrokerState    KafkaBrokerState    `json:"KafkaBrokerState"`
	Version             int                 `json:"version"`
}

// KafkaPartitionState represents the states of the partitions on the cluster.
type KafkaPartitionState struct {
	Offline             []PartitionState `json:"offline"`
	URP                 []PartitionState `json:"urp"`
	WithOfflineReplicas []PartitionState `json:"with-offline-replicas"`
	UnderMinISR         []PartitionState `json:"under-min-isr"`
}

// KafkaBrokerState represents the state of a single Kafka broker.
type KafkaBrokerState struct {
	OfflineLogDirsByBrokerID      map[string][]string `json:"OfflineLogDirsByBrokerId"`
	OnlineLogDirsByBrokerID       map[string][]string `json:"OnlineLogDirsByBrokerId"`
	ReplicaCountByBrokerID        map[string]int      `json:"ReplicaCountByBrokerId"`
	OfflineReplicaCountByBrokerID map[string]int      `json:"OfflineReplicaCountByBrokerId"`
	OutOfSyncCountByBrokerID      map[string]int      `json:"OutOfSyncCountByBrokerId"`
	LeaderCountByBrokerID         map[string]int      `json:"LeaderCountByBrokerId"`
}

// PartitionState represents the state of a single partition.
type PartitionState struct {
	InSync    []int  `json:"in-sync"`
	OutOfSync []int  `json:"out-of-sync"`
	Leader    int    `json:"leader"`
	Offline   []int  `json:"offline"`
	Replicas  []int  `json:"replicas"`
	Topic     string `json:"topic"`
	Partition int    `json:"partition"`
}
