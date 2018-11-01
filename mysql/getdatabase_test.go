package mysql_test

import (
	"testing"

	"github.com/lucashtc/gobackup/mysql"
)

func Test_GetDataBase(t *testing.T) {
	db := mysql.DB{}
	bases, err := db.GetDatabase()
	if err != nil {
		t.Errorf("Falha os obter name database >> %s", err)
	}
	//Funciona com go test -v mysql
	t.Logf("Result %s \n", bases)
}

func Test_GetTable(t *testing.T) {
	db := mysql.DB{}
	bases, err := db.GetDatabase()
	if err != nil {
		t.Errorf("Falha ao obter nome das bases de dados >> %s", err)
	}

	for _, v := range bases {
		out, err := db.GetTable(v)
		if err != nil {
			t.Errorf("Falha ao obter tables das bases  >> %s", err)
		}
		t.Logf("Resultado %s", out)

	}
}

func Test_GetProcedure(t *testing.T) {
	db := mysql.DB{}
	r, err := db.GetProcedure("acompanhamento")
	if err != nil {
		t.Errorf("Error >>> %s", err)
	}
	t.Logf("Resultado %s", r)
}
