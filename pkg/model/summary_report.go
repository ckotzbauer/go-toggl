package model

type GetSummaryReportRequest struct {
	EndDate   string `json:"end_date"`
	StartTime string `json:"start_time"`
	StartDate string `json:"start_date"`
}

type GetSummaryReportResponse struct {
	BillableSeconds int `json:"billable_seconds"`
	ProjectID       int `json:"project_id"`
	TrackedSeconds  int `json:"tracked_seconds"`
	UserID          int `json:"user_id"`
}
