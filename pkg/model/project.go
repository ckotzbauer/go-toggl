package model

type GetProjectUsersQuery struct {
	ProjectIDs       string `json:"project_ids,omitempty"`
	WithGroupMembers bool   `json:"with_group_members,omitempty"`
}

type GetProjectUsersResponse struct {
	LabourCost      int    `json:"labour_cost"`
	Manager         bool   `json:"manager"`
	ProjectID       int    `json:"project_id"`
	Rate            int    `json:"rate"`
	RateChangeMode  string `json:"rate_change_mode"`
	UserID          int    `json:"user_id"`
	At              string `json:"at"`
	Gid             int    `json:"gid"`
	GroupID         int    `json:"group_id"`
	ID              int    `json:"id"`
	RateLastUpdated string `json:"rate_last_updated"`
	WorkspaceID     int    `json:"workspace_id"`
}

type AddProjectUserRequest struct {
	LabourCost     int    `json:"labour_cost"`
	Manager        bool   `json:"manager"`
	ProjectID      int    `json:"project_id"`
	Rate           int    `json:"rate"`
	RateChangeMode string `json:"rate_change_mode"`
	UserID         int    `json:"user_id"`
}

type AddProjectUserResponse struct {
	LabourCost      int    `json:"labour_cost"`
	Manager         bool   `json:"manager"`
	ProjectID       int    `json:"project_id"`
	Rate            int    `json:"rate"`
	RateChangeMode  string `json:"rate_change_mode"`
	UserID          int    `json:"user_id"`
	At              string `json:"at"`
	Gid             int    `json:"gid"`
	GroupID         int    `json:"group_id"`
	ID              int    `json:"id"`
	RateLastUpdated string `json:"rate_last_updated"`
	WorkspaceID     int    `json:"workspace_id"`
}

type UpdateMultipleProjectUsersResponse struct {
	Failure []struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	} `json:"failure"`
	Success []int `json:"success"`
}

type UpdateProjectUserRequest struct {
	LabourCost     int    `json:"labour_cost"`
	Manager        bool   `json:"manager"`
	ProjectID      int    `json:"project_id"`
	Rate           int    `json:"rate"`
	RateChangeMode string `json:"rate_change_mode"`
	UserID         int    `json:"user_id"`
}

type UpdateProjectUserResponse struct {
	LabourCost      int    `json:"labour_cost"`
	Manager         bool   `json:"manager"`
	ProjectID       int    `json:"project_id"`
	Rate            int    `json:"rate"`
	RateChangeMode  string `json:"rate_change_mode"`
	UserID          int    `json:"user_id"`
	At              string `json:"at"`
	Gid             int    `json:"gid"`
	GroupID         int    `json:"group_id"`
	ID              int    `json:"id"`
	RateLastUpdated string `json:"rate_last_updated"`
	WorkspaceID     int    `json:"workspace_id"`
}

type GetWorkspaceProjectsQuery struct {
	Active        *bool  `url:"active,omitempty"`
	Since         *int   `url:"since,omitempty"`
	Billable      *bool  `url:"billable,omitempty"`
	UserIDs       []int  `url:"user_ids,omitempty"`
	ClientIDs     []int  `url:"client_ids,omitempty"`
	GroupIDs      []int  `url:"group_ids,omitempty"`
	Name          string `url:"name,omitempty"`
	Page          *int   `url:"page,omitempty"`
	SortField     string `url:"sort_field,omitempty"`
	SortOrder     string `url:"sort_order,omitempty"`
	OnlyTemplates *bool  `url:"only_templates,omitempty"`
	PerPage       *int   `url:"per_page,omitempty"`
}

type GetWorkspaceProjectsResponse struct {
	ID                  int    `json:"id"`
	ClientID            int    `json:"client_id"`
	Cid                 int    `json:"cid"`
	Name                string `json:"name,omitempty"`
	Wid                 int    `json:"wid"`
	WorkspaceID         int    `json:"workspace_id"`
	At                  string `json:"at"`
	IsPrivate           bool   `json:"is_private"`
	Active              bool   `json:"active"`
	CreatedAt           string `json:"created_at"`
	ServerDeletedAt     string `json:"server_deleted_at,omitempty"`
	Color               string `json:"color"`
	Billable            bool   `json:"billable,omitempty"`
	Template            bool   `json:"template,omitempty"`
	AutoEstimates       bool   `json:"auto_estimates,omitempty"`
	EstimatedHours      int    `json:"estimated_hours,omitempty"`
	Rate                int    `json:"rate"`
	RateLastUpdated     string `json:"rate_last_updated,omitempty"`
	Currency            string `json:"currency,omitempty"`
	Recurring           bool   `json:"recurring"`
	RecurringParameters []struct {
		CustomPeriod       int    `json:"custom_period"`
		EstimatedSeconds   int    `json:"estimated_seconds"`
		ParameterEndDate   string `json:"parameter_end_date,omitempty"`
		ParameterStartDate string `json:"parameter_start_date"`
		Period             string `json:"period"`
		ProjectStartDate   string `json:"project_start_date"`
	} `json:"recurring_parameters"`
	CurrentPeriod struct {
		EndDate   string `json:"end_date"`
		StartDate string `json:"start_date"`
	} `json:"current_period"`
	FixedFee    int `json:"fixed_fee,omitempty"`
	ActualHours int `json:"actual_hours,omitempty"`
}

type CreateWorkspaceProjectRequest struct {
	Active              bool   `json:"active"`
	AutoEstimates       bool   `json:"auto_estimates,omitempty"`
	Billable            bool   `json:"billable,omitempty"`
	ClientID            int    `json:"cid,omitempty"`
	Color               string `json:"color"`
	Currency            string `json:"currency,omitempty"`
	EndDate             string `json:"end_date"`
	EstimatedHours      int    `json:"estimated_hours,omitempty"`
	FixedFee            int    `json:"fixed_fee,omitempty"`
	IsPrivate           bool   `json:"is_private"`
	Name                string `json:"name"`
	Rate                int    `json:"rate"`
	RateChangeMode      string `json:"rate_change_mode,omitempty"`
	Recurring           bool   `json:"recurring"`
	RecurringParameters []struct {
		CustomPeriod       int    `json:"custom_period"`
		EstimatedSeconds   int    `json:"estimated_seconds"`
		ParameterEndDate   string `json:"parameter_end_date,omitempty"`
		ParameterStartDate string `json:"parameter_start_date"`
		Period             string `json:"period"`
		ProjectStartDate   string `json:"project_start_date"`
	} `json:"recurring_parameters"`
	StartDate   string `json:"start_date"`
	Template    bool   `json:"template,omitempty"`
	TemplateID  int    `json:"template_id,omitempty"`
	WorkspaceID int    `json:"wid"`
}

type CreateWorkspaceProjectResponse struct {
	ActualHours   int    `json:"actual_hours,omitempty"`
	At            string `json:"at"`
	CreatedAt     string `json:"created_at"`
	CurrentPeriod struct {
		EndDate   string `json:"end_date"`
		StartDate string `json:"start_date"`
	} `json:"current_period"`
	FirstTimeEntry  string `json:"first_time_entry"`
	ID              int    `json:"id"`
	RateLastUpdated string `json:"rate_last_updated,omitempty"`
	ServerDeletedAt string `json:"server_deleted_at,omitempty"`
	Wid             int    `json:"wid"`
	WorkspaceID     int    `json:"workspace_id"`
}

type UpdateMultipleWorkspaceProjectsRequest struct {
	Operator string `json:"op"`
	Path     string `json:"path"`
	Value    string `json:"value"`
}

type UpdateMultipleWorkspaceProjectsResponse struct {
	Failure []struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	} `json:"failure"`
	Success []int `json:"success"`
}

type GetWorkspaceProjectQuery struct {
	WithFirstTimeEntry bool `json:"with_first_time_entry,omitempty"`
}

type GetWorkspaceProjectResponse struct {
	ActualHours   int    `json:"actual_hours,omitempty"`
	At            string `json:"at"`
	CreatedAt     string `json:"created_at"`
	CurrentPeriod struct {
		EndDate   string `json:"end_date"`
		StartDate string `json:"start_date"`
	} `json:"current_period"`
	FirstTimeEntry  string `json:"first_time_entry"`
	ID              int    `json:"id"`
	RateLastUpdated string `json:"rate_last_updated,omitempty"`
	ServerDeletedAt string `json:"server_deleted_at,omitempty"`
	Wid             int    `json:"wid"`
	WorkspaceID     int    `json:"workspace_id"`
}

type UpdateWorkspaceProjectRequest struct {
	Active              bool   `json:"active"`
	AutoEstimates       bool   `json:"auto_estimates,omitempty"`
	Billable            bool   `json:"billable,omitempty"`
	Cid                 int    `json:"cid"`
	ClientID            int    `json:"client_id,omitempty"`
	ClientName          string `json:"client_name,omitempty"`
	Color               string `json:"color"`
	Currency            string `json:"currency,omitempty"`
	EndDate             string `json:"end_date"`
	EstimatedHours      int    `json:"estimated_hours,omitempty"`
	FixedFee            int    `json:"fixed_fee,omitempty"`
	IsPrivate           bool   `json:"is_private"`
	Name                string `json:"name"`
	Rate                int    `json:"rate,omitempty"`
	RateChangeMode      string `json:"rate_change_mode,omitempty"`
	Recurring           bool   `json:"recurring,omitempty"`
	RecurringParameters []struct {
		CustomPeriod       int    `json:"custom_period"`
		EstimatedSeconds   int    `json:"estimated_seconds"`
		ParameterEndDate   string `json:"parameter_end_date,omitempty"`
		ParameterStartDate string `json:"parameter_start_date"`
		Period             string `json:"period"`
		ProjectStartDate   string `json:"project_start_date"`
	} `json:"recurring_parameters,omitempty"`
	StartDate  string `json:"start_date"`
	Template   bool   `json:"template,omitempty"`
	TemplateID int    `json:"template_id,omitempty"`
}

type UpdateWorkspaceProjectResponse struct {
	ActualHours   int    `json:"actual_hours,omitempty"`
	At            string `json:"at"`
	CreatedAt     string `json:"created_at"`
	CurrentPeriod struct {
		EndDate   string `json:"end_date"`
		StartDate string `json:"start_date"`
	} `json:"current_period"`
	FirstTimeEntry  string `json:"first_time_entry"`
	ID              int    `json:"id"`
	RateLastUpdated string `json:"rate_last_updated,omitempty"`
	ServerDeletedAt string `json:"server_deleted_at,omitempty"`
	Wid             int    `json:"wid"`
	WorkspaceID     int    `json:"workspace_id"`
}
