package types

import "time"

// GetLoadRequest represents the parameters for a load request.
type GetLoadRequest struct {
	// Start is an optional start-time of the cluster load.
	Start time.Time `param:"start"`
	// End is an optional end-time for the cluster load.
	End time.Time `param:"end"`
	// AllowCapacityEstimation specifies whether to allow broker capacity to be
	// estimated from other brokers in the cluster.
	AllowCapacityEstimation bool `param:"allow_capacity_estimation"`
	// PopulateDiskInfo specifies whether to show the load of each disk on a
	// broker.
	PopulateDiskInfo bool `param:"populate_disk_info"`
}

// GetLoadDefaults returns a request with the default load parameters.
func GetLoadDefaults() *GetLoadRequest {
	return &GetLoadRequest{
		AllowCapacityEstimation: true,
	}
}

// GetLoadResponse represents a response to a get load request.
type GetLoadResponse struct {
	Hosts   []HostLoad   `json:"hosts"`
	Brokers []BrokerLoad `json:"brokers"`
	Version int          `json:"version"`
}
