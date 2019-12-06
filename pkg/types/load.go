package types

import "time"

type GetLoadRequest struct {
	Start                   time.Time `param:"start"`
	End                     time.Time `param:"end"`
	AllowCapacityEstimation bool      `param:"allow_capacity_estimation"`
	// PopulateDiskInfo not tested
	PopulateDiskInfo bool `param:"populate_disk_info"`
}

func GetLoadDefaults() *GetLoadRequest {
	return &GetLoadRequest{
		AllowCapacityEstimation: true,
	}
}

type GetLoadResponse struct {
	Hosts   []HostLoad   `json:"hosts"`
	Brokers []BrokerLoad `json:"brokers"`
	Version int          `json:"version"`
}
