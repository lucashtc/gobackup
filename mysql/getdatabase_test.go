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
	t.Logf("Result %v \n", bases)
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

func Test_Conn(t *testing.T) {
	conf := mysql.DataBase{}
	conf.User = "root"

	db, err := mysql.Conn(conf)
	if err != nil {
		t.Fatalf("Falha ao obter nome dos Schemas. Error: %s", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Fatalf("Falha ao Pinga servidor de banco. Error: %s", err)
	}

}
