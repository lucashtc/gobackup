package mysql

import (
	"fmt"

	"github.com/lucashtc/gobackup/dir"
)

// CreateDir vai criar as pastas com base no nome dos bancos
// que serão passados por paramentros
func CreateDir([]string) error {
	var (
		f  *dir.File
		db *DB
	)

	_, err := db.GetDatabase()
	if err != nil {
		return err
	}

	for _, b := range db.Database {

		fmt.Printf("Criando pasta para o banco - %s \n", b)
		_, err = dir.CreateDir(b, "/backup/")
		if err != nil {
			return err
		}
	}

}
