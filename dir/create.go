package dir

import (
	"fmt"
	"os"

	"github.com/lucashtc/gobackup/helper"
)

// File Data for created dir and files
type File struct {
	Name    string
	Dir     string
	Ext     string
	Created string
}

//CreateDir Diretory for save files sql
// name nome do banco
// path é a pasta onde ficará os backups
func CreateDir(name string, path string) (*File, error) {
	var f *File
	// If not exist define default
	if path == "" {
		path = "/backup/"
	}
	// Get date e hour
	created := helper.GetCurrentTime()
	//f.Created = created

	// Validate if  f.Dir is empty
	if name == "" {
		return f, fmt.Errorf("Param name não pode ser vazio")
	}

	// If not exist folder go create :-)
	exists := pathNotExist(path)
	if exists {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return f, fmt.Errorf("Falha ao criar path>>> %s", err)
		}
	}

	dirStmt := fmt.Sprintf("%s/%s_%s", path, name, created)
	fmt.Println(dirStmt)
	err := os.MkdirAll(dirStmt, os.ModePerm)
	if err != nil {
		return f, fmt.Errorf("Falha ao criar pasta %s >> %s", dirStmt, err)
	}
	//f.Dir = dirStmt
	return f, nil
}

func pathNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
