package model

type GetProjectUsersQuery struct {
	ProjectIDs       *string `json:"project_ids"`
	WithGroupMembers *bool   `json:"with_group_members"`
}

type baseProjectUser struct {
	LabourCost     int    `json:"labour_cost"`
	Manager        bool   `json:"manager"`
	ProjectID      int    `json:"project_id"`
	Rate           *int   `json:"rate"`
	RateChangeMode string `json:"rate_change_mode"`
	UserID         int    `json:"user_id"`
}

type projectUser struct {
	baseProjectUser
	At              string `json:"at"`
	Gid             int    `json:"gid"`
	GroupID         int    `json:"group_id"`
	ID              int    `json:"id"`
	RateLastUpdated string `json:"rate_last_updated"`
	WorkspaceID     int    `json:"workspace_id"`
}

type GetProjectUsersResponse struct {
	projectUser
}

type AddProjectUserRequest struct {
	baseProjectUser
}

type AddProjectUserResponse struct {
	projectUser
}

type UpdateMultipleProjectUsersResponse struct {
	Failure []struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	} `json:"failure"`
	Success []int `json:"success"`
}

type UpdateProjectUserRequest struct {
	baseProjectUser
}

type UpdateProjectUserResponse struct {
	projectUser
}

type GetWorkspaceProjectsQuery struct {
	Active        *bool   `json:"active"`
	Since         *int    `json:"since"`
	Billable      *bool   `json:"billable"`
	UserIDs       *[]int  `json:"user_ids"`
	ClientIDs     *[]int  `json:"client_ids"`
	GroupIDs      *[]int  `json:"group_ids"`
	Name          *string `json:"name"`
	Page          *int    `json:"page"`
	SortField     *string `json:"sort_field"`
	SortOrder     *string `json:"sort_order"`
	OnlyTemplates *bool   `json:"only_templates"`
	PerPage       *int    `json:"per_page"`
}

type baseProject struct {
	Active              bool    `json:"active"`
	AutoEstimates       *bool   `json:"auto_estimates"`
	Billable            *bool   `json:"billable"`
	Cid                 int     `json:"cid"`
	ClientID            *int    `json:"client_id"`
	ClientName          string  `json:"client_name"`
	Color               string  `json:"color"`
	Currency            *string `json:"currency"`
	EndDate             string  `json:"end_date"`
	EstimatedHours      *int    `json:"estimated_hours"`
	FixedFee            int     `json:"fixed_fee"`
	IsPrivate           bool    `json:"is_private"`
	Name                string  `json:"name"`
	Rate                int     `json:"rate"`
	RateChangeMode      string  `json:"rate_change_mode"`
	Recurring           bool    `json:"recurring"`
	RecurringParameters []struct {
		CustomPeriod       int     `json:"custom_period"`
		EstimatedSeconds   int     `json:"estimated_seconds"`
		ParameterEndDate   *string `json:"parameter_end_date"`
		ParameterStartDate string  `json:"parameter_start_date"`
		Period             string  `json:"period"`
		ProjectStartDate   string  `json:"project_start_date"`
	} `json:"recurring_parameters"`
	StartDate  string `json:"start_date"`
	Template   *bool  `json:"template"`
	TemplateID *int   `json:"template_id"`
}

type project struct {
	ActualHours   *int   `json:"actual_hours"`
	At            string `json:"at"`
	CreatedAt     string `json:"created_at"`
	CurrentPeriod struct {
		EndDate   string `json:"end_date"`
		StartDate string `json:"start_date"`
	} `json:"current_period"`
	FirstTimeEntry  string  `json:"first_time_entry"`
	ID              int     `json:"id"`
	RateLastUpdated *string `json:"rate_last_updated"`
	ServerDeletedAt *string `json:"server_deleted_at"`
	Wid             int     `json:"wid"`
	WorkspaceID     int     `json:"workspace_id"`
}

type GetWorkspaceProjectsResponse struct {
	project
}

type CreateWorkspaceProjectRequest struct {
	baseProject
}

type CreateWorkspaceProjectResponse struct {
	project
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
	WithFirstTimeEntry *bool `json:"with_first_time_entry"`
}

type GetWorkspaceProjectResponse struct {
	project
}

type UpdateWorkspaceProjectRequest struct {
	baseProject
}

type UpdateWorkspaceProjectResponse struct {
	project
}
