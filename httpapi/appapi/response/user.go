package response

type UserInfoDetail struct {
	BossUserId int    `json:"boss_user_id"`
	Phone      string `json:"phone"`
	IdCard     string `json:"id_card"`
	CardType   int    `json:"card_type"`
	Username   string `json:"username"`
}