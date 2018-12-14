// Package mysql  vai egar todas as infrmações das base de dados e fazer dump
package mysql

import (
	"fmt"
	"os"

	"github.com/lucashtc/gobackup/dir"
	"github.com/lucashtc/gobackup/execmysql"
	"github.com/lucashtc/gobackup/helper"
	"github.com/mholt/archiver"
	//"github.com/lucashtc/gobackup/mysql"
)

// DirDump function create directory for save files backup
// Separando por pastas cadas tipo de objeto do schema
func DirDump(local, time, dataBase string) (string, error) {
	d, err := dir.CreateDir(time, dataBase, local)
	if err != nil {
		return "", err
	}
	return d, nil
}

// DumpAll function vai realizar dump de toda o Schema encontrado no servidor
func DumpAll(cf DataBase) {

	time := helper.GetCurrentTime()

	Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("=================================================== \n"))
	Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("=================================================== \n"))

	Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Buscando informações dos schemas no servidor...\n"))
	db, err := GetData(cf)
	if err != nil {
		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Falha ao obter dados da bases >>>> \n %s", err))
	}

	Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Criando dumps... \n\n"))
	for _, d := range db {
		if d.Name == "" {
			//Caso o Name seja empty pula para a proxima execução do For
			continue
		}
		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("=================================================== \n"))

		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Fazendo backup da base %s\n", d.Name))
		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Criando pasta da base %s\n", d.Name))

		dirName, err := DirDump(cf.Dir, time, d.Name)
		if err != nil {
			Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Falha ao criar pasta %s \n error: >>>> error %s \n", d.Name, err))
			continue
		}

		dirNameFile := fmt.Sprintf("%s/%s.sql", dirName, d.Name)

		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Realizando dump da base %s \n", d.Name))
		if cf.Password == "" {
			pass := fmt.Sprintf("-p%s", cf.Password)
		} else {
			pass := ""
		}

		param := []string{"-u", cf.User, pass, "--no-create-db", "--skip-add-drop-table", d.Name, "-r", dirNameFile}
		_, err = execmysql.ExecDump(param)
		if err != nil {
			Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Falha ao executar dump da base %s \n Error >> %s", d.Name, err))
			continue
		}

		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Compactando arquivo da base \n"))

		err = archiver.Archive([]string{dirNameFile}, dirNameFile+".zip")
		if err != nil {
			Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Ocorreu um error ao compactar o arquivo %s \n", dirNameFile))
			continue
		}

		if err := dir.Delete(dirNameFile); err != nil {
			Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Error: %s", err))
			continue
		}
		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("=================================================== \n"))
	}

	Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf(" \n\n\nFim da execução do script de backup\n\n"))
}

// Log ...
func Log(name, time, text string) {
	fmt.Print(text)

	if err := os.MkdirAll(name, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	local := fmt.Sprintf("%s/log.txt", name)
	if err := dir.CreateFile(local); err != nil {
		fmt.Printf("Falha ao criar arquivo. Error %s\n", err)
	}

	text = fmt.Sprintf("%s========%s\r\n", time, text)
	if err := dir.WriteFile(local, text); err != nil {
		fmt.Printf("Falha ao escrever no arquivo %s. Error %s\n", local, err)
	}
}
