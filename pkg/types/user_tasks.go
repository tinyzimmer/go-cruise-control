package types

// GetUserTasksRequest represents the parameters for a user_tasks request.
type GetUserTasksRequest struct {
	// UserTaskIDs filters the response by a list of specific user tasks IDs.
	UserTaskIDs []string `param:"user_task_ids"`
	// ClientIDs filters the response by a list of client IP addresses.
	ClientIDs []string `param:"client_ids"`
	// Entries specifies the number of entries to return in the response.
	Entries int `param:"entries"`
	// Endpoints filters the results based off specific API endpoints.
	Endpoints []string `param:"endpoints"`
	// RequestTypes filters the results based off the given request types.
	RequestTypes []string `param:"types"`
	// FetchCompletedTask specifies whether to return the original request's response.
	FetchCompletedTask bool `param:"fetch_completed_task"`
}

// GetUserTasksDefaults returns a request with the default user_tasks parameters.
func GetUserTasksDefaults() *GetUserTasksRequest {
	return &GetUserTasksRequest{}
}

// GetUserTasksResponse represents the response to a user_tasks request.
type GetUserTasksResponse struct {
	UserTasks []UserTask `json:"userTasks"`
	Version   int        `json:"version"`
}

// UserTask contains the details for a single user-initiated task.
type UserTask struct {
	Status         string `json:"Status"`
	UserTaskID     string `json:"UserTaskId"`
	StartMs        string `json:"StartMs"`
	ClientIdentity string `json:"ClientIdentity"`
	RequestURL     string `json:"RequestURL"`
}
