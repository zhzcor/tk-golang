// Code generated by entc, DO NOT EDIT.

package activity

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the activity type in the database.
	Label = "activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldSubTitle holds the string denoting the sub_title field in the database.
	FieldSubTitle = "sub_title"
	// FieldCoverImgID holds the string denoting the cover_img_id field in the database.
	FieldCoverImgID = "cover_img_id"
	// FieldNotice holds the string denoting the notice field in the database.
	FieldNotice = "notice"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldPlace holds the string denoting the place field in the database.
	FieldPlace = "place"
	// FieldIsFree holds the string denoting the is_free field in the database.
	FieldIsFree = "is_free"
	// FieldIsPublish holds the string denoting the is_publish field in the database.
	FieldIsPublish = "is_publish"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldApplyStartAt holds the string denoting the apply_start_at field in the database.
	FieldApplyStartAt = "apply_start_at"
	// FieldApplyEndAt holds the string denoting the apply_end_at field in the database.
	FieldApplyEndAt = "apply_end_at"
	// FieldIsHot holds the string denoting the is_hot field in the database.
	FieldIsHot = "is_hot"
	// FieldIsAutoPublish holds the string denoting the is_auto_publish field in the database.
	FieldIsAutoPublish = "is_auto_publish"
	// FieldApplyCount holds the string denoting the apply_count field in the database.
	FieldApplyCount = "apply_count"
	// FieldJoinCount holds the string denoting the join_count field in the database.
	FieldJoinCount = "join_count"
	// FieldIsLimitJoinCount holds the string denoting the is_limit_join_count field in the database.
	FieldIsLimitJoinCount = "is_limit_join_count"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldSignRemark holds the string denoting the sign_remark field in the database.
	FieldSignRemark = "sign_remark"
	// FieldActivityTypeID holds the string denoting the activity_type_id field in the database.
	FieldActivityTypeID = "activity_type_id"
	// FieldCreatedAdminID holds the string denoting the created_admin_id field in the database.
	FieldCreatedAdminID = "created_admin_id"
	// EdgeActivityType holds the string denoting the activity_type edge name in mutations.
	EdgeActivityType = "activity_type"
	// EdgeApplyActivities holds the string denoting the apply_activities edge name in mutations.
	EdgeApplyActivities = "apply_activities"
	// EdgeAdmin holds the string denoting the admin edge name in mutations.
	EdgeAdmin = "admin"
	// Table holds the table name of the activity in the database.
	Table = "activities"
	// ActivityTypeTable is the table the holds the activity_type relation/edge.
	ActivityTypeTable = "activities"
	// ActivityTypeInverseTable is the table name for the ActivityType entity.
	// It exists in this package in order to avoid circular dependency with the "activitytype" package.
	ActivityTypeInverseTable = "activity_types"
	// ActivityTypeColumn is the table column denoting the activity_type relation/edge.
	ActivityTypeColumn = "activity_type_id"
	// ApplyActivitiesTable is the table the holds the apply_activities relation/edge.
	ApplyActivitiesTable = "activity_apply_infos"
	// ApplyActivitiesInverseTable is the table name for the ActivityApplyInfo entity.
	// It exists in this package in order to avoid circular dependency with the "activityapplyinfo" package.
	ApplyActivitiesInverseTable = "activity_apply_infos"
	// ApplyActivitiesColumn is the table column denoting the apply_activities relation/edge.
	ApplyActivitiesColumn = "activity_id"
	// AdminTable is the table the holds the admin relation/edge.
	AdminTable = "activities"
	// AdminInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	AdminInverseTable = "admins"
	// AdminColumn is the table column denoting the admin relation/edge.
	AdminColumn = "created_admin_id"
)

// Columns holds all SQL columns for activity fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTitle,
	FieldSubTitle,
	FieldCoverImgID,
	FieldNotice,
	FieldDetail,
	FieldPlace,
	FieldIsFree,
	FieldIsPublish,
	FieldAmount,
	FieldStartAt,
	FieldEndAt,
	FieldApplyStartAt,
	FieldApplyEndAt,
	FieldIsHot,
	FieldIsAutoPublish,
	FieldApplyCount,
	FieldJoinCount,
	FieldIsLimitJoinCount,
	FieldBirthday,
	FieldSignRemark,
	FieldActivityTypeID,
	FieldCreatedAdminID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "tkserver/internal/store/ent/runtime"
//
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultSubTitle holds the default value on creation for the "sub_title" field.
	DefaultSubTitle string
	// DefaultCoverImgID holds the default value on creation for the "cover_img_id" field.
	DefaultCoverImgID int
	// DefaultPlace holds the default value on creation for the "place" field.
	DefaultPlace string
	// DefaultIsFree holds the default value on creation for the "is_free" field.
	DefaultIsFree uint8
	// DefaultIsPublish holds the default value on creation for the "is_publish" field.
	DefaultIsPublish uint8
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int
	// DefaultIsHot holds the default value on creation for the "is_hot" field.
	DefaultIsHot uint8
	// DefaultIsAutoPublish holds the default value on creation for the "is_auto_publish" field.
	DefaultIsAutoPublish uint8
	// DefaultApplyCount holds the default value on creation for the "apply_count" field.
	DefaultApplyCount int
	// DefaultJoinCount holds the default value on creation for the "join_count" field.
	DefaultJoinCount int
	// DefaultIsLimitJoinCount holds the default value on creation for the "is_limit_join_count" field.
	DefaultIsLimitJoinCount uint8
)
