package model

type GetTimeEntriesQuery struct {
	Since     *int    `json:"since"`
	Before    *string `json:"before"`
	StartDate *string `json:"start_date"`
	EndDate   *string `json:"end_date"`
}

type baseTimeEntry struct {
	Billable    bool     `json:"billable"`
	CreatedWith string   `json:"created_with"`
	Description *string  `json:"description"`
	Duration    int      `json:"duration"`
	Duronly     bool     `json:"duronly"`
	Pid         int      `json:"pid"`
	ProjectID   *int     `json:"project_id"`
	StartDate   string   `json:"start_date"`
	Start       string   `json:"start"`
	Stop        string   `json:"stop"`
	TagAction   string   `json:"tag_action"`
	TagIDs      []int    `json:"tag_ids"`
	Tags        []string `json:"tags"`
	TaskID      *int     `json:"task_id"`
	Tid         int      `json:"tid"`
	Uid         int      `json:"uid"`
	UserID      int      `json:"user_id"`
	Wid         int      `json:"wid"`
	WorkspaceID int      `json:"workspace_id"`
}

type timeEntry struct {
	baseTimeEntry
	At              string  `json:"at"`
	ID              int     `json:"id"`
	ServerDeletedAt *string `json:"server_deleted_at"`
}

type GetTimeEntriesResponse struct {
	timeEntry
}

type GetCurrentTimeEntryResponse struct {
	timeEntry
}

type CreateTimeEntryRequest struct {
	baseTimeEntry
}

type CreateTimeEntryResponse struct {
	timeEntry
}

type UpdateMultipleTimeEntriesRequest struct {
	Operator string `json:"op"`
	Path     string `json:"path"`
	Value    string `json:"value"`
}

type UpdateMultipleTimeEntriesResponse struct {
	Failure []struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	} `json:"failure"`
	Success []int `json:"success"`
}

type UpdateTimeEntryRequest struct {
	baseTimeEntry
}

type UpdateTimeEntryResponse struct {
	timeEntry
}

type StopTimeEntryResponse struct {
	timeEntry
}
