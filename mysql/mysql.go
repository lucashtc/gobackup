// Package mysql ...
package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	database []string
	param    []string
}

// Define name executable
const (
	executabledump  string = "mysqldump"
	executablemysql string = "mysql"
)

func (m *mysql) GetDatabase() {

}

// Conn connetion database
func (m *mysql) Conn(*sql.DB, error) {
	pas := fmt.Sprintf(":%s", conf.Password)
	db, err := sql.Open("mysql", fmt.Sprintf("%s%s@/information_schema", conf.User, pas))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetDatabase excute command for getting name all databases
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
