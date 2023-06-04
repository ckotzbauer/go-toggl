package api

import (
	"fmt"
	"net/http"

	"github.com/ckotzbauer/go-toggl/pkg/model"
)

func (ctx *TogglContext) GetSummaryReport(workspaceId int, req *model.GetSummaryReportRequest) (*[]model.GetSummaryReportResponse, error) {
	return DoRequestWithBody[model.GetSummaryReportRequest, []model.GetSummaryReportResponse](http.MethodPost, fmt.Sprintf("/reports/api/v3/workspace/%v/projects/summary", workspaceId), ctx, req)
}
