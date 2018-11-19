package dir

// package responsavel por criar pastas onde será armazenado os backups.
// Cria pasta root e pastas separadas por data de backup e pastas para cada banco
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
func CreateDir(name string, path string) (string, error) {
	subFolder := []string{"tabela", "function"}
	// If not exist, define default
	if path == "" {
		path = "/backup/"
	}
	// Get date and hour for name folder
	created := helper.GetCurrentTime()

	// Validate if f.Dir is empty
	if name == "" {
		return "", fmt.Errorf("Parametro name não pode ser vazio")
	}

	dirStmt := fmt.Sprintf("%s/backup_%s/%s", path, created, name)
	err := os.MkdirAll(dirStmt, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Falha ao criar pasta %s >> %s", dirStmt, err)
	}

	for _, v := range subFolder {
		SubFolderStmt := fmt.Sprintf("%s/%s", dirStmt, v)
		err := os.MkdirAll(SubFolderStmt, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("Falha ao criar pasta %s >> %s", v, err)
		}
	}

	return dirStmt, nil
}
