package main

import (
	"os"

	"github.com/lucashtc/gobackup/arg"
)

//. "github.com/logrusorgru/aurora"

func main() {
	arg.Arg(os.Args)
}
