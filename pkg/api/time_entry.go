package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ckotzbauer/go-toggl/pkg/model"
)

func (ctx *TogglContext) GetTimeEntries(query model.GetTimeEntriesQuery) (*[]model.GetTimeEntriesResponse, error) {
	return doRequest[[]model.GetTimeEntriesResponse](http.MethodGet, fmt.Sprintf("/api/v9/me/time_entries%s", formatQuery(query)), ctx)
}

func (ctx *TogglContext) GetCurrentTimeEntry(query model.GetTimeEntriesQuery) (*[]model.GetCurrentTimeEntryResponse, error) {
	return doRequest[[]model.GetCurrentTimeEntryResponse](http.MethodGet, "/api/v9/me/time_entries/current", ctx)
}

func (ctx *TogglContext) CreateTimeEntry(workspaceId int, req *model.CreateTimeEntryRequest) (*model.CreateTimeEntryResponse, error) {
	return doRequestWithBody[model.CreateTimeEntryRequest, model.CreateTimeEntryResponse](http.MethodPost, fmt.Sprintf("/api/v9/workspaces/%v/time_entries", workspaceId), ctx, req)
}

func (ctx *TogglContext) UpdateMultipleTimeEntries(workspaceId int, timeEntryIds []int, req *[]model.UpdateMultipleTimeEntriesRequest) (*[]model.UpdateMultipleTimeEntriesResponse, error) {
	ids := strings.Trim(strings.Replace(fmt.Sprint(timeEntryIds), " ", ",", -1), "[]")
	return doRequestWithBody[[]model.UpdateMultipleTimeEntriesRequest, []model.UpdateMultipleTimeEntriesResponse](http.MethodPatch, fmt.Sprintf("/api/v9/workspaces/%v/time_entries/%v", workspaceId, ids), ctx, req)
}

func (ctx *TogglContext) UpdateTimeEntries(workspaceId, timeEntryId int, req *model.UpdateTimeEntryRequest) (*model.UpdateTimeEntryResponse, error) {
	return doRequestWithBody[model.UpdateTimeEntryRequest, model.UpdateTimeEntryResponse](http.MethodPut, fmt.Sprintf("/api/v9/workspaces/%v/time_entries/%v", workspaceId, timeEntryId), ctx, req)
}

func (ctx *TogglContext) DeleteTimeEntry(workspaceId, timeEntryId int) (*any, error) {
	return doRequest[any](http.MethodDelete, fmt.Sprintf("/api/v9/workspaces/%v/time_entries/%v", workspaceId, timeEntryId), ctx)
}

func (ctx *TogglContext) StopTimeEntry(workspaceId, timeEntryId int) (*model.StopTimeEntryResponse, error) {
	return doRequest[model.StopTimeEntryResponse](http.MethodPatch, fmt.Sprintf("/api/v9/workspaces/%v/time_entries/%v/stop", workspaceId, timeEntryId), ctx)
}
