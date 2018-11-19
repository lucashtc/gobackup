// Package mysql  vai egar todas as infrmações das base de dados e fazer dump
package mysql

import (
	"github.com/lucashtc/gobackup/dir"
)

// DirDump function create directory for save files backup
// Separando por pastas cadas tido de objeto do schema
func DirDump(dataBase string) (string, error) {

	d, err := dir.CreateDir(dataBase, "")
	if err != nil {
		return "", err
	}

	return d, nil
}

// DumpAll function vai realizar dump de toda o Schema encontrado na base
// func DumpAll(bd []DataBase) error {

// }
