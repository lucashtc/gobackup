package mysql

// DataBase store info by database
type DataBase struct {
	Name      []string
	User      string
	Password  string
	Param     []string
	Table     []string
	Procedure []string
}
