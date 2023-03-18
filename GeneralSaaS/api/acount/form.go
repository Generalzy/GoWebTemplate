package acount

type GetUserInfoListForm struct {
	PageSize  int    `form:"page_size" json:"page_size" binding:"required"`
	PageIndex int    `form:"page_index" json:"page_index" binding:"required"`
	SortField string `form:"sort_field" json:"sort_field" binding:"required"`
	SortOrder string `form:"sort_order" json:"sort_order" binding:"required"`
}
