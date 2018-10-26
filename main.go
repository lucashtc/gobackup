package main

import (
	"fmt"

	"github.com/lucashtc/gobackup/mysql"
)

//. "github.com/logrusorgru/aurora"

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

	// file, err := ioutil.ReadFile("./config.yml")
	// if err != nil {
	// 	log.Fatal("Error ao encotrar o arquivo =>", err)
	// }
	// t := []config.Database{
	// 	{
	// 		Server: "SGE",
	// 		Param: []config.Param{
	// 			{
	// 				DB:    "acompanhamento",
	// 				Table: []string{"tabela1", "table2"},
	// 			},
	// 		},
	// 	}, {
	// 		Server: "SGE",
	// 		Param: []config.Param{
	// 			{
	// 				DB:    "acompanhamento",
	// 				Table: []string{"tabela1", "table2"},
	// 			},
	// 		},
	// 	},
	// }
	// New, _ := yaml.Marshal(t)

	// arg := os.Args
	// fmt.Println(arg)
	// file, _ := ioutil.ReadFile("config.yml")
	// t := []config.Database{}

	// err := yaml.Unmarshal(file, &t)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(t)
	//arg.Arg(os.Args)

	m := mysql.DB{}
	_, err := m.GetTables("acompanhamento")
	if err != nil {
		fmt.Printf("Deu ruin %s", err)
	}
	fmt.Println(m.Table)
}
