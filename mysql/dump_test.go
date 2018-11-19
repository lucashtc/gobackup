package mysql

import (
	"fmt"
	"testing"
)

func Test_DirDump(t *testing.T) {
	bd := make([]DataBase, 2)
	bd[0] = DataBase{
		Name: "banco1",
		//Table:     []string{"table1", "table2"},
		//Procedure: []string{"rot_faz_nada", "rot_update_sem_where"},
	}
	bd[1] = DataBase{
		Name: "banco2",
		//Table:     []string{"tabela1", "tabela2"},
		//Procedure: []string{"rot_pedeu_tudo", "rot_sem_backup"},
	}

	for _, v := range bd {

		d, err := DirDump(v.Name)
		if err != nil {
			t.Logf("Deu ruin >> %s", err)
		}
		fmt.Printf("Diretorio Criado >> %s \n", d)
	}
}
