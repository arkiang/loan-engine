package dto

type CommonFilter struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}
