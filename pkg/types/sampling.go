package types

type TriggerSamplingRequest struct {
	Reason string `param:"reason"`
}

func TriggerSamplingDefaults() *TriggerSamplingRequest {
	return &TriggerSamplingRequest{}
}
