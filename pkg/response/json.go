package response

type Meta struct {
	Page     int `json:"page"`
	Pages    int `json:"pages"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
}

type Pagination struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    Meta        `json:"meta"`
	Total   int         `json:"total"`
}
