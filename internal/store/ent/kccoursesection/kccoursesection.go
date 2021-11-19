// Code generated by entc, DO NOT EDIT.

package kccoursesection

const (
	// Label holds the string label denoting the kccoursesection type in the database.
	Label = "kc_course_section"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCourseChapterID holds the string denoting the course_chapter_id field in the database.
	FieldCourseChapterID = "course_chapter_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// EdgeChapter holds the string denoting the chapter edge name in mutations.
	EdgeChapter = "chapter"
	// EdgeCourseSmallSections holds the string denoting the course_small_sections edge name in mutations.
	EdgeCourseSmallSections = "course_small_sections"
	// Table holds the table name of the kccoursesection in the database.
	Table = "kc_course_sections"
	// ChapterTable is the table the holds the chapter relation/edge.
	ChapterTable = "kc_course_sections"
	// ChapterInverseTable is the table name for the KcCourseChapter entity.
	// It exists in this package in order to avoid circular dependency with the "kccoursechapter" package.
	ChapterInverseTable = "kc_course_chapters"
	// ChapterColumn is the table column denoting the chapter relation/edge.
	ChapterColumn = "course_chapter_id"
	// CourseSmallSectionsTable is the table the holds the course_small_sections relation/edge.
	CourseSmallSectionsTable = "kc_course_small_categories"
	// CourseSmallSectionsInverseTable is the table name for the KcCourseSmallCategory entity.
	// It exists in this package in order to avoid circular dependency with the "kccoursesmallcategory" package.
	CourseSmallSectionsInverseTable = "kc_course_small_categories"
	// CourseSmallSectionsColumn is the table column denoting the course_small_sections relation/edge.
	CourseSmallSectionsColumn = "section_id"
)

// Columns holds all SQL columns for kccoursesection fields.
var Columns = []string{
	FieldID,
	FieldCourseChapterID,
	FieldTitle,
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

var (
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
)
