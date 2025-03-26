package apis

// Pager request 分頁資訊
type Pager struct {
	// 頁碼 page index
	Index int `form:"pi"`
	// 筆數 page size
	Size int `form:"ps"`
}

// Pagination response 分頁資訊
type Pagination struct {
	// 頁碼 page index
	Index int `json:"pi"`
	// 筆數 page size
	Size int `json:"ps"`
	// 總頁數 total pages
	TotalPage int `json:"total_page"`
	// 總筆數 total items
	TotalRow int `json:"total_row"`
}
