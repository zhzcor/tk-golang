package request

import (
	"time"
)

type SetClass struct {
	IdPtrNillable
	ClassTitle       *string  `json:"class_title" form:"class_title" binding:"required"`
	ClassCode        *string  `json:"class_code" form:"class_code"`
	CateId           *int     `json:"cate_id" form:"cate_id"`
	MajorIds         []int    `json:"major_ids" form:"major_ids"`
	CityId           *int     `json:"city_id" form:"city_id"`
	ClassCoverImgId  *int     `json:"class_cover_img_id" form:"class_cover_img_id"`
	ClassDesc        *string  `json:"class_desc" form:"class_desc"`
	ClassPeriodsType *int     `json:"class_periods_type" form:"class_periods_type"`
	Price            *float64 `json:"price" form:"price"`
	IsDisplay        *int     `json:"is_display" form:"is_display"`
}

//设置班级状态
type SetClassStatus struct {
	Id     *int `json:"id"  form:"id" binding:"required"`
	Status *int `json:"status" form:"status" binding:"required,min=1,max=3"`
}

//班级列表（分页）
type ClassPageListRequest struct {
	ClassTitle *string `json:"class_title" form:"class_title"`
	Status     *int    `json:"status" form:"status"`
	PageInfo
}

//班级课程列表
type ClassCoursePageList struct {
	IdPtrOnly
	PageInfo
	CourseName *string `json:"course_name"`
}

//设置班级课程
type SetClassCourse struct {
	CourseIds []int `json:"course_ids"  form:"course_ids" binding:"required"`
	ClassId   *int  `json:"class_id" form:"class_id" binding:"required"`
}

//移除班级课程
type RemoveClassCourse struct {
	CourseId *int `json:"course_id"  form:"course_id" binding:"required"`
	ClassId  *int `json:"class_id" form:"class_id" binding:"required"`
}

//班级id查询
type ClassIdRequest struct {
	ClassId *int `json:"class_id" form:"class_id" binding:"required"`
	PageInfo
}

//设置班级老师
type SetClassTeacher struct {
	ClassId    *int `json:"class_id"  form:"class_id" binding:"required"`
	TeacherId  *int `json:"teacher_id" form:"teacher_id" binding:"required"`
	ShowStatus *int `json:"show_status" form:"show_status"`
}

//设置班级学员
type SetClassUser struct {
	ClassId *int `json:"class_id"  form:"class_id" binding:"required"`
	UserId  *int `json:"user_id" form:"user_id" binding:"required"`
}

//班级学员有效期
type SetClassUserValidity struct {
	UserId      *int       `json:"user_id" form:"user_id" binding:"required"`
	ClassId     *int       `json:"class_id" form:"class_id" binding:"required"`
	ClosingDate *time.Time `json:"closing_date" form:"closing_date" form:"required"`
}
