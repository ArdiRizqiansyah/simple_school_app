package dto

type WebResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type Page struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	From        int `json:"from"`
	To          int `json:"to"`
	TotalData   int `json:"total_data"`
	TotalPage   int `json:"total_page"`
	LastPage    int `json:"last_page"`
}
