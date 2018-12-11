// Package mysql  vai egar todas as infrmações das base de dados e fazer dump
package mysql

import (
	"fmt"

	"github.com/lucashtc/gobackup/dir"
	"github.com/lucashtc/gobackup/execmysql"
	"github.com/mholt/archiver"
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
	var dbName string
	fmt.Printf("Buscando informações dos schemas no servidor...\n")
	db, err := GetData()
	if err != nil {
		fmt.Printf("Falha ao obter dados da bases >>>> \n %s", err)
	}

	fmt.Printf("Criando dumps... \n\n")
	for _, d := range db {
		dbName = d.Name
		if dbName == '' {
		fmt.Printf("=================================================== \n")

		fmt.Printf("Fazendo backup da base %s\n", d.Name)
		fmt.Printf("Criando pasta da base %s\n", d.Name)

		dirName, err := DirDump(d.Name)
		if err != nil {
			fmt.Printf("Falha ao criar pasta %s \n error: >>>> error %s \n", d.Name, err)
		}

		dirName = fmt.Sprintf("%s/%s.sql", dirName, d.Name)

		fmt.Printf("Realizando dump da base %s \n", d.Name)

		param := []string{"-u", "root", "--no-create-db", d.Name, "-r", dirName}
		_, err = execmysql.ExecDump(param)
		if err != nil {
			fmt.Printf("Falha ao executar dump da base %s \n Error >> %s", d.Name, err)
		}

		fmt.Printf("Compactando arquivo da base \n")

		err = archiver.Archive([]string{dirName}, dirName+"test.zip")
		if err != nil {
			fmt.Printf("Ocorreu um error ao compactar o arquivo %s \n", dirName)
		}
		fmt.Printf("=================================================== \n")
	}
	}
}
