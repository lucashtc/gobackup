package arg

import (
	"fmt"
)

// Arg statment param, define option backup
func Arg(arg []string) {
	args := statmentArgs(arg)
	v := args[0]
	switch v {
	case "HELP":
		fmt.Print(HELP, "\n")
		break
	case "SOME":
		fmt.Printf("Vai chamar func para executar a SOME %s", v)
		break
	case "ALL":
		fmt.Printf("Vai executar o que foi definido no arquivo config.yml %s", v)
		break
	default:
		fmt.Printf(ERROR, v, HELP)
	}
}

// statmentArgs vai pegar os parametros e tratar
// retorna param removendo o nome do programa verificando se esta vazio
func statmentArgs(arg []string) []string {
	args := arg[1:]
	if len(args) == 0 {
		fmt.Printf(EMPTY, HELP)
		return []string{}
	}

	return args
}
