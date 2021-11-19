package response

import "gserver/internal/store/ent"

type ActivityTypePageListResponse struct {
	List ent.ActivityTypes `json:"list"`
	Page PageResponse      `json:"page"`
}

type InformationClassifyListResponse struct {
	List ent.InformationClassifies `json:"list"`
	Page PageResponse              `json:"page"`
}

type MsgTypeListResponse struct {
	List ent.MessageTypes `json:"list"`
	Page PageResponse     `json:"page"`
}

type HotSearchListResponse struct {
	List ent.HotSearches `json:"list"`
	Page PageResponse    `json:"page"`
}
