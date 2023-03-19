package acount

type GetUserInfoListForm struct {
	PageSize  int    `form:"page_size" binding:"required,gte=10,lte=100" json:"page_size"`
	PageIndex int    `form:"page_index" binding:"required,gte=1" json:"page_index"`
	SortField string `form:"sort_field" binding:"required" json:"sort_field"`
	SortOrder string `form:"sort_order" binding:"required" json:"sort_order"`
}
