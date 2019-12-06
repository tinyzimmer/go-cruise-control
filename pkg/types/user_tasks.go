package types

type GetUserTasksRequest struct {
	UserTaskIDs        []string `param:"user_task_ids"`
	ClientIDs          []string `param:"client_ids"`
	Entries            int      `param:"entries"`
	Endpoints          []string `param:"endpoints"`
	RequestTypes       []string `param:"types"`
	FetchCompletedTask bool     `param:"fetch_completed_task"`
}

func GetUserTasksDefaults() *GetUserTasksRequest {
	return &GetUserTasksRequest{}
}

type GetUserTasksResponse struct {
	UserTasks []UserTask `json:"userTasks"`
	Version   int        `json:"version"`
}

type UserTask struct {
	Status         string `json:"Status"`
	UserTaskID     string `json:"UserTaskId"`
	StartMs        string `json:"StartMs"`
	ClientIdentity string `json:"ClientIdentity"`
	RequestURL     string `json:"RequestURL"`
}
