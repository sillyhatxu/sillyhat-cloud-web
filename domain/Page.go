package domain

type Page struct {
	Content       interface{} `json:"content"`
	CurrentPage   int64       `json:"current_page"`
	PageSize      int64       `json:"page_size"`
	TotalElements int64       `json:"total_elements"`
	TotalPage     int64       `json:"total_page"`
	First         bool        `json:"first"`
	Last          bool        `json:"last"`
	Sort          interface{} `json:"sort"`
}

func BuildPage(content interface{}, currentPage int64, pageSize int64, totalElements int64, sort interface{}) *Page {
	totalPage := (totalElements + pageSize - 1) / pageSize
	return &Page{
		Content:       content,
		CurrentPage:   currentPage,
		PageSize:      pageSize,
		TotalElements: totalElements,
		TotalPage:     totalPage,
		First:         currentPage == 0,
		Last:          currentPage == totalPage-1,
		Sort:          sort,
	}
}
