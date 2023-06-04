package api

import (
	"fmt"
	"net/http"

	"github.com/ckotzbauer/go-toggl/pkg/model"
)

func (ctx *TogglContext) AcceptInvitation(invitationCode string) (*any, error) {
	return DoRequest[any](http.MethodPost, fmt.Sprintf("/api/v9/organizations/invitations/%s/accept", invitationCode), ctx)
}

func (ctx *TogglContext) RejectInvitation(invitationCode string) (*any, error) {
	return DoRequest[any](http.MethodPost, fmt.Sprintf("/api/v9/organizations/invitations/%s/reject", invitationCode), ctx)
}

func (ctx *TogglContext) ResendInvitation(organizationId int, invitationCode string) (*any, error) {
	return DoRequest[any](http.MethodPost, fmt.Sprintf("/api/v9/organizations/%v/invitations/%s/resend", organizationId, invitationCode), ctx)
}

func (ctx *TogglContext) CreateInvitation(organizationId int, req *model.CreateInvitationRequest) (*model.CreateInvitationResponse, error) {
	return DoRequestWithBody[model.CreateInvitationRequest, model.CreateInvitationResponse](http.MethodPost, fmt.Sprintf("/api/v9/organizations/%v/invitations", organizationId), ctx, req)
}
