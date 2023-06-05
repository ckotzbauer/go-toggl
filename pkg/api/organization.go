package api

import (
	"fmt"
	"net/http"

	"github.com/ckotzbauer/go-toggl/pkg/model"
)

func (ctx *TogglContext) CreateOrganization(req *model.CreateOrganizationRequest) (*model.CreateOrganizationResponse, error) {
	return doRequestWithBody[model.CreateOrganizationRequest, model.CreateOrganizationResponse](http.MethodPost, "/api/v9/organizations", ctx, req)
}

func (ctx *TogglContext) GetOrganization(organizationId int) (*model.GetOrganizationResponse, error) {
	return doRequest[model.GetOrganizationResponse](http.MethodGet, fmt.Sprintf("/api/v9/organizations/%v", organizationId), ctx)
}

func (ctx *TogglContext) UpdateOrganization(organizationId int, req *model.UpdateOrganizationRequest) (*any, error) {
	return doRequestWithBody[model.UpdateOrganizationRequest, any](http.MethodPut, fmt.Sprintf("/api/v9/organizations/%v", organizationId), ctx, req)
}

func (ctx *TogglContext) GetOrganizationUsers(organizationId int, query model.GetOrganizationUsersQuery) (*[]model.GetOrganizationUsersResponse, error) {
	return doRequest[[]model.GetOrganizationUsersResponse](http.MethodGet, fmt.Sprintf("/api/v9/organizations/%v/users%s", organizationId, formatQuery(query)), ctx)
}

func (ctx *TogglContext) UpdateOrganizationUsers(organizationId int, req *model.UpdateOrganizationUsersRequest) (*any, error) {
	return doRequestWithBody[model.UpdateOrganizationUsersRequest, any](http.MethodPatch, fmt.Sprintf("/api/v9/organizations/%v/users", organizationId), ctx, req)
}

func (ctx *TogglContext) LeaveOrganization(organizationId int) (*any, error) {
	return doRequest[any](http.MethodDelete, fmt.Sprintf("/api/v9/organizations/%v/users/leave", organizationId), ctx)
}

func (ctx *TogglContext) UpdateOrganizationUser(organizationId, organizationUserId int, req *model.UpdateOrganizationUserRequest) (*any, error) {
	return doRequestWithBody[model.UpdateOrganizationUserRequest, any](http.MethodPatch, fmt.Sprintf("/api/v9/organizations/%v/users/%v", organizationId, organizationUserId), ctx, req)
}

func (ctx *TogglContext) GetWorkspaceStatistics(organizationId int) (*any, error) {
	return doRequest[any](http.MethodGet, fmt.Sprintf("/api/v9/organizations/%v/workspaces/statistics", organizationId), ctx)
}

func (ctx *TogglContext) UpdateUserAssignments(organizationId, workspaceId int, req *model.UpdateUserAssignmentsRequest) (*any, error) {
	return doRequestWithBody[model.UpdateUserAssignmentsRequest, any](http.MethodPut, fmt.Sprintf("/api/v9/organizations/%v/workspaces/%v/assignments", organizationId, workspaceId), ctx, req)
}
