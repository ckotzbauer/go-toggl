package model

type CreateOrganizationRequest struct {
	Name          string `json:"name"`
	WorkspaceName string `json:"workspace_name"`
}

type CreateOrganizationResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	WorkspaceName string `json:"workspace_name"`
	WorkspaceID   int    `json:"workspace_id"`
}

type GetOrganizationResponse struct {
	Admin                   bool    `json:"admin"`
	At                      string  `json:"at"`
	CreatedAt               string  `json:"created_at"`
	ID                      int     `json:"id"`
	IsChargify              bool    `json:"is_chargify"`
	IsMultiWorkspaceEnabled bool    `json:"is_multi_workspace_enabled"`
	IsUnified               bool    `json:"is_unified"`
	MaxWorkspaces           int     `json:"max_workspaces"`
	Name                    string  `json:"name"`
	Owner                   bool    `json:"owner"`
	PaymentMethods          string  `json:"payment_methods"`
	ServerDeletedAt         *string `json:"server_deleted_at"`
	SuspendedAt             string  `json:"suspended_at"`
	TrialInfo               struct {
		LastPricingPlanID *int    `json:"last_pricing_plan_id"`
		NextPaymentDate   *string `json:"next_payment_date"`
		Trial             bool    `json:"trial"`
		TrialAvailable    bool    `json:"trial_available"`
		TrialEndDate      *string `json:"trial_end_date"`
	} `json:"trial_info"`
	UserAccount int `json:"user_account"`
}

type UpdateOrganizationRequest struct {
	Name string `json:"name"`
}

type GetOrganizationUsersQuery struct {
	Filter       string `json:"filter"`
	ActiveStatus string `json:"active_status"`
	OnlyAdmins   string `json:"only_admins"`
	Groups       string `json:"groups"`
	Workspaces   string `json:"workspaces"`
	Page         int    `json:"page"`
	PerPage      int    `json:"per_page"`
	SortDir      string `json:"sort_dir"`
}

type GetOrganizationUsersResponse struct {
	Admin        bool   `json:"admin"`
	AvatarUrl    string `json:"avatar_url"`
	CanEditEmail bool   `json:"can_edit_email"`
	Email        string `json:"email"`
	Groups       struct {
		GroupID int    `json:"group_id"`
		Name    string `json:"name"`
	} `json:"groups"`
	ID             int    `json:"id"`
	Inactive       bool   `json:"inactive"`
	InvitationCode string `json:"invitation_code"`
	Joined         bool   `json:"joined"`
	Name           string `json:"name"`
	Owner          bool   `json:"owner"`
	UserID         int    `json:"user_id"`
	Workspaces     struct {
		Admin       bool   `json:"admin"`
		Name        string `json:"name"`
		WorkspaceID int    `json:"workspace_id"`
	} `json:"workspaces"`
}

type UpdateOrganizationUsersRequest struct {
	Delete []int `json:"delete"`
}

type UpdateOrganizationUserRequest struct {
	Email              string `json:"email"`
	Groups             []int  `json:"groups"`
	Inactive           bool   `json:"inactive"`
	Name               string `json:"name"`
	Organization_admin bool   `json:"organization_admin"`
	Workspaces         struct {
		Admin       bool   `json:"admin"`
		Name        string `json:"name"`
		WorkspaceID int    `json:"workspace_id"`
	} `json:"workspaces"`
}

type UpdateUserAssignmentsRequest struct {
	GroupID   int    `json:"group_id"`
	Joined    bool   `json:"joined"`
	Operation string `json:"operation"`
	UserID    int    `json:"user_id"`
}
