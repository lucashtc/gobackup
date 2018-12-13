// Package mysql resposanvel por pegar nome das bases, tabelas e functions/procedures no servidor mysql
// Vai pegar todas as bases do servidor local
package mysql

// package responsável por obter informações do banco, como nome das bases e tabelas nas bases
import (
	"fmt"
	"strings"

	"github.com/lucashtc/gobackup/execmysql"
)

// GetDatabase excute command for getting name all databases
func GetDatabase() ([]string, error) {
	var newStringStmt []string
	command := []string{"-B", "-s", "-u", "root", "-e", "SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT IN ('mysql','information_schema','performance_schema') "}
	out, err := execmysql.Exec(command)
	if err != nil {
		return []string{}, fmt.Errorf("Falha ao pegar name das bases >> %s", err)
	}

	stringStmt := strings.Split(string(out), "\r\n")

	for _, v := range stringStmt {
		if v != "" {
			newStringStmt = append(newStringStmt, v)
		}
	}
	return newStringStmt, nil
}

// GetData function get name databse
// retorna um array com essas informações
func GetData() ([]DataBase, error) {

	base, err := GetDatabase()
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
