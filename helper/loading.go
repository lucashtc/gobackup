package helper

import (
	"fmt"
	"time"
)

// Loading vai mostrar mensagem de carregando
// Stop loading call function param boolean
func Loading(stop chan bool) {
	for <-stop {
		fmt.Print("Carregando |\r")
		time.Sleep(500 * time.Millisecond)

		fmt.Print("Carregando /\r")
		time.Sleep(500 * time.Millisecond)

		fmt.Print("Carregando -\r")
		time.Sleep(500 * time.Millisecond)

		fmt.Print("Carregando \\")
		fmt.Print("\r")
		time.Sleep(500 * time.Millisecond)

		fmt.Print("Carregando |\r")
		time.Sleep(500 * time.Millisecond)

		fmt.Print("Carregando /\r")
		time.Sleep(500 * time.Millisecond)

		fmt.Print("Carregando -\r")
		time.Sleep(500 * time.Millisecond)

	}
}
