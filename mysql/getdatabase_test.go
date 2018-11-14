package mysql_test

import (
	"testing"

	"github.com/lucashtc/gobackup/mysql"
)

func Test_GetDataBase(t *testing.T) {
	bases, err := mysql.GetDatabase()
	if err != nil {
		t.Errorf("Falha os obter name database >> %s", err)
	}
	//Funciona com go test -v mysql
	t.Logf("Result %s \n", bases)
}

func Test_GetTable(t *testing.T) {
	bases, err := mysql.GetDatabase()
	if err != nil {
		t.Errorf("Falha ao obter nome das bases de dados >> %s", err)
	}

	for _, v := range bases {
		out, err := mysql.GetTable(v)
		if err != nil {
			t.Errorf("Falha ao obter tables das bases  >> %s", err)
		}
		t.Log("\n")
		t.Logf("Base >> %s \n", v)
		for _, ta := range out {
			t.Logf("Tabela >>>>> %s \n", ta)
		}
	}
}

func Test_GetProcedure(t *testing.T) {
	bases, err := mysql.GetDatabase()
	if err != nil {
		t.Errorf("Falha ao obter nome das bases de dados >> %s", err)
	}

	for _, v := range bases {
		r, err := mysql.GetProcedure(v)
		if err != nil {
			t.Errorf("Error >>> %s", err)
		}

		t.Logf("\n")
		t.Logf("Base > %s \n", v)
		for _, po := range r {
			t.Logf("Function or Procedure >> %s", po)
		}
	}

}

func Test_GetData(t *testing.T) {
	//var db []mysql.DataBase
	r, err := mysql.GetData()
	if err != nil {
		t.Errorf("Falha ao obter dados do banco de dados >> %s", err)
	}

	t.Logf("Resultado>>>> \n\n %s \n\n", r)
}
