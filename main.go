package main

//. "github.com/logrusorgru/aurora"
import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	// stop := make(chan bool, 1)
	// stop <- true
	// go helper.Loading(stop)

	// time.Sleep(5 * time.Second)

	// stop <- true
	// file := dir.File{
	// 	Dir:     "C:/Users/oi404011/teste",
	// 	Created: helper.GetCurrentTime(),
	// }

	// err := file.CreateDir()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Param database e tables
	type Param struct {
		BD    string   `yaml:"BD"`
		Table []string `yaml:",Table"`
	}

	// Database info databse
	type Database struct {
		Server string  `yaml:"Server"`
		Param  []Param `yaml:",Param"`
	}

	// file, err := ioutil.ReadFile("./config.yml")
	// if err != nil {
	// 	log.Fatal("Error ao encotrar o arquivo =>", err)
	// }
	t := Database{
		Server : "SGE",
		Param: [0]{
			DB: "acompanhamento",
			Param: []
		}
	}

	yaml.Unmarshal([]byte(file), &t)

	fmt.Println(t)

}
