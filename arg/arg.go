package arg

import (
	"fmt"
	"log"

	"github.com/lucashtc/gobackup/mysql"
)

// Arg statement param, define option backup
func Arg(arg []string) {
	args := statementArgs(arg)
	v := args[0]
	switch v {
	case "HELP":
		fmt.Print(HELP, "\n")
		break
	case "SOME":
		fmt.Printf("Vai chamar func para executar a SOME %s", v)
		break
	case "ALL":
		mysql.DumpAll()
		break
	default:
		fmt.Printf(ERROR, v, HELP)
	}
}

// statementArgs vai pegar os parametros e tratar
// retorna param removendo o nome do programa verificando se esta vazio
func statementArgs(arg []string) []string {
	args := arg[1:]
	if len(args) == 0 {
		log.Fatal(HELP)
	}

	return args
}
