package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Database info databse
type Database struct {
	Server string  `yaml:"Server"`
	Param  []Param `yaml:"Param"`
}

// Param database e tables
type Param struct {
	DB    string   `yaml:"DB"`
	Table []string `yaml:"Table"`
}

// InitConfig initilize config white file yml
// case not found use config default
func InitConfig() error {
	var conf []Database

	file, err := ioutil.ReadFile("./../config.yml")
	if err != nil {
		return fmt.Errorf("Falha ao ler arquivos de configuração >> %s", err)
	}

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return fmt.Errorf("Falha ao Unmarshal para struct >> %s", err)
	}
	return nil
}
