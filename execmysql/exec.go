package execmysql

import (
	"fmt"
	"os/exec"
)

//Exec function return exec command
func Exec(param []string) ([]byte, error) {

	cmd := exec.Command("C:/xampp/mysql/bin/mysql", param...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, fmt.Errorf("Falha ao executar comando >> %s error >> %s", err, out)
	}
	return out, nil
	//fmt.Printf("%+q", strings.Split(string(out), "\n"))
}
