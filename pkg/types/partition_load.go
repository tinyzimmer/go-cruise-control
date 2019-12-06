package types

import "time"

type LoadType string

const (
	DiskResource       LoadType = "DISK"
	CPUResource        LoadType = "CPU"
	NetworkInResource  LoadType = "NW_IN"
	NetworkOutResource LoadType = "NW_OUT"
)

type GetPartitionLoadRequest struct {
	Resource                LoadType  `param:"resource"`
	Start                   time.Time `param:"start"`
	End                     time.Time `param:"end"`
	Entries                 int       `param:"entries"`
	AllowCapacityEstimation bool      `param:"allow_capacity_estimation"`
	MaxLoad                 bool      `param:"max_load"`
	AvgLoad                 bool      `param:"avg_load"`
	TopicRegex              string    `param:"topic"`
	PartitionRange          string    `param:"partition"`
	MinValidPartitionRatio  float64   `param:"min_valid_partition_ratio"`
	BrokerID                int       `param:"broker_id"`
}

func GetPartitionLoadDefaults() *GetPartitionLoadRequest {
	return &GetPartitionLoadRequest{
		Resource:                DiskResource,
		AllowCapacityEstimation: true,
	}
}

type GetPartitionLoadResponse struct {
	Records []PartitionLoadRecord `json:"records"`
	Version int                   `json:"version"`
}

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
