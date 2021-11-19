package request

//添加(编辑)活动类型
type SetActivityTypeRequest struct {
	Id     *int    `json:"id"  form:"id"`
	Name   *string `json:"name"  form:"name" binding:"required"`
	Status *int    `json:"status" form:"status" binding:"required"`
}

//设置活动类型状态
type SetActivityTypeStatusRequest struct {
	Id     *int `json:"id"  form:"id" binding:"required"`
	Status *int `json:"status" form:"status" binding:"required,min=1,max=2"`
}

//活动类型分页查询
type ActivityTypePageListRequest struct {
	Name   string `json:"name"  form:"name"`
	Status *int   `json:"status" form:"status"`
	PageInfo
}

//添加(编辑)热门搜索
type SetHotSearchRequest struct {
	Id          *int    `json:"id"  form:"id"`
	Name        *string `json:"name"  form:"name" binding:"required"`
	SearchCount *int    `json:"search_count"  form:"search_count" binding:"required"`
	Status      *int    `json:"status" form:"status" binding:"required"`
}

//热门搜索分页查询
type HotSearchPageListRequest struct {
	Name        string  `json:"name"  form:"name"`
	Status      *int    `json:"status" form:"status"`
	SearchCount *string `json:"search_count"`
	PageInfo
}

//添加（编辑）活动
