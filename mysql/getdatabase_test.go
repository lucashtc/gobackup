package mysql_test

import (
	"testing"

	"github.com/lucashtc/gobackup/helper"
	"github.com/lucashtc/gobackup/mysql"
)

func Test_GetDataBase(t *testing.T) {
	cf := mysql.DataBase{}
	cf.Dir = "/backup/"
	cf.User = "root"
	bases, err := mysql.GetDatabase(cf)
	if err != nil {
		t.Errorf("Falha os obter name database >> %s", err)
	}
	//Funciona com go test -v mysql
	t.Logf("Result %s \n", bases)
}

func Test_GetData(t *testing.T) {
	cf := mysql.DataBase{}
	cf.Dir = "/backup/"
	cf.User = "root"
	//var db []mysql.DataBase
	r, err := mysql.GetData(cf)
	if err != nil {
		t.Errorf("Falha ao obter dados do banco de dados >> %s", err)
	}

	helper.PrettyPrint(r)
}
