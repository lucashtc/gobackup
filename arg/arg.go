package arg

import (
	"log"
	"os"

	"github.com/lucashtc/gobackup/mysql"

	"github.com/urfave/cli"
)

// Arg statement param, define option backup
func Arg() {
	var conf mysql.DataBase

	app := cli.NewApp()

	app.Name = "gobackup"
	app.Usage = "Script para fazer backup da base de dados local"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "local, l",
			Usage:       "Define local onde o backup será salvo, por padrão será /'backup/'",
			Destination: &conf.Dir,
		},
		cli.StringFlag{
			Name:        "user, u",
			Usage:       "Usuario",
			Destination: &conf.User,
		},
		cli.StringFlag{
			Name:        "password, p",
			Usage:       "Password das bases",
			Destination: &conf.Password,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "all",
			Usage: "Será feito backup de toda a base de dados. Essa Opção é padrão :-)",
			Action: func(c *cli.Context) error {
				mysql.DumpAll(&conf)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
