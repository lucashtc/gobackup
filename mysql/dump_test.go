package mysql

import (
	"fmt"
	"testing"
)

func Test_DirDump(t *testing.T) {
	bd := []DataBase{
		DataBase{Name: "banco1"},
		DataBase{Name: "banco2"},
		DataBase{Name: "banco3"},
		DataBase{Name: "banco4"},
		DataBase{Name: "banco5"},
		DataBase{Name: "banco6"},
		DataBase{Name: "banco7"},
		DataBase{Name: "banco8"},
		DataBase{Name: "banco9"},
	}
	for _, v := range bd {

		d, err := DirDump(v.Name)
		if err != nil {
			t.Logf("Deu ruin >> %s", err)
		}
		fmt.Printf("Diretorio Criado >> %s \n", d)
	}
}
