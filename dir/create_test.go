package dir

import (
	"fmt"
	"testing"
)

// Melhorar codigo para verificar se os arquivos foram criados como esperado
func Test_CreatDir(t *testing.T) {
	bd := []string{"user", "produtos", "umGrandeVazio"}
	path := "/backup"

	for _, v := range bd {
		d, err := CreateDir(v, path)
		if err != nil {
			t.Errorf("<<< Falha ao criar pasta da base >>>%v ", err)
		}
		fmt.Printf("Pasta Criada %s >>> \n", d)
	}

}
