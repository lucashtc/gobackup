package exec

import (
	"fmt"
	"os/exec"
)

//Exec function return exec command
func Exec(program string, param []string) ([]byte, error) {

	cmd := exec.Command("mysql", param...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, fmt.Errorf("Falha ao executar command >> %s error >> %s", err, out)
	}
	return out, nil
}
