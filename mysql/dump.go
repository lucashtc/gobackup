// Package mysql  vai egar todas as infrmações das base de dados e fazer dump
package mysql

import (
	"fmt"

	"github.com/lucashtc/gobackup/dir"
	"github.com/lucashtc/gobackup/execmysql"
	"github.com/lucashtc/gobackup/helper"
	"github.com/mholt/archiver"
	//"github.com/lucashtc/gobackup/mysql"
)

// DirDump function create directory for save files backup
// Separando por pastas cadas tipo de objeto do schema
func DirDump(time, dataBase string) (string, error) {
	d, err := dir.CreateDir(time, dataBase, "")
	if err != nil {
		return "", err
	}

	return d, nil
}

// DumpAll function vai realizar dump de toda o Schema encontrado no servidor
func DumpAll() {

	time := helper.GetCurrentTime()

	fmt.Printf("Buscando informações dos schemas no servidor...\n")
	db, err := GetData()
	if err != nil {
		fmt.Printf("Falha ao obter dados da bases >>>> \n %s", err)
	}

	fmt.Printf("Criando dumps... \n\n")
	for _, d := range db {
		if d.Name == "" {
			//Caso o Name seja empty pula para a proxima execução do For
			continue
		}
		fmt.Printf("=================================================== \n")

		fmt.Printf("Fazendo backup da base %s\n", d.Name)
		fmt.Printf("Criando pasta da base %s\n", d.Name)

		dirName, err := DirDump(time, d.Name)
		if err != nil {
			fmt.Printf("Falha ao criar pasta %s \n error: >>>> error %s \n", d.Name, err)
		}

		dirNameFile := fmt.Sprintf("%s/%s.sql", dirName, d.Name)

		fmt.Printf("Realizando dump da base %s \n", d.Name)

		param := []string{"-u", "root", "--no-create-db", "--skip-add-drop-table", d.Name, "-r", dirNameFile}
		_, err = execmysql.ExecDump(param)
		if err != nil {
			fmt.Printf("Falha ao executar dump da base %s \n Error >> %s", d.Name, err)
		}

		fmt.Printf("Compactando arquivo da base \n")

		err = archiver.Archive([]string{dirNameFile}, dirNameFile+".zip")
		if err != nil {
			fmt.Printf("Ocorreu um error ao compactar o arquivo %s \n", dirNameFile)
		}

		if err := dir.Delete(dirNameFile); err != nil {
			fmt.Printf("Error: %s", err)
		}
		fmt.Printf("=================================================== \n")
	}
}

// Log ...
func Log(name, time, text string) {
	fmt.Println(text)
	local := fmt.Sprintf("%s/log.txt", name)
	if err := dir.CreateFile(local); err != nil {
		fmt.Printf("Falha ao criar arquivo. Error %s\n", err)
	}

	text = fmt.Sprintf("%s========%s", time, text)
	if err := dir.WriteFile(local, text); err != nil {
		fmt.Printf("Falha ao escrever no arquivo %s. Error %s\n", local, err)
	}
}
