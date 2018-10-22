package config

// Database info databse
type Database struct {
	Server string  `yaml:"Server"`
	Param  []Param `yaml:",Param"`
}

// Param database e tables
type Param struct {
	BD    string   `yaml:"BD"`
	Table []string `yaml:",Table"`
}

func (d Database) GetConfig() {

}
