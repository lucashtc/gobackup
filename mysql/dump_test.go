package mysql

import (
	"fmt"
	"testing"
)

func Test_DirDump(t *testing.T) {
	bd := []DataBase{
		{Name: "banco1"},
		{Name: "banco2"},
		{Name: "banco3"},
		{Name: "banco4"},
		{Name: "banco5"},
		{Name: "banco6"},
		{Name: "banco7"},
		{Name: "banco8"},
		{Name: "banco9"},
	}
	for _, v := range bd {

		d, err := DirDump(v.Name)
		if err != nil {
			t.Logf("Deu ruin >> %s", err)
		}
		fmt.Printf("Diretorio Criado >> %s \n", d)
	}
}
