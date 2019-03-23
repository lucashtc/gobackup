package helper

import (
	"fmt"
	"os/exec"
)

//Exec function return exec command
func Exec(program string, param []string) ([]byte, error) {

	cmd := exec.Command(program, param...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Fail e command >> %s error >> %s", err, out)
	}
	return out, nil
}
