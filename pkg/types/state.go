package types

type Substate string

const (
	SubstateAnalyzer        Substate = "analyzer"
	SubstateMonitor         Substate = "monitor"
	SubstateExecutor        Substate = "executor"
	SubstateAnomalyDetector Substate = "anomaly_detector"
)

type GetStateRequest struct {
	Substates    []Substate `param:"substates"`
	Verbose      bool       `param:"verbose"`
	SuperVerbose bool       `param:"super_verbose"`
}

func GetStateDefaults() *GetStateRequest {
	return &GetStateRequest{}
}

type GetStateResponse struct {
	AnalyzerState        AnalyzerState        `json:"AnalyzerState"`
	MonitorState         MonitorState         `json:"MonitorState"`
	ExecutorState        ExecutorState        `json:"ExecutorState"`
	AnomalyDetectorState AnomalyDetectorState `json:"AnomalyDetectorState"`
	Version              int                  `json:"version"`
}

type AnalyzerState struct {
	IsProposalReady bool            `json:"isProposalReady"`
	ReadyGoals      []string        `json:"readyGoals"`
	GoalReadiness   []GoalReadiness `json:"goalReadiness"`
}

type GoalReadiness struct {
	Name                     string                   `json:"name"`
	ModelCompleteRequirement ModelCompleteRequirement `json:"modelCompleteRequirement"`
	Status                   string                   `json:"status"`
}

type ModelCompleteRequirement struct {
	IncludeAllTopics                 bool    `json:"includeAllTopics"`
	MinMonitoredPartitionsPercentage float64 `json:"minMonitoredPartitionsPercentage"`
	RequiredNumSnapshots             int     `json:"requiredNumSnapshots"`
}

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

type ExecutorState struct {
	State string `json:"state"`
}

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

type Anomaly struct {
	Description    string `json:"description"`
	AnomalyID      string `json:"anomalyId"`
	DetectionMs    int64  `json:"detectionMs"`
	StatusUpdateMs int64  `json:"statusUpdateMs"`
	Status         string `json:"status"`
}

type AnomalyDetectorMetrics struct {
	MeanTimeToStartFixMs       float64            `json:"meanTimeToStartFixMs"`
	NumSelfHealingStarted      int                `json:"numSelfHealingStarted"`
	OngoingAnomalyDurationMs   float64            `json:"ongoingAnomalyDurationMs"`
	MeanTimeBetweenAnomaliesMs map[string]float64 `json:"meanTimeBetweenAnomaliesMs"`
}
