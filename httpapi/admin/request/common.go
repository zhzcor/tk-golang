package request

type IdOnly struct {
	Id int `json:"id"  form:"id" binding:"required"`
}
type IdPtrOnly struct {
	Id *int `json:"id"  form:"id" binding:"required"`
}
type IdPtrNillable struct {
	Id *int `json:"id"  form:"id"`
}

type NameOnly struct {
	Name string `json:"name"  form:"name" binding:"required"`
}
type NamePtrOnly struct {
	Name *string `json:"name"  form:"name" binding:"required"`
}

type PageInfo struct {
	Page     *int `json:"current"  form:"current"`
	PageSize *int `json:"page_size"  form:"page_size"`
}
