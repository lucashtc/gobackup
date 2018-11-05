package mysql

import (
	"fmt"
	
)

// Database info databse
type Database struct {
	Server string  `yaml:"Server"`
	Param  []Param `yaml:"Param"`
}

// Param database e tables
type Param struct {
	BD        string   `yaml:"BD"`
	Table     []string `yaml:"Table"`
	Procedure []string `yaml:"Procedure"`
}

// DumpAll dump all database
func (d Database) DumpAll() error {

	go func(){
		err := CreateDir()
		if err != nil {
			return err
		}error
	}
}

// Dump databse and tables
func (d Database) Dump() {

}
