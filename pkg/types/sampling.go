package types

// TriggerSamplingRequest represents a request to pause or resume metrics
// sampling
type TriggerSamplingRequest struct {
	// Reason provides a message for why the request was made.
	Reason string `param:"reason"`
}

// TriggerSamplingDefaults returns a TriggerSamplingRequest with the default
// parameters.
func TriggerSamplingDefaults() *TriggerSamplingRequest {
	return &TriggerSamplingRequest{}
}
