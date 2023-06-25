package api

import (
	"fmt"
	"net/http"

	"github.com/ckotzbauer/go-toggl/pkg/model"
)

func (ctx *TogglContext) GetSummaryReport(workspaceId int, req *model.GetSummaryReportRequest) (*[]model.GetSummaryReportResponse, error) {
	return doRequestWithBody[model.GetSummaryReportRequest, []model.GetSummaryReportResponse](http.MethodPost, fmt.Sprintf("/reports/api/v3/workspace/%v/projects/summary", workspaceId), ctx, req)
}

func (ctx *TogglContext) SearchTimeEntries(workspaceId int, req *model.SearchTimeEntriesRequest) (*[]model.SearchTimeEntriesResponse, error) {
	return doRequestWithBody[model.SearchTimeEntriesRequest, []model.SearchTimeEntriesResponse](http.MethodPost, fmt.Sprintf("/reports/api/v3/workspace/%v/search/time_entries", workspaceId), ctx, req)
}
