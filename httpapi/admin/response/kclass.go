package response

import "gserver/internal/store/ent"

//班级详情
type ClassDetail struct {
	IdSuccess
	ClassTitle       string     `json:"class_title"`
	ClassCode        string     `json:"class_code"`
	CateId           int        `json:"cate_id"`
	CateName         string     `json:"cate_name"`
	CityId           int        `json:"city_id"`
	CityName         string     `json:"city_name"`
	Majors           ent.Majors `json:"majors"`
	ClassCoverImgId  int        `json:"class_cover_img_id"`
	ClassCoverImgUrl string     `json:"class_cover_img_url"`
	ClassDesc        string     `json:"class_desc"`
	Price            float64    `json:"price"`
	ClassPeriodsType int        `json:"class_periods_type"`
	IsDisplay        int        `json:"is_display"`
	Status           int        `json:"status"`
	StudentCount     int        `json:"student_count"`
	CourseCount      int        `json:"course_count"`
}

//班级列表（分页）
type ClassPageListResponse struct {
	List      []ClassPageDetail `json:"list"`
	Page      PageResponse      `json:"page"`
	Statistic CourseStatistic   `json:"statistic"`
}
type ClassPageDetail struct {
	ItemCategoryName string `json:"item_category_name"`
	CreatedAdminName string `json:"created_admin_name"`
	ent.KcClass
}

//班级课程列表
type ClassCoursePageListResponse struct {
	List []ClassCourseDetail `json:"list"`
	Page PageResponse        `json:"page"`
}
type ClassCourseDetail struct {
	CourseId          int     `json:"course_id"`
	CourseName        string  `json:"course_name"`
	Price             float64 `json:"price"`
	Type              int     `json:"type"`
	CourseCoverImgId  int     `json:"course_cover_img_id"`
	CourseCoverImgUrl string  `json:"course_cover_img_url"`
	TeacherId         int     `json:"teacher_id"`
	TeacherName       string  `json:"teacher_name"`
	TeacherAvatar     string  `json:"teacher_avatar"`
	PeriodsType       int     `json:"periods_type"`
	IsSelected        int     `json:"is_selected"`
}

//班级老师列表
type ClassTeacherPageListResponse struct {
	List []ClassTeacherDetail `json:"list"`
	Page PageResponse         `json:"page"`
}

type ClassTeacherDetail struct {
	TeacherName string `json:"teacher_name"`
	Avatar      string `json:"avatar"`
	ent.KcClassTeacher
}

//获取班级已添加teacher id列表
type GetTeacherIdsInClass struct {
	TeacherIds []int `json:"teacher_ids"`
}

//获取班级已添加user id列表
type GetUserIdsInClass struct {
	UserIds []int `json:"user_ids"`
}

//班级用户列表
type ClassUserPageListResponse struct {
	List []ClassUserDetail `json:"list"`
	Page PageResponse      `json:"page"`
}

type ClassUserDetail struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	ent.KcUserClass
}

//获取班级班主任
type GetClassMasterTeacher struct {
	ClassId             int    `json:"class_id"`
	ClassHeadMasterId   int    `json:"class_head_master_id"`
	ClassHeadMasterName string `json:"class_head_master_name"`
	Avatar              string `json:"avatar"`
}
