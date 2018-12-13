package mysql

import (
	"fmt"
	"testing"

	"github.com/lucashtc/gobackup/helper"
)

func Test_DirDump(t *testing.T) {
	time := helper.GetCurrentTime()
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

		d, err := DirDump(time, v.Name)
		if err != nil {
			t.Logf("Deu ruin >> %s", err)
		}
		fmt.Printf("Diretorio Criado >> %s \n", d)
	}
}

// func Test_DumpAll(t *testing.T) {
// 	DumpAll()
// }

func Test_Log(t *testing.T) {
	time := helper.GetCurrentTime()
	text := "Criando dumper"
	local := "C:/backup/"
	Log(local, time, text)
}
