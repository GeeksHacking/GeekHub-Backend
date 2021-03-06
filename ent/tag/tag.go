// Code generated by entc, DO NOT EDIT.

package tag

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// ProjectsTable is the table that holds the projects relation/edge. The primary key declared below.
	ProjectsTable = "project_tags"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// ProjectsPrimaryKey and ProjectsColumn2 are the table columns denoting the
	// primary key for the projects relation (M2M).
	ProjectsPrimaryKey = []string{"project_id", "tag_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
