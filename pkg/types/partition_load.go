package types

import "time"

// GetPartitionLoadRequest represents the parameters for a partition_load request.
type GetPartitionLoadRequest struct {
	// Resource specifies what to sort results by.
	Resource LoadType `param:"resource"`
	// Start is an optional start-time of the partition load.
	Start time.Time `param:"start"`

	// End is an optional end-time for the partition load.
	End time.Time `param:"end"`

	// Entries specifies the number of load entries to return in the response.
	Entries int `param:"entries"`

	// AllowCapacityEstimation specifies whether to allow broker capacity to be
	// estimated from other brokers in the cluster.
	AllowCapacityEstimation bool `param:"allow_capacity_estimation"`

	// MaxLoad specifies whether to report the maximum load for partition in windows.
	MaxLoad bool `param:"max_load"`

	// AvgLoad specifies whether to report the average load for partition in windows.
	AvgLoad bool `param:"avg_load"`

	// TopicRegex is a regex to filter partition load based on partition's topic.
	TopicRegex string `param:"topic"`

	// PartitionRange is a number or range to filter partitions with.
	PartitionRange string `param:"partition"`

	// MinValidPartitionRatio is the minimal valid partition ratio requirement for
	// cluster models.
	MinValidPartitionRatio float64 `param:"min_valid_partition_ratio"`

	// BrokerID is a broker to filter results by.
	BrokerID int `param:"broker_id"`
}

// GetPartitionLoadDefaults returns a request with the default partitiion_load
// parameters.
func GetPartitionLoadDefaults() *GetPartitionLoadRequest {
	return &GetPartitionLoadRequest{
		Resource:                DiskResource,
		AllowCapacityEstimation: true,
	}
}

// GetPartitionLoadResponse represents a response to a get partition load request.
type GetPartitionLoadResponse struct {
	Records []PartitionLoadRecord `json:"records"`
	Version int                   `json:"version"`
}

// PartitionLoadRecord represents the load details for a single partition.
type PartitionLoadRecord struct {
	Leader          int     `json:"leader"`
	Disk            float64 `json:"disk"`
	Partition       int     `json:"partition"`
	Followers       []int   `json:"followers"`
	MsgIn           float64 `json:"msg_in"`
	Topic           string  `json:"topic"`
	CPU             float64 `json:"cpu"`
	NetworkOutbound float64 `json:"networkOutbound"`
	NetworkInbound  float64 `json:"networkInbound"`
}
