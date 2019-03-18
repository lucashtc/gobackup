// Package mysql ...
package mysql

import (
	"database/sql"
	"fmt"

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
	conf := mysql.NewConfig()

	conf.User = m.User
	conf.Passwd = m.Password
	conf.DBName = m.DBName
	conf.Net = "tcp"
	conf.Addr = "localhost:3306"
	conf.Params = map[string]string{"charset": "utf8"}

	configParse := conf.FormatDSN()

	db, err := sql.Open("mysql", configParse)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetDatabase get name schemas
func (m *my) GetDatabase() ([]string, error) {
	DB, err := m.Conn()
	var dataBase []string
	if err != nil {
		return nil, err
	}
	defer DB.Close()

	stmt, err := DB.Query("SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT IN ('information_schema','performance_schema')")
	if err != nil {
		return nil, err
	}
	for stmt.Next() {
		var schema string
		err := stmt.Scan(&schema)
		if err != err {
			return nil, err
		}
		dataBase = append(dataBase, schema)
	}

	if len(dataBase) == 0 {
		return nil, fmt.Errorf("Seu banco de dados nao possui Schemas. Total bases encontradas: %d", len(dataBase))
	}

	m.Databases = dataBase

	return dataBase, nil
}

func (m *my) dump() {
	
}
