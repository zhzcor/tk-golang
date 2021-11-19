package response

import (
	"time"
	"tkserver/internal/store/ent"
)

type RegisterSuccess struct {
	Token  string `json:"token"`
	UserId string `json:"user_id"`
}

type AdminLoginSuccess struct {
	Token         string   `json:"token"`
	AdminId       int      `json:"admin_id"`
	Email         string   `json:"email"`
	RealName      string   `json:"real_name"`
	Phone         string   `json:"phone"`
	RoleId        int      `json:"role_id"`
	PermissionIds []string `json:"permission_ids"`
	AdminAvatarId int      `json:"admin_avatar_id"`
	AvatarUrl     string   `json:"avatar_url"`
}

type SimpleUserInfoByPage struct {
	List []SimpleUserInfoResponse `json:"list"`
	Page PageResponse             `json:"page"`
}
type SimpleUserInfoResponse struct {
	Id             int       `json:"id"`
	Username       string    `json:"username"`
	Phone          string    `json:"phone"`
	AvatarUrl      string    `json:"avatar_url"`
	CreatedAt      time.Time `json:"created_at"`
	CardType       int       `json:"card_type"`
	IdCard         string    `json:"id_card"`
	Sex            int       `json:"sex"`
	Birthday       time.Time `json:"birthday"`
	CityId         int       `json:"city_id"`
	CityName       string    `json:"city_name"`
	CateId         int       `json:"cate_id"`
	CateName       string    `json:"cate_name"`
	ParentCateId   int       `json:"parent_cate_id"`
	ParentCateName string    `json:"parent_catename"`
}

type SimpleTeacherInfoByPage struct {
	List []SimpleTeacherInfoResponse `json:"list"`
	Page PageResponse                `json:"page"`
}
type SimpleTeacherInfoResponse struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Nickname  string         `json:"nickname"`
	Phone     string         `json:"phone"`
	AvatarUrl string         `json:"avatar_url"`
	CreatedAt *time.Time     `json:"created_at"`
	Sex       int            `json:"sex"`
	Majors    ent.Majors     `json:"majors"`
	Courses   []ent.KcCourse `json:"courses"`
	Status    int            `json:"status"`
}

type GetAdminPageListResponse struct {
	List []AdminDetail `json:"list"`
	Page PageResponse  `json:"page"`
}

type AdminDetail struct {
	ent.Admin
	RoleName string `json:"role_name"`
	RoleId   int    `json:"role_id"`
	Password string `json:"password"`
}

type RoleDetail struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type RolePageList struct {
	List []RoleDetail `json:"list"`
	Page PageResponse `json:"page"`
}

type RolePermissionsResponse struct {
	PermissionIds []string `json:"permission_ids"`
}

//操作日志
type OperationLogListResponse struct {
	List []OperationLogDetail `json:"list"`
	Page PageResponse         `json:"page"`
}
type OperationLogDetail struct {
	Id        int        `json:"id"`
	Ip        string     `json:"ip"`
	Record    string     `json:"record"`
	AdminId   int        `json:"admin_id"`
	AdminName string     `json:"admin_name"`
	CreatedAt *time.Time `json:"created_at"`
}

//app登录日志
type LoginUserListResponse struct {
	List []LoginUserDetail `json:"list"`
	Page PageResponse      `json:"page"`
}
type LoginUserDetail struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	UserId   int    `json:"user_id"`
	ent.UserLoginLog
}

//id获取老师信息
type TeacherDetailResponse struct {
	ent.Teacher
	AvatarUrl   string          `json:"avatar_url"`
	Majors      ent.Majors      `json:"majors"`
	Courses     []ent.KcCourse  `json:"courses"`
	TeacherTags ent.TeacherTags `json:"teacher_tags"`
}

type UserInfoDetail struct {
	BossUserId int    `json:"boss_user_id"`
	Phone      string `json:"phone"`
	IdCard     string `json:"id_card"`
	CardType   int    `json:"card_type"`
	Username   string `json:"username"`
}
