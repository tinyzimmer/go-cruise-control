package types

type GetKafkaClusterStateRequest struct {
	Topic   string `param:"topic"`
	Verbose bool   `param:"verbose"`
}

func GetKafkaClusterStateDefaults() *GetKafkaClusterStateRequest {
	return &GetKafkaClusterStateRequest{}
}

type GetKafkaClusterStateResponse struct {
	KafkaPartitionState KafkaPartitionState `json:"KafkaPartitionState"`
	KafkaBrokerState    KafkaBrokerState    `json:"KafkaBrokerState"`
	Version             int                 `json:"version"`
}

type KafkaPartitionState struct {
	Offline             []PartitionState `json:"offline"`
	URP                 []PartitionState `json:"urp"`
	WithOfflineReplicas []PartitionState `json:"with-offline-replicas"`
	UnderMinISR         []PartitionState `json:"under-min-isr"`
}

type KafkaBrokerState struct {
	OfflineLogDirsByBrokerID      map[string][]string `json:"OfflineLogDirsByBrokerId"`
	OnlineLogDirsByBrokerID       map[string][]string `json:"OnlineLogDirsByBrokerId"`
	ReplicaCountByBrokerID        map[string]int      `json:"ReplicaCountByBrokerId"`
	OfflineReplicaCountByBrokerID map[string]int      `json:"OfflineReplicaCountByBrokerId"`
	OutOfSyncCountByBrokerID      map[string]int      `json:"OutOfSyncCountByBrokerId"`
	LeaderCountByBrokerID         map[string]int      `json:"LeaderCountByBrokerId"`
}

type PartitionState struct {
	InSync    []int  `json:"in-sync"`
	OutOfSync []int  `json:"out-of-sync"`
	Leader    int    `json:"leader"`
	Offline   []int  `json:"offline"`
	Replicas  []int  `json:"replicas"`
	Topic     string `json:"topic"`
	Partition int    `json:"partition"`
}
