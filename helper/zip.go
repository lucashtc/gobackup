package helper

import (
	"fmt"

	"github.com/mholt/archiver"
)

// Zip func zip files
// Param f local and name file
func Zip(f string) error {

	err := archiver.Archive([]string{f}, f+".zip")
	if err != nil {
		return fmt.Errorf("Falha ao Zipar file %s. Error:: %s", f, err)
	}

	return nil
}
