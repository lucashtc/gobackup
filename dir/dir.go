package dir

// package responsavel por criar pastas onde será armazenado os backups.
// Cria pasta root e pastas separadas por data de backup e pastas para cada banco
import (
	"fmt"
	"os"
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
func CreateDir(time, name, path string) (string, error) {

	// If not exist, define default
	if path == "" {
		path = "/backup/"
	}
	// Validate if f.Dir is empty
	if name == "" {
		return "", fmt.Errorf("Parametro name não pode ser vazio")
	}

	dirStmt := fmt.Sprintf("%s/backup_%s/%s", path, time, name)
	dirStmt = isExist(dirStmt)

	err := os.MkdirAll(dirStmt, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Falha ao criar pasta %s >> %s", dirStmt, err)
	}

	return dirStmt, nil
}

// isExist verify folder exist and define new name
func isExist(n string) string {
	if _, err := os.Stat(n); !os.IsNotExist(err) {
		n = fmt.Sprintf("%s_1", n)
		return n
	}
	return n
}

// Delete ...
func Delete(name string) error {
	if err := os.Remove(name); err != nil {
		return fmt.Errorf("Falha ao deletar arquivo: %s", err)
	}
	return nil
}

// CreateFile ...
func CreateFile(name string) error {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

// WriteFile ...
func WriteFile(name, s string) error {
	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(s)
	if err != nil {
		return err
	}
	if err = file.Sync(); err != nil {
		return err
	}
	return nil
}
