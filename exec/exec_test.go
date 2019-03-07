package exec

import (
	"testing"
)

func TestExec(t *testing.T) {

	exemple := []struct {
		program string
		param   []string
	}{
		{"ping", []string{"8.8.8.8"}},
		{"go", []string{"version"}},
	}

	for _, e := range exemple {
		_, err := Exec(e.program, e.param)
		if err != nil {
			t.Error(err)
		}
	}
}
