package model

type CreateInvitationRequest struct {
	Emails     []string `json:"emails"`
	Workspaces []struct {
		Admin       bool `json:"admin"`
		WorkspaceID int  `json:"workspace_id"`
	} `json:"workspaces"`
}

type CreateInvitationResponse struct {
	Data []struct {
		Email          string `json:"email"`
		InvitationID   int    `json:"invitation_id"`
		Invite_url     string `json:"invite_url"`
		OrganizationID int    `json:"organization_id"`
		RecipientID    int    `json:"recipient_id"`
		SenderID       int    `json:"sender_id"`
	} `json:"data"`
	Messages []string `json:"emails"`
}
