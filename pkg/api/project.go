package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ckotzbauer/go-toggl/pkg/model"
)

func (ctx *TogglContext) GetProjectUsers(workspaceId int, query model.GetProjectUsersQuery) (*[]model.GetOrganizationUsersResponse, error) {
	return doRequest[[]model.GetOrganizationUsersResponse](http.MethodGet, fmt.Sprintf("/api/v9/workspaces/%v/project_users%s", workspaceId, formatQuery(query)), ctx)
}

func (ctx *TogglContext) AddProjectUser(workspaceId int, req *model.AddProjectUserRequest) (*model.AddProjectUserResponse, error) {
	return doRequestWithBody[model.AddProjectUserRequest, model.AddProjectUserResponse](http.MethodPost, fmt.Sprintf("/api/v9/workspaces/%v/project_users", workspaceId), ctx, req)
}

func (ctx *TogglContext) UpdateMultipleProjectUsers(workspaceId int, projectUserIds []int) (*model.UpdateMultipleProjectUsersResponse, error) {
	ids := strings.Trim(strings.Replace(fmt.Sprint(projectUserIds), " ", ",", -1), "[]")
	return doRequest[model.UpdateMultipleProjectUsersResponse](http.MethodPatch, fmt.Sprintf("/api/v9/workspaces/%v/project_users/%s", workspaceId, ids), ctx)
}

func (ctx *TogglContext) UpdateProjectUser(workspaceId, projectUserId int, req *model.UpdateProjectUserRequest) (*model.UpdateProjectUserResponse, error) {
	return doRequestWithBody[model.UpdateProjectUserRequest, model.UpdateProjectUserResponse](http.MethodPut, fmt.Sprintf("/api/v9/workspaces/%v/project_users/%v", workspaceId, projectUserId), ctx, req)
}

func (ctx *TogglContext) DeleteProjectUsers(workspaceId, projectUserId int) (*any, error) {
	return doRequest[any](http.MethodDelete, fmt.Sprintf("/api/v9/workspaces/%v/project_users/%v", workspaceId, projectUserId), ctx)
}

func (ctx *TogglContext) GetWorkspaceProjects(workspaceId int, query model.GetWorkspaceProjectsQuery) (*[]model.GetWorkspaceProjectsResponse, error) {
	return doRequest[[]model.GetWorkspaceProjectsResponse](http.MethodGet, fmt.Sprintf("/api/v9/workspaces/%v/projects%s", workspaceId, formatQuery(query)), ctx)
}

func (ctx *TogglContext) CreateWorkspaceProject(workspaceId int, req *model.CreateWorkspaceProjectRequest) (*model.CreateWorkspaceProjectResponse, error) {
	return doRequestWithBody[model.CreateWorkspaceProjectRequest, model.CreateWorkspaceProjectResponse](http.MethodPost, fmt.Sprintf("/api/v9/workspaces/%v/projects", workspaceId), ctx, req)
}

func (ctx *TogglContext) UpdateMultipleWorkspaceProjects(workspaceId int, projectIds []int, req *model.UpdateMultipleWorkspaceProjectsRequest) (*model.UpdateMultipleWorkspaceProjectsResponse, error) {
	ids := strings.Trim(strings.Replace(fmt.Sprint(projectIds), " ", ",", -1), "[]")
	return doRequestWithBody[model.UpdateMultipleWorkspaceProjectsRequest, model.UpdateMultipleWorkspaceProjectsResponse](http.MethodPatch, fmt.Sprintf("/api/v9/workspaces/%v/project/%s", workspaceId, ids), ctx, req)
}

func (ctx *TogglContext) GetWorkspaceProject(workspaceId, projectId int, query model.GetWorkspaceProjectQuery) (*model.GetWorkspaceProjectResponse, error) {
	return doRequest[model.GetWorkspaceProjectResponse](http.MethodGet, fmt.Sprintf("/api/v9/workspaces/%v/projects/%v%s", workspaceId, projectId, formatQuery(query)), ctx)
}

func (ctx *TogglContext) UpdateWorkspaceProject(workspaceId, projectId int, req *model.UpdateWorkspaceProjectRequest) (*model.UpdateWorkspaceProjectResponse, error) {
	return doRequestWithBody[model.UpdateWorkspaceProjectRequest, model.UpdateWorkspaceProjectResponse](http.MethodPut, fmt.Sprintf("/api/v9/workspaces/%v/projects/%v", workspaceId, projectId), ctx, req)
}

func (ctx *TogglContext) DeleteWorkspaceProject(workspaceId, projectId int) (*any, error) {
	return doRequest[any](http.MethodDelete, fmt.Sprintf("/api/v9/workspaces/%v/projects/%v", workspaceId, projectId), ctx)
}
