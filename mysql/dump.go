// Package mysql  vai egar todas as infrmações das base de dados e fazer dump
package mysql

import (
	"fmt"

	"github.com/lucashtc/gobackup/dir"
	//"github.com/lucashtc/gobackup/mysql"
)

// DirDump function create directory for save files backup
// Separando por pastas cadas tipo de objeto do schema
func DirDump(dataBase string) (string, error) {

	d, err := dir.CreateDir(dataBase, "")
	if err != nil {
		return "", err
	}

	return d, nil
}

// DumpAll function vai realizar dump de toda o Schema encontrado no servidor
func DumpAll() {

	fmt.Printf("Buscando ifnromações dos schemas no servidor...\n")
	db, err := GetData()
	if err != nil {
		fmt.Printf("Falha ao obter dados da bases >>>> \n %s", err)
	}

	fmt.Printf("Criando dumps backups ... \n\n")
	for _, d := range db {

		fmt.Printf("Fazendo backup da base %s", d.Name)
		_, err = DirDump(d.Name)
		if err != nil {
			fmt.Printf("Falha ao criar pasta %s \n error: >>>> err or%s \n", d.Name, err)
		}
	}
}
