package arg

import "testing"

func Test_statmentArgs(t *testing.T) {
	v := make([][]string, 4)
	v[0] = []string{"Nome/Programa", "HELP"}
	v[1] = []string{"Nome/Programa", "NADAs asd a "}
	v[2] = []string{"Nome/Programa", "SOME"}
	v[3] = []string{"Nome/Programa", "HELP"}

	for _, e := range v {
		r := statmentArgs(e)
		t.Logf("GG >> %s ", r)
	}
}

func Test_Arg(t *testing.T) {
	v := make([][]string, 4)
	v[0] = []string{"Nome/Programa", "HELP"}
	v[1] = []string{"Nome/Programa", "NADAs asd a "}
	v[2] = []string{"Nome/Programa", "SOME"}
	v[3] = []string{"Nome/Programa", "HELP"}

	for _, va := range v {

		//Sim, preciso melhorar os testes
		Arg(va)
	}
}
