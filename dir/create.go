package dir

import (
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora"
)

// File Data for created dir and files
type File struct {
	Name    string
	Dir     string
	Ext     string
	Created string
}

//CreateDir Diretory for save files sql
func (f File) CreateDir() error {
	// Validate if  f.Dir is empty
	if f.Dir == "" {
		return fmt.Errorf("Atributo Dir não pode ser vazio\n")
	}

	dir := fmt.Sprintf("%s_%s", f.Dir, f.Created)
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("Pasta %s criada!\n", Green(f.Dir))
	return nil
}
