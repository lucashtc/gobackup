package dir

import (
	"testing"
)

// Melhorar codigo para verificar se os arquivos firam criados como esperado
func Test_CreatDir(t *testing.T) {
	bd := []string{"user", "produtos", "umGrandeVazio"}
	path := "/backup"
	for _, v := range bd {
		_, err := CreateDir(v, path)
		if err != nil {
			t.Errorf("<<< Falha ao criar pasta da base >>>%v ", err)
		}
	}
}
