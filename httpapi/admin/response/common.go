package response

type IdSuccess struct {
	Id int `json:"id"`
}

type PidNameSuccess struct {
	Name string `json:"name"`
	Pid  uint32 `json:"pid"`
	Id   int    `json:"id"`
}

type PageResponse struct {
	Total int `json:"total"`
}

type CreatedAdminName struct {
	CreatedAdminName string `json:"created_admin_name"`
}
