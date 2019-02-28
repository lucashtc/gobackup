// Package mysql ...
package mysql

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type my struct {
	Databases []string
	DBName    string
	Param     []string
	Password  string
	User      string
	Dir       string
	File      string
}

// Define name executable
const (
	executabledump  string = "mysqldump"
	executablemysql string = "mysql"
)

// Conn connetion database
func (m *my) Conn() (*sql.DB, error) {
	var conf = mysql.Config{
		User:   m.User,
		Passwd: m.Password,
		DBName: m.DBName,
		Net:    "tcp",
		Addr:   "localhost:3306",
		Params: map[string]string{"charset": "utf8"},
	}

	configParse := conf.FormatDSN()

	db, err := sql.Open("mysql", configParse)
	if err != nil {
		return nil, err
	}

	return db, nil
}

/*
func (m *mysql) GetDatabase() ([]string, error) {

}

// GetDatabase execute command for getting name all databases
func GetDatabase(m *mysql) ([]string, error) {
	var dataBases []string
	db, err := Conn(m)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err := db.Query("SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT IN ('mysql','information_schema','performance_schema') ")
	if err != nil {
		return nil, err
	}
	for result.Next() {
		var schema string
		err := result.Scan(&schema)
		if err != nil {
			return nil, err
		}
		dataBases = append(dataBases, schema)
	}

	return dataBases, nil
}
*/
