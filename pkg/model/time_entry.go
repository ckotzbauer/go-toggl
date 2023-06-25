package model

type GetTimeEntriesQuery struct {
	Since     int    `json:"since,omitempty"`
	Before    string `json:"before,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type GetTimeEntriesResponse struct {
	Billable        bool     `json:"billable"`
	CreatedWith     string   `json:"created_with"`
	Description     string   `json:"description,omitempty"`
	Duration        int      `json:"duration"`
	Duronly         bool     `json:"duronly"`
	Pid             int      `json:"pid"`
	ProjectID       int      `json:"project_id,omitempty"`
	StartDate       string   `json:"start_date"`
	Start           string   `json:"start"`
	Stop            string   `json:"stop"`
	TagAction       string   `json:"tag_action"`
	TagIDs          []int    `json:"tag_ids"`
	Tags            []string `json:"tags"`
	TaskID          int      `json:"task_id,omitempty"`
	Tid             int      `json:"tid"`
	Uid             int      `json:"uid"`
	UserID          int      `json:"user_id"`
	Wid             int      `json:"wid"`
	WorkspaceID     int      `json:"workspace_id"`
	At              string   `json:"at"`
	ID              int      `json:"id"`
	ServerDeletedAt string   `json:"server_deleted_at,omitempty"`
}

type GetCurrentTimeEntryResponse struct {
	Billable        bool     `json:"billable"`
	CreatedWith     string   `json:"created_with"`
	Description     string   `json:"description,omitempty"`
	Duration        int      `json:"duration"`
	Duronly         bool     `json:"duronly"`
	Pid             int      `json:"pid"`
	ProjectID       int      `json:"project_id,omitempty"`
	StartDate       string   `json:"start_date"`
	Start           string   `json:"start"`
	Stop            string   `json:"stop"`
	TagAction       string   `json:"tag_action"`
	TagIDs          []int    `json:"tag_ids"`
	Tags            []string `json:"tags"`
	TaskID          int      `json:"task_id,omitempty"`
	Tid             int      `json:"tid"`
	Uid             int      `json:"uid"`
	UserID          int      `json:"user_id"`
	Wid             int      `json:"wid"`
	WorkspaceID     int      `json:"workspace_id"`
	At              string   `json:"at"`
	ID              int      `json:"id"`
	ServerDeletedAt string   `json:"server_deleted_at,omitempty"`
}

type CreateTimeEntryRequest struct {
	Billable    bool     `json:"billable"`
	CreatedWith string   `json:"created_with"`
	Description string   `json:"description,omitempty"`
	Duration    int      `json:"duration"`
	Duronly     bool     `json:"duronly"`
	Pid         int      `json:"pid"`
	ProjectID   int      `json:"project_id,omitempty"`
	StartDate   string   `json:"start_date"`
	Start       string   `json:"start"`
	Stop        string   `json:"stop"`
	TagAction   string   `json:"tag_action"`
	TagIDs      []int    `json:"tag_ids"`
	Tags        []string `json:"tags"`
	TaskID      int      `json:"task_id,omitempty"`
	Tid         int      `json:"tid"`
	Uid         int      `json:"uid"`
	UserID      int      `json:"user_id"`
	Wid         int      `json:"wid"`
	WorkspaceID int      `json:"workspace_id"`
}

type CreateTimeEntryResponse struct {
	Billable        bool     `json:"billable"`
	CreatedWith     string   `json:"created_with"`
	Description     string   `json:"description,omitempty"`
	Duration        int      `json:"duration"`
	Duronly         bool     `json:"duronly"`
	Pid             int      `json:"pid"`
	ProjectID       int      `json:"project_id,omitempty"`
	StartDate       string   `json:"start_date"`
	Start           string   `json:"start"`
	Stop            string   `json:"stop"`
	TagAction       string   `json:"tag_action"`
	TagIDs          []int    `json:"tag_ids"`
	Tags            []string `json:"tags"`
	TaskID          int      `json:"task_id,omitempty"`
	Tid             int      `json:"tid"`
	Uid             int      `json:"uid"`
	UserID          int      `json:"user_id"`
	Wid             int      `json:"wid"`
	WorkspaceID     int      `json:"workspace_id"`
	At              string   `json:"at"`
	ID              int      `json:"id"`
	ServerDeletedAt string   `json:"server_deleted_at,omitempty"`
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
	Billable    bool     `json:"billable"`
	CreatedWith string   `json:"created_with"`
	Description string   `json:"description,omitempty"`
	Duration    int      `json:"duration"`
	Duronly     bool     `json:"duronly"`
	Pid         int      `json:"pid"`
	ProjectID   int      `json:"project_id,omitempty"`
	StartDate   string   `json:"start_date"`
	Start       string   `json:"start"`
	Stop        string   `json:"stop"`
	TagAction   string   `json:"tag_action"`
	TagIDs      []int    `json:"tag_ids"`
	Tags        []string `json:"tags"`
	TaskID      int      `json:"task_id,omitempty"`
	Tid         int      `json:"tid"`
	Uid         int      `json:"uid"`
	UserID      int      `json:"user_id"`
	Wid         int      `json:"wid"`
	WorkspaceID int      `json:"workspace_id"`
}

type UpdateTimeEntryResponse struct {
	Billable        bool     `json:"billable"`
	CreatedWith     string   `json:"created_with"`
	Description     string   `json:"description,omitempty"`
	Duration        int      `json:"duration"`
	Duronly         bool     `json:"duronly"`
	Pid             int      `json:"pid"`
	ProjectID       int      `json:"project_id,omitempty"`
	StartDate       string   `json:"start_date"`
	Start           string   `json:"start"`
	Stop            string   `json:"stop"`
	TagAction       string   `json:"tag_action"`
	TagIDs          []int    `json:"tag_ids"`
	Tags            []string `json:"tags"`
	TaskID          int      `json:"task_id,omitempty"`
	Tid             int      `json:"tid"`
	Uid             int      `json:"uid"`
	UserID          int      `json:"user_id"`
	Wid             int      `json:"wid"`
	WorkspaceID     int      `json:"workspace_id"`
	At              string   `json:"at"`
	ID              int      `json:"id"`
	ServerDeletedAt string   `json:"server_deleted_at,omitempty"`
}

type StopTimeEntryResponse struct {
	Billable        bool     `json:"billable"`
	CreatedWith     string   `json:"created_with"`
	Description     string   `json:"description,omitempty"`
	Duration        int      `json:"duration"`
	Duronly         bool     `json:"duronly"`
	Pid             int      `json:"pid"`
	ProjectID       int      `json:"project_id,omitempty"`
	StartDate       string   `json:"start_date"`
	Start           string   `json:"start"`
	Stop            string   `json:"stop"`
	TagAction       string   `json:"tag_action"`
	TagIDs          []int    `json:"tag_ids"`
	Tags            []string `json:"tags"`
	TaskID          int      `json:"task_id,omitempty"`
	Tid             int      `json:"tid"`
	Uid             int      `json:"uid"`
	UserID          int      `json:"user_id"`
	Wid             int      `json:"wid"`
	WorkspaceID     int      `json:"workspace_id"`
	At              string   `json:"at"`
	ID              int      `json:"id"`
	ServerDeletedAt string   `json:"server_deleted_at,omitempty"`
}
