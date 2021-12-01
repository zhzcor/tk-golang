package httpapi

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"tkserver/httpapi/admin"
	"tkserver/httpapi/appapi"
	"tkserver/internal/app"
	"tkserver/internal/errorno"
	"tkserver/pkg/log"
)

type H func(ctx *gin.Context) (interface{}, error)

func Use(engine *gin.Engine) {

	// 在这里注册路由
	engine.POST("/pong", ResponseToJSON(appapi.Pong))
	ugroup := engine.Group("/user")
	ugroup.POST("/register", ResponseToJSON(appapi.UserRegister))
	/*	ugroup.POST("/login", ResponseToJSON(appapi.UserLogin))
	 */ /*************appApi接口******************************************************/
	//app首页接不要验证
	appIndexRoute := engine.Group("/api/auth")
	appIndexRoute.POST("/get_index_info", ResponseToJSON(appapi.GetIndexInfo))
	//公开课体验课不用验证
	appIndexRoute.POST("/get_live_open_courses", ResponseToJSON(appapi.GetLiveOpenCourses)) //公开课
	appIndexRoute.POST("/get_record_courses", ResponseToJSON(appapi.GetRecordCourses))      //精品课程
	appIndexRoute.POST("/get_cate_small_detail", ResponseToJSON(appapi.GetCateSmallDetail)) //小节课程详情
	appIndexRoute.POST("/get_course_details", ResponseToJSON(appapi.GetCourseDetails))      //获取课程详情目录课时

	apiRoute := engine.Group("/api")
	apiRoute.POST("/sms_login", ResponseToJSON(appapi.UserSmsLogin))
	apiRoute.POST("/pwd_login", ResponseToJSON(appapi.UserPwdLogin))
	apiRoute.POST("/applet_wechat_login", ResponseToJSON(appapi.AppletWeChatLogin))
	apiRoute.POST("/user_check_code", ResponseToJSON(appapi.UserCheckCode)) //验证验证码
	apiRoute.POST("/send_sms", ResponseToJSON(appapi.SendSms))
	apiRoute.POST("/set_pwd", ResponseToJSON(appapi.SetPwd))
	apiRoute.POST("/get_app_agreement", ResponseToJSON(appapi.GetAppAgreement)) //协议
	apiRoute.POST("/get_app_version", ResponseToJSON(appapi.GetAppVersion))     //版本
	apiRoute.POST("/get_course_question_bank_info", ResponseToJSON(appapi.GetCourseQuestionBankInfo))//题库分类详情

	apiRoute.POST("/get_question_bank_tag", ResponseToJSON(appapi.GetQuestionBankTag))//题库标签tag
	apiRoute.POST("/get_major_question_bank_list", ResponseToJSON(appapi.GetMajorQuestionBankList))//题库列表
	apiRoute.POST("/get_group_card_list", ResponseToJSON(appapi.GetGroupCardList))//题库列表
	/*apiRoute.POST("/get_index_info", ResponseToJSON(appapi.GetIndexInfo))      */ //app首页

	apiRouteLogin := apiRoute.Group("/auth").Use(AuthMiddleware())

	apiRouteLogin.POST("/personal_data", ResponseToJSON(appapi.ScreenItems))                               //个人中心项目，地区，筛选
	apiRouteLogin.POST("/update_user_info", ResponseToJSON(appapi.UpdateUserInfo))                         //个人中心资料修改
	apiRouteLogin.POST("/update_phone", ResponseToJSON(appapi.UpdatePhone))                                //修改手机号
	apiRouteLogin.POST("/user_update_pwd", ResponseToJSON(appapi.UserUpdatePwd))                           //pc原密码修改密码
	apiRouteLogin.POST("/get_question_index", ResponseToJSON(appapi.GetQuestionIndex))                     //题库首页
/*	apiRouteLogin.POST("/get_course_question_bank_info", ResponseToJSON(appapi.GetCourseQuestionBankInfo)) //题库首页
*/	apiRouteLogin.POST("/get_question_bank_exam_list", ResponseToJSON(appapi.GetQuestionBankExamList))     //题库试卷分类下列表

	apiRouteLogin.POST("/get_exa_question_list", ResponseToJSON(appapi.GetQuestionList))                 //试卷分类下列表
	apiRouteLogin.POST("/get_tk_cs_question_list", ResponseToJSON(appapi.GetTkCsQuestionList))           //章节下题目
	apiRouteLogin.POST("/add_user_question_recode", ResponseToJSON(appapi.AddUserQuestionRecode))        //提交试卷题目
	apiRouteLogin.POST("/get_user_answer_sheet", ResponseToJSON(appapi.GetUserAnswerSheet))              //答题报告
	apiRouteLogin.POST("/get_user_wrong_question_list", ResponseToJSON(appapi.GetUserWrongQuestionList)) //错题本首页
	apiRouteLogin.POST("/get_wrong_detail", ResponseToJSON(appapi.GetWrongList))                         //错题题目详情
	apiRouteLogin.POST("/get_question_detail", ResponseToJSON(appapi.GetQuestionDetail))                 //题目详情

	apiRouteLogin.POST("/get_question_list_detail", ResponseToJSON(appapi.GetQuestionListDetail))          //批量题目详情
	apiRouteLogin.POST("/add_wrong_question", ResponseToJSON(appapi.AddWrongQuestion))                     //记录单个错题
	apiRouteLogin.POST("/add_question_feed_back", ResponseToJSON(appapi.AddQuestionFeedBack))              //题目纠错
	apiRouteLogin.POST("/add_collection", ResponseToJSON(appapi.AddCollection))                            //收藏
	apiRouteLogin.POST("/get_collection_list", ResponseToJSON(appapi.GetCollectionList))                   //收藏首页列表
	apiRouteLogin.POST("/get_make_user_question_recode", ResponseToJSON(appapi.GetMakeUserQuestionRecode)) //做题记录
	apiRouteLogin.POST("/remove_wrong_question", ResponseToJSON(appapi.RemoveWrongQuestion))               //移除做题记录
	apiRouteLogin.POST("/get_simulation_exam_status", ResponseToJSON(appapi.GetSimulationExamStatus))      //获取模拟考试状态跟新
	//学习
	apiRouteLogin.POST("/course_calendar", ResponseToJSON(appapi.CourseCalendar))           //学习日历
	apiRouteLogin.POST("/get_calendar_detail", ResponseToJSON(appapi.GetCalendarDetail))    //学习日历详情
	apiRouteLogin.POST("/get_user_live_course", ResponseToJSON(appapi.GetUserLiveCourse))   //获取用户直播课程
	apiRouteLogin.POST("/get_class_course_list", ResponseToJSON(appapi.GetClassCourseList)) //获取班级下课程
	apiRouteLogin.POST("/get_course_file", ResponseToJSON(appapi.GetCourseFile))            //获取课程资料
	apiRouteLogin.POST("/get_practice_in_class", ResponseToJSON(appapi.GetPracticeInClass)) //获取随堂练习
	apiRouteLogin.POST("/get_teacher_list", ResponseToJSON(appapi.GetTeacherList))          //老师
	apiRouteLogin.POST("/user_ask_answer", ResponseToJSON(appapi.UserAskAnswer))            //学员提问老师

	apiRouteLogin.POST("/add_user_course_appraise", ResponseToJSON(appapi.AddUserCourseAppraise))         //学员提交评价
	apiRouteLogin.POST("/get_user_course_appraiseList", ResponseToJSON(appapi.GetUserCourseAppraiseList)) //获取课程评价
	apiRouteLogin.POST("/get_pc_index_info", ResponseToJSON(appapi.GetPcIndex))                           //pc首页
	apiRouteLogin.POST("/get_mess_info", ResponseToJSON(appapi.GetMessInfo))                              //消息
	apiRouteLogin.POST("/add_video_record", ResponseToJSON(appapi.AddVideoRecord))                        //添加观看记录
	apiRouteLogin.POST("/get_video_record", ResponseToJSON(appapi.GetVideoRecord))                        //查看观看记录

	apiRouteLogin.POST("/oss", ResponseToJSON(appapi.AuthOss))

	/*	apiConfig := loginGroup.Group("/api")
	 */

	/*************app后台管理系统接口******************************************************/

	//app后台接口
	//no auth
	adminGroup := engine.Group("/admin") //不带auth认证
	//adminGroup.POST("/other/sync_user", ResponseToJSON(admin.SyncUserFromBoss))         //boss后台同步学员
	//adminGroup.POST("/other/sync_admin", ResponseToJSON(admin.SyncAdminFromBoss))       //boss后台同步管理员
	adminGroup.POST("/other/login", ResponseToJSON(admin.LoginAdmin)) //登录
	//adminGroup.POST("/other/bjy_callback_set", ResponseToJSON(admin.SetBjyCallbackUrl)) //设置百家云回调
	//adminGroup.POST("/other/bjy_class_callback", admin.ClassCallbackHandler())          //上下课回调
	//adminGroup.POST("/other/bjy_order_callback", admin.VideosCallbackHandler())         //点播（回放）回调
	adminGroup.POST("/other/oss_callback", ResponseToJSON(admin.OssCallbackHandler)) //oss回调
	//adminGroup.POST("/other/third_login", ResponseToJSON(admin.ThirdLoginAdmin))        //钉钉登录

	//基础配置(字典管理)
	basicConfigGroup := adminGroup.Group("/basic_config").Use(AdminAuthMiddleware())
	basicConfigGroup.POST("/oss/auth", ResponseToJSON(admin.AuthOss))
	//地区、项目、专业
	basicConfigGroup.POST("/city_set", ResponseToJSON(admin.SetCity)) //添加（编辑）地区
	basicConfigGroup.POST("/city_del", ResponseToJSON(admin.DelCity)) //删除地区
	//basicConfigGroup.POST("/city_status", ResponseToJSON(admin.SetCityStatus))            //设置地区状态
	basicConfigGroup.POST("/city_page_list", ResponseToJSON(admin.GetCityByPage)) //地区列表(分页)
	basicConfigGroup.POST("/city_list", ResponseToJSON(admin.GetCityAll))         //地区列表(无分页)
	basicConfigGroup.POST("/item_set", ResponseToJSON(admin.SetItemCategory))     //添加（编辑）项目
	basicConfigGroup.POST("/item_del", ResponseToJSON(admin.DelItemCategory))     //删除项目
	//basicConfigGroup.POST("/item_status", ResponseToJSON(admin.SetItemStatus))            //设置项目状态
	basicConfigGroup.POST("/item_page_list", ResponseToJSON(admin.GetItemCategoryByPage)) //项目列表(分页)
	basicConfigGroup.POST("/item_list", ResponseToJSON(admin.GetItemCategoryAll))         //项目列表(无分页)
	basicConfigGroup.POST("/major_set", ResponseToJSON(admin.SetMajor))                   //添加（编辑）专业
	basicConfigGroup.POST("/major_del", ResponseToJSON(admin.DelMajor))                   //删除专业
	//basicConfigGroup.POST("/major_status", ResponseToJSON(admin.SetMajorStatus))          //设置专业状态
	basicConfigGroup.POST("/major_page_list", ResponseToJSON(admin.GetMajorByPage)) //专业列表(分页)
	basicConfigGroup.POST("/major_list", ResponseToJSON(admin.GetMajorAll))         //专业列表(无分页)
	basicConfigGroup.POST("/level_set", ResponseToJSON(admin.SetLevel))             //添加（编辑）层次
	basicConfigGroup.POST("/level_del", ResponseToJSON(admin.DelLevel))             //删除层次
	basicConfigGroup.POST("/level_page_list", ResponseToJSON(admin.GetLevelByPage)) //层次列表(分页)
	basicConfigGroup.POST("/level_list", ResponseToJSON(admin.GetLevelAll))         //层次列表(无分页)

	basicConfigGroup.POST("/agreement_set", ResponseToJSON(admin.SetAppAgreement))             //添加app协议
	basicConfigGroup.POST("/agreement_get", ResponseToJSON(admin.GetAppAgreement))             //根据协议类型获取协议详情
	basicConfigGroup.POST("/app_version_set", ResponseToJSON(admin.SetAppVersion))             //添加(编辑)app版本
	basicConfigGroup.POST("/app_version_page_list", ResponseToJSON(admin.GetAppVersionByPage)) //app版本列表（分页）
	basicConfigGroup.POST("/app_version_del", ResponseToJSON(admin.DelAppVersion))             //app版本删除

	basicConfigGroup.POST("/group_card_set", ResponseToJSON(admin.SetGroupCard))             //添加（编辑）群名片
	basicConfigGroup.POST("/group_card_del", ResponseToJSON(admin.DelGroupCard))             //删除群名片
	basicConfigGroup.POST("/group_card_status", ResponseToJSON(admin.SetGroupCardStatus))    //设置群名片状态
	basicConfigGroup.POST("/group_card_page_list", ResponseToJSON(admin.GetGroupCardByPage)) //群名片列表(分页)

	//管理员管理
	adminManagerGroup := adminGroup.Group("/manager").Use(AdminAuthMiddleware())
	//adminManagerGroup.POST("/index_statistic", ResponseToJSON(admin.GetIndexStatistic))     //首页统计
	//adminManagerGroup.POST("/role_add", ResponseToJSON(admin.AddRole))                      //添加角色
	//adminManagerGroup.POST("/role_status", ResponseToJSON(admin.SetRoleStatus))             //设置管理员状态
	//adminManagerGroup.POST("/role_list", ResponseToJSON(admin.GetRoleList))                 //角色列表（无分页）
	//adminManagerGroup.POST("/role_page_list", ResponseToJSON(admin.GetRoleListByPage))      //角色列表（分页）
	//adminManagerGroup.POST("/role_del", ResponseToJSON(admin.DelRole))                      //删除角色
	adminManagerGroup.POST("/admin_set", ResponseToJSON(admin.SetAdmin))                 //编辑管理员
	adminManagerGroup.POST("/admin_page_list", ResponseToJSON(admin.GetAdminListByPage)) //管理员列表（分页）
	adminManagerGroup.POST("/admin_status", ResponseToJSON(admin.SetAdminStatus))        //设置管理员状态
	//adminManagerGroup.POST("/role_permission_add", ResponseToJSON(admin.AddRolePermission)) //角色添加权限
	//adminManagerGroup.POST("/role_permission_get", ResponseToJSON(admin.GetRolePermission)) //获取角色权限

	adminManagerGroup.POST("/password_reset", ResponseToJSON(admin.ResetPassword))                      //重置密码
	adminManagerGroup.POST("/avatar_upload", ResponseToJSON(admin.UpdateAdminAvatar))                   //上传头像
	adminManagerGroup.POST("/login_log_page_list", ResponseToJSON(admin.GetAdminLoginLogByPage))        //后台登录日志
	adminManagerGroup.POST("/operator_log_page_list", ResponseToJSON(admin.GetAdminOperationLogByPage)) //操作日志
	adminManagerGroup.POST("/login_user_page_list", ResponseToJSON(admin.GetLoginUserByPage))           //app登录日志

	adminManagerGroup.POST("/user_list", ResponseToJSON(admin.GetSimpleUserInfo))         //根据条件获取学员姓名手机号列表(分页)
	adminManagerGroup.POST("/user_detail_by_id", ResponseToJSON(admin.GetUserDetailById)) //id获取用户详情
	//adminManagerGroup.POST("/boss_user_by_phone", ResponseToJSON(admin.GetBossUserByPhone)) //手机号获取工作台用户信息
	adminManagerGroup.POST("/user_set", ResponseToJSON(admin.SetUser)) //添加（编辑）用户

	adminManagerGroup.POST("/teacher_list", ResponseToJSON(admin.GetSimpleTeacherInfo))  //根据条件获取老师姓名手机号列表（分页）
	adminManagerGroup.POST("/teacher_set", ResponseToJSON(admin.SetTeacher))             //添加（编辑）老师
	adminManagerGroup.POST("/teacher_status", ResponseToJSON(admin.SetTeacherStatus))    //设置老师状态
	adminManagerGroup.POST("/teacher_del", ResponseToJSON(admin.DelTeacher))             //删除老师
	adminManagerGroup.POST("/teacher_by_id", ResponseToJSON(admin.GetTeacherDetailById)) //id获取老师信息

	//用户管理
	adminManagerGroup.POST("/ask_page_list", ResponseToJSON(admin.GetAskAnswerPageList))           //用户问答列表（分页）
	adminManagerGroup.POST("/ask_reply", ResponseToJSON(admin.ReplyAskAnswer))                     //回复用户问答
	adminManagerGroup.POST("/ask_status", ResponseToJSON(admin.SetAskAnswerShowStatus))            //设置用户问答显示状态
	adminManagerGroup.POST("/ask_del", ResponseToJSON(admin.DelAskAnswer))                         //用户问答删除
	adminManagerGroup.POST("/ask_reply_status", ResponseToJSON(admin.SetAskAnswerReplyShowStatus)) //设置用户问答回复显示状态
	adminManagerGroup.POST("/ask_reply_del", ResponseToJSON(admin.DelAskAnswerReply))              //用户问答回复删除

	adminManagerGroup.POST("/course_appraise_page_list", ResponseToJSON(admin.GetUserCourseAppraisePageList)) //用户课程评论列表（分页）
	adminManagerGroup.POST("/course_appraise_reply", ResponseToJSON(admin.ReplyUserCourseAppraise))           //回复用户课程评论
	adminManagerGroup.POST("/course_appraise_status", ResponseToJSON(admin.SetUserCourseAppraiseStatus))      //设置用户课程评论显示状态
	adminManagerGroup.POST("/course_appraise_del", ResponseToJSON(admin.DelUserCourseAppraise))               //用户课程评论删除

	//题库
	tkGroup := adminGroup.Group("/tk").Use(AdminAuthMiddleware())
	tkGroup.POST("/question_bank_page_list", ResponseToJSON(admin.GetTkQuestionBankByPage)) //题库
	tkGroup.POST("/question_bank_set", ResponseToJSON(admin.SetTkQuestionBank))
	tkGroup.POST("/question_bank_del", ResponseToJSON(admin.DelTkQuestionBank))
	tkGroup.POST("/question_bank_status", ResponseToJSON(admin.SetTkQuestionBankStatus))
	tkGroup.POST("/bank_question_type_count", ResponseToJSON(admin.GetTkBankQuestionTypeCount)) //获取题库各类型题目数量
	tkGroup.POST("/bank_question_simple_list", ResponseToJSON(admin.GetTkSimpleQuestionBank))   //获取题库id name列表

	tkGroup.POST("/chapter_page_list", ResponseToJSON(admin.GetTkQuestionBankChapterSectionByPage)) //章节
	tkGroup.POST("/chapter_set", ResponseToJSON(admin.SetTkQuestionBankChapter))
	tkGroup.POST("/chapter_del", ResponseToJSON(admin.DelTkQuestionBankChapter))
	tkGroup.POST("/section_set", ResponseToJSON(admin.SetTkQuestionBankSection))
	tkGroup.POST("/section_del", ResponseToJSON(admin.DelTkQuestionBankSection))
	tkGroup.POST("/cs_question_list", ResponseToJSON(admin.GetTkCsQuestionList))         //章节题目详情（带分页）
	tkGroup.POST("/question_ids_in_cs", ResponseToJSON(admin.GetQuestionIdsInSection))   //获取章节下已有题目id数组
	tkGroup.POST("/section_question_set", ResponseToJSON(admin.SetQuestionToSection))    //为章节添加题目
	tkGroup.POST("/section_question_del", ResponseToJSON(admin.RemoveQuestionToSection)) //获取章节下已有题目id数组

	tkGroup.POST("/knowledge_point_page_list", ResponseToJSON(admin.GetTkKnowledgePointsByPage)) //题库-知识点(带分页)
	tkGroup.POST("/knowledge_point_list", ResponseToJSON(admin.GetTkKnowledgePoints))            //知识点(不带分页)
	tkGroup.POST("/knowledge_point_set", ResponseToJSON(admin.SetTkKnowledgePoint))
	tkGroup.POST("/knowledge_point_del", ResponseToJSON(admin.DelTkKnowledgePoint))

	tkGroup.POST("/question_page_list", ResponseToJSON(admin.GetTkQuestionsByPage)) //题目（带分页）
	tkGroup.POST("/question_by_id", ResponseToJSON(admin.GetTkQuestionsById))       //id获取题目详情
	tkGroup.POST("/question_set", ResponseToJSON(admin.SetTkQuestion))
	tkGroup.POST("/question_del", ResponseToJSON(admin.DelTkQuestion))
	tkGroup.POST("/question_import", ResponseToJSON(admin.UploadQuestion))          //导入题目
	tkGroup.POST("/question_import_template", ResponseToJSON(admin.UploadQuestion)) //导入题目模版

	tkGroup.POST("/feedback_list", ResponseToJSON(admin.GetQuestionErrorFeedbackList)) //错题反馈列表
	tkGroup.POST("/deal_feedback", ResponseToJSON(admin.DealTkErrorFeedback))          //错题反馈处理

	tkGroup.POST("/exam_paper_set", ResponseToJSON(admin.SetTkExamPaper))              //添加（编辑）试卷
	tkGroup.POST("/exam_paper_status", ResponseToJSON(admin.SetExamPaperStatus))       //设置试卷状态
	tkGroup.POST("/exam_paper_del", ResponseToJSON(admin.DelTkExamPaper))              //删除试卷
	tkGroup.POST("/exam_paper_page_list", ResponseToJSON(admin.GetTkExamPapersByPage)) //试卷列表(带分页)
	tkGroup.POST("/exam_paper_all_list", ResponseToJSON(admin.GetTkExamPapersListAll)) //试卷列表(不带分页)
	tkGroup.POST("/exam_paper_by_id", ResponseToJSON(admin.GetTkExamPaperById))        //id获取试卷详情
	tkGroup.POST("/exam_rank", ResponseToJSON(admin.RankTkExamPaper))                  //用户模拟考试排名

	tkGroup.POST("/simulation_page_list", ResponseToJSON(admin.GetUserSimulationExamPageList))              //模拟考试明细
	tkGroup.POST("/user_subjective_simulation", ResponseToJSON(admin.GetUserSimulationSubjectiveQuestions)) //用户模拟考试主观题记录
	tkGroup.POST("/user_subjective_set", ResponseToJSON(admin.SetUserSubjectiveScore))                      //用户模拟考试主观题判分

	tkGroup.POST("/student_statistic_list", ResponseToJSON(admin.GetStudentStatisticList))                      //学生数据列表
	tkGroup.POST("/user_bank_statistic_list", ResponseToJSON(admin.GetQuestionBankStatisticList))               //获取做题记录统计列表
	tkGroup.POST("/user_question_statistic_list", ResponseToJSON(admin.GetUserQuestionBankStatisticDetailList)) //获取用户做题记录统计详细列表

	//课程
	//kcGroup := adminGroup.Group("/kc").Use(AdminAuthMiddleware())
	//kcGroup.POST("/course_set", ResponseToJSON(admin.SetCourse))                //添加（编辑）课程
	//kcGroup.POST("/course_page_list", ResponseToJSON(admin.CoursePageList))     //课程列表（分页）
	//kcGroup.POST("/course_del", ResponseToJSON(admin.DelCourse))                //删除课程
	//kcGroup.POST("/course_status", ResponseToJSON(admin.SetCourseStatus))       //设置课程状态
	//kcGroup.POST("/course_by_id", ResponseToJSON(admin.GetCourseSmallCateList)) //获取课程下所有章节课时（无分页）
	//kcGroup.POST("/course_detail", ResponseToJSON(admin.GetCourseDetailById))   //id获取课程详情
	//kcGroup.POST("/course_live_list", ResponseToJSON(admin.GetLivePageList))    //直播课程列表
	//kcGroup.POST("/course_simple_list", ResponseToJSON(admin.CourseSimpleList)) //课程id name列表
	//
	//kcGroup.POST("/chapter_set", ResponseToJSON(admin.SetKcCourseChapter)) //添加（编辑）课程章
	//kcGroup.POST("/chapter_del", ResponseToJSON(admin.DelKcCourseChapter)) //删除课程章
	//kcGroup.POST("/section_set", ResponseToJSON(admin.SetKcCourseSection)) //添加（编辑）课节
	//kcGroup.POST("/section_del", ResponseToJSON(admin.DelKcCourseSection)) //删除课节
	//
	//kcGroup.POST("/course_teacher_add", ResponseToJSON(admin.AddCourseTeacher))                          //添加课程老师
	//kcGroup.POST("/course_teacher_status", ResponseToJSON(admin.SetCourseTeacherStatus))                 //添加课程老师状态
	//kcGroup.POST("/course_teacher_remove", ResponseToJSON(admin.RemoveCourseTeacher))                    //移除课程老师
	//kcGroup.POST("/course_teacher_page_list", ResponseToJSON(admin.GetCourseTeacherPageList))            //课程老师列表（分页¬）
	//kcGroup.POST("/course_teacher_selected_ids", ResponseToJSON(admin.GetSelectedCourserTeacherIdsList)) //获取课程已添加teacher id列表
	//
	//kcGroup.POST("/course_user_add", ResponseToJSON(admin.AddCourseUser))                          //添加课程学员
	//kcGroup.POST("/course_user_remove", ResponseToJSON(admin.RemoveCourseUser))                    //移除课程学员
	//kcGroup.POST("/course_user_page_list", ResponseToJSON(admin.GetCourseUserPageList))            //获取课程用户列表（分页）
	//kcGroup.POST("/user_detail_by_id", ResponseToJSON(admin.GetUserDetailById))                    //id获取用户详情
	//kcGroup.POST("/user_course_validity", ResponseToJSON(admin.SetCourseUserValidity))             //设置用户课程有效期
	//kcGroup.POST("/course_user_selected_ids", ResponseToJSON(admin.GetSelectedCourserUserIdsList)) //获取课程已添加user id列表
	//
	//kcGroup.POST("/course_question_bank_add", ResponseToJSON(admin.AddCourseQuestionBank))       //添加课程题库
	//kcGroup.POST("/course_question_bank_remove", ResponseToJSON(admin.RemoveCourseQuestionBank)) //移除课程题库
	//kcGroup.POST("/course_question_detail", ResponseToJSON(admin.GetCourseQuestionBankDetail))   //获取课程下题库详情
	//
	//kcGroup.POST("/small_cate_set", ResponseToJSON(admin.SetKcCourseSmallCategory))                          //添加(编辑)课时
	//kcGroup.POST("/small_cate_del", ResponseToJSON(admin.DelKcCourseSmallCategory))                          //删除课时
	//kcGroup.POST("/small_cate_by_id", ResponseToJSON(admin.GetSmallCateById))                                //id获取课时详情
	//kcGroup.POST("/small_cate_exam_set", ResponseToJSON(admin.SetKcSmallCategoryExamPaper))                  //添加课时试卷（作业）
	//kcGroup.POST("/small_cate_exam_del", ResponseToJSON(admin.RemoveKcSmallCategoryExamPaper))               //移除课时试卷（作业）
	//kcGroup.POST("/small_cate_exam_list", ResponseToJSON(admin.GetSmallCourseExamPageList))                  //课时试卷、作业列表（分页）
	//kcGroup.POST("/small_cate_selected_exam_list", ResponseToJSON(admin.GetSelectedSmallCourseExamPageList)) //已选中 课时试卷、作业列表（分页）
	//
	//kcGroup.POST("/small_cate_false_set", ResponseToJSON(admin.SetFalseVideo))          //添加(编辑)伪直播
	//kcGroup.POST("/small_cate_false_Del", ResponseToJSON(admin.DelFalseVideo))          //删除伪直播
	//kcGroup.POST("/small_cate_false_list", ResponseToJSON(admin.GetFalseVideoPageList)) //课程伪直播列表（分页）
	//
	//kcGroup.POST("/small_cate_back_list", ResponseToJSON(admin.GetLiveBackPageList)) //课程直播回放列表（分页）
	//kcGroup.POST("/small_cate_back_reload", ResponseToJSON(admin.UpdateLiveBack))    //课程直播回放重新录制
	//kcGroup.POST("/small_cate_back_replace", ResponseToJSON(admin.ReplaceLiveBack))  //点播替换回放
	//
	//kcGroup.POST("/small_cate_question_set", ResponseToJSON(admin.SetKcSmallCategoryQuestion))                       //添加课时练习
	//kcGroup.POST("/small_cate_question_del", ResponseToJSON(admin.RemoveKcSmallCategoryQuestion))                    //移除课时练习
	//kcGroup.POST("/small_cate_question_list", ResponseToJSON(admin.GetSmallCourseQuestionPageList))                  //课时练习列表（分页）
	//kcGroup.POST("/small_cate_selected_question_list", ResponseToJSON(admin.GetSelectedSmallCourseQuestionPageList)) //已选中 课时练习列表（分页）
	//
	//kcGroup.POST("/small_cate_attachment_set", ResponseToJSON(admin.SetKcSmallCategoryAttachment))                       //添加课时资料
	//kcGroup.POST("/small_cate_attachment_del", ResponseToJSON(admin.RemoveKcSmallCategoryAttachment))                    //移除课时资料
	//kcGroup.POST("/small_cate_selected_attachment_list", ResponseToJSON(admin.GetSelectedSmallCourseAttachmentPageList)) //已选中 课时资料列表（分页）
	//
	////班级
	//clGroup := adminGroup.Group("/cl").Use(AdminAuthMiddleware())
	//clGroup.POST("/class_set", ResponseToJSON(admin.SetKcClass))          //添加（编辑）班级
	//clGroup.POST("/class_del", ResponseToJSON(admin.DelClass))            //删除班级
	//clGroup.POST("/class_status", ResponseToJSON(admin.SetClassStatus))   //设置班级状态
	//clGroup.POST("/class_detail", ResponseToJSON(admin.ClassDetailById))  //id获取班级详情
	//clGroup.POST("/class_page_list", ResponseToJSON(admin.ClassPageList)) //班级列表（分页）
	//
	//clGroup.POST("/class_course_page_list", ResponseToJSON(admin.GetClassCoursePageList))                  //班级下课程列表（分页）
	//clGroup.POST("/class_course_selected_page_list", ResponseToJSON(admin.GetSelectedClassCoursePageList)) //已选中 班级下课程列表（分页）
	//clGroup.POST("/class_course_set", ResponseToJSON(admin.AddClassCourse))                                //添加班级课程
	//clGroup.POST("/class_course_remove", ResponseToJSON(admin.RemoveClassCourse))                          //移除班级课程
	//
	//clGroup.POST("/class_teacher_add", ResponseToJSON(admin.AddClassTeacher))                         //添加班级老师
	//clGroup.POST("/class_teacher_status", ResponseToJSON(admin.SetClassTeacherStatus))                //添加班级老师状态
	//clGroup.POST("/class_teacher_remove", ResponseToJSON(admin.RemoveClassTeacher))                   //移除班级老师
	//clGroup.POST("/class_teacher_page_list", ResponseToJSON(admin.GetClassTeacherPageList))           //班级老师列表（分页¬）
	//clGroup.POST("/class_teacher_selected_ids", ResponseToJSON(admin.GetSelectedClassTeacherIdsList)) //获取班级已添加teacher id列表
	//
	//clGroup.POST("/class_master_add", ResponseToJSON(admin.AddClassMasterTeacher)) //添加班级班主任
	//clGroup.POST("/class_master_remove", ResponseToJSON(admin.RemoveClassMaster))  //移除班级班主任
	//clGroup.POST("/class_master_get", ResponseToJSON(admin.GetClassMasterTeacher)) //获取班级班主任
	//
	//clGroup.POST("/class_user_add", ResponseToJSON(admin.AddClassUser))               //添加班级学员
	//clGroup.POST("/class_user_remove", ResponseToJSON(admin.RemoveClassUser))         //移除班级学员
	//clGroup.POST("/class_user_page_list", ResponseToJSON(admin.GetClassUserPageList)) //获取班级用户列表（分页）
	////clGroup.POST("/user_detail_by_id", ResponseToJSON(admin.GetUserDetailById))                 //id获取用户详情
	//clGroup.POST("/user_class_validity", ResponseToJSON(admin.AddClassUserValidity))            //设置用户班级有效期
	//clGroup.POST("/class_user_selected_ids", ResponseToJSON(admin.GetSelectedClassUserIdsList)) //获取班级已添加user id列表

	//社区
	//communityGroup := adminGroup.Group("/community").Use(AdminAuthMiddleware())
	//communityGroup.POST("/activity_type_set", ResponseToJSON(admin.SetActivityType))             //添加（编辑）活动类型
	//communityGroup.POST("/activity_type_del", ResponseToJSON(admin.DelActivityType))             //删除活动类型
	//communityGroup.POST("/activity_type_status", ResponseToJSON(admin.SetActivityTypeStatus))    //设置活动类型状态
	//communityGroup.POST("/activity_type_page_list", ResponseToJSON(admin.GetActivityTypeByPage)) //活动类型列表(分页)
	//communityGroup.POST("/activity_type_list", ResponseToJSON(admin.GetActivityTypeAll))         //活动类型列表(无分页)
	//
	//communityGroup.POST("/info_classify_set", ResponseToJSON(admin.SetInformationClassify))             //添加（编辑）资讯分类
	//communityGroup.POST("/info_classify_del", ResponseToJSON(admin.DelInformationClassify))             //删除资讯分类
	//communityGroup.POST("/info_classify_status", ResponseToJSON(admin.SetInformationClassifyStatus))    //设置资讯分类状态
	//communityGroup.POST("/info_classify_page_list", ResponseToJSON(admin.GetInformationClassifyByPage)) //资讯分类列表(分页)
	//communityGroup.POST("/info_classify_list", ResponseToJSON(admin.GetInformationClassifyAll))         //资讯分类列表(无分页)
	//
	//communityGroup.POST("/msg_type_set", ResponseToJSON(admin.SetMsgType))             //添加（编辑）消息类型
	//communityGroup.POST("/msg_type_del", ResponseToJSON(admin.DelMsgType))             //删除消息类型
	//communityGroup.POST("/msg_type_status", ResponseToJSON(admin.SetMsgTypeStatus))    //设置消息类型状态
	//communityGroup.POST("/msg_type_page_list", ResponseToJSON(admin.GetMsgTypeByPage)) //消息类型列表(分页)
	//communityGroup.POST("/msg_type_list", ResponseToJSON(admin.GetMsgTypeAll))         //消息类型列表(无分页)
	//
	//communityGroup.POST("/hot_search_set", ResponseToJSON(admin.SetHotSearch))             //添加（编辑）热门搜索
	//communityGroup.POST("/hot_search_del", ResponseToJSON(admin.DelHotSearch))             //删除热门搜索
	//communityGroup.POST("/hot_search_status", ResponseToJSON(admin.SetHotSearchStatus))    //设置热门搜索状态
	//communityGroup.POST("/hot_search_page_list", ResponseToJSON(admin.GetHotSearchByPage)) //热门搜索列表(分页)
	//communityGroup.POST("/hot_search_list", ResponseToJSON(admin.GetHotSearchAll))         //热门搜索列表(无分页)

	//推广管理
	//popGroup := adminGroup.Group("/pop").Use(AdminAuthMiddleware())
	//popGroup.POST("/advertise_set", ResponseToJSON(admin.SetAdvertise))             //添加（编辑）广告
	//popGroup.POST("/advertise_del", ResponseToJSON(admin.DelAdvertise))             //删除广告
	//popGroup.POST("/advertise_status", ResponseToJSON(admin.SetAdvertiseStatus))    //设置广告状态
	//popGroup.POST("/advertise_page_list", ResponseToJSON(admin.GetAdvertiseByPage)) //广告列表(分页)
	//
	//popGroup.POST("/poster_set", ResponseToJSON(admin.SetSharePoster))             //添加（编辑）分享海报
	//popGroup.POST("/poster_del", ResponseToJSON(admin.DelSharePoster))             //删除分享海报
	//popGroup.POST("/poster_status", ResponseToJSON(admin.SetSharePosterStatus))    //设置分享海报状态
	//popGroup.POST("/poster_page_list", ResponseToJSON(admin.GetSharePosterByPage)) //分享海报列表(分页)
	//
	//popGroup.POST("/message_set", ResponseToJSON(admin.SetMessage))             //添加（编辑）消息
	//popGroup.POST("/message_del", ResponseToJSON(admin.DelMessage))             //删除消息
	//popGroup.POST("/message_status", ResponseToJSON(admin.SetMessageStatus))    //设置消息状态
	//popGroup.POST("/message_page_list", ResponseToJSON(admin.GetMessageByPage)) //消息列表(分页)
}

func ResponseToJSON(h H) gin.HandlerFunc {
	return func(c *gin.Context) {
		defaultResp := gin.H{"error": nil, "code": 0, "data": nil}
		resp, err := h(c)
		if err != nil {
			validationErrors := validator.ValidationErrors{}
			errno := errorno.ErrAs(err)
			switch {
			case errors.As(err, &validationErrors):
				var msg string
				for _, v := range validationErrors.Translate(trans) {
					msg = msg + v + ","
				}
				defaultResp["error"] = msg
				defaultResp["code"] = http.StatusBadRequest
			case errno != nil:
				log.Error("on http request error", zap.Error(errno), zap.Errors("errors", errno.Errored))
				defaultResp["error"] = errno.Message
				defaultResp["code"] = errno.Code
			default:
				defaultResp["error"] = err.Error()
				defaultResp["code"] = http.StatusInternalServerError
			}
			c.AbortWithStatusJSON(http.StatusOK, defaultResp)
			return

		}
		defaultResp["data"] = resp
		c.JSON(http.StatusOK, defaultResp)
	}
}

func AuthMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		auth := context.GetHeader("Authorization")
		if auth != "" {
			authInfo := strings.Split(auth, " ")
			if len(authInfo) > 1 {
				authType := authInfo[0]
				tk := authInfo[1]
				switch {
				case authType == "Bearer":
					userService := app.UserCenter{}
					uid, err := userService.CheckToken(tk)
					if err != nil {
						errno := errorno.ErrAs(err)
						if errno == nil {
							errno = errorno.NewInternalErr(err)
						}
						defaultResp := gin.H{"error": errno.Message, "code": -1, "data": nil}
						context.AbortWithStatusJSON(http.StatusOK, defaultResp)
						return
					}
					// do something
					context.Set("uid", uid)
					log.Info("user login :", uid)
				}
			}
		} else {
			err := errorno.NewErr(errorno.TokenError)
			defaultResp := gin.H{"error": err.Message, "code": -1, "data": nil}
			context.AbortWithStatusJSON(http.StatusOK, defaultResp)
			return
		}
	}
}

//app后台auth
func AdminAuthMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		auth := context.GetHeader("Authorization")
		if auth != "" {
			authInfo := strings.Split(auth, " ")
			if len(authInfo) > 1 {
				authType := authInfo[0]
				tk := authInfo[1]
				switch {
				case authType == "Bearer":
					adminManager := app.AdminManager{}
					adminId, err := adminManager.CheckToken(tk)
					if err != nil {
						errno := errorno.ErrAs(err)
						if errno == nil {
							errno = errorno.NewInternalErr(err)
						}
						defaultResp := gin.H{"error": errno.Message, "code": -1, "data": nil}
						context.AbortWithStatusJSON(200, defaultResp)
						return
					}
					context.Set("GlobalAdminId", adminId)
					// do something
					//log.Info("user login :", adminId)
				}
			}
		} else {
			err := errorno.NewErr(errorno.TokenError)
			defaultResp := gin.H{"error": err.Message, "code": -1, "data": nil}
			context.AbortWithStatusJSON(200, defaultResp)
			return
		}
	}
}
