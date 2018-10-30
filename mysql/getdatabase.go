package mysql

import (
	"fmt"
	"strings"

	"github.com/lucashtc/gobackup/execmysql"
)

// DB store data databases
type DB struct {
	Database  []string
	Table     []string
	Param     []string
	Trigger   []string
	View      []string
	Procedure []string
}

// GetDatabase excute command for getting name all databases
func (db *DB) GetDatabase() ([]string, error) {
	var newStringStmt []string
	command := []string{"-B", "-s", "-u", "root", "-e", "SHOW DATABASES"}
	out, err := execmysql.Exec(command)
	if err != nil {
		return []string{}, fmt.Errorf("Falha ao getting name databases >> %s", err)
	}

	stringStmt := strings.Split(string(out), "\r\n")

	for _, v := range stringStmt {
		newStringStmt = append(newStringStmt, v)
	}
	db.Database = newStringStmt
	return newStringStmt, nil
}

// GetTables execute command for getting all name tables and views
func (db *DB) GetTable(database string) ([]string, error) {
	var newStringStmt []string
	stmtCommand := fmt.Sprintf("USE %s; SHOW TABLES;", database)
	command := []string{"-B", "-s", "-u", "root", "-e", stmtCommand}
	out, err := execmysql.Exec(command)
	if err != nil {
		return []string{}, fmt.Errorf("Falha ao executar comando %s, errror >>> %s", command, err)
	}
	stringStmt := strings.Split(string(out), "\r\n")
	for _, v := range stringStmt {
		if v != "" {
			newStringStmt = append(newStringStmt, v)
		}
	}

	db.Table = newStringStmt
	return newStringStmt, nil
}

// // GetProcedure getProcedures by database
// func (db *DB) GetProcedure(database string) ([]string, error) {
// 	var newStringStmt []string
// 	stmtCommand := fmt.Sprintf("USE %s; SHOW PROCEDURES")
// 	return []string{}, nil
// }
