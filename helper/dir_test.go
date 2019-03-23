package helper

import (
	"fmt"
	"testing"
)

// Melhorar codigo para verificar se os arquivos foram criados como esperado
func Test_CreatDir(t *testing.T) {
	path := []string{"backup", "teste"}
	time := GetCurrentTime()

	for _, p := range path {
		d, err := CreateDir(time, p)
		if err != nil {
			t.Errorf("<<< Falha ao criar pasta da base >>>%v ", err)
		}
		fmt.Printf("Pasta Criada %s >>> \n", d)
	}

}
