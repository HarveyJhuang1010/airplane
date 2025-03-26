package dto

type PaginationFilter struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type StandardResponse struct {
	// response meta
	Meta ResponseMeta `json:"meta"`
	// server response code
	Code string `json:"code,omitempty"`
	// server response message
	Message string `json:"message,omitempty"`
	// server response data
	Data interface{} `json:"data,omitempty"`
}

type ResponseMeta struct {
	// request id
	RequestID string `json:"request_id"`
	// process time
	ProcessTime string `json:"process_time,omitempty"`
	// time logs
	TimeLogs []string `json:"time_logs,omitempty"`
}

type ErrorResponse struct {
	// server response error code
	Code string `json:"code,omitempty" example:"2T-001"`
	// server response error message
	Message string `json:"message,omitempty" example:"unexpected error"`
}
