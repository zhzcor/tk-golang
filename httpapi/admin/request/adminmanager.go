package request

import "time"

type LoginAdmin struct {
	Phone    string `json:"phone"  form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	CityName string `json:"city_name" form:"CityName"`
}

type AddRole struct {
	IdPtrNillable
	RoleName string `json:"name"  form:"name" binding:"required"`
	Status   int    `json:"status" form:"status" binding:"required"`
}

type AddPermission struct {
	PermissionName *string `json:"name"  form:"name" binding:"required"`
	Desc           *string `json:"desc" form:"desc"`
	Pid            *int    `json:"pid" form:"pid" binding:"required,gte=0"`
}

type AddAdminRole struct {
	AdminId *int  `json:"admin_id"  form:"admin_id" binding:"required,gte=0"`
	RoleIds []int `json:"role_ids" form:"desc"`
}

type AddRolePermission struct {
	RoleId        *int     `json:"role_id"  form:"role_id" binding:"required,gte=0"`
	PermissionIds []string `json:"permission_ids" form:"permission_ids"`
}

type SyncUserFromBoss struct {
	UserName   string `json:"username"  form:"username" binding:"required"`
	BossUserId int    `json:"boss_user_id" form:"boss_user_id" binding:"required"`
	CardType   int    `json:"card_type" form:"card_type" binding:"required`
	IdCard     string `json:"id_card" form:"id_card" binding:"required"`
	Phone      string `json:"phone" form:"phone" binding:"required"`
}

type SyncAdminFromBoss struct {
	RealName    *string `json:"real_name"  form:"real_name" binding:"required"`
	BossAdminId *int    `json:"boss_admin_id" form:"boss_admin_id" binding:"required"`
	Email       *string `json:"email" form:"email" binding:"required`
	Phone       *string `json:"phone" form:"phone" binding:"required"`
	ThirdOpenId *string `json:"third_open_id" form:"third_open_id" binding:"required"`
	Platform    *int    `json:"platform" form:"platform" binding:"required"`
	IsActive    *int    `json:"is_active" form:"is_active" binding:"required"`
	SyncType    int     `json:"sync_type" form:"sync_type"`
}

type GetUserInfo struct {
	Phone    string `json:"phone" form:"phone"`
	Username string `json:"username" form:"username"`
	PageInfo
}

type GetTeacherInfo struct {
	Phone  string `json:"phone" form:"phone"`
	Name   string `json:"name" form:"name"`
	Status *int   `json:"status" form:"status"`
	PageInfo
}

type GetAdminPageList struct {
	Phone    *string `json:"phone" form:"phone"`
	RealName *string `json:"real_name" form:"real_name"`
	Status   *int    `json:"status" form:"status"`
	RoleId   *int    `json:"role_id" form:"role_id"`
	PageInfo
}

type SetAdminRequest struct {
	Id *int `json:"id" form:"id" binding:"required"`
	//Phone    *string `json:"phone" form:"phone" binding:"required"`
	//RealName *string `json:"real_name" form:"real_name" binding:"required"`
	Password *string `json:"password" form:"password"`
	RoleId   *int    `json:"role_id" form:"role_id" binding:"required"`
	Remark   *string `json:"remark" form:"remark`
	//Email    *string `json:"email" form:"email" binding:"required"`
}

type SetAdminStatus struct {
	IdPtrOnly
	Status *int `json:"status" form:"status" binding:"required"`
}

type GetRoleList struct {
	Name   *string `json:"name" form:"name"`
	Status *int    `json:"status" form:"status"`
	PageInfo
}

type RoleId struct {
	RoleId *int `json:"role_id" form:"role_id" binding:"required"`
}

//修改密码
type ResetPassword struct {
	IdPtrOnly
	OldPassword     *string `json:"old_password" form:"old_password" binding:"required"`
	ConfirmPassword *string `json:"confirm_password" form:"confirm_password" binding:"required"`
}

//上传头像
type UpdateAdminAvatar struct {
	IdPtrOnly
	AdminAvatarId *int `json:"admin_avatar_id" form:"admin_avatar_id" binding:"required"`
}

//添加（编辑）老师
type SetTeacher struct {
	IdPtrNillable
	NamePtrOnly
	Phone       *string  `json:"phone" form:"phone" binding:"required"`
	Nickname    string   `json:"nickname" form:"nickname"`
	Sex         *int     `json:"sex" form:"sex" binding:"required"`
	MajorIds    []int    `json:"major_ids" form:"major_ids"`
	CourseIds   []int    `json:"course_ids" form:"course_ids"`
	Email       string   `json:"email" form:"email"`
	SortOrder   int      `json:"sort_order" form:"sort_order"`
	AvatarId    int      `json:"avatar_id" form:"avatar_id" binding:"required"`
	TeachingAge int      `json:"teaching_age" form:"teaching_age"`
	TeacherTags []string `json:"teacher_tags" form:"teacher_tags"`
	SubTitle    string   `json:"sub_title" form:"sub_title"`
	Detail      string   `json:"detail" form:"detail"`
	Status      int      `json:"status" form:"status" binding:"required"`
}

type ScanLogin struct {
	Code     *string `json:"code" form:"code" binding:"required"`
	Password string
}

type PhoneRequest struct {
	Phone *string `json:"phone"  form:"phone" binding:"required"`
}

type SetUser struct {
	IdPtrNillable
	Phone      *string    `json:"phone"  form:"phone" binding:"required"`
	Username   *string    `json:"username" form:"username" binding:"required"`
	IdCard     *string    `json:"id_card" form:"id_card"`
	CardType   *int       `json:"card_type" form:"card_type"`
	BossUserId *int       `json:"boss_user_id" form:"boss_user_id"`
	Sex        *int       `json:"sex" form:"sex"`
	Birthday   *time.Time `json:"birthday" form:"birthday"`
	CityId     int        `json:"city_id" form:"city_id"`
	CateId     int        `json:"cate_id" form:"cate_id"`
}
