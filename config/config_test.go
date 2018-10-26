package config

import (
	"testing"
)

// Tentaiva de testing
func Test_InitConfig(t *testing.T) {
	err := InitConfig()
	if err != nil {
		t.Errorf("Deu ruin %s \n", err)
	}
}
