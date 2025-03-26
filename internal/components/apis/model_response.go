package apis

type Meta struct {
	RequestID   string   `json:"requestID" example:"38df1107-94b9-498e-b24f-a0e82b77031b"`
	ProcessTime string   `json:"processTime,omitempty" swaggerignore:"true"`
	TimeLogs    []string `json:"timeLogs,omitempty" swaggerignore:"true"`
}

type Response struct {
	Status      int
	ContentType string
	Headers     map[string]string
	Data        []byte
	Error       error

	// TraceNamed
	// 會帶入 logger.Named 再根據 Named 設定的 level 映出對應內容
	// debug 所有 response ，內容 meta data
	// info 所有 response ，內容 meta
	// error 判斷 response error !=nil ，內容 response status meta data error
	TraceNamed string
}

type EmptyData struct{}

type StandardResponse struct {
	// Response 的 Meta 資訊
	Meta *Meta `json:"meta"`

	// PM 定義的 code
	Code string `json:"code"`

	// 資料內容
	Data any `json:"data"`

	// Error 資訊
	Error any `json:"error,omitempty"`
}

type StandardListResponse struct {
	StandardResponse
	Pagination `json:"pagination"`
}

type StandardError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
