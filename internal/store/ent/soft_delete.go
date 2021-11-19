// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"
	"tkserver/internal/store/ent/activity"
	"tkserver/internal/store/ent/activityapplyinfo"
	"tkserver/internal/store/ent/activitytype"
	"tkserver/internal/store/ent/admin"
	"tkserver/internal/store/ent/adminloginlog"
	"tkserver/internal/store/ent/adminoperationlog"
	"tkserver/internal/store/ent/advertise"
	"tkserver/internal/store/ent/appagreement"
	"tkserver/internal/store/ent/appversion"
	"tkserver/internal/store/ent/attachment"
	"tkserver/internal/store/ent/city"
	"tkserver/internal/store/ent/collection"
	"tkserver/internal/store/ent/groupcard"
	"tkserver/internal/store/ent/hotsearch"
	"tkserver/internal/store/ent/importtask"
	"tkserver/internal/store/ent/informationclassify"
	"tkserver/internal/store/ent/itemcategory"
	"tkserver/internal/store/ent/kcclass"
	"tkserver/internal/store/ent/kccourse"
	"tkserver/internal/store/ent/kccoursesmallcategory"
	"tkserver/internal/store/ent/kcsmallcategoryattachment"
	"tkserver/internal/store/ent/kcsmallcategoryexampaper"
	"tkserver/internal/store/ent/kcsmallcategoryquestion"
	"tkserver/internal/store/ent/kcuserclass"
	"tkserver/internal/store/ent/kcusercourse"
	"tkserver/internal/store/ent/kcvideouploadtask"
	"tkserver/internal/store/ent/level"
	"tkserver/internal/store/ent/major"
	"tkserver/internal/store/ent/majordetail"
	"tkserver/internal/store/ent/majordetailtag"
	"tkserver/internal/store/ent/makeuserquestionrecord"
	"tkserver/internal/store/ent/message"
	"tkserver/internal/store/ent/messagetype"
	"tkserver/internal/store/ent/permission"
	"tkserver/internal/store/ent/role"
	"tkserver/internal/store/ent/shareposter"
	"tkserver/internal/store/ent/teacher"
	"tkserver/internal/store/ent/teachertag"
	"tkserver/internal/store/ent/tkchapter"
	"tkserver/internal/store/ent/tkexampaper"
	"tkserver/internal/store/ent/tkexampaperpartition"
	"tkserver/internal/store/ent/tkexampaperpartitionscore"
	"tkserver/internal/store/ent/tkexampapersimulation"
	"tkserver/internal/store/ent/tkexampartitionquestionlink"
	"tkserver/internal/store/ent/tkexamquestiontype"
	"tkserver/internal/store/ent/tkknowledgepoint"
	"tkserver/internal/store/ent/tkquestion"
	"tkserver/internal/store/ent/tkquestionansweroption"
	"tkserver/internal/store/ent/tkquestionbank"
	"tkserver/internal/store/ent/tkquestionbankcity"
	"tkserver/internal/store/ent/tkquestionbankmajor"
	"tkserver/internal/store/ent/tkquestionerrorfeedback"
	"tkserver/internal/store/ent/tkquestionsection"
	"tkserver/internal/store/ent/tksection"
	"tkserver/internal/store/ent/tkuserexamscorerecord"
	"tkserver/internal/store/ent/tkuserquestionbankrecord"
	"tkserver/internal/store/ent/tkuserquestionrecord"
	"tkserver/internal/store/ent/tkuserrandomexamrecode"
	"tkserver/internal/store/ent/tkusersimulationteachermark"
	"tkserver/internal/store/ent/tkuserwrongquestionrecode"
	"tkserver/internal/store/ent/user"
	"tkserver/internal/store/ent/useraskanswer"
	"tkserver/internal/store/ent/usercourseappraise"
	"tkserver/internal/store/ent/userloginlog"
	"tkserver/internal/store/ent/videorecord"
)

func (ac *ActivityCreate) SoftDelete() *ActivityCreate {
	ac.SetDeletedAt(time.Now())
	return ac
}

func (au *ActivityUpdate) SoftDelete() *ActivityUpdate {
	au.SetDeletedAt(time.Now())
	return au
}

func (auo *ActivityUpdateOne) SoftDelete() *ActivityUpdateOne {
	auo.SetDeletedAt(time.Now())
	return auo
}

func (aq *ActivityQuery) SoftDelete() *ActivityQuery {
	return aq.Where(activity.DeletedAtIsNil())
}

func (aaic *ActivityApplyInfoCreate) SoftDelete() *ActivityApplyInfoCreate {
	aaic.SetDeletedAt(time.Now())
	return aaic
}

func (aaiu *ActivityApplyInfoUpdate) SoftDelete() *ActivityApplyInfoUpdate {
	aaiu.SetDeletedAt(time.Now())
	return aaiu
}

func (aaiuo *ActivityApplyInfoUpdateOne) SoftDelete() *ActivityApplyInfoUpdateOne {
	aaiuo.SetDeletedAt(time.Now())
	return aaiuo
}

func (aaiq *ActivityApplyInfoQuery) SoftDelete() *ActivityApplyInfoQuery {
	return aaiq.Where(activityapplyinfo.DeletedAtIsNil())
}

func (atc *ActivityTypeCreate) SoftDelete() *ActivityTypeCreate {
	atc.SetDeletedAt(time.Now())
	return atc
}

func (atu *ActivityTypeUpdate) SoftDelete() *ActivityTypeUpdate {
	atu.SetDeletedAt(time.Now())
	return atu
}

func (atuo *ActivityTypeUpdateOne) SoftDelete() *ActivityTypeUpdateOne {
	atuo.SetDeletedAt(time.Now())
	return atuo
}

func (atq *ActivityTypeQuery) SoftDelete() *ActivityTypeQuery {
	return atq.Where(activitytype.DeletedAtIsNil())
}

func (ac *AdminCreate) SoftDelete() *AdminCreate {
	ac.SetDeletedAt(time.Now())
	return ac
}

func (au *AdminUpdate) SoftDelete() *AdminUpdate {
	au.SetDeletedAt(time.Now())
	return au
}

func (auo *AdminUpdateOne) SoftDelete() *AdminUpdateOne {
	auo.SetDeletedAt(time.Now())
	return auo
}

func (aq *AdminQuery) SoftDelete() *AdminQuery {
	return aq.Where(admin.DeletedAtIsNil())
}

func (allc *AdminLoginLogCreate) SoftDelete() *AdminLoginLogCreate {
	allc.SetDeletedAt(time.Now())
	return allc
}

func (allu *AdminLoginLogUpdate) SoftDelete() *AdminLoginLogUpdate {
	allu.SetDeletedAt(time.Now())
	return allu
}

func (alluo *AdminLoginLogUpdateOne) SoftDelete() *AdminLoginLogUpdateOne {
	alluo.SetDeletedAt(time.Now())
	return alluo
}

func (allq *AdminLoginLogQuery) SoftDelete() *AdminLoginLogQuery {
	return allq.Where(adminloginlog.DeletedAtIsNil())
}

func (aolc *AdminOperationLogCreate) SoftDelete() *AdminOperationLogCreate {
	aolc.SetDeletedAt(time.Now())
	return aolc
}

func (aolu *AdminOperationLogUpdate) SoftDelete() *AdminOperationLogUpdate {
	aolu.SetDeletedAt(time.Now())
	return aolu
}

func (aoluo *AdminOperationLogUpdateOne) SoftDelete() *AdminOperationLogUpdateOne {
	aoluo.SetDeletedAt(time.Now())
	return aoluo
}

func (aolq *AdminOperationLogQuery) SoftDelete() *AdminOperationLogQuery {
	return aolq.Where(adminoperationlog.DeletedAtIsNil())
}

func (ac *AdvertiseCreate) SoftDelete() *AdvertiseCreate {
	ac.SetDeletedAt(time.Now())
	return ac
}

func (au *AdvertiseUpdate) SoftDelete() *AdvertiseUpdate {
	au.SetDeletedAt(time.Now())
	return au
}

func (auo *AdvertiseUpdateOne) SoftDelete() *AdvertiseUpdateOne {
	auo.SetDeletedAt(time.Now())
	return auo
}

func (aq *AdvertiseQuery) SoftDelete() *AdvertiseQuery {
	return aq.Where(advertise.DeletedAtIsNil())
}

func (aac *AppAgreementCreate) SoftDelete() *AppAgreementCreate {
	aac.SetDeletedAt(time.Now())
	return aac
}

func (aau *AppAgreementUpdate) SoftDelete() *AppAgreementUpdate {
	aau.SetDeletedAt(time.Now())
	return aau
}

func (aauo *AppAgreementUpdateOne) SoftDelete() *AppAgreementUpdateOne {
	aauo.SetDeletedAt(time.Now())
	return aauo
}

func (aaq *AppAgreementQuery) SoftDelete() *AppAgreementQuery {
	return aaq.Where(appagreement.DeletedAtIsNil())
}

func (avc *AppVersionCreate) SoftDelete() *AppVersionCreate {
	avc.SetDeletedAt(time.Now())
	return avc
}

func (avu *AppVersionUpdate) SoftDelete() *AppVersionUpdate {
	avu.SetDeletedAt(time.Now())
	return avu
}

func (avuo *AppVersionUpdateOne) SoftDelete() *AppVersionUpdateOne {
	avuo.SetDeletedAt(time.Now())
	return avuo
}

func (avq *AppVersionQuery) SoftDelete() *AppVersionQuery {
	return avq.Where(appversion.DeletedAtIsNil())
}

func (ac *AttachmentCreate) SoftDelete() *AttachmentCreate {
	ac.SetDeletedAt(time.Now())
	return ac
}

func (au *AttachmentUpdate) SoftDelete() *AttachmentUpdate {
	au.SetDeletedAt(time.Now())
	return au
}

func (auo *AttachmentUpdateOne) SoftDelete() *AttachmentUpdateOne {
	auo.SetDeletedAt(time.Now())
	return auo
}

func (aq *AttachmentQuery) SoftDelete() *AttachmentQuery {
	return aq.Where(attachment.DeletedAtIsNil())
}

func (cc *CityCreate) SoftDelete() *CityCreate {
	cc.SetDeletedAt(time.Now())
	return cc
}

func (cu *CityUpdate) SoftDelete() *CityUpdate {
	cu.SetDeletedAt(time.Now())
	return cu
}

func (cuo *CityUpdateOne) SoftDelete() *CityUpdateOne {
	cuo.SetDeletedAt(time.Now())
	return cuo
}

func (cq *CityQuery) SoftDelete() *CityQuery {
	return cq.Where(city.DeletedAtIsNil())
}

func (cc *CollectionCreate) SoftDelete() *CollectionCreate {
	cc.SetDeletedAt(time.Now())
	return cc
}

func (cu *CollectionUpdate) SoftDelete() *CollectionUpdate {
	cu.SetDeletedAt(time.Now())
	return cu
}

func (cuo *CollectionUpdateOne) SoftDelete() *CollectionUpdateOne {
	cuo.SetDeletedAt(time.Now())
	return cuo
}

func (cq *CollectionQuery) SoftDelete() *CollectionQuery {
	return cq.Where(collection.DeletedAtIsNil())
}

func (gcc *GroupCardCreate) SoftDelete() *GroupCardCreate {
	gcc.SetDeletedAt(time.Now())
	return gcc
}

func (gcu *GroupCardUpdate) SoftDelete() *GroupCardUpdate {
	gcu.SetDeletedAt(time.Now())
	return gcu
}

func (gcuo *GroupCardUpdateOne) SoftDelete() *GroupCardUpdateOne {
	gcuo.SetDeletedAt(time.Now())
	return gcuo
}

func (gcq *GroupCardQuery) SoftDelete() *GroupCardQuery {
	return gcq.Where(groupcard.DeletedAtIsNil())
}

func (hsc *HotSearchCreate) SoftDelete() *HotSearchCreate {
	hsc.SetDeletedAt(time.Now())
	return hsc
}

func (hsu *HotSearchUpdate) SoftDelete() *HotSearchUpdate {
	hsu.SetDeletedAt(time.Now())
	return hsu
}

func (hsuo *HotSearchUpdateOne) SoftDelete() *HotSearchUpdateOne {
	hsuo.SetDeletedAt(time.Now())
	return hsuo
}

func (hsq *HotSearchQuery) SoftDelete() *HotSearchQuery {
	return hsq.Where(hotsearch.DeletedAtIsNil())
}

func (itc *ImportTaskCreate) SoftDelete() *ImportTaskCreate {
	itc.SetDeletedAt(time.Now())
	return itc
}

func (itu *ImportTaskUpdate) SoftDelete() *ImportTaskUpdate {
	itu.SetDeletedAt(time.Now())
	return itu
}

func (ituo *ImportTaskUpdateOne) SoftDelete() *ImportTaskUpdateOne {
	ituo.SetDeletedAt(time.Now())
	return ituo
}

func (itq *ImportTaskQuery) SoftDelete() *ImportTaskQuery {
	return itq.Where(importtask.DeletedAtIsNil())
}

func (icc *InformationClassifyCreate) SoftDelete() *InformationClassifyCreate {
	icc.SetDeletedAt(time.Now())
	return icc
}

func (icu *InformationClassifyUpdate) SoftDelete() *InformationClassifyUpdate {
	icu.SetDeletedAt(time.Now())
	return icu
}

func (icuo *InformationClassifyUpdateOne) SoftDelete() *InformationClassifyUpdateOne {
	icuo.SetDeletedAt(time.Now())
	return icuo
}

func (icq *InformationClassifyQuery) SoftDelete() *InformationClassifyQuery {
	return icq.Where(informationclassify.DeletedAtIsNil())
}

func (icc *ItemCategoryCreate) SoftDelete() *ItemCategoryCreate {
	icc.SetDeletedAt(time.Now())
	return icc
}

func (icu *ItemCategoryUpdate) SoftDelete() *ItemCategoryUpdate {
	icu.SetDeletedAt(time.Now())
	return icu
}

func (icuo *ItemCategoryUpdateOne) SoftDelete() *ItemCategoryUpdateOne {
	icuo.SetDeletedAt(time.Now())
	return icuo
}

func (icq *ItemCategoryQuery) SoftDelete() *ItemCategoryQuery {
	return icq.Where(itemcategory.DeletedAtIsNil())
}

func (kcc *KcClassCreate) SoftDelete() *KcClassCreate {
	kcc.SetDeletedAt(time.Now())
	return kcc
}

func (kcu *KcClassUpdate) SoftDelete() *KcClassUpdate {
	kcu.SetDeletedAt(time.Now())
	return kcu
}

func (kcuo *KcClassUpdateOne) SoftDelete() *KcClassUpdateOne {
	kcuo.SetDeletedAt(time.Now())
	return kcuo
}

func (kcq *KcClassQuery) SoftDelete() *KcClassQuery {
	return kcq.Where(kcclass.DeletedAtIsNil())
}

func (kcc *KcCourseCreate) SoftDelete() *KcCourseCreate {
	kcc.SetDeletedAt(time.Now())
	return kcc
}

func (kcu *KcCourseUpdate) SoftDelete() *KcCourseUpdate {
	kcu.SetDeletedAt(time.Now())
	return kcu
}

func (kcuo *KcCourseUpdateOne) SoftDelete() *KcCourseUpdateOne {
	kcuo.SetDeletedAt(time.Now())
	return kcuo
}

func (kcq *KcCourseQuery) SoftDelete() *KcCourseQuery {
	return kcq.Where(kccourse.DeletedAtIsNil())
}

func (kcscc *KcCourseSmallCategoryCreate) SoftDelete() *KcCourseSmallCategoryCreate {
	kcscc.SetDeletedAt(time.Now())
	return kcscc
}

func (kcscu *KcCourseSmallCategoryUpdate) SoftDelete() *KcCourseSmallCategoryUpdate {
	kcscu.SetDeletedAt(time.Now())
	return kcscu
}

func (kcscuo *KcCourseSmallCategoryUpdateOne) SoftDelete() *KcCourseSmallCategoryUpdateOne {
	kcscuo.SetDeletedAt(time.Now())
	return kcscuo
}

func (kcscq *KcCourseSmallCategoryQuery) SoftDelete() *KcCourseSmallCategoryQuery {
	return kcscq.Where(kccoursesmallcategory.DeletedAtIsNil())
}

func (kscac *KcSmallCategoryAttachmentCreate) SoftDelete() *KcSmallCategoryAttachmentCreate {
	kscac.SetDeletedAt(time.Now())
	return kscac
}

func (kscau *KcSmallCategoryAttachmentUpdate) SoftDelete() *KcSmallCategoryAttachmentUpdate {
	kscau.SetDeletedAt(time.Now())
	return kscau
}

func (kscauo *KcSmallCategoryAttachmentUpdateOne) SoftDelete() *KcSmallCategoryAttachmentUpdateOne {
	kscauo.SetDeletedAt(time.Now())
	return kscauo
}

func (kscaq *KcSmallCategoryAttachmentQuery) SoftDelete() *KcSmallCategoryAttachmentQuery {
	return kscaq.Where(kcsmallcategoryattachment.DeletedAtIsNil())
}

func (kscepc *KcSmallCategoryExamPaperCreate) SoftDelete() *KcSmallCategoryExamPaperCreate {
	kscepc.SetDeletedAt(time.Now())
	return kscepc
}

func (kscepu *KcSmallCategoryExamPaperUpdate) SoftDelete() *KcSmallCategoryExamPaperUpdate {
	kscepu.SetDeletedAt(time.Now())
	return kscepu
}

func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SoftDelete() *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.SetDeletedAt(time.Now())
	return kscepuo
}

func (kscepq *KcSmallCategoryExamPaperQuery) SoftDelete() *KcSmallCategoryExamPaperQuery {
	return kscepq.Where(kcsmallcategoryexampaper.DeletedAtIsNil())
}

func (kscqc *KcSmallCategoryQuestionCreate) SoftDelete() *KcSmallCategoryQuestionCreate {
	kscqc.SetDeletedAt(time.Now())
	return kscqc
}

func (kscqu *KcSmallCategoryQuestionUpdate) SoftDelete() *KcSmallCategoryQuestionUpdate {
	kscqu.SetDeletedAt(time.Now())
	return kscqu
}

func (kscquo *KcSmallCategoryQuestionUpdateOne) SoftDelete() *KcSmallCategoryQuestionUpdateOne {
	kscquo.SetDeletedAt(time.Now())
	return kscquo
}

func (kscqq *KcSmallCategoryQuestionQuery) SoftDelete() *KcSmallCategoryQuestionQuery {
	return kscqq.Where(kcsmallcategoryquestion.DeletedAtIsNil())
}

func (kucc *KcUserClassCreate) SoftDelete() *KcUserClassCreate {
	kucc.SetDeletedAt(time.Now())
	return kucc
}

func (kucu *KcUserClassUpdate) SoftDelete() *KcUserClassUpdate {
	kucu.SetDeletedAt(time.Now())
	return kucu
}

func (kucuo *KcUserClassUpdateOne) SoftDelete() *KcUserClassUpdateOne {
	kucuo.SetDeletedAt(time.Now())
	return kucuo
}

func (kucq *KcUserClassQuery) SoftDelete() *KcUserClassQuery {
	return kucq.Where(kcuserclass.DeletedAtIsNil())
}

func (kucc *KcUserCourseCreate) SoftDelete() *KcUserCourseCreate {
	kucc.SetDeletedAt(time.Now())
	return kucc
}

func (kucu *KcUserCourseUpdate) SoftDelete() *KcUserCourseUpdate {
	kucu.SetDeletedAt(time.Now())
	return kucu
}

func (kucuo *KcUserCourseUpdateOne) SoftDelete() *KcUserCourseUpdateOne {
	kucuo.SetDeletedAt(time.Now())
	return kucuo
}

func (kucq *KcUserCourseQuery) SoftDelete() *KcUserCourseQuery {
	return kucq.Where(kcusercourse.DeletedAtIsNil())
}

func (kvutc *KcVideoUploadTaskCreate) SoftDelete() *KcVideoUploadTaskCreate {
	kvutc.SetDeletedAt(time.Now())
	return kvutc
}

func (kvutu *KcVideoUploadTaskUpdate) SoftDelete() *KcVideoUploadTaskUpdate {
	kvutu.SetDeletedAt(time.Now())
	return kvutu
}

func (kvutuo *KcVideoUploadTaskUpdateOne) SoftDelete() *KcVideoUploadTaskUpdateOne {
	kvutuo.SetDeletedAt(time.Now())
	return kvutuo
}

func (kvutq *KcVideoUploadTaskQuery) SoftDelete() *KcVideoUploadTaskQuery {
	return kvutq.Where(kcvideouploadtask.DeletedAtIsNil())
}

func (lc *LevelCreate) SoftDelete() *LevelCreate {
	lc.SetDeletedAt(time.Now())
	return lc
}

func (lu *LevelUpdate) SoftDelete() *LevelUpdate {
	lu.SetDeletedAt(time.Now())
	return lu
}

func (luo *LevelUpdateOne) SoftDelete() *LevelUpdateOne {
	luo.SetDeletedAt(time.Now())
	return luo
}

func (lq *LevelQuery) SoftDelete() *LevelQuery {
	return lq.Where(level.DeletedAtIsNil())
}

func (mc *MajorCreate) SoftDelete() *MajorCreate {
	mc.SetDeletedAt(time.Now())
	return mc
}

func (mu *MajorUpdate) SoftDelete() *MajorUpdate {
	mu.SetDeletedAt(time.Now())
	return mu
}

func (muo *MajorUpdateOne) SoftDelete() *MajorUpdateOne {
	muo.SetDeletedAt(time.Now())
	return muo
}

func (mq *MajorQuery) SoftDelete() *MajorQuery {
	return mq.Where(major.DeletedAtIsNil())
}

func (mdc *MajorDetailCreate) SoftDelete() *MajorDetailCreate {
	mdc.SetDeletedAt(time.Now())
	return mdc
}

func (mdu *MajorDetailUpdate) SoftDelete() *MajorDetailUpdate {
	mdu.SetDeletedAt(time.Now())
	return mdu
}

func (mduo *MajorDetailUpdateOne) SoftDelete() *MajorDetailUpdateOne {
	mduo.SetDeletedAt(time.Now())
	return mduo
}

func (mdq *MajorDetailQuery) SoftDelete() *MajorDetailQuery {
	return mdq.Where(majordetail.DeletedAtIsNil())
}

func (mdtc *MajorDetailTagCreate) SoftDelete() *MajorDetailTagCreate {
	mdtc.SetDeletedAt(time.Now())
	return mdtc
}

func (mdtu *MajorDetailTagUpdate) SoftDelete() *MajorDetailTagUpdate {
	mdtu.SetDeletedAt(time.Now())
	return mdtu
}

func (mdtuo *MajorDetailTagUpdateOne) SoftDelete() *MajorDetailTagUpdateOne {
	mdtuo.SetDeletedAt(time.Now())
	return mdtuo
}

func (mdtq *MajorDetailTagQuery) SoftDelete() *MajorDetailTagQuery {
	return mdtq.Where(majordetailtag.DeletedAtIsNil())
}

func (muqrc *MakeUserQuestionRecordCreate) SoftDelete() *MakeUserQuestionRecordCreate {
	muqrc.SetDeletedAt(time.Now())
	return muqrc
}

func (muqru *MakeUserQuestionRecordUpdate) SoftDelete() *MakeUserQuestionRecordUpdate {
	muqru.SetDeletedAt(time.Now())
	return muqru
}

func (muqruo *MakeUserQuestionRecordUpdateOne) SoftDelete() *MakeUserQuestionRecordUpdateOne {
	muqruo.SetDeletedAt(time.Now())
	return muqruo
}

func (muqrq *MakeUserQuestionRecordQuery) SoftDelete() *MakeUserQuestionRecordQuery {
	return muqrq.Where(makeuserquestionrecord.DeletedAtIsNil())
}

func (mc *MessageCreate) SoftDelete() *MessageCreate {
	mc.SetDeletedAt(time.Now())
	return mc
}

func (mu *MessageUpdate) SoftDelete() *MessageUpdate {
	mu.SetDeletedAt(time.Now())
	return mu
}

func (muo *MessageUpdateOne) SoftDelete() *MessageUpdateOne {
	muo.SetDeletedAt(time.Now())
	return muo
}

func (mq *MessageQuery) SoftDelete() *MessageQuery {
	return mq.Where(message.DeletedAtIsNil())
}

func (mtc *MessageTypeCreate) SoftDelete() *MessageTypeCreate {
	mtc.SetDeletedAt(time.Now())
	return mtc
}

func (mtu *MessageTypeUpdate) SoftDelete() *MessageTypeUpdate {
	mtu.SetDeletedAt(time.Now())
	return mtu
}

func (mtuo *MessageTypeUpdateOne) SoftDelete() *MessageTypeUpdateOne {
	mtuo.SetDeletedAt(time.Now())
	return mtuo
}

func (mtq *MessageTypeQuery) SoftDelete() *MessageTypeQuery {
	return mtq.Where(messagetype.DeletedAtIsNil())
}

func (pc *PermissionCreate) SoftDelete() *PermissionCreate {
	pc.SetDeletedAt(time.Now())
	return pc
}

func (pu *PermissionUpdate) SoftDelete() *PermissionUpdate {
	pu.SetDeletedAt(time.Now())
	return pu
}

func (puo *PermissionUpdateOne) SoftDelete() *PermissionUpdateOne {
	puo.SetDeletedAt(time.Now())
	return puo
}

func (pq *PermissionQuery) SoftDelete() *PermissionQuery {
	return pq.Where(permission.DeletedAtIsNil())
}

func (rc *RoleCreate) SoftDelete() *RoleCreate {
	rc.SetDeletedAt(time.Now())
	return rc
}

func (ru *RoleUpdate) SoftDelete() *RoleUpdate {
	ru.SetDeletedAt(time.Now())
	return ru
}

func (ruo *RoleUpdateOne) SoftDelete() *RoleUpdateOne {
	ruo.SetDeletedAt(time.Now())
	return ruo
}

func (rq *RoleQuery) SoftDelete() *RoleQuery {
	return rq.Where(role.DeletedAtIsNil())
}

func (spc *SharePosterCreate) SoftDelete() *SharePosterCreate {
	spc.SetDeletedAt(time.Now())
	return spc
}

func (spu *SharePosterUpdate) SoftDelete() *SharePosterUpdate {
	spu.SetDeletedAt(time.Now())
	return spu
}

func (spuo *SharePosterUpdateOne) SoftDelete() *SharePosterUpdateOne {
	spuo.SetDeletedAt(time.Now())
	return spuo
}

func (spq *SharePosterQuery) SoftDelete() *SharePosterQuery {
	return spq.Where(shareposter.DeletedAtIsNil())
}

func (tc *TeacherCreate) SoftDelete() *TeacherCreate {
	tc.SetDeletedAt(time.Now())
	return tc
}

func (tu *TeacherUpdate) SoftDelete() *TeacherUpdate {
	tu.SetDeletedAt(time.Now())
	return tu
}

func (tuo *TeacherUpdateOne) SoftDelete() *TeacherUpdateOne {
	tuo.SetDeletedAt(time.Now())
	return tuo
}

func (tq *TeacherQuery) SoftDelete() *TeacherQuery {
	return tq.Where(teacher.DeletedAtIsNil())
}

func (ttc *TeacherTagCreate) SoftDelete() *TeacherTagCreate {
	ttc.SetDeletedAt(time.Now())
	return ttc
}

func (ttu *TeacherTagUpdate) SoftDelete() *TeacherTagUpdate {
	ttu.SetDeletedAt(time.Now())
	return ttu
}

func (ttuo *TeacherTagUpdateOne) SoftDelete() *TeacherTagUpdateOne {
	ttuo.SetDeletedAt(time.Now())
	return ttuo
}

func (ttq *TeacherTagQuery) SoftDelete() *TeacherTagQuery {
	return ttq.Where(teachertag.DeletedAtIsNil())
}

func (tcc *TkChapterCreate) SoftDelete() *TkChapterCreate {
	tcc.SetDeletedAt(time.Now())
	return tcc
}

func (tcu *TkChapterUpdate) SoftDelete() *TkChapterUpdate {
	tcu.SetDeletedAt(time.Now())
	return tcu
}

func (tcuo *TkChapterUpdateOne) SoftDelete() *TkChapterUpdateOne {
	tcuo.SetDeletedAt(time.Now())
	return tcuo
}

func (tcq *TkChapterQuery) SoftDelete() *TkChapterQuery {
	return tcq.Where(tkchapter.DeletedAtIsNil())
}

func (tepc *TkExamPaperCreate) SoftDelete() *TkExamPaperCreate {
	tepc.SetDeletedAt(time.Now())
	return tepc
}

func (tepu *TkExamPaperUpdate) SoftDelete() *TkExamPaperUpdate {
	tepu.SetDeletedAt(time.Now())
	return tepu
}

func (tepuo *TkExamPaperUpdateOne) SoftDelete() *TkExamPaperUpdateOne {
	tepuo.SetDeletedAt(time.Now())
	return tepuo
}

func (tepq *TkExamPaperQuery) SoftDelete() *TkExamPaperQuery {
	return tepq.Where(tkexampaper.DeletedAtIsNil())
}

func (teppc *TkExamPaperPartitionCreate) SoftDelete() *TkExamPaperPartitionCreate {
	teppc.SetDeletedAt(time.Now())
	return teppc
}

func (teppu *TkExamPaperPartitionUpdate) SoftDelete() *TkExamPaperPartitionUpdate {
	teppu.SetDeletedAt(time.Now())
	return teppu
}

func (teppuo *TkExamPaperPartitionUpdateOne) SoftDelete() *TkExamPaperPartitionUpdateOne {
	teppuo.SetDeletedAt(time.Now())
	return teppuo
}

func (teppq *TkExamPaperPartitionQuery) SoftDelete() *TkExamPaperPartitionQuery {
	return teppq.Where(tkexampaperpartition.DeletedAtIsNil())
}

func (teppsc *TkExamPaperPartitionScoreCreate) SoftDelete() *TkExamPaperPartitionScoreCreate {
	teppsc.SetDeletedAt(time.Now())
	return teppsc
}

func (teppsu *TkExamPaperPartitionScoreUpdate) SoftDelete() *TkExamPaperPartitionScoreUpdate {
	teppsu.SetDeletedAt(time.Now())
	return teppsu
}

func (teppsuo *TkExamPaperPartitionScoreUpdateOne) SoftDelete() *TkExamPaperPartitionScoreUpdateOne {
	teppsuo.SetDeletedAt(time.Now())
	return teppsuo
}

func (teppsq *TkExamPaperPartitionScoreQuery) SoftDelete() *TkExamPaperPartitionScoreQuery {
	return teppsq.Where(tkexampaperpartitionscore.DeletedAtIsNil())
}

func (tepsc *TkExamPaperSimulationCreate) SoftDelete() *TkExamPaperSimulationCreate {
	tepsc.SetDeletedAt(time.Now())
	return tepsc
}

func (tepsu *TkExamPaperSimulationUpdate) SoftDelete() *TkExamPaperSimulationUpdate {
	tepsu.SetDeletedAt(time.Now())
	return tepsu
}

func (tepsuo *TkExamPaperSimulationUpdateOne) SoftDelete() *TkExamPaperSimulationUpdateOne {
	tepsuo.SetDeletedAt(time.Now())
	return tepsuo
}

func (tepsq *TkExamPaperSimulationQuery) SoftDelete() *TkExamPaperSimulationQuery {
	return tepsq.Where(tkexampapersimulation.DeletedAtIsNil())
}

func (tepqlc *TkExamPartitionQuestionLinkCreate) SoftDelete() *TkExamPartitionQuestionLinkCreate {
	tepqlc.SetDeletedAt(time.Now())
	return tepqlc
}

func (tepqlu *TkExamPartitionQuestionLinkUpdate) SoftDelete() *TkExamPartitionQuestionLinkUpdate {
	tepqlu.SetDeletedAt(time.Now())
	return tepqlu
}

func (tepqluo *TkExamPartitionQuestionLinkUpdateOne) SoftDelete() *TkExamPartitionQuestionLinkUpdateOne {
	tepqluo.SetDeletedAt(time.Now())
	return tepqluo
}

func (tepqlq *TkExamPartitionQuestionLinkQuery) SoftDelete() *TkExamPartitionQuestionLinkQuery {
	return tepqlq.Where(tkexampartitionquestionlink.DeletedAtIsNil())
}

func (teqtc *TkExamQuestionTypeCreate) SoftDelete() *TkExamQuestionTypeCreate {
	teqtc.SetDeletedAt(time.Now())
	return teqtc
}

func (teqtu *TkExamQuestionTypeUpdate) SoftDelete() *TkExamQuestionTypeUpdate {
	teqtu.SetDeletedAt(time.Now())
	return teqtu
}

func (teqtuo *TkExamQuestionTypeUpdateOne) SoftDelete() *TkExamQuestionTypeUpdateOne {
	teqtuo.SetDeletedAt(time.Now())
	return teqtuo
}

func (teqtq *TkExamQuestionTypeQuery) SoftDelete() *TkExamQuestionTypeQuery {
	return teqtq.Where(tkexamquestiontype.DeletedAtIsNil())
}

func (tkpc *TkKnowledgePointCreate) SoftDelete() *TkKnowledgePointCreate {
	tkpc.SetDeletedAt(time.Now())
	return tkpc
}

func (tkpu *TkKnowledgePointUpdate) SoftDelete() *TkKnowledgePointUpdate {
	tkpu.SetDeletedAt(time.Now())
	return tkpu
}

func (tkpuo *TkKnowledgePointUpdateOne) SoftDelete() *TkKnowledgePointUpdateOne {
	tkpuo.SetDeletedAt(time.Now())
	return tkpuo
}

func (tkpq *TkKnowledgePointQuery) SoftDelete() *TkKnowledgePointQuery {
	return tkpq.Where(tkknowledgepoint.DeletedAtIsNil())
}

func (tqc *TkQuestionCreate) SoftDelete() *TkQuestionCreate {
	tqc.SetDeletedAt(time.Now())
	return tqc
}

func (tqu *TkQuestionUpdate) SoftDelete() *TkQuestionUpdate {
	tqu.SetDeletedAt(time.Now())
	return tqu
}

func (tquo *TkQuestionUpdateOne) SoftDelete() *TkQuestionUpdateOne {
	tquo.SetDeletedAt(time.Now())
	return tquo
}

func (tqq *TkQuestionQuery) SoftDelete() *TkQuestionQuery {
	return tqq.Where(tkquestion.DeletedAtIsNil())
}

func (tqaoc *TkQuestionAnswerOptionCreate) SoftDelete() *TkQuestionAnswerOptionCreate {
	tqaoc.SetDeletedAt(time.Now())
	return tqaoc
}

func (tqaou *TkQuestionAnswerOptionUpdate) SoftDelete() *TkQuestionAnswerOptionUpdate {
	tqaou.SetDeletedAt(time.Now())
	return tqaou
}

func (tqaouo *TkQuestionAnswerOptionUpdateOne) SoftDelete() *TkQuestionAnswerOptionUpdateOne {
	tqaouo.SetDeletedAt(time.Now())
	return tqaouo
}

func (tqaoq *TkQuestionAnswerOptionQuery) SoftDelete() *TkQuestionAnswerOptionQuery {
	return tqaoq.Where(tkquestionansweroption.DeletedAtIsNil())
}

func (tqbc *TkQuestionBankCreate) SoftDelete() *TkQuestionBankCreate {
	tqbc.SetDeletedAt(time.Now())
	return tqbc
}

func (tqbu *TkQuestionBankUpdate) SoftDelete() *TkQuestionBankUpdate {
	tqbu.SetDeletedAt(time.Now())
	return tqbu
}

func (tqbuo *TkQuestionBankUpdateOne) SoftDelete() *TkQuestionBankUpdateOne {
	tqbuo.SetDeletedAt(time.Now())
	return tqbuo
}

func (tqbq *TkQuestionBankQuery) SoftDelete() *TkQuestionBankQuery {
	return tqbq.Where(tkquestionbank.DeletedAtIsNil())
}

func (tqbcc *TkQuestionBankCityCreate) SoftDelete() *TkQuestionBankCityCreate {
	tqbcc.SetDeletedAt(time.Now())
	return tqbcc
}

func (tqbcu *TkQuestionBankCityUpdate) SoftDelete() *TkQuestionBankCityUpdate {
	tqbcu.SetDeletedAt(time.Now())
	return tqbcu
}

func (tqbcuo *TkQuestionBankCityUpdateOne) SoftDelete() *TkQuestionBankCityUpdateOne {
	tqbcuo.SetDeletedAt(time.Now())
	return tqbcuo
}

func (tqbcq *TkQuestionBankCityQuery) SoftDelete() *TkQuestionBankCityQuery {
	return tqbcq.Where(tkquestionbankcity.DeletedAtIsNil())
}

func (tqbmc *TkQuestionBankMajorCreate) SoftDelete() *TkQuestionBankMajorCreate {
	tqbmc.SetDeletedAt(time.Now())
	return tqbmc
}

func (tqbmu *TkQuestionBankMajorUpdate) SoftDelete() *TkQuestionBankMajorUpdate {
	tqbmu.SetDeletedAt(time.Now())
	return tqbmu
}

func (tqbmuo *TkQuestionBankMajorUpdateOne) SoftDelete() *TkQuestionBankMajorUpdateOne {
	tqbmuo.SetDeletedAt(time.Now())
	return tqbmuo
}

func (tqbmq *TkQuestionBankMajorQuery) SoftDelete() *TkQuestionBankMajorQuery {
	return tqbmq.Where(tkquestionbankmajor.DeletedAtIsNil())
}

func (tqefc *TkQuestionErrorFeedbackCreate) SoftDelete() *TkQuestionErrorFeedbackCreate {
	tqefc.SetDeletedAt(time.Now())
	return tqefc
}

func (tqefu *TkQuestionErrorFeedbackUpdate) SoftDelete() *TkQuestionErrorFeedbackUpdate {
	tqefu.SetDeletedAt(time.Now())
	return tqefu
}

func (tqefuo *TkQuestionErrorFeedbackUpdateOne) SoftDelete() *TkQuestionErrorFeedbackUpdateOne {
	tqefuo.SetDeletedAt(time.Now())
	return tqefuo
}

func (tqefq *TkQuestionErrorFeedbackQuery) SoftDelete() *TkQuestionErrorFeedbackQuery {
	return tqefq.Where(tkquestionerrorfeedback.DeletedAtIsNil())
}

func (tqsc *TkQuestionSectionCreate) SoftDelete() *TkQuestionSectionCreate {
	tqsc.SetDeletedAt(time.Now())
	return tqsc
}

func (tqsu *TkQuestionSectionUpdate) SoftDelete() *TkQuestionSectionUpdate {
	tqsu.SetDeletedAt(time.Now())
	return tqsu
}

func (tqsuo *TkQuestionSectionUpdateOne) SoftDelete() *TkQuestionSectionUpdateOne {
	tqsuo.SetDeletedAt(time.Now())
	return tqsuo
}

func (tqsq *TkQuestionSectionQuery) SoftDelete() *TkQuestionSectionQuery {
	return tqsq.Where(tkquestionsection.DeletedAtIsNil())
}

func (tsc *TkSectionCreate) SoftDelete() *TkSectionCreate {
	tsc.SetDeletedAt(time.Now())
	return tsc
}

func (tsu *TkSectionUpdate) SoftDelete() *TkSectionUpdate {
	tsu.SetDeletedAt(time.Now())
	return tsu
}

func (tsuo *TkSectionUpdateOne) SoftDelete() *TkSectionUpdateOne {
	tsuo.SetDeletedAt(time.Now())
	return tsuo
}

func (tsq *TkSectionQuery) SoftDelete() *TkSectionQuery {
	return tsq.Where(tksection.DeletedAtIsNil())
}

func (tuesrc *TkUserExamScoreRecordCreate) SoftDelete() *TkUserExamScoreRecordCreate {
	tuesrc.SetDeletedAt(time.Now())
	return tuesrc
}

func (tuesru *TkUserExamScoreRecordUpdate) SoftDelete() *TkUserExamScoreRecordUpdate {
	tuesru.SetDeletedAt(time.Now())
	return tuesru
}

func (tuesruo *TkUserExamScoreRecordUpdateOne) SoftDelete() *TkUserExamScoreRecordUpdateOne {
	tuesruo.SetDeletedAt(time.Now())
	return tuesruo
}

func (tuesrq *TkUserExamScoreRecordQuery) SoftDelete() *TkUserExamScoreRecordQuery {
	return tuesrq.Where(tkuserexamscorerecord.DeletedAtIsNil())
}

func (tuqbrc *TkUserQuestionBankRecordCreate) SoftDelete() *TkUserQuestionBankRecordCreate {
	tuqbrc.SetDeletedAt(time.Now())
	return tuqbrc
}

func (tuqbru *TkUserQuestionBankRecordUpdate) SoftDelete() *TkUserQuestionBankRecordUpdate {
	tuqbru.SetDeletedAt(time.Now())
	return tuqbru
}

func (tuqbruo *TkUserQuestionBankRecordUpdateOne) SoftDelete() *TkUserQuestionBankRecordUpdateOne {
	tuqbruo.SetDeletedAt(time.Now())
	return tuqbruo
}

func (tuqbrq *TkUserQuestionBankRecordQuery) SoftDelete() *TkUserQuestionBankRecordQuery {
	return tuqbrq.Where(tkuserquestionbankrecord.DeletedAtIsNil())
}

func (tuqrc *TkUserQuestionRecordCreate) SoftDelete() *TkUserQuestionRecordCreate {
	tuqrc.SetDeletedAt(time.Now())
	return tuqrc
}

func (tuqru *TkUserQuestionRecordUpdate) SoftDelete() *TkUserQuestionRecordUpdate {
	tuqru.SetDeletedAt(time.Now())
	return tuqru
}

func (tuqruo *TkUserQuestionRecordUpdateOne) SoftDelete() *TkUserQuestionRecordUpdateOne {
	tuqruo.SetDeletedAt(time.Now())
	return tuqruo
}

func (tuqrq *TkUserQuestionRecordQuery) SoftDelete() *TkUserQuestionRecordQuery {
	return tuqrq.Where(tkuserquestionrecord.DeletedAtIsNil())
}

func (turerc *TkUserRandomExamRecodeCreate) SoftDelete() *TkUserRandomExamRecodeCreate {
	turerc.SetDeletedAt(time.Now())
	return turerc
}

func (tureru *TkUserRandomExamRecodeUpdate) SoftDelete() *TkUserRandomExamRecodeUpdate {
	tureru.SetDeletedAt(time.Now())
	return tureru
}

func (tureruo *TkUserRandomExamRecodeUpdateOne) SoftDelete() *TkUserRandomExamRecodeUpdateOne {
	tureruo.SetDeletedAt(time.Now())
	return tureruo
}

func (turerq *TkUserRandomExamRecodeQuery) SoftDelete() *TkUserRandomExamRecodeQuery {
	return turerq.Where(tkuserrandomexamrecode.DeletedAtIsNil())
}

func (tustmc *TkUserSimulationTeacherMarkCreate) SoftDelete() *TkUserSimulationTeacherMarkCreate {
	tustmc.SetDeletedAt(time.Now())
	return tustmc
}

func (tustmu *TkUserSimulationTeacherMarkUpdate) SoftDelete() *TkUserSimulationTeacherMarkUpdate {
	tustmu.SetDeletedAt(time.Now())
	return tustmu
}

func (tustmuo *TkUserSimulationTeacherMarkUpdateOne) SoftDelete() *TkUserSimulationTeacherMarkUpdateOne {
	tustmuo.SetDeletedAt(time.Now())
	return tustmuo
}

func (tustmq *TkUserSimulationTeacherMarkQuery) SoftDelete() *TkUserSimulationTeacherMarkQuery {
	return tustmq.Where(tkusersimulationteachermark.DeletedAtIsNil())
}

func (tuwqrc *TkUserWrongQuestionRecodeCreate) SoftDelete() *TkUserWrongQuestionRecodeCreate {
	tuwqrc.SetDeletedAt(time.Now())
	return tuwqrc
}

func (tuwqru *TkUserWrongQuestionRecodeUpdate) SoftDelete() *TkUserWrongQuestionRecodeUpdate {
	tuwqru.SetDeletedAt(time.Now())
	return tuwqru
}

func (tuwqruo *TkUserWrongQuestionRecodeUpdateOne) SoftDelete() *TkUserWrongQuestionRecodeUpdateOne {
	tuwqruo.SetDeletedAt(time.Now())
	return tuwqruo
}

func (tuwqrq *TkUserWrongQuestionRecodeQuery) SoftDelete() *TkUserWrongQuestionRecodeQuery {
	return tuwqrq.Where(tkuserwrongquestionrecode.DeletedAtIsNil())
}

func (uc *UserCreate) SoftDelete() *UserCreate {
	uc.SetDeletedAt(time.Now())
	return uc
}

func (uu *UserUpdate) SoftDelete() *UserUpdate {
	uu.SetDeletedAt(time.Now())
	return uu
}

func (uuo *UserUpdateOne) SoftDelete() *UserUpdateOne {
	uuo.SetDeletedAt(time.Now())
	return uuo
}

func (uq *UserQuery) SoftDelete() *UserQuery {
	return uq.Where(user.DeletedAtIsNil())
}

func (uaac *UserAskAnswerCreate) SoftDelete() *UserAskAnswerCreate {
	uaac.SetDeletedAt(time.Now())
	return uaac
}

func (uaau *UserAskAnswerUpdate) SoftDelete() *UserAskAnswerUpdate {
	uaau.SetDeletedAt(time.Now())
	return uaau
}

func (uaauo *UserAskAnswerUpdateOne) SoftDelete() *UserAskAnswerUpdateOne {
	uaauo.SetDeletedAt(time.Now())
	return uaauo
}

func (uaaq *UserAskAnswerQuery) SoftDelete() *UserAskAnswerQuery {
	return uaaq.Where(useraskanswer.DeletedAtIsNil())
}

func (ucac *UserCourseAppraiseCreate) SoftDelete() *UserCourseAppraiseCreate {
	ucac.SetDeletedAt(time.Now())
	return ucac
}

func (ucau *UserCourseAppraiseUpdate) SoftDelete() *UserCourseAppraiseUpdate {
	ucau.SetDeletedAt(time.Now())
	return ucau
}

func (ucauo *UserCourseAppraiseUpdateOne) SoftDelete() *UserCourseAppraiseUpdateOne {
	ucauo.SetDeletedAt(time.Now())
	return ucauo
}

func (ucaq *UserCourseAppraiseQuery) SoftDelete() *UserCourseAppraiseQuery {
	return ucaq.Where(usercourseappraise.DeletedAtIsNil())
}

func (ullc *UserLoginLogCreate) SoftDelete() *UserLoginLogCreate {
	ullc.SetDeletedAt(time.Now())
	return ullc
}

func (ullu *UserLoginLogUpdate) SoftDelete() *UserLoginLogUpdate {
	ullu.SetDeletedAt(time.Now())
	return ullu
}

func (ulluo *UserLoginLogUpdateOne) SoftDelete() *UserLoginLogUpdateOne {
	ulluo.SetDeletedAt(time.Now())
	return ulluo
}

func (ullq *UserLoginLogQuery) SoftDelete() *UserLoginLogQuery {
	return ullq.Where(userloginlog.DeletedAtIsNil())
}

func (vrc *VideoRecordCreate) SoftDelete() *VideoRecordCreate {
	vrc.SetDeletedAt(time.Now())
	return vrc
}

func (vru *VideoRecordUpdate) SoftDelete() *VideoRecordUpdate {
	vru.SetDeletedAt(time.Now())
	return vru
}

func (vruo *VideoRecordUpdateOne) SoftDelete() *VideoRecordUpdateOne {
	vruo.SetDeletedAt(time.Now())
	return vruo
}

func (vrq *VideoRecordQuery) SoftDelete() *VideoRecordQuery {
	return vrq.Where(videorecord.DeletedAtIsNil())
}
