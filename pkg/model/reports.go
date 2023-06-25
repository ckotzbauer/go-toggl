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

type SearchTimeEntriesRequest struct {
	Billable           bool     `json:"billable,omitempty"`
	ClientIDs          []int    `json:"client_ids,omitempty"`
	Description        string   `json:"description,omitempty"`
	EndDate            string   `json:"end_date,omitempty"`
	FirstID            int      `json:"first_id,omitempty"`
	FirstRowNumber     int      `json:"first_row_number,omitempty"`
	FirstTimestamp     int      `json:"first_timestamp,omitempty"`
	GroupIDs           []int    `json:"group_ids,omitempty"`
	Grouped            bool     `json:"grouped,omitempty"`
	HideAmounts        bool     `json:"hide_amounts,omitempty"`
	MaxDurationSeconds int      `json:"max_duration_seconds,omitempty"`
	MinDurationSeconds int      `json:"min_duration_seconds,omitempty"`
	OrderBy            string   `json:"order_by,omitempty"`
	OrderDir           string   `json:"order_dir,omitempty"`
	PostedFields       []string `json:"postedFields,omitempty"`
	ProjectIDs         []int    `json:"project_ids,omitempty"`
	Rounding           int      `json:"rounding,omitempty"`
	RoundingMinutes    int      `json:"rounding_minutes,omitempty"`
	StartTime          string   `json:"startTime,omitempty"`
	StartDate          string   `json:"start_date,omitempty"`
	TagIDs             []int    `json:"tag_ids,omitempty"`
	TaskIDs            []int    `json:"task_ids,omitempty"`
	TimeEntryIDs       []int    `json:"time_entry_ids,omitempty"`
	UserIDs            []int    `json:"user_ids,omitempty"`
}

type SearchTimeEntriesResponse struct {
	UserID                int    `json:"user_id"`
	UserName              string `json:"username"`
	ProjectID             int    `json:"project_id"`
	TaskID                *int   `json:"task_id"`
	Billable              bool   `json:"billable"`
	Description           string `json:"description"`
	TagIDs                []int  `json:"tag_ids"`
	BillableAmountInCents int    `json:"billable_amount_in_cents"`
	HourlyRateInCents     int    `json:"hourly_rate_in_cents"`
	Currency              string `json:"currency"`
	RowNumber             int    `json:"row_number"`
	TimeEntries           []struct {
		ID      int    `json:"id"`
		Seconds int    `json:"seconds"`
		Start   string `json:"start"`
		Stop    string `json:"stop"`
		At      string `json:"at"`
	} `json:"time_entries"`
}
