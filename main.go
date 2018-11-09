package main

import (
	"github.com/lucashtc/gobackup/mysql"
)

//. "github.com/logrusorgru/aurora"

func main() {
	var dump mysql.Database

	dump.DumpAll()

}
