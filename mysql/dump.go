package mysql

// Database info databse
type Database struct {
	Server string  `yaml:"Server"`
	Param  []Param `yaml:"Param"`
}

// Param database e tables
type Param struct {
	BD    string   `yaml:"BD"`
	Table []string `yaml:"Table"`
}

// DumpAll dump all database
func (d Database) DumpAll() {
	// cmd := exec.Command("mysqldump", "d.")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
}

// Dump databse and tables
func (d Database) Dump() {

}
