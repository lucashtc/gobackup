package main

import (
	"runtime"

	"github.com/lucashtc/gobackup/arg"
)

//. "github.com/logrusorgru/aurora"

func main() {
	maxProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(maxProcs)
	arg.Arg()
}
