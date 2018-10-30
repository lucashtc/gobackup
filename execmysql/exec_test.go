package execmysql

import (
	"testing"
)

func Test_Exec(t *testing.T) {
	co := []string{"-B", "-s", "-u", "root", "-e", "SHOW DATABASES"}

	_, err := Exec(co)
	if err != nil {
		t.Fatalf("Error %s \n", err)
	}

}
