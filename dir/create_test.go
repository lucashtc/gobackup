package dir

import (
	"testing"
)

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
