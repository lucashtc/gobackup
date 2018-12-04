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
	command := []string{"-B", "-s", "-u", "root", "-e", "SHOW DATABASES"}
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

// GetTable execute command for getting all name tables and views
func GetTable(database string) ([]string, error) {
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

	return newStringStmt, nil
}

// GetProcedure getProcedures by database local
func GetProcedure(database string) ([]string, error) {
	var newStringStmt []string
	command := fmt.Sprintf("use mysql; select name from mysql.proc where db = '%s';", database)

	stmtCommand := []string{"-B", "-s", "-u", "root", "-e", command}
	out, err := execmysql.Exec(stmtCommand)
	if err != nil {
		return []string{}, fmt.Errorf("Falha ao executar comando %s, Errror >> %s", command, err)
	}

	stringStmt := strings.Split(string(out), "\r\n")
	for _, v := range stringStmt {
		newStringStmt = append(newStringStmt, v)
	}

	return newStringStmt, nil
}

// GetData function vao pegar todas os dados do banco, nome da base, nome das tabelas e procedures
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

		// get name table by database
		table, err := GetTable(v)
		if err != nil {
			return []DataBase{}, err
		}
		db[i].Table = table

		// get Procedure/function
		proc, err := GetProcedure(v)
		if err != nil {
			return nil, err
		}
		db[i].Procedure = proc
	}

	return db, nil

}
