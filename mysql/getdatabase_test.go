package mysql

import (
	"testing"
)

func Test_GetDataBase(t *testing.T) {
	db := DB{}
	bases, err := db.GetDatabase()
	if err != nil {
		t.Errorf("Falha os obter name database >> %s", err)
	}
	//Funciona com go test -v mysql
	t.Logf("Result %s \n", bases)
}

func Test_GetTable(t *testing.T) {
	db := DB{}
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
