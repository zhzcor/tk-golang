// Code generated by entc, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
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
	"tkserver/internal/store/ent/kcclassteacher"
	"tkserver/internal/store/ent/kccourse"
	"tkserver/internal/store/ent/kccoursechapter"
	"tkserver/internal/store/ent/kccoursesection"
	"tkserver/internal/store/ent/kccoursesmallcategory"
	"tkserver/internal/store/ent/kccourseteacher"
	"tkserver/internal/store/ent/kccoursevideo"
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
	"tkserver/internal/store/ent/rolepermission"
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
	"tkserver/internal/store/ent/useraskanswerattachment"
	"tkserver/internal/store/ent/usercourseappraise"
	"tkserver/internal/store/ent/userloginlog"
	"tkserver/internal/store/ent/videorecord"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ent aliases to avoid import conflicts in user's code.
type (
	Op         = ent.Op
	Hook       = ent.Hook
	Value      = ent.Value
	Query      = ent.Query
	Policy     = ent.Policy
	Mutator    = ent.Mutator
	Mutation   = ent.Mutation
	MutateFunc = ent.MutateFunc
)

// OrderFunc applies an ordering on the sql selector.
type OrderFunc func(*sql.Selector)

// columnChecker returns a function indicates if the column exists in the given column.
func columnChecker(table string) func(string) error {
	checks := map[string]func(string) bool{
		activity.Table:                    activity.ValidColumn,
		activityapplyinfo.Table:           activityapplyinfo.ValidColumn,
		activitytype.Table:                activitytype.ValidColumn,
		admin.Table:                       admin.ValidColumn,
		adminloginlog.Table:               adminloginlog.ValidColumn,
		adminoperationlog.Table:           adminoperationlog.ValidColumn,
		advertise.Table:                   advertise.ValidColumn,
		appagreement.Table:                appagreement.ValidColumn,
		appversion.Table:                  appversion.ValidColumn,
		attachment.Table:                  attachment.ValidColumn,
		city.Table:                        city.ValidColumn,
		collection.Table:                  collection.ValidColumn,
		groupcard.Table:                   groupcard.ValidColumn,
		hotsearch.Table:                   hotsearch.ValidColumn,
		importtask.Table:                  importtask.ValidColumn,
		informationclassify.Table:         informationclassify.ValidColumn,
		itemcategory.Table:                itemcategory.ValidColumn,
		kcclass.Table:                     kcclass.ValidColumn,
		kcclassteacher.Table:              kcclassteacher.ValidColumn,
		kccourse.Table:                    kccourse.ValidColumn,
		kccoursechapter.Table:             kccoursechapter.ValidColumn,
		kccoursesection.Table:             kccoursesection.ValidColumn,
		kccoursesmallcategory.Table:       kccoursesmallcategory.ValidColumn,
		kccourseteacher.Table:             kccourseteacher.ValidColumn,
		kccoursevideo.Table:               kccoursevideo.ValidColumn,
		kcsmallcategoryattachment.Table:   kcsmallcategoryattachment.ValidColumn,
		kcsmallcategoryexampaper.Table:    kcsmallcategoryexampaper.ValidColumn,
		kcsmallcategoryquestion.Table:     kcsmallcategoryquestion.ValidColumn,
		kcuserclass.Table:                 kcuserclass.ValidColumn,
		kcusercourse.Table:                kcusercourse.ValidColumn,
		kcvideouploadtask.Table:           kcvideouploadtask.ValidColumn,
		level.Table:                       level.ValidColumn,
		major.Table:                       major.ValidColumn,
		majordetail.Table:                 majordetail.ValidColumn,
		majordetailtag.Table:              majordetailtag.ValidColumn,
		makeuserquestionrecord.Table:      makeuserquestionrecord.ValidColumn,
		message.Table:                     message.ValidColumn,
		messagetype.Table:                 messagetype.ValidColumn,
		permission.Table:                  permission.ValidColumn,
		role.Table:                        role.ValidColumn,
		rolepermission.Table:              rolepermission.ValidColumn,
		shareposter.Table:                 shareposter.ValidColumn,
		teacher.Table:                     teacher.ValidColumn,
		teachertag.Table:                  teachertag.ValidColumn,
		tkchapter.Table:                   tkchapter.ValidColumn,
		tkexampaper.Table:                 tkexampaper.ValidColumn,
		tkexampaperpartition.Table:        tkexampaperpartition.ValidColumn,
		tkexampaperpartitionscore.Table:   tkexampaperpartitionscore.ValidColumn,
		tkexampapersimulation.Table:       tkexampapersimulation.ValidColumn,
		tkexampartitionquestionlink.Table: tkexampartitionquestionlink.ValidColumn,
		tkexamquestiontype.Table:          tkexamquestiontype.ValidColumn,
		tkknowledgepoint.Table:            tkknowledgepoint.ValidColumn,
		tkquestion.Table:                  tkquestion.ValidColumn,
		tkquestionansweroption.Table:      tkquestionansweroption.ValidColumn,
		tkquestionbank.Table:              tkquestionbank.ValidColumn,
		tkquestionbankcity.Table:          tkquestionbankcity.ValidColumn,
		tkquestionbankmajor.Table:         tkquestionbankmajor.ValidColumn,
		tkquestionerrorfeedback.Table:     tkquestionerrorfeedback.ValidColumn,
		tkquestionsection.Table:           tkquestionsection.ValidColumn,
		tksection.Table:                   tksection.ValidColumn,
		tkuserexamscorerecord.Table:       tkuserexamscorerecord.ValidColumn,
		tkuserquestionbankrecord.Table:    tkuserquestionbankrecord.ValidColumn,
		tkuserquestionrecord.Table:        tkuserquestionrecord.ValidColumn,
		tkuserrandomexamrecode.Table:      tkuserrandomexamrecode.ValidColumn,
		tkusersimulationteachermark.Table: tkusersimulationteachermark.ValidColumn,
		tkuserwrongquestionrecode.Table:   tkuserwrongquestionrecode.ValidColumn,
		user.Table:                        user.ValidColumn,
		useraskanswer.Table:               useraskanswer.ValidColumn,
		useraskanswerattachment.Table:     useraskanswerattachment.ValidColumn,
		usercourseappraise.Table:          usercourseappraise.ValidColumn,
		userloginlog.Table:                userloginlog.ValidColumn,
		videorecord.Table:                 videorecord.ValidColumn,
	}
	check, ok := checks[table]
	if !ok {
		return func(string) error {
			return fmt.Errorf("unknown table %q", table)
		}
	}
	return func(column string) error {
		if !check(column) {
			return fmt.Errorf("unknown column %q for table %q", column, table)
		}
		return nil
	}
}

// Asc applies the given fields in ASC order.
func Asc(fields ...string) OrderFunc {
	return func(s *sql.Selector) {
		check := columnChecker(s.TableName())
		for _, f := range fields {
			if err := check(f); err != nil {
				s.AddError(&ValidationError{Name: f, err: fmt.Errorf("ent: %w", err)})
			}
			s.OrderBy(sql.Asc(s.C(f)))
		}
	}
}

// Desc applies the given fields in DESC order.
func Desc(fields ...string) OrderFunc {
	return func(s *sql.Selector) {
		check := columnChecker(s.TableName())
		for _, f := range fields {
			if err := check(f); err != nil {
				s.AddError(&ValidationError{Name: f, err: fmt.Errorf("ent: %w", err)})
			}
			s.OrderBy(sql.Desc(s.C(f)))
		}
	}
}

// AggregateFunc applies an aggregation step on the group-by traversal/selector.
type AggregateFunc func(*sql.Selector) string

// As is a pseudo aggregation function for renaming another other functions with custom names. For example:
//
//	GroupBy(field1, field2).
//	Aggregate(ent.As(ent.Sum(field1), "sum_field1"), (ent.As(ent.Sum(field2), "sum_field2")).
//	Scan(ctx, &v)
//
func As(fn AggregateFunc, end string) AggregateFunc {
	return func(s *sql.Selector) string {
		return sql.As(fn(s), end)
	}
}

// Count applies the "count" aggregation function on each group.
func Count() AggregateFunc {
	return func(s *sql.Selector) string {
		return sql.Count("*")
	}
}

// Max applies the "max" aggregation function on the given field of each group.
func Max(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Max(s.C(field))
	}
}

// Mean applies the "mean" aggregation function on the given field of each group.
func Mean(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Avg(s.C(field))
	}
}

// Min applies the "min" aggregation function on the given field of each group.
func Min(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Min(s.C(field))
	}
}

// Sum applies the "sum" aggregation function on the given field of each group.
func Sum(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Sum(s.C(field))
	}
}

// ValidationError returns when validating a field fails.
type ValidationError struct {
	Name string // Field or edge name.
	err  error
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return e.err.Error()
}

// Unwrap implements the errors.Wrapper interface.
func (e *ValidationError) Unwrap() error {
	return e.err
}

// IsValidationError returns a boolean indicating whether the error is a validaton error.
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}
	var e *ValidationError
	return errors.As(err, &e)
}

// NotFoundError returns when trying to fetch a specific entity and it was not found in the database.
type NotFoundError struct {
	label string
}

// Error implements the error interface.
func (e *NotFoundError) Error() string {
	return "ent: " + e.label + " not found"
}

// IsNotFound returns a boolean indicating whether the error is a not found error.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *NotFoundError
	return errors.As(err, &e)
}

// MaskNotFound masks not found error.
func MaskNotFound(err error) error {
	if IsNotFound(err) {
		return nil
	}
	return err
}

// NotSingularError returns when trying to fetch a singular entity and more then one was found in the database.
type NotSingularError struct {
	label string
}

// Error implements the error interface.
func (e *NotSingularError) Error() string {
	return "ent: " + e.label + " not singular"
}

// IsNotSingular returns a boolean indicating whether the error is a not singular error.
func IsNotSingular(err error) bool {
	if err == nil {
		return false
	}
	var e *NotSingularError
	return errors.As(err, &e)
}

// NotLoadedError returns when trying to get a node that was not loaded by the query.
type NotLoadedError struct {
	edge string
}

// Error implements the error interface.
func (e *NotLoadedError) Error() string {
	return "ent: " + e.edge + " edge was not loaded"
}

// IsNotLoaded returns a boolean indicating whether the error is a not loaded error.
func IsNotLoaded(err error) bool {
	if err == nil {
		return false
	}
	var e *NotLoadedError
	return errors.As(err, &e)
}

// ConstraintError returns when trying to create/update one or more entities and
// one or more of their constraints failed. For example, violation of edge or
// field uniqueness.
type ConstraintError struct {
	msg  string
	wrap error
}

// Error implements the error interface.
func (e ConstraintError) Error() string {
	return "ent: constraint failed: " + e.msg
}

// Unwrap implements the errors.Wrapper interface.
func (e *ConstraintError) Unwrap() error {
	return e.wrap
}

// IsConstraintError returns a boolean indicating whether the error is a constraint failure.
func IsConstraintError(err error) bool {
	if err == nil {
		return false
	}
	var e *ConstraintError
	return errors.As(err, &e)
}

func isSQLConstraintError(err error) (*ConstraintError, bool) {
	if sqlgraph.IsConstraintError(err) {
		return &ConstraintError{err.Error(), err}, true
	}
	return nil, false
}

// rollback calls tx.Rollback and wraps the given error with the rollback error if present.
func rollback(tx dialect.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	if err, ok := isSQLConstraintError(err); ok {
		return err
	}
	return err
}
