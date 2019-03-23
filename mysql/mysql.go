// Package mysql ...
package mysql

import (
	"database/sql"
	"fmt"

	"github.com/lucashtc/gobackup/helper"

	"github.com/go-sql-driver/mysql"
)

// My ...
type My struct {
	Databases       []string // Stored name databases
	DBName          string
	Param           []string // Parameters for dumps
	Password        string
	User            string
	Port            int
	Dir             string // Path for directory dumps current by database
	DirPath         string // Directory for stored dumps
	File            string // Name of the current dump
	DirRelative     string
	Executabledump  string
	Executablemysql string
	CurrentTime     string
}

// New return struct instance
func New() *My {
	return &My{
		Executabledump:  "mysqldump",
		Executablemysql: "mysql",
		DirPath:         "/backup/",
		CurrentTime:     helper.GetCurrentTime(),
		Port:            3306,
	}
}

// Conn connetion database
func (m *My) Conn() (*sql.DB, error) {
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
func (m *My) GetDatabase() ([]string, error) {
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

// CreateParam ...
func (m *My) CreateParam() {
	var param []string
	m.Param = []string{}
	param = append(param, "-u", m.User)
	if m.Password != "" {
		param = append(param, "-p", m.Password)
	}
	param = append(param, "-r", m.DirRelative)
	param = append(param, m.Param...)
	param = append(param, m.DBName)
	m.Param = param

}

// Exec ...
func (m *My) Exec() error {
	m.CreateParam()
	_, err := helper.Exec(m.Executabledump, m.Param)
	if err != nil {
		return err
	}
	return nil
}

// Dump ...
func (m *My) Dump() error {
	_, err := m.GetDatabase()
	if err != nil {
		return err
	}

	m.Dir, err = helper.CreateDir(m.CurrentTime, m.DirPath)
	for _, b := range m.Databases {
		m.DBName = b

		m.File = fmt.Sprintf("%s.sql", b)
		m.DirRelative = fmt.Sprintf("%s/%s", m.Dir, m.File)

		err = helper.CreateFile(m.DirRelative)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		err = m.Exec()
		if err != nil {
			return err
		}
	}
	return nil
}
