package types

// Substate represents a substate for a GET state request.
type Substate string

// DataSource represents where to retrieve data for an operation from.
type DataSource string

// ReplicaMovementStrategy represents a replica movement strategy.
type ReplicaMovementStrategy string

// LoadType represents a type of resource load to sort results by.
type LoadType string

const (
	// SubstateAnalyzer contains if proposals are cached and/or ready.
	SubstateAnalyzer Substate = "analyzer"
	// SubstateMonitor contains details about the cruise control monitoring.
	SubstateMonitor Substate = "monitor"
	// SubstateExecutor contains the status of executions in cruise control.
	SubstateExecutor Substate = "executor"
	// SubstateAnomalyDetector contains the details of the anomaly detectors.
	SubstateAnomalyDetector Substate = "anomaly_detector"

	// ValidWindows when specified as a data_from argument, signals to handle
	// the operation based on the information in the available valid snapshot windows.
	// A valid snapshot window is a window whose valid monitored partitions coverage
	// meets the requirements of all the goals.
	ValidWindows DataSource = "VALID_WINDOWS"
	// ValidPartitions signals to handle the operation based on all the available
	// valid partitions. All the snapshot windows will be included in this case.
	ValidPartitions DataSource = "VALID_PARTITIONS"

	// PrioritizeSmallReplica will first move small replicas.
	PrioritizeSmallReplica ReplicaMovementStrategy = "PrioritizeSmallReplicaMovementStrategy"
	// PrioritizeLargeReplica will first move large replicas.
	PrioritizeLargeReplica ReplicaMovementStrategy = "PrioritizeLargeReplicaMovementStrategy"
	// PostponeUrpReplica will first move replicas for partitions having no
	// out-of-sync replica.
	PostponeUrpReplica ReplicaMovementStrategy = "PostponeUrpReplicaMovementStrategy"

	// DiskResource sorts results by disk load.
	DiskResource LoadType = "DISK"
	// CPUResource sorts results by CPU load.
	CPUResource LoadType = "CPU"
	// NetworkInResource sorts results by network-in load.
	NetworkInResource LoadType = "NW_IN"
	// NetworkOutResource sorts results by network-out load.
	NetworkOutResource LoadType = "NW_OUT"
)

// ClientOptions contains the parameters for creating a new Cruise Control client.
type ClientOptions struct {
	URL       string
	BasicAuth *BasicAuthCredentials
}

// BasicAuthCredentials contains the username and password to use for HTTP
// basic authentication.
type BasicAuthCredentials struct {
	Username string
	Password string
}

// ClusterLoad represents the load data for hosts and brokers in a cluster.
type ClusterLoad struct {
	Hosts   []HostLoad   `json:"hosts"`
	Brokers []BrokerLoad `json:"brokers"`
}

// HostLoad represents the load of a server running a Kafka broker.
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

// BrokerLoad represents the load from a single Kafka broker.
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

// GenericResponse represents a server response that just consists of a message
// and version identifier.
type GenericResponse struct {
	Message string `json:"message"`
	Version int    `json:"version"`
}
