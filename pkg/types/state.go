package types

// GetStateRequest represents the parameters for a state request.
type GetStateRequest struct {
	// Substates is a list of substates to report on. If empty, all substate will
	// be returned.
	Substates []Substate `param:"substates"`
	// Verbose returns detailed state information.
	Verbose bool `param:"verbose"`
	// SuperVerbose returns even more detailed state information.
	SuperVerbose bool `param:"super_verbose"`
}

// GetStateDefaults returns the defaults for a state request.
func GetStateDefaults() *GetStateRequest {
	return &GetStateRequest{}
}

// GetStateResponse represents the response to a state request.
type GetStateResponse struct {
	AnalyzerState        AnalyzerState        `json:"AnalyzerState"`
	MonitorState         MonitorState         `json:"MonitorState"`
	ExecutorState        ExecutorState        `json:"ExecutorState"`
	AnomalyDetectorState AnomalyDetectorState `json:"AnomalyDetectorState"`
	Version              int                  `json:"version"`
}

// AnalyzerState represents the state of the analyzer.
type AnalyzerState struct {
	IsProposalReady bool            `json:"isProposalReady"`
	ReadyGoals      []string        `json:"readyGoals"`
	GoalReadiness   []GoalReadiness `json:"goalReadiness"`
}

// GoalReadiness represents the status of a goal.
type GoalReadiness struct {
	Name                     string                   `json:"name"`
	ModelCompleteRequirement ModelCompleteRequirement `json:"modelCompleteRequirement"`
	Status                   string                   `json:"status"`
}

// ModelCompleteRequirement represents the requirements for which a model is complete.
type ModelCompleteRequirement struct {
	IncludeAllTopics                 bool    `json:"includeAllTopics"`
	MinMonitoredPartitionsPercentage float64 `json:"minMonitoredPartitionsPercentage"`
	RequiredNumSnapshots             int     `json:"requiredNumSnapshots"`
}

// MonitorState represents the state of the monitor.
type MonitorState struct {
	TrainingPct                 float64            `json:"trainingPct"`
	Trained                     bool               `json:"trained"`
	NumFlawedPartitions         int                `json:"numFlawedPartitions"`
	State                       string             `json:"state"`
	NumTotalPartitions          int                `json:"numTotalPartitions"`
	NumMonitoredWindows         int                `json:"numMonitoredWindows"`
	MonitoredWindows            map[string]float64 `json:"monitoredWindows"`
	MonitoringCoveragePct       float64            `json:"monitoringCoveragePct"`
	ReasonOfLatestPauseOrResume string             `json:"reasonOfLatestPauseOrResume"`
	NumValidPartitions          int                `json:"numValidPartitions"`
}

// ExecutorState represents the state of the executor.
type ExecutorState struct {
	State string `json:"state"`
}

// AnomalyDetectorState represents the state of the anomaly detector.
type AnomalyDetectorState struct {
	RecentBrokerFailures    []Anomaly              `json:"recentBrokerFailures"`
	RecentGoalViolations    []Anomaly              `json:"recentGoalViolations"`
	SelfHealingDisabled     []string               `json:"selfHealingDisabled"`
	SelfHealingEnabled      []string               `json:"selfHealingEnabled"`
	RecentDiskFailures      []Anomaly              `json:"recentDiskFailures"`
	Metrics                 AnomalyDetectorMetrics `json:"metrics"`
	RecentMetricAnomalies   []Anomaly              `json:"recentMetricAnomalies"`
	SelfHealingEnabledRatio map[string]float64     `json:"selfHealingEnabledRatio"`
}

// Anomaly represents the details about an anomaly in the cluster.
type Anomaly struct {
	Description    string `json:"description"`
	AnomalyID      string `json:"anomalyId"`
	DetectionMs    int64  `json:"detectionMs"`
	StatusUpdateMs int64  `json:"statusUpdateMs"`
	Status         string `json:"status"`
}

// AnomalyDetectorMetrics represents the metrics for the anomaly detector.
type AnomalyDetectorMetrics struct {
	MeanTimeToStartFixMs       float64            `json:"meanTimeToStartFixMs"`
	NumSelfHealingStarted      int                `json:"numSelfHealingStarted"`
	OngoingAnomalyDurationMs   float64            `json:"ongoingAnomalyDurationMs"`
	MeanTimeBetweenAnomaliesMs map[string]float64 `json:"meanTimeBetweenAnomaliesMs"`
}
