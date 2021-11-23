package cache

import (
	"time"
	"tkserver/pkg/cache"
)

var (
	SmsCacheKey         = "sms:login:ssmd:"
	OssCredentialsKey   = "oss:credentials:ssmd:"
	DingScanAccessToken = "ding:scantoken:ssmd:"

	GetQuestionKey        = "question:details:ssmd:"  //单个题目详情
	ExamPartition         = "exam:partition:ssmd:"    //试卷子部分信息
	UserCalendarKey       = "user:Calendar:ssmd:"     //用户直播课日历
	UserLiveCourseDateKey = "user:live:date:ssmd:"    //用户直播课单个日期
	SecQuestionList       = "sec:question:list:"      //章节练习
	AppIndexKey           = "app:index:ssmd"          //首页
	OpenCourseKey         = "app:opencourse:ssmd"     //公开课
	RecommendCourse       = "app:recodeCourse:ssmd:"  //精品课
	TeacherListKey        = "app:teacherlist:ssmd:"   //问老师列表
	CourseChapterKey      = "CourseChapter:ssmd:"     //课程目录
	ClassCourseListKey    = "class:course:list:"      //班级下课程
	QuestionBankTag       = "question:tag:list:ssmd"  //题库标签
	QuestionBankCheckTag  = "question:bank:list:ssmd:" //题库筛选列表

)

var MemoryCache = cache.NewLRU(5*time.Minute, 5*time.Minute, 500)

var QuestionCache = cache.NewLRU(1*time.Minute, 20*time.Second, 100)
var SectionExamList = cache.NewLRU(20*time.Minute, 30*time.Minute, 100)
var CommCache = cache.NewLRU(10*time.Second, 20*time.Second, 100)
