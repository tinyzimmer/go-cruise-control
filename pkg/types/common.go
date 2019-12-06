package types

type DataSource string
type ReplicaMovementStrategy string

const (
	ValidWindows    DataSource = "VALID_WINDOWS"
	ValidPartitions DataSource = "VALID_PARTITIONS"

	PrioritizeSmallReplica ReplicaMovementStrategy = "PrioritizeSmallReplicaMovementStrategy"
	PrioritizeLargeReplica ReplicaMovementStrategy = "PrioritizeLargeReplicaMovementStrategy"
	PostponeUrpReplica     ReplicaMovementStrategy = "PostponeUrpReplicaMovementStrategy"
)

// ClientOptions contains the parameters for creating a new Cruise Control client
type ClientOptions struct {
	URL string
}

type HostLoad struct {
	FollowerNwInRate float64 `json:"FollowerNwInRate"`
	Leaders          int     `json:"Leaders"`
	DiskMB           float64 `json:"DiskMB"`
	PnwOutRate       float64 `json:"PnwOutRate"`
	NwOutRate        float64 `json:"NwOutRate"`
	Host             string  `json:"Host"`
	CPUPct           float64 `json:"CpuPct"`
	Replicas         int     `json:"Replicas"`
	LeaderNwInRate   float64 `json:"LeaderNwInRate"`
	DiskPct          float64 `json:"DiskPct"`
}

type BrokerLoad struct {
	Broker           int     `json:"Broker"`
	BrokerState      string  `json:"BrokerState"`
	FollowerNwInRate float64 `json:"FollowerNwInRate"`
	Leaders          int     `json:"Leaders"`
	DiskMB           float64 `json:"DiskMB"`
	PnwOutRate       float64 `json:"PnwOutRate"`
	NwOutRate        float64 `json:"NwOutRate"`
	Host             string  `json:"Host"`
	CPUPct           float64 `json:"CpuPct"`
	Replicas         int     `json:"Replicas"`
	LeaderNwInRate   float64 `json:"LeaderNwInRate"`
	DiskPct          float64 `json:"DiskPct"`
}

type GenericResponse struct {
	Message string `json:"message"`
	Version int    `json:"version"`
}
