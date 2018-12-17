// Package mysql resposanvel por pegar nome das bases, tabelas e functions/procedures no servidor mysql
// Vai pegar todas as bases do servidor local
package mysql

// package responsável por obter informações do banco, como nome das bases e tabelas nas bases
import (
	"database/sql"
	"fmt"

	// Interface
	_ "github.com/go-sql-driver/mysql"
)

// Conn connetion database
func Conn(conf *DataBase) (*sql.DB, error) {
	pas := fmt.Sprintf(":%s", conf.Password)
	db, err := sql.Open("mysql", fmt.Sprintf("%s%s@/information_schema", conf.User, pas))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetDatabase excute command for getting name all databases
func GetDatabase(cf *DataBase) ([]string, error) {
	var dataBases []string
	db, err := Conn(cf)
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

// GetData function get name databse
// retorna um array com essas informações
func GetData(cf *DataBase) ([]DataBase, error) {

	base, err := GetDatabase(cf)
	if err != nil {
		return []DataBase{}, err
	}

	//Alimenta cada posição do slice indexado com name dos schemas, com nome das tabelas e procedures/functions
	db := make([]DataBase, len(base))

	for i, v := range base {
		db[i].Name = v
	}

	return db, nil

}
